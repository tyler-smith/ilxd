// Copyright (c) 2022 The illium developers
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package blockchain

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/project-illium/ilxd/params"
	"github.com/project-illium/ilxd/repo"
	"github.com/project-illium/ilxd/types"
	"github.com/project-illium/ilxd/types/blocks"
	"github.com/project-illium/ilxd/types/transactions"
	"github.com/project-illium/weightedrand/v2"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	// TODO: decide on a value for this
	maxTimeBetweenFlushes = time.Minute * 15

	ValidatorExpiration = time.Hour * 24 * 7 * 26
)

// setConsistencyStatus (SCS) codes are used to indicate the
// consistency status of the validator set state in the database.
type setConsistencyStatus uint8

const (
	// scsEmpty is used as a return value to indicate that no status was
	// stored.  The zero value should not be stored in the database.
	scsEmpty setConsistencyStatus = iota

	// scsConsistent indicates that the set is consistent with the last
	// flush hash stored in the database.
	scsConsistent

	// scsFlushOngoing indicates a flush is ongoing. If a node states with this
	// state it means it must have crashed in the middle of a flush.
	scsFlushOngoing

	// scsNbCodes is the number of valid consistency status codes.
	scsNbCodes
)

// Stake represents a given staked utxo and the time at which it
// was added to the validator set.
type Stake struct {
	Amount     types.Amount
	Blockstamp time.Time
}

// Validator holds all the information about a validator in the ValidatorSet
// that is needed to validate blocks.
type Validator struct {
	PeerID           peer.ID
	TotalStake       types.Amount
	Nullifiers       map[types.Nullifier]Stake
	UnclaimedCoins   types.Amount
	EpochBlocks      uint32
	stakeAccumulator float64
}

// Clone returns a copy of the validator
func (v *Validator) Clone() *Validator {
	ret := &Validator{
		PeerID:           v.PeerID,
		TotalStake:       v.TotalStake,
		Nullifiers:       make(map[types.Nullifier]Stake),
		UnclaimedCoins:   v.UnclaimedCoins,
		stakeAccumulator: v.stakeAccumulator,
		EpochBlocks:      v.EpochBlocks,
	}
	for n, s := range v.Nullifiers {
		ret.Nullifiers[n.Clone()] = Stake{
			Amount:     s.Amount,
			Blockstamp: s.Blockstamp,
		}
	}
	return ret
}

// ValidatorSet maintains the current state of the set of validators in
// the network.
//
// We maintain a memory cache and periodically flush changes to disk.
// This is done because a change to the set may cancel out a previous
// change and committing one to disk only to have to go to disk and undo
// the change would increase overhead.
type ValidatorSet struct {
	params        *params.NetworkParams
	ds            repo.Datastore
	validators    map[peer.ID]*Validator
	nullifierMap  map[types.Nullifier]*Validator
	toDelete      map[peer.ID]struct{}
	chooser       *weightedrand.Chooser[peer.ID, types.Amount]
	EpochBlocks   uint32
	lastFlush     time.Time
	notifications []NotificationCallback
	notifMtx      sync.RWMutex
	mtx           sync.RWMutex
}

// NewValidatorSet returns a new, uninitialized, ValidatorSet.
func NewValidatorSet(params *params.NetworkParams, ds repo.Datastore) *ValidatorSet {
	vs := &ValidatorSet{
		params:       params,
		ds:           ds,
		validators:   make(map[peer.ID]*Validator),
		nullifierMap: make(map[types.Nullifier]*Validator),
		toDelete:     make(map[peer.ID]struct{}),
		mtx:          sync.RWMutex{},
	}
	return vs
}

// Init initializes the validator set. We check to make sure that the set is consistent
// with the tip of the blockchain. In the event of a hard shutdown or shutdown in the middle
// of a flush, the state could be inconsistent. If this is the case, Init will attempt
// to repair the validator set. Depending on the severity of the problem, repair could
// take a while as we may need to rebuild the set from genesis.
func (vs *ValidatorSet) Init(tip *blockNode) error {
	consistencyStatus, err := dsFetchValidatorSetConsistencyStatus(vs.ds)
	if err != nil {
		return err
	}
	lastFlushHeight, err := dsFetchValidatorLastFlushHeight(vs.ds)
	if err != nil {
		return err
	}

	switch consistencyStatus {
	case scsConsistent:
		validators, err := dsFetchValidators(vs.ds)
		if err != nil {
			return err
		}
		for _, val := range validators {
			vs.validators[val.PeerID] = val
			for nullifier := range val.Nullifiers {
				vs.nullifierMap[nullifier] = val
			}
			vs.EpochBlocks += val.EpochBlocks
		}
		if lastFlushHeight == tip.Height() {
			// We're good
			break
		} else if lastFlushHeight < tip.Height() {
			// Load the missing blocks from disk and
			// apply any changes to the validator set.
			// Build the nullifier map.
			var (
				node = tip
				err  error
			)
			for {
				if node.height == lastFlushHeight+1 {
					break
				}
				node, err = node.Parent()
				if err != nil {
					return err
				}
			}
			for {
				blk, err := node.Block()
				if err != nil {
					return err
				}

				validatorReward := types.Amount(0)

				if blk.Header.Height > 0 {
					parent, err := node.Parent()
					if err != nil {
						return err
					}
					prevHeader, err := parent.Header()
					if err != nil {
						return err
					}

					prevEpoch := (prevHeader.Timestamp - vs.params.GenesisBlock.Header.Timestamp) / vs.params.EpochLength
					blkEpoch := (blk.Header.Timestamp - vs.params.GenesisBlock.Header.Timestamp) / vs.params.EpochLength

					if blkEpoch > prevEpoch {
						validatorReward = calculateNextValidatorReward(vs.params, blkEpoch)
					}
				}

				if err := vs.CommitBlock(blk, validatorReward, FlushPeriodic); err != nil {
					return err
				}
				if node.height == tip.height {
					break
				}
				node, err = node.Child()
				if err != nil {
					return err
				}
			}
			if err := vs.Flush(FlushRequired, tip.height); err != nil {
				return err
			}
		} else if lastFlushHeight > tip.Height() {
			// This really should never happen.
			// If we're here it's unlikely the tip node
			// has any attached children that we can use
			// to load the blocks and remove the validator
			// changes from the set. Panic?
			log.Fatal("Validator set last flush height ahead of blockchain tip")
		}
	case scsFlushOngoing:
		// Unfortunately we can't recover from this without rebuilding
		// from genesis.
		log.Warn("Reconstructing validator set after unclean shutdown. This may take a while.")
		dbtx, err := vs.ds.NewTransaction(context.Background(), false)
		if err != nil {
			return err
		}
		if err := dsDeleteValidatorSet(dbtx); err != nil {
			return err
		}
		if err := dsPutValidatorLastFlushHeight(dbtx, 0); err != nil {
			return err
		}
		if err := dbtx.Commit(context.Background()); err != nil {
			return err
		}
		if err := dsPutValidatorSetConsistencyStatus(vs.ds, scsEmpty); err != nil {
			return err
		}
		if err := vs.CommitBlock(vs.params.GenesisBlock, 0, FlushRequired); err != nil {
			return err
		}
		return vs.Init(tip)
	case scsEmpty:
		// Nothing to do here
	}

	b := make([]byte, 8)
	crand.Read(b)
	rand.Seed(int64(binary.BigEndian.Uint64(b)))

	choices := make([]weightedrand.Choice[peer.ID, types.Amount], 0, len(vs.validators))
	for peerID, validator := range vs.validators {
		choices = append(choices, weightedrand.NewChoice(peerID, validator.TotalStake))
	}
	// The chooser errors and returns nil when either:
	// - The total weight exceeds a MaxInt64 (The total coins in the network
	//   won't exceed this value for 105 years).
	// - There are zero validators.
	//
	// So we just ignore the error here and let it be nil if there are zero validators.
	// In the WeightedRandomValidator() method we will check for nil before accessing it.
	vs.chooser, _ = weightedrand.NewChooser(choices...)

	return err
}

// GetValidator returns the validator given the ID.
func (vs *ValidatorSet) GetValidator(id peer.ID) (*Validator, error) {
	vs.mtx.RLock()
	defer vs.mtx.RUnlock()

	val, ok := vs.validators[id]
	if !ok {
		return val, errors.New("not found")
	}
	cpy := &Validator{}
	copyValidator(cpy, val)
	return cpy, nil
}

// ValidatorExists returns whether the validator exists in the set.
func (vs *ValidatorSet) ValidatorExists(id peer.ID) bool {
	vs.mtx.RLock()
	defer vs.mtx.RUnlock()

	_, ok := vs.validators[id]
	return ok
}

func (vs *ValidatorSet) SubscribeEvents(callback NotificationCallback) {
	vs.notifMtx.Lock()
	vs.notifications = append(vs.notifications, callback)
	vs.notifMtx.Unlock()
}

func (vs *ValidatorSet) sendNotification(peerID peer.ID, typ NotificationType) {
	vs.notifMtx.RLock()
	n := Notification{Type: typ, Data: peerID}
	for _, callback := range vs.notifications {
		callback(&n)
	}
	vs.notifMtx.RUnlock()
}

// NullifierExists returns whether or not a nullifier exists in the set.
func (vs *ValidatorSet) NullifierExists(nullifier types.Nullifier) bool {
	vs.mtx.RLock()
	defer vs.mtx.RUnlock()

	_, ok := vs.nullifierMap[nullifier]
	return ok
}

// TotalStaked returns the total staked by all validators.
//
// This method is safe for concurrent access.
func (vs *ValidatorSet) TotalStaked() types.Amount {
	vs.mtx.RLock()
	defer vs.mtx.RUnlock()

	return vs.totalStaked()
}

// totalStaked returns the total staked by all validators.
//
// This method is NOT safe for concurrent access.
func (vs *ValidatorSet) totalStaked() types.Amount {
	total := types.Amount(0)
	for _, val := range vs.validators {
		total += val.TotalStake
	}
	return total
}

// CommitBlock commits the changes to the validator set found in the block into the set.
// This function is fully atomic, if an error is returned, no changes are committed.
// It is expected that the block is fully validated before calling this method.
func (vs *ValidatorSet) CommitBlock(blk *blocks.Block, validatorReward types.Amount, flushMode flushMode) error {
	vs.mtx.Lock()
	defer vs.mtx.Unlock()

	updates := make(map[peer.ID]*Validator)
	nullifiersToAdd := make(map[types.Nullifier]peer.ID)
	nullifiersToDelete := make(map[types.Nullifier]struct{})
	blockTime := time.Unix(blk.Header.Timestamp, 0)

	var (
		producerID peer.ID
		err        error
	)
	if blk.Header.Height > 0 {
		producerID, err = peer.IDFromBytes(blk.Header.Producer_ID)
		if err != nil {
			return err
		}
	}

	for _, t := range blk.GetTransactions() {
		switch tx := t.GetTx().(type) {
		case *transactions.Transaction_CoinbaseTransaction:
			if blk.Header.Height > 0 {
				validatorID, err := peer.IDFromBytes(tx.CoinbaseTransaction.Validator_ID)
				if err != nil {
					return err
				}

				valNew, ok := updates[validatorID]
				if !ok {
					valOld, ok := vs.validators[validatorID]
					if !ok {
						log.Warn("Coinbase transaction for validator not in set")
					}
					valNew = &Validator{}
					copyValidator(valNew, valOld)
				}
				if types.Amount(tx.CoinbaseTransaction.NewCoins) >= valNew.UnclaimedCoins {
					valNew.UnclaimedCoins = 0
				} else {
					valNew.UnclaimedCoins -= types.Amount(tx.CoinbaseTransaction.NewCoins)
				}
				updates[validatorID] = valNew
			}
		case *transactions.Transaction_StakeTransaction:
			validatorID, err := peer.IDFromBytes(tx.StakeTransaction.Validator_ID)
			if err != nil {
				return err
			}

			valNew, ok := updates[validatorID]
			if !ok {
				valOld, ok := vs.validators[validatorID]
				if !ok {
					valNew = &Validator{
						PeerID:         validatorID,
						TotalStake:     0,
						Nullifiers:     make(map[types.Nullifier]Stake),
						UnclaimedCoins: 0,
						EpochBlocks:    0,
					}
				} else {
					valNew = &Validator{}
					copyValidator(valNew, valOld)
				}
			}
			if _, ok := valNew.Nullifiers[types.NewNullifier(tx.StakeTransaction.Nullifier)]; !ok {
				valNew.TotalStake += types.Amount(tx.StakeTransaction.Amount)
			}
			valNew.Nullifiers[types.NewNullifier(tx.StakeTransaction.Nullifier)] = Stake{
				Amount:     types.Amount(tx.StakeTransaction.Amount),
				Blockstamp: blockTime,
			}
			updates[validatorID] = valNew
			nullifiersToAdd[types.NewNullifier(tx.StakeTransaction.Nullifier)] = validatorID
		case *transactions.Transaction_StandardTransaction:
			for _, nullifier := range tx.StandardTransaction.Nullifiers {
				valOld, ok := vs.nullifierMap[types.NewNullifier(nullifier)]
				if !ok {
					continue
				}
				v, ok := updates[valOld.PeerID]
				if ok {
					valOld = v
				}
				valNew := &Validator{}
				copyValidator(valNew, valOld)
				stake, ok := valNew.Nullifiers[types.NewNullifier(nullifier)]
				if !ok {
					return errors.New("nullifier not found with validator")
				}
				valNew.TotalStake -= stake.Amount
				delete(valNew.Nullifiers, types.NewNullifier(nullifier))
				nullifiersToDelete[types.NewNullifier(nullifier)] = struct{}{}
				updates[valNew.PeerID] = valNew
			}
		case *transactions.Transaction_MintTransaction:
			for _, nullifier := range tx.MintTransaction.Nullifiers {
				valOld, ok := vs.nullifierMap[types.NewNullifier(nullifier)]
				if !ok {
					continue
				}
				v, ok := updates[valOld.PeerID]
				if ok {
					valOld = v
				}
				valNew := &Validator{}
				copyValidator(valNew, valOld)
				stake, ok := valNew.Nullifiers[types.NewNullifier(nullifier)]
				if !ok {
					return errors.New("nullifier not found with validator")
				}
				valNew.TotalStake -= stake.Amount
				delete(valNew.Nullifiers, types.NewNullifier(nullifier))
				nullifiersToDelete[types.NewNullifier(nullifier)] = struct{}{}
				updates[valNew.PeerID] = valNew
			}
		}
	}

	totalStaked := vs.totalStaked()
	if validatorReward > 0 {
		for _, valOld := range vs.validators {
			valNew, ok := updates[valOld.PeerID]
			if !ok {
				valNew = &Validator{}
				copyValidator(valNew, valOld)
			}

			expectedBlocks := valNew.stakeAccumulator
			if expectedBlocks < 1 {
				expectedBlocks = 1
			}
			valTotal := types.Amount(0)
			for nullifier, stake := range valNew.Nullifiers {
				timeSinceStake := blockTime.Sub(stake.Blockstamp)

				if timeSinceStake >= ValidatorExpiration {
					nullifiersToDelete[nullifier] = struct{}{}
					delete(valNew.Nullifiers, nullifier)
					continue
				}

				epochLength := float64(vs.params.EpochLength)
				if timeSinceStake.Seconds() >= epochLength {
					valTotal += stake.Amount
				} else {
					valTotal += types.Amount(float64(stake.Amount) * (float64(timeSinceStake.Seconds()) / epochLength))
				}
			}
			if valTotal > 0 {
				valNew.UnclaimedCoins = valNew.UnclaimedCoins + types.Amount(float64(validatorReward)*(float64(valTotal)/float64(totalStaked)))
			}

			if valNew.EpochBlocks > blockProductionLimit(float64(vs.EpochBlocks), expectedBlocks/float64(vs.EpochBlocks)) {
				valNew.UnclaimedCoins = 0
			}

			if valNew.EpochBlocks < blockProductionFloor(float64(vs.EpochBlocks), expectedBlocks/float64(vs.EpochBlocks)) {
				valNew.UnclaimedCoins = 0
			}

			valNew.EpochBlocks = 0
			valNew.stakeAccumulator = 0
			updates[valNew.PeerID] = valNew
		}
		vs.EpochBlocks = 0
	}

	if blk.Header.Height > 0 {
		blockProducer, ok := vs.validators[producerID]
		if ok {
			blockProducer.EpochBlocks++
		}
		for _, val := range vs.validators {
			val.stakeAccumulator += float64(val.TotalStake) / float64(totalStaked)
		}
	}

	for _, val := range updates {
		if _, ok := vs.validators[val.PeerID]; !ok {
			vs.sendNotification(val.PeerID, NTAddValidator)
		}
		vs.validators[val.PeerID] = val
	}

	for nullifier, peerID := range nullifiersToAdd {
		val := vs.validators[peerID]
		vs.nullifierMap[nullifier] = val
	}

	for nullifier := range nullifiersToDelete {
		delete(vs.nullifierMap, nullifier)
	}

	for _, val := range updates {
		if len(val.Nullifiers) == 0 {
			vs.toDelete[val.PeerID] = struct{}{}
			delete(vs.validators, val.PeerID)
			vs.sendNotification(val.PeerID, NTRemoveValidator)
		}
	}

	if len(updates) > 0 || len(nullifiersToAdd) > 0 || len(nullifiersToDelete) > 0 {
		choices := make([]weightedrand.Choice[peer.ID, types.Amount], 0, len(vs.validators))
		for peerID, validator := range vs.validators {
			choices = append(choices, weightedrand.NewChoice(peerID, validator.TotalStake))
		}
		vs.chooser, _ = weightedrand.NewChooser(choices...)
	}

	vs.EpochBlocks++

	return vs.flush(flushMode, blk.Header.Height)
}

// WeightedRandomValidator returns a validator weighted by their current stake.
//
// NOTE: If there are no validators then "" will be returned for the peer ID.
func (vs *ValidatorSet) WeightedRandomValidator() peer.ID {
	vs.mtx.RLock()
	defer vs.mtx.RUnlock()

	if vs.totalStaked() == 0 {
		return ""
	}

	return vs.chooser.Pick()
}

// BlockProductionLimit returns the maximum blocks that a validator can produce without losing
// this coinbase. This is based on the current snapshot state of the set and changes every
// block.
func (vs *ValidatorSet) BlockProductionLimit(validatorID peer.ID) (uint32, uint32, error) {
	vs.mtx.RLock()
	defer vs.mtx.RUnlock()

	val, ok := vs.validators[validatorID]
	if !ok {
		return 0, 0, errors.New("not found")
	}
	if vs.EpochBlocks == 0 {
		return val.EpochBlocks, 1, nil
	}
	expectedBlocks := val.stakeAccumulator
	if expectedBlocks < 1 {
		expectedBlocks = 1
	}
	return val.EpochBlocks, blockProductionLimit(float64(vs.EpochBlocks), expectedBlocks/float64(vs.EpochBlocks)), nil
}

// Flush flushes changes from the memory cache to disk.
//
// This method is safe for concurrent access.
func (vs *ValidatorSet) Flush(mode flushMode, chainHeight uint32) error {
	vs.mtx.Lock()
	defer vs.mtx.Unlock()

	return vs.flush(mode, chainHeight)
}

// flush flushes changes from the memory cache to disk.
//
// This method is NOT safe for concurrent access.
func (vs *ValidatorSet) flush(mode flushMode, chainHeight uint32) error {
	switch mode {
	case FlushRequired:
		return vs.flushToDisk(chainHeight)
	case FlushPeriodic:
		if vs.lastFlush.Add(maxTimeBetweenFlushes).Before(time.Now()) {
			return vs.flushToDisk(chainHeight)
		}
		return nil
	case FlushNop:
		return nil
	default:
		return fmt.Errorf("unsupported flushmode for the validator set")
	}
}

func (vs *ValidatorSet) flushToDisk(chainHeight uint32) error {
	if err := dsPutValidatorSetConsistencyStatus(vs.ds, scsFlushOngoing); err != nil {
		return err
	}
	dbtx, err := vs.ds.NewTransaction(context.Background(), false)
	if err != nil {
		return err
	}
	defer dbtx.Discard(context.Background())
	for _, val := range vs.validators {
		if err := dsPutValidator(dbtx, val); err != nil {
			return err
		}
	}
	for id := range vs.toDelete {
		if err := dsDeleteValidator(dbtx, id); err != nil {
			return err
		}
	}

	if err := dsPutValidatorLastFlushHeight(dbtx, chainHeight); err != nil {
		return err
	}

	if err := dbtx.Commit(context.Background()); err != nil {
		return err
	}

	if err := dsPutValidatorSetConsistencyStatus(vs.ds, scsConsistent); err != nil {
		return err
	}

	vs.toDelete = make(map[peer.ID]struct{})

	vs.lastFlush = time.Now()

	return nil
}

// Six standard deviations from the expected number of blocks.
func blockProductionLimit(EpochBlocks float64, stakePercentage float64) uint32 {
	x := float64(EpochBlocks * stakePercentage)
	y := float64(math.Sqrt(EpochBlocks*stakePercentage*(1-stakePercentage)) * 6)
	if stakePercentage == 1 {
		y = 1
	}
	return uint32(x + y)
}

// Six standard deviations from the expected number of blocks.
func blockProductionFloor(EpochBlocks float64, stakePercentage float64) uint32 {
	expectedBlocks := uint32(EpochBlocks * stakePercentage)
	dev := uint32(math.Sqrt(EpochBlocks*stakePercentage*(1-stakePercentage)) * 6)
	if dev > expectedBlocks {
		return 0
	}
	return (expectedBlocks - dev) / 3
}

func copyValidator(dest *Validator, src *Validator) {
	dest.PeerID = src.PeerID
	dest.TotalStake = src.TotalStake
	dest.EpochBlocks = src.EpochBlocks
	dest.UnclaimedCoins = src.UnclaimedCoins
	dest.Nullifiers = make(map[types.Nullifier]Stake)
	for k, v := range src.Nullifiers {
		dest.Nullifiers[k.Clone()] = v
	}
}

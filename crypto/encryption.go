// Copyright (c) 2022 Project Illium
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/rand"
	"errors"
	"github.com/libp2p/go-libp2p/core/crypto"
	"golang.org/x/crypto/nacl/box"
)

const (
	// Length of nacl nonce
	NonceBytes = 24

	// Length of nacl ephemeral public key
	EphemeralPublicKeyBytes = 32
)

// ErrBoxDecryption Nacl box decryption failed
var ErrBoxDecryption = errors.New("failed to decrypt curve25519")

// Encrypt encrypts an output with the public key.
func Encrypt(pubKey crypto.PubKey, plaintext []byte) ([]byte, error) {
	curve25519PubKey, ok := pubKey.(*Curve25519PublicKey)
	if !ok {
		return nil, errors.New("pubkey must be of type Curve25519PublicKey")
	}

	// Generate ephemeral key pair
	ephemPub, ephemPriv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Encrypt with nacl
	var (
		ciphertext []byte
		pt         = make([]byte, len(plaintext))
		nonce      [24]byte
		n          = make([]byte, 24)
	)
	_, err = rand.Read(n)
	if err != nil {
		return nil, err
	}
	copy(nonce[:], n)
	copy(pt, plaintext)

	ciphertext = box.Seal(ciphertext, pt, &nonce, curve25519PubKey.k, ephemPriv)

	// Prepend the ephemeral public key
	ciphertext = append(ephemPub[:], ciphertext...)

	// Prepend nonce
	ciphertext = append(nonce[:], ciphertext...)
	return ciphertext, nil
}

// Decrypt decrypts an output using a private key.
func Decrypt(privKey crypto.PrivKey, ciphertext []byte) ([]byte, error) {
	curve25519PrivKey, ok := privKey.(*Curve25519PrivateKey)
	if !ok {
		return nil, errors.New("privkey must be of type Curve25519PrivateKey")
	}
	n := ciphertext[:NonceBytes]
	ephemPubkeyBytes := ciphertext[NonceBytes : NonceBytes+EphemeralPublicKeyBytes]
	ct := ciphertext[NonceBytes+EphemeralPublicKeyBytes:]

	var (
		plaintext   []byte
		priv        [32]byte
		ephemPubkey [32]byte
		nonce       [24]byte
	)
	copy(ephemPubkey[:], ephemPubkeyBytes)
	copy(nonce[:], n)
	copy(priv[:], curve25519PrivKey.k[:Curve25519PrivateKeySize])

	plaintext, success := box.Open(plaintext, ct, &nonce, &ephemPubkey, &priv)
	if !success {
		return nil, ErrBoxDecryption
	}
	return plaintext, nil
}

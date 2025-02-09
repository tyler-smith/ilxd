// Copyright (c) 2022 Project Illium
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: blocks.proto

package blocks

import (
	transactions "github.com/project-illium/ilxd/types/transactions"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Height      uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Parent      []byte `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Timestamp   int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TxRoot      []byte `protobuf:"bytes,5,opt,name=tx_root,json=txRoot,proto3" json:"tx_root,omitempty"`
	Producer_ID []byte `protobuf:"bytes,6,opt,name=producer_ID,json=producerID,proto3" json:"producer_ID,omitempty"`
	Signature   []byte `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_blocks_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHeader) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockHeader) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockHeader) GetParent() []byte {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *BlockHeader) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockHeader) GetTxRoot() []byte {
	if x != nil {
		return x.TxRoot
	}
	return nil
}

func (x *BlockHeader) GetProducer_ID() []byte {
	if x != nil {
		return x.Producer_ID
	}
	return nil
}

func (x *BlockHeader) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BlockHeader                `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions []*transactions.Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blocks_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*transactions.Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type BlockTxs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*transactions.Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *BlockTxs) Reset() {
	*x = BlockTxs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockTxs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockTxs) ProtoMessage() {}

func (x *BlockTxs) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockTxs.ProtoReflect.Descriptor instead.
func (*BlockTxs) Descriptor() ([]byte, []int) {
	return file_blocks_proto_rawDescGZIP(), []int{2}
}

func (x *BlockTxs) GetTransactions() []*transactions.Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type XThinnerBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BlockHeader                          `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TxCount      uint32                                `protobuf:"varint,2,opt,name=tx_count,json=txCount,proto3" json:"tx_count,omitempty"`
	Pops         []byte                                `protobuf:"bytes,3,opt,name=pops,proto3" json:"pops,omitempty"`
	Pushes       []byte                                `protobuf:"bytes,4,opt,name=pushes,proto3" json:"pushes,omitempty"`
	PushBytes    []byte                                `protobuf:"bytes,5,opt,name=push_bytes,json=pushBytes,proto3" json:"push_bytes,omitempty"`
	PrefilledTxs []*XThinnerBlock_PrefilledTransaction `protobuf:"bytes,6,rep,name=prefilled_txs,json=prefilledTxs,proto3" json:"prefilled_txs,omitempty"`
}

func (x *XThinnerBlock) Reset() {
	*x = XThinnerBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XThinnerBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XThinnerBlock) ProtoMessage() {}

func (x *XThinnerBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XThinnerBlock.ProtoReflect.Descriptor instead.
func (*XThinnerBlock) Descriptor() ([]byte, []int) {
	return file_blocks_proto_rawDescGZIP(), []int{3}
}

func (x *XThinnerBlock) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *XThinnerBlock) GetTxCount() uint32 {
	if x != nil {
		return x.TxCount
	}
	return 0
}

func (x *XThinnerBlock) GetPops() []byte {
	if x != nil {
		return x.Pops
	}
	return nil
}

func (x *XThinnerBlock) GetPushes() []byte {
	if x != nil {
		return x.Pushes
	}
	return nil
}

func (x *XThinnerBlock) GetPushBytes() []byte {
	if x != nil {
		return x.PushBytes
	}
	return nil
}

func (x *XThinnerBlock) GetPrefilledTxs() []*XThinnerBlock_PrefilledTransaction {
	if x != nil {
		return x.PrefilledTxs
	}
	return nil
}

type CompressedBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height     uint32                 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Nullifiers [][]byte               `protobuf:"bytes,2,rep,name=nullifiers,proto3" json:"nullifiers,omitempty"`
	Outputs    []*transactions.Output `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *CompressedBlock) Reset() {
	*x = CompressedBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedBlock) ProtoMessage() {}

func (x *CompressedBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedBlock.ProtoReflect.Descriptor instead.
func (*CompressedBlock) Descriptor() ([]byte, []int) {
	return file_blocks_proto_rawDescGZIP(), []int{4}
}

func (x *CompressedBlock) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *CompressedBlock) GetNullifiers() [][]byte {
	if x != nil {
		return x.Nullifiers
	}
	return nil
}

func (x *CompressedBlock) GetOutputs() []*transactions.Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type XThinnerBlock_PrefilledTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *transactions.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Index       uint32                    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *XThinnerBlock_PrefilledTransaction) Reset() {
	*x = XThinnerBlock_PrefilledTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XThinnerBlock_PrefilledTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XThinnerBlock_PrefilledTransaction) ProtoMessage() {}

func (x *XThinnerBlock_PrefilledTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_blocks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XThinnerBlock_PrefilledTransaction.ProtoReflect.Descriptor instead.
func (*XThinnerBlock_PrefilledTransaction) Descriptor() ([]byte, []int) {
	return file_blocks_proto_rawDescGZIP(), []int{3, 0}
}

func (x *XThinnerBlock_PrefilledTransaction) GetTransaction() *transactions.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *XThinnerBlock_PrefilledTransaction) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_blocks_proto protoreflect.FileDescriptor

var file_blocks_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x5f, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x78, 0x73, 0x12,
	0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xc3, 0x02, 0x0a, 0x0d, 0x58, 0x54, 0x68, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x78, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x70, 0x6f, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x73, 0x68,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x73, 0x68, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x73, 0x68, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x48, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x78, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x58, 0x54, 0x68, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x54, 0x78, 0x73, 0x1a, 0x5c, 0x0a, 0x14, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x6c, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocks_proto_rawDescOnce sync.Once
	file_blocks_proto_rawDescData = file_blocks_proto_rawDesc
)

func file_blocks_proto_rawDescGZIP() []byte {
	file_blocks_proto_rawDescOnce.Do(func() {
		file_blocks_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocks_proto_rawDescData)
	})
	return file_blocks_proto_rawDescData
}

var file_blocks_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_blocks_proto_goTypes = []interface{}{
	(*BlockHeader)(nil),                        // 0: BlockHeader
	(*Block)(nil),                              // 1: Block
	(*BlockTxs)(nil),                           // 2: BlockTxs
	(*XThinnerBlock)(nil),                      // 3: XThinnerBlock
	(*CompressedBlock)(nil),                    // 4: CompressedBlock
	(*XThinnerBlock_PrefilledTransaction)(nil), // 5: XThinnerBlock.PrefilledTransaction
	(*transactions.Transaction)(nil),           // 6: Transaction
	(*transactions.Output)(nil),                // 7: Output
}
var file_blocks_proto_depIdxs = []int32{
	0, // 0: Block.header:type_name -> BlockHeader
	6, // 1: Block.transactions:type_name -> Transaction
	6, // 2: BlockTxs.transactions:type_name -> Transaction
	0, // 3: XThinnerBlock.header:type_name -> BlockHeader
	5, // 4: XThinnerBlock.prefilled_txs:type_name -> XThinnerBlock.PrefilledTransaction
	7, // 5: CompressedBlock.outputs:type_name -> Output
	6, // 6: XThinnerBlock.PrefilledTransaction.transaction:type_name -> Transaction
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_blocks_proto_init() }
func file_blocks_proto_init() {
	if File_blocks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blocks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blocks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blocks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockTxs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blocks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XThinnerBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blocks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blocks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XThinnerBlock_PrefilledTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blocks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blocks_proto_goTypes,
		DependencyIndexes: file_blocks_proto_depIdxs,
		MessageInfos:      file_blocks_proto_msgTypes,
	}.Build()
	File_blocks_proto = out.File
	file_blocks_proto_rawDesc = nil
	file_blocks_proto_goTypes = nil
	file_blocks_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: domain/block/model/block.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockStatus int32

const (
	BlockStatus_BLOCK_STATUS_UNSPECIFIED BlockStatus = 0
	BlockStatus_BLOCK_STATUS_STABLE      BlockStatus = 1
	BlockStatus_BLOCK_STATUS_UNSTABLE    BlockStatus = 2
	BlockStatus_BLOCK_STATUS_REORG       BlockStatus = 3
)

// Enum value maps for BlockStatus.
var (
	BlockStatus_name = map[int32]string{
		0: "BLOCK_STATUS_UNSPECIFIED",
		1: "BLOCK_STATUS_STABLE",
		2: "BLOCK_STATUS_UNSTABLE",
		3: "BLOCK_STATUS_REORG",
	}
	BlockStatus_value = map[string]int32{
		"BLOCK_STATUS_UNSPECIFIED": 0,
		"BLOCK_STATUS_STABLE":      1,
		"BLOCK_STATUS_UNSTABLE":    2,
		"BLOCK_STATUS_REORG":       3,
	}
)

func (x BlockStatus) Enum() *BlockStatus {
	p := new(BlockStatus)
	*p = x
	return p
}

func (x BlockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_block_model_block_proto_enumTypes[0].Descriptor()
}

func (BlockStatus) Type() protoreflect.EnumType {
	return &file_domain_block_model_block_proto_enumTypes[0]
}

func (x BlockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockStatus.Descriptor instead.
func (BlockStatus) EnumDescriptor() ([]byte, []int) {
	return file_domain_block_model_block_proto_rawDescGZIP(), []int{0}
}

type BlockRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"block_num"
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"block_num"`
	// @gotags: json:"block_hash"
	Hash       string `protobuf:"bytes,2,opt,name=hash,proto3" json:"block_hash"`
	ParentHash string `protobuf:"bytes,3,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	// @gotags: json:"transactions,omitempty"
	Transactions []*Transaction `protobuf:"bytes,4,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// @gotags: json:"-"
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"-"`
	Depth     int64                  `protobuf:"varint,6,opt,name=depth,proto3" json:"depth,omitempty"`
	Status    BlockStatus            `protobuf:"varint,7,opt,name=status,proto3,enum=block.BlockStatus" json:"status,omitempty"`
}

func (x *BlockRecord) Reset() {
	*x = BlockRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_model_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRecord) ProtoMessage() {}

func (x *BlockRecord) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_model_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRecord.ProtoReflect.Descriptor instead.
func (*BlockRecord) Descriptor() ([]byte, []int) {
	return file_domain_block_model_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRecord) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockRecord) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockRecord) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockRecord) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *BlockRecord) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *BlockRecord) GetDepth() int64 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *BlockRecord) GetStatus() BlockStatus {
	if x != nil {
		return x.Status
	}
	return BlockStatus_BLOCK_STATUS_UNSPECIFIED
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_model_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_model_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_domain_block_model_block_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Event) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash string `protobuf:"bytes,8,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// @gotags: json:"tx_hash"
	Hash  string `protobuf:"bytes,1,opt,name=hash,proto3" json:"tx_hash"`
	From  string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Data  string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Value string `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// @gotags: json:"logs,omitempty"
	Events []*Event `protobuf:"bytes,7,rep,name=events,proto3" json:"logs,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_block_model_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_domain_block_model_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_domain_block_model_block_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Transaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Transaction) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_domain_block_model_block_proto protoreflect.FileDescriptor

var file_domain_block_model_block_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xca, 0x01, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x77, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4f, 0x52, 0x47,
	0x10, 0x03, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79, 0x61, 0x2f, 0x65, 0x74,
	0x68, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_block_model_block_proto_rawDescOnce sync.Once
	file_domain_block_model_block_proto_rawDescData = file_domain_block_model_block_proto_rawDesc
)

func file_domain_block_model_block_proto_rawDescGZIP() []byte {
	file_domain_block_model_block_proto_rawDescOnce.Do(func() {
		file_domain_block_model_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_block_model_block_proto_rawDescData)
	})
	return file_domain_block_model_block_proto_rawDescData
}

var file_domain_block_model_block_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_block_model_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_domain_block_model_block_proto_goTypes = []interface{}{
	(BlockStatus)(0),              // 0: block.BlockStatus
	(*BlockRecord)(nil),           // 1: block.BlockRecord
	(*Event)(nil),                 // 2: block.Event
	(*Transaction)(nil),           // 3: block.Transaction
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_domain_block_model_block_proto_depIdxs = []int32{
	3, // 0: block.BlockRecord.transactions:type_name -> block.Transaction
	4, // 1: block.BlockRecord.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: block.BlockRecord.status:type_name -> block.BlockStatus
	2, // 3: block.Transaction.events:type_name -> block.Event
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_domain_block_model_block_proto_init() }
func file_domain_block_model_block_proto_init() {
	if File_domain_block_model_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_block_model_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRecord); i {
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
		file_domain_block_model_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_domain_block_model_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_domain_block_model_block_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_block_model_block_proto_goTypes,
		DependencyIndexes: file_domain_block_model_block_proto_depIdxs,
		EnumInfos:         file_domain_block_model_block_proto_enumTypes,
		MessageInfos:      file_domain_block_model_block_proto_msgTypes,
	}.Build()
	File_domain_block_model_block_proto = out.File
	file_domain_block_model_block_proto_rawDesc = nil
	file_domain_block_model_block_proto_goTypes = nil
	file_domain_block_model_block_proto_depIdxs = nil
}

//*
// Broadcast service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: broadcast.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Requests
type GetBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchNumber uint64 `protobuf:"varint,1,opt,name=batch_number,json=batchNumber,proto3" json:"batch_number,omitempty"`
}

func (x *GetBatchRequest) Reset() {
	*x = GetBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchRequest) ProtoMessage() {}

func (x *GetBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchRequest.ProtoReflect.Descriptor instead.
func (*GetBatchRequest) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *GetBatchRequest) GetBatchNumber() uint64 {
	if x != nil {
		return x.BatchNumber
	}
	return 0
}

// Responses
type GetBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchNumber             uint64         `protobuf:"varint,1,opt,name=batch_number,json=batchNumber,proto3" json:"batch_number,omitempty"`
	GlobalExitRoot          string         `protobuf:"bytes,2,opt,name=global_exit_root,json=globalExitRoot,proto3" json:"global_exit_root,omitempty"`
	LocalExitRoot           string         `protobuf:"bytes,3,opt,name=local_exit_root,json=localExitRoot,proto3" json:"local_exit_root,omitempty"`
	StateRoot               string         `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	MainnetExitRoot         string         `protobuf:"bytes,5,opt,name=mainnet_exit_root,json=mainnetExitRoot,proto3" json:"mainnet_exit_root,omitempty"`
	RollupExitRoot          string         `protobuf:"bytes,6,opt,name=rollup_exit_root,json=rollupExitRoot,proto3" json:"rollup_exit_root,omitempty"`
	GlobalExitRootTimestamp uint64         `protobuf:"varint,7,opt,name=global_exit_root_timestamp,json=globalExitRootTimestamp,proto3" json:"global_exit_root_timestamp,omitempty"`
	Timestamp               uint64         `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sequencer               string         `protobuf:"bytes,9,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
	ForcedBatchNumber       uint64         `protobuf:"varint,10,opt,name=forced_batch_number,json=forcedBatchNumber,proto3" json:"forced_batch_number,omitempty"`
	Transactions            []*Transaction `protobuf:"bytes,11,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetBatchResponse) Reset() {
	*x = GetBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchResponse) ProtoMessage() {}

func (x *GetBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchResponse.ProtoReflect.Descriptor instead.
func (*GetBatchResponse) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{1}
}

func (x *GetBatchResponse) GetBatchNumber() uint64 {
	if x != nil {
		return x.BatchNumber
	}
	return 0
}

func (x *GetBatchResponse) GetGlobalExitRoot() string {
	if x != nil {
		return x.GlobalExitRoot
	}
	return ""
}

func (x *GetBatchResponse) GetLocalExitRoot() string {
	if x != nil {
		return x.LocalExitRoot
	}
	return ""
}

func (x *GetBatchResponse) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *GetBatchResponse) GetMainnetExitRoot() string {
	if x != nil {
		return x.MainnetExitRoot
	}
	return ""
}

func (x *GetBatchResponse) GetRollupExitRoot() string {
	if x != nil {
		return x.RollupExitRoot
	}
	return ""
}

func (x *GetBatchResponse) GetGlobalExitRootTimestamp() uint64 {
	if x != nil {
		return x.GlobalExitRootTimestamp
	}
	return 0
}

func (x *GetBatchResponse) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetBatchResponse) GetSequencer() string {
	if x != nil {
		return x.Sequencer
	}
	return ""
}

func (x *GetBatchResponse) GetForcedBatchNumber() uint64 {
	if x != nil {
		return x.ForcedBatchNumber
	}
	return 0
}

func (x *GetBatchResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

// Common
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encoded string `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[2]
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
	return file_broadcast_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetEncoded() string {
	if x != nil {
		return x.Encoded
	}
	return ""
}

var File_broadcast_proto protoreflect.FileDescriptor

var file_broadcast_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0xe4, 0x03, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x78, 0x69, 0x74,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78,
	0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x45,
	0x78, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x6c, 0x75,
	0x70, 0x5f, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70, 0x45, 0x78, 0x69, 0x74, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x69, 0x74,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x78, 0x69,
	0x74, 0x52, 0x6f, 0x6f, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x27, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x32, 0xa5, 0x01, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1e, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x30, 0x78, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x7a, 0x2f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x2d, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_broadcast_proto_rawDescOnce sync.Once
	file_broadcast_proto_rawDescData = file_broadcast_proto_rawDesc
)

func file_broadcast_proto_rawDescGZIP() []byte {
	file_broadcast_proto_rawDescOnce.Do(func() {
		file_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_broadcast_proto_rawDescData)
	})
	return file_broadcast_proto_rawDescData
}

var file_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_broadcast_proto_goTypes = []interface{}{
	(*GetBatchRequest)(nil),  // 0: broadcast.v1.GetBatchRequest
	(*GetBatchResponse)(nil), // 1: broadcast.v1.GetBatchResponse
	(*Transaction)(nil),      // 2: broadcast.v1.Transaction
	(*emptypb.Empty)(nil),    // 3: google.protobuf.Empty
}
var file_broadcast_proto_depIdxs = []int32{
	2, // 0: broadcast.v1.GetBatchResponse.transactions:type_name -> broadcast.v1.Transaction
	3, // 1: broadcast.v1.BroadcastService.GetLastBatch:input_type -> google.protobuf.Empty
	0, // 2: broadcast.v1.BroadcastService.GetBatch:input_type -> broadcast.v1.GetBatchRequest
	1, // 3: broadcast.v1.BroadcastService.GetLastBatch:output_type -> broadcast.v1.GetBatchResponse
	1, // 4: broadcast.v1.BroadcastService.GetBatch:output_type -> broadcast.v1.GetBatchResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_broadcast_proto_init() }
func file_broadcast_proto_init() {
	if File_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchRequest); i {
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
		file_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchResponse); i {
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
		file_broadcast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_broadcast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_broadcast_proto_goTypes,
		DependencyIndexes: file_broadcast_proto_depIdxs,
		MessageInfos:      file_broadcast_proto_msgTypes,
	}.Build()
	File_broadcast_proto = out.File
	file_broadcast_proto_rawDesc = nil
	file_broadcast_proto_goTypes = nil
	file_broadcast_proto_depIdxs = nil
}

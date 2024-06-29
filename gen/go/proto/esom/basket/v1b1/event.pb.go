// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: esom/basket/v1beta1/event.proto

package basketpb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type BasketAbandoned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasketId   string                 `protobuf:"bytes,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	OccurredAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
}

func (x *BasketAbandoned) Reset() {
	*x = BasketAbandoned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esom_basket_v1beta1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketAbandoned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketAbandoned) ProtoMessage() {}

func (x *BasketAbandoned) ProtoReflect() protoreflect.Message {
	mi := &file_esom_basket_v1beta1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketAbandoned.ProtoReflect.Descriptor instead.
func (*BasketAbandoned) Descriptor() ([]byte, []int) {
	return file_esom_basket_v1beta1_event_proto_rawDescGZIP(), []int{0}
}

func (x *BasketAbandoned) GetBasketId() string {
	if x != nil {
		return x.BasketId
	}
	return ""
}

func (x *BasketAbandoned) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

type BasketUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasketId   string                 `protobuf:"bytes,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	OccurredAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
}

func (x *BasketUpdated) Reset() {
	*x = BasketUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esom_basket_v1beta1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketUpdated) ProtoMessage() {}

func (x *BasketUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_esom_basket_v1beta1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketUpdated.ProtoReflect.Descriptor instead.
func (*BasketUpdated) Descriptor() ([]byte, []int) {
	return file_esom_basket_v1beta1_event_proto_rawDescGZIP(), []int{1}
}

func (x *BasketUpdated) GetBasketId() string {
	if x != nil {
		return x.BasketId
	}
	return ""
}

func (x *BasketUpdated) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

type BasketViewed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasketId   string                 `protobuf:"bytes,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	OccurredAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
}

func (x *BasketViewed) Reset() {
	*x = BasketViewed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_esom_basket_v1beta1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketViewed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketViewed) ProtoMessage() {}

func (x *BasketViewed) ProtoReflect() protoreflect.Message {
	mi := &file_esom_basket_v1beta1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketViewed.ProtoReflect.Descriptor instead.
func (*BasketViewed) Descriptor() ([]byte, []int) {
	return file_esom_basket_v1beta1_event_proto_rawDescGZIP(), []int{2}
}

func (x *BasketViewed) GetBasketId() string {
	if x != nil {
		return x.BasketId
	}
	return ""
}

func (x *BasketViewed) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

var File_esom_basket_v1beta1_event_proto protoreflect.FileDescriptor

var file_esom_basket_v1beta1_event_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x73, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x65, 0x73, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x41,
	0x62, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08,
	0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0x48, 0x05, 0xb2, 0x01, 0x02, 0x38, 0x01, 0x52, 0x0a, 0x6f,
	0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0x48, 0x05, 0xb2, 0x01, 0x02, 0x38, 0x01,
	0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7f, 0x0a, 0x0c,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x09,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0x48, 0x05, 0xb2, 0x01, 0x02, 0x38,
	0x01, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x42, 0x54, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x53, 0x68, 0x6f,
	0x70, 0x4f, 0x6e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x73, 0x6f, 0x6d, 0x2f, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x31, 0x3b, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_esom_basket_v1beta1_event_proto_rawDescOnce sync.Once
	file_esom_basket_v1beta1_event_proto_rawDescData = file_esom_basket_v1beta1_event_proto_rawDesc
)

func file_esom_basket_v1beta1_event_proto_rawDescGZIP() []byte {
	file_esom_basket_v1beta1_event_proto_rawDescOnce.Do(func() {
		file_esom_basket_v1beta1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_esom_basket_v1beta1_event_proto_rawDescData)
	})
	return file_esom_basket_v1beta1_event_proto_rawDescData
}

var file_esom_basket_v1beta1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_esom_basket_v1beta1_event_proto_goTypes = []any{
	(*BasketAbandoned)(nil),       // 0: esom.basket.v1beta1.BasketAbandoned
	(*BasketUpdated)(nil),         // 1: esom.basket.v1beta1.BasketUpdated
	(*BasketViewed)(nil),          // 2: esom.basket.v1beta1.BasketViewed
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_esom_basket_v1beta1_event_proto_depIdxs = []int32{
	3, // 0: esom.basket.v1beta1.BasketAbandoned.occurred_at:type_name -> google.protobuf.Timestamp
	3, // 1: esom.basket.v1beta1.BasketUpdated.occurred_at:type_name -> google.protobuf.Timestamp
	3, // 2: esom.basket.v1beta1.BasketViewed.occurred_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_esom_basket_v1beta1_event_proto_init() }
func file_esom_basket_v1beta1_event_proto_init() {
	if File_esom_basket_v1beta1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_esom_basket_v1beta1_event_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BasketAbandoned); i {
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
		file_esom_basket_v1beta1_event_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BasketUpdated); i {
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
		file_esom_basket_v1beta1_event_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BasketViewed); i {
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
			RawDescriptor: file_esom_basket_v1beta1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_esom_basket_v1beta1_event_proto_goTypes,
		DependencyIndexes: file_esom_basket_v1beta1_event_proto_depIdxs,
		MessageInfos:      file_esom_basket_v1beta1_event_proto_msgTypes,
	}.Build()
	File_esom_basket_v1beta1_event_proto = out.File
	file_esom_basket_v1beta1_event_proto_rawDesc = nil
	file_esom_basket_v1beta1_event_proto_goTypes = nil
	file_esom_basket_v1beta1_event_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: action.proto

package publish

import (
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

type DouyinPublishActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"` // 用户鉴权token
	Data  []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`   // 视频数据
	Title *string `protobuf:"bytes,3,req,name=title" json:"title,omitempty"` // 视频标题
}

func (x *DouyinPublishActionRequest) Reset() {
	*x = DouyinPublishActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishActionRequest) ProtoMessage() {}

func (x *DouyinPublishActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinPublishActionRequest) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinPublishActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinPublishActionRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DouyinPublishActionRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type DouyinPublishActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *DouyinPublishActionResponse) Reset() {
	*x = DouyinPublishActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishActionResponse) ProtoMessage() {}

func (x *DouyinPublishActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinPublishActionResponse) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinPublishActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinPublishActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

var File_action_proto protoreflect.FileDescriptor

var file_action_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x5f, 0x0a, 0x1d, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x60, 0x0a, 0x1e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x32, 0x7a,
	0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x68, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e,
}

var (
	file_action_proto_rawDescOnce sync.Once
	file_action_proto_rawDescData = file_action_proto_rawDesc
)

func file_action_proto_rawDescGZIP() []byte {
	file_action_proto_rawDescOnce.Do(func() {
		file_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_proto_rawDescData)
	})
	return file_action_proto_rawDescData
}

var file_action_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_action_proto_goTypes = []interface{}{
	(*DouyinPublishActionRequest)(nil),  // 0: douyin.core.douyin_publish_action_request
	(*DouyinPublishActionResponse)(nil), // 1: douyin.core.douyin_publish_action_response
}
var file_action_proto_depIdxs = []int32{
	0, // 0: douyin.core.PublishService.PublishAction:input_type -> douyin.core.douyin_publish_action_request
	1, // 1: douyin.core.PublishService.PublishAction:output_type -> douyin.core.douyin_publish_action_response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_action_proto_init() }
func file_action_proto_init() {
	if File_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishActionRequest); i {
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
		file_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishActionResponse); i {
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
			RawDescriptor: file_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_action_proto_goTypes,
		DependencyIndexes: file_action_proto_depIdxs,
		MessageInfos:      file_action_proto_msgTypes,
	}.Build()
	File_action_proto = out.File
	file_action_proto_rawDesc = nil
	file_action_proto_goTypes = nil
	file_action_proto_depIdxs = nil
}

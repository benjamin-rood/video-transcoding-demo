// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: video_streaming.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VideoChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data          []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	SegmentStart  bool   `protobuf:"varint,2,opt,name=segment_start,json=segmentStart,proto3" json:"segment_start,omitempty"`
	SegmentEnd    bool   `protobuf:"varint,3,opt,name=segment_end,json=segmentEnd,proto3" json:"segment_end,omitempty"`
	VideoId       string `protobuf:"bytes,4,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	SegmentNumber uint32 `protobuf:"varint,5,opt,name=segment_number,json=segmentNumber,proto3" json:"segment_number,omitempty"`
}

func (x *VideoChunk) Reset() {
	*x = VideoChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoChunk) ProtoMessage() {}

func (x *VideoChunk) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoChunk.ProtoReflect.Descriptor instead.
func (*VideoChunk) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *VideoChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *VideoChunk) GetSegmentStart() bool {
	if x != nil {
		return x.SegmentStart
	}
	return false
}

func (x *VideoChunk) GetSegmentEnd() bool {
	if x != nil {
		return x.SegmentEnd
	}
	return false
}

func (x *VideoChunk) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *VideoChunk) GetSegmentNumber() uint32 {
	if x != nil {
		return x.SegmentNumber
	}
	return 0
}

var File_video_streaming_proto protoreflect.FileDescriptor

var file_video_streaming_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x32, 0x60, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1b, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x28, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x65, 0x6e, 0x6a, 0x61, 0x6d, 0x69, 0x6e, 0x2d, 0x72, 0x6f, 0x6f, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_video_streaming_proto_rawDescOnce sync.Once
	file_video_streaming_proto_rawDescData = file_video_streaming_proto_rawDesc
)

func file_video_streaming_proto_rawDescGZIP() []byte {
	file_video_streaming_proto_rawDescOnce.Do(func() {
		file_video_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_streaming_proto_rawDescData)
	})
	return file_video_streaming_proto_rawDescData
}

var file_video_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_video_streaming_proto_goTypes = []interface{}{
	(*VideoChunk)(nil),  // 0: video_streaming.VideoChunk
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_video_streaming_proto_depIdxs = []int32{
	0, // 0: video_streaming.StreamingVideoIngestor.UploadVideo:input_type -> video_streaming.VideoChunk
	1, // 1: video_streaming.StreamingVideoIngestor.UploadVideo:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_streaming_proto_init() }
func file_video_streaming_proto_init() {
	if File_video_streaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_streaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoChunk); i {
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
			RawDescriptor: file_video_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_streaming_proto_goTypes,
		DependencyIndexes: file_video_streaming_proto_depIdxs,
		MessageInfos:      file_video_streaming_proto_msgTypes,
	}.Build()
	File_video_streaming_proto = out.File
	file_video_streaming_proto_rawDesc = nil
	file_video_streaming_proto_goTypes = nil
	file_video_streaming_proto_depIdxs = nil
}

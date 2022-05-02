// To compile, run the following command
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\graphgrpc\graph_grpc.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: graphgrpc/graph_grpc.proto

package graphgrpc

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

type Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: HOW TO DEFINE GRAPHDATA???
	VertexID int64 `protobuf:"varint,1,opt,name=vertexID,proto3" json:"vertexID,omitempty"`
}

func (x *Vertex) Reset() {
	*x = Vertex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphgrpc_graph_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertex) ProtoMessage() {}

func (x *Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_graphgrpc_graph_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertex.ProtoReflect.Descriptor instead.
func (*Vertex) Descriptor() ([]byte, []int) {
	return file_graphgrpc_graph_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Vertex) GetVertexID() int64 {
	if x != nil {
		return x.VertexID
	}
	return 0
}

type GraphID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: What should the ID be? A randomly generated number?
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GraphID) Reset() {
	*x = GraphID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphgrpc_graph_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphID) ProtoMessage() {}

func (x *GraphID) ProtoReflect() protoreflect.Message {
	mi := &file_graphgrpc_graph_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphID.ProtoReflect.Descriptor instead.
func (*GraphID) Descriptor() ([]byte, []int) {
	return file_graphgrpc_graph_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GraphID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type VertexDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Vertex1 *Vertex `protobuf:"bytes,2,opt,name=vertex1,proto3" json:"vertex1,omitempty"`
	Vertex2 *Vertex `protobuf:"bytes,3,opt,name=vertex2,proto3" json:"vertex2,omitempty"`
}

func (x *VertexDescription) Reset() {
	*x = VertexDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphgrpc_graph_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VertexDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VertexDescription) ProtoMessage() {}

func (x *VertexDescription) ProtoReflect() protoreflect.Message {
	mi := &file_graphgrpc_graph_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VertexDescription.ProtoReflect.Descriptor instead.
func (*VertexDescription) Descriptor() ([]byte, []int) {
	return file_graphgrpc_graph_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *VertexDescription) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VertexDescription) GetVertex1() *Vertex {
	if x != nil {
		return x.Vertex1
	}
	return nil
}

func (x *VertexDescription) GetVertex2() *Vertex {
	if x != nil {
		return x.Vertex2
	}
	return nil
}

type Distance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance int64 `protobuf:"varint,1,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *Distance) Reset() {
	*x = Distance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphgrpc_graph_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distance) ProtoMessage() {}

func (x *Distance) ProtoReflect() protoreflect.Message {
	mi := &file_graphgrpc_graph_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distance.ProtoReflect.Descriptor instead.
func (*Distance) Descriptor() ([]byte, []int) {
	return file_graphgrpc_graph_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Distance) GetDistance() int64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type DeleteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteResult) Reset() {
	*x = DeleteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphgrpc_graph_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResult) ProtoMessage() {}

func (x *DeleteResult) ProtoReflect() protoreflect.Message {
	mi := &file_graphgrpc_graph_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResult.ProtoReflect.Descriptor instead.
func (*DeleteResult) Descriptor() ([]byte, []int) {
	return file_graphgrpc_graph_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_graphgrpc_graph_grpc_proto protoreflect.FileDescriptor

var file_graphgrpc_graph_grpc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x22, 0x24, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x49, 0x44, 0x22, 0x19, 0x0a,
	0x07, 0x47, 0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7d, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x52, 0x07, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x31, 0x12, 0x2b, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x78, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x32, 0x22, 0x26, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x26, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xc7, 0x01, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x22, 0x00, 0x28, 0x01, 0x12, 0x43,
	0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graphgrpc_graph_grpc_proto_rawDescOnce sync.Once
	file_graphgrpc_graph_grpc_proto_rawDescData = file_graphgrpc_graph_grpc_proto_rawDesc
)

func file_graphgrpc_graph_grpc_proto_rawDescGZIP() []byte {
	file_graphgrpc_graph_grpc_proto_rawDescOnce.Do(func() {
		file_graphgrpc_graph_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_graphgrpc_graph_grpc_proto_rawDescData)
	})
	return file_graphgrpc_graph_grpc_proto_rawDescData
}

var file_graphgrpc_graph_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_graphgrpc_graph_grpc_proto_goTypes = []interface{}{
	(*Vertex)(nil),            // 0: graphgrpc.Vertex
	(*GraphID)(nil),           // 1: graphgrpc.GraphID
	(*VertexDescription)(nil), // 2: graphgrpc.VertexDescription
	(*Distance)(nil),          // 3: graphgrpc.Distance
	(*DeleteResult)(nil),      // 4: graphgrpc.DeleteResult
}
var file_graphgrpc_graph_grpc_proto_depIdxs = []int32{
	0, // 0: graphgrpc.VertexDescription.vertex1:type_name -> graphgrpc.Vertex
	0, // 1: graphgrpc.VertexDescription.vertex2:type_name -> graphgrpc.Vertex
	0, // 2: graphgrpc.Operations.PostGraph:input_type -> graphgrpc.Vertex
	2, // 3: graphgrpc.Operations.ShortestPath:input_type -> graphgrpc.VertexDescription
	1, // 4: graphgrpc.Operations.DeleteGraph:input_type -> graphgrpc.GraphID
	1, // 5: graphgrpc.Operations.PostGraph:output_type -> graphgrpc.GraphID
	3, // 6: graphgrpc.Operations.ShortestPath:output_type -> graphgrpc.Distance
	4, // 7: graphgrpc.Operations.DeleteGraph:output_type -> graphgrpc.DeleteResult
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_graphgrpc_graph_grpc_proto_init() }
func file_graphgrpc_graph_grpc_proto_init() {
	if File_graphgrpc_graph_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graphgrpc_graph_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertex); i {
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
		file_graphgrpc_graph_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphID); i {
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
		file_graphgrpc_graph_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VertexDescription); i {
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
		file_graphgrpc_graph_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distance); i {
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
		file_graphgrpc_graph_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResult); i {
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
			RawDescriptor: file_graphgrpc_graph_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_graphgrpc_graph_grpc_proto_goTypes,
		DependencyIndexes: file_graphgrpc_graph_grpc_proto_depIdxs,
		MessageInfos:      file_graphgrpc_graph_grpc_proto_msgTypes,
	}.Build()
	File_graphgrpc_graph_grpc_proto = out.File
	file_graphgrpc_graph_grpc_proto_rawDesc = nil
	file_graphgrpc_graph_grpc_proto_goTypes = nil
	file_graphgrpc_graph_grpc_proto_depIdxs = nil
}

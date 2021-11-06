// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: namenode.proto

package nmproto

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

type Playersmoves struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round       string `protobuf:"bytes,1,opt,name=round,proto3" json:"round,omitempty"`
	PlayerMoves string `protobuf:"bytes,2,opt,name=playerMoves,proto3" json:"playerMoves,omitempty"`
}

func (x *Playersmoves) Reset() {
	*x = Playersmoves{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playersmoves) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playersmoves) ProtoMessage() {}

func (x *Playersmoves) ProtoReflect() protoreflect.Message {
	mi := &file_namenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playersmoves.ProtoReflect.Descriptor instead.
func (*Playersmoves) Descriptor() ([]byte, []int) {
	return file_namenode_proto_rawDescGZIP(), []int{0}
}

func (x *Playersmoves) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *Playersmoves) GetPlayerMoves() string {
	if x != nil {
		return x.PlayerMoves
	}
	return ""
}

type Playermove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moves  string `protobuf:"bytes,1,opt,name=moves,proto3" json:"moves,omitempty"`
	Round  string `protobuf:"bytes,2,opt,name=round,proto3" json:"round,omitempty"`
	Player string `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *Playermove) Reset() {
	*x = Playermove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playermove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playermove) ProtoMessage() {}

func (x *Playermove) ProtoReflect() protoreflect.Message {
	mi := &file_namenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playermove.ProtoReflect.Descriptor instead.
func (*Playermove) Descriptor() ([]byte, []int) {
	return file_namenode_proto_rawDescGZIP(), []int{1}
}

func (x *Playermove) GetMoves() string {
	if x != nil {
		return x.Moves
	}
	return ""
}

func (x *Playermove) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *Playermove) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namenode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_namenode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_namenode_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_namenode_proto protoreflect.FileDescriptor

var file_namenode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76,
	0x65, 0x73, 0x22, 0x50, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x6d, 0x6f, 0x76, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x95, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x19, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x4d, 0x6f,
	0x76, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x4d, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x6d, 0x6f, 0x76, 0x65, 0x1a, 0x10, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2f, 0x6e, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_namenode_proto_rawDescOnce sync.Once
	file_namenode_proto_rawDescData = file_namenode_proto_rawDesc
)

func file_namenode_proto_rawDescGZIP() []byte {
	file_namenode_proto_rawDescOnce.Do(func() {
		file_namenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_namenode_proto_rawDescData)
	})
	return file_namenode_proto_rawDescData
}

var file_namenode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_namenode_proto_goTypes = []interface{}{
	(*Playersmoves)(nil), // 0: namenode.Playersmoves
	(*Playermove)(nil),   // 1: namenode.Playermove
	(*Status)(nil),       // 2: namenode.Status
}
var file_namenode_proto_depIdxs = []int32{
	0, // 0: namenode.StartServer.nameNodeStorePlayersMoves:input_type -> namenode.Playersmoves
	1, // 1: namenode.StartServer.dataNodeStoreMove:input_type -> namenode.Playermove
	2, // 2: namenode.StartServer.nameNodeStorePlayersMoves:output_type -> namenode.Status
	2, // 3: namenode.StartServer.dataNodeStoreMove:output_type -> namenode.Status
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_namenode_proto_init() }
func file_namenode_proto_init() {
	if File_namenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_namenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playersmoves); i {
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
		file_namenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playermove); i {
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
		file_namenode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_namenode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_namenode_proto_goTypes,
		DependencyIndexes: file_namenode_proto_depIdxs,
		MessageInfos:      file_namenode_proto_msgTypes,
	}.Build()
	File_namenode_proto = out.File
	file_namenode_proto_rawDesc = nil
	file_namenode_proto_goTypes = nil
	file_namenode_proto_depIdxs = nil
}

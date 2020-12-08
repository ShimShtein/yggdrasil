// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: protocol/yggdrasil.proto

package protocol

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

type WorkerStatus_Status int32

const (
	WorkerStatus_READY WorkerStatus_Status = 0
	WorkerStatus_BUSY  WorkerStatus_Status = 1
)

// Enum value maps for WorkerStatus_Status.
var (
	WorkerStatus_Status_name = map[int32]string{
		0: "READY",
		1: "BUSY",
	}
	WorkerStatus_Status_value = map[string]int32{
		"READY": 0,
		"BUSY":  1,
	}
)

func (x WorkerStatus_Status) Enum() *WorkerStatus_Status {
	p := new(WorkerStatus_Status)
	*p = x
	return p
}

func (x WorkerStatus_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkerStatus_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_yggdrasil_proto_enumTypes[0].Descriptor()
}

func (WorkerStatus_Status) Type() protoreflect.EnumType {
	return &file_protocol_yggdrasil_proto_enumTypes[0]
}

func (x WorkerStatus_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerStatus_Status.Descriptor instead.
func (WorkerStatus_Status) EnumDescriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{5, 0}
}

// An Empty message.
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{0}
}

// A RegisterRequest message contains information necessary for a client to
// register with the dispatcher for a specified work type.
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of work the worker is capable of handling.
	Handler string `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

// A RegisterResponse message contains the result of a registration request.
type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the dispatcher accepted the registration request or not.
	Registered bool `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	// The address on which the worker can be dialed to assign work.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// A message from the dispatch with more information about the request.
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *RegisterResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// A Work message contains information about a work assignment.
type Work struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The work assignment identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The work assignment payload.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Work) Reset() {
	*x = Work{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Work) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Work) ProtoMessage() {}

func (x *Work) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Work.ProtoReflect.Descriptor instead.
func (*Work) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{3}
}

func (x *Work) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Work) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// A StartResponse message contains the result of a work assignment request.
type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the worker accepted the work assignment or not.
	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{4}
}

func (x *StartResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

// A WorkerStatus message contains the status of a worker.
type WorkerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status WorkerStatus_Status `protobuf:"varint,1,opt,name=status,proto3,enum=yggdrasil.WorkerStatus_Status" json:"status,omitempty"`
}

func (x *WorkerStatus) Reset() {
	*x = WorkerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_yggdrasil_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerStatus) ProtoMessage() {}

func (x *WorkerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_yggdrasil_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerStatus.ProtoReflect.Descriptor instead.
func (*WorkerStatus) Descriptor() ([]byte, []int) {
	return file_protocol_yggdrasil_proto_rawDescGZIP(), []int{5}
}

func (x *WorkerStatus) GetStatus() WorkerStatus_Status {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_READY
}

var File_protocol_yggdrasil_proto protoreflect.FileDescriptor

var file_protocol_yggdrasil_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x79, 0x67, 0x67, 0x64, 0x72,
	0x61, 0x73, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x79, 0x67, 0x67, 0x64,
	0x72, 0x61, 0x73, 0x69, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x22, 0x2a, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x65, 0x0a, 0x0c, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x79, 0x67, 0x67,
	0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x1d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59, 0x10,
	0x01, 0x32, 0x82, 0x01, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x79,
	0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72,
	0x61, 0x73, 0x69, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x12, 0x0f, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x1a, 0x10, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0x75, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x79, 0x67, 0x67, 0x64,
	0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x1a, 0x18, 0x2e, 0x79, 0x67, 0x67,
	0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x10, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x17, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x68,
	0x61, 0x74, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x79, 0x67, 0x67, 0x64, 0x72,
	0x61, 0x73, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_yggdrasil_proto_rawDescOnce sync.Once
	file_protocol_yggdrasil_proto_rawDescData = file_protocol_yggdrasil_proto_rawDesc
)

func file_protocol_yggdrasil_proto_rawDescGZIP() []byte {
	file_protocol_yggdrasil_proto_rawDescOnce.Do(func() {
		file_protocol_yggdrasil_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_yggdrasil_proto_rawDescData)
	})
	return file_protocol_yggdrasil_proto_rawDescData
}

var file_protocol_yggdrasil_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_yggdrasil_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protocol_yggdrasil_proto_goTypes = []interface{}{
	(WorkerStatus_Status)(0), // 0: yggdrasil.WorkerStatus.Status
	(*Empty)(nil),            // 1: yggdrasil.Empty
	(*RegisterRequest)(nil),  // 2: yggdrasil.RegisterRequest
	(*RegisterResponse)(nil), // 3: yggdrasil.RegisterResponse
	(*Work)(nil),             // 4: yggdrasil.Work
	(*StartResponse)(nil),    // 5: yggdrasil.StartResponse
	(*WorkerStatus)(nil),     // 6: yggdrasil.WorkerStatus
}
var file_protocol_yggdrasil_proto_depIdxs = []int32{
	0, // 0: yggdrasil.WorkerStatus.status:type_name -> yggdrasil.WorkerStatus.Status
	2, // 1: yggdrasil.Dispatcher.Register:input_type -> yggdrasil.RegisterRequest
	4, // 2: yggdrasil.Dispatcher.Finish:input_type -> yggdrasil.Work
	4, // 3: yggdrasil.Worker.Start:input_type -> yggdrasil.Work
	1, // 4: yggdrasil.Worker.Status:input_type -> yggdrasil.Empty
	3, // 5: yggdrasil.Dispatcher.Register:output_type -> yggdrasil.RegisterResponse
	1, // 6: yggdrasil.Dispatcher.Finish:output_type -> yggdrasil.Empty
	5, // 7: yggdrasil.Worker.Start:output_type -> yggdrasil.StartResponse
	6, // 8: yggdrasil.Worker.Status:output_type -> yggdrasil.WorkerStatus
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_yggdrasil_proto_init() }
func file_protocol_yggdrasil_proto_init() {
	if File_protocol_yggdrasil_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_yggdrasil_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protocol_yggdrasil_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_protocol_yggdrasil_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_protocol_yggdrasil_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Work); i {
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
		file_protocol_yggdrasil_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_protocol_yggdrasil_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerStatus); i {
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
			RawDescriptor: file_protocol_yggdrasil_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protocol_yggdrasil_proto_goTypes,
		DependencyIndexes: file_protocol_yggdrasil_proto_depIdxs,
		EnumInfos:         file_protocol_yggdrasil_proto_enumTypes,
		MessageInfos:      file_protocol_yggdrasil_proto_msgTypes,
	}.Build()
	File_protocol_yggdrasil_proto = out.File
	file_protocol_yggdrasil_proto_rawDesc = nil
	file_protocol_yggdrasil_proto_goTypes = nil
	file_protocol_yggdrasil_proto_depIdxs = nil
}

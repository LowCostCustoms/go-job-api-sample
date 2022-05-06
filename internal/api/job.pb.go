// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: job.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type JobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Schedules []*JobScheduleReply `protobuf:"bytes,3,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *JobReply) Reset() {
	*x = JobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobReply) ProtoMessage() {}

func (x *JobReply) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobReply.ProtoReflect.Descriptor instead.
func (*JobReply) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobReply) GetSchedules() []*JobScheduleReply {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type JobScheduleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JobId string `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId,omitempty"`
	Cron  string `protobuf:"bytes,3,opt,name=cron,proto3" json:"cron,omitempty"`
}

func (x *JobScheduleReply) Reset() {
	*x = JobScheduleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobScheduleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobScheduleReply) ProtoMessage() {}

func (x *JobScheduleReply) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobScheduleReply.ProtoReflect.Descriptor instead.
func (*JobScheduleReply) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobScheduleReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobScheduleReply) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobScheduleReply) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

type JobPageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Items []*JobReply `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *JobPageReply) Reset() {
	*x = JobPageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPageReply) ProtoMessage() {}

func (x *JobPageReply) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobPageReply.ProtoReflect.Descriptor instead.
func (*JobPageReply) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobPageReply) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *JobPageReply) GetItems() []*JobReply {
	if x != nil {
		return x.Items
	}
	return nil
}

type JobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schedules []*JobScheduleRequest `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobRequest) GetSchedules() []*JobScheduleRequest {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type JobScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cron string `protobuf:"bytes,1,opt,name=cron,proto3" json:"cron,omitempty"`
}

func (x *JobScheduleRequest) Reset() {
	*x = JobScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobScheduleRequest) ProtoMessage() {}

func (x *JobScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobScheduleRequest.ProtoReflect.Descriptor instead.
func (*JobScheduleRequest) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{4}
}

func (x *JobScheduleRequest) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{5}
}

func (x *ListJobsRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListJobsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{6}
}

func (x *GetJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_job_proto protoreflect.FileDescriptor

var file_job_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63,
	0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33,
	0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f,
	0x6e, 0x22, 0x49, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x57, 0x0a, 0x0a,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x4a, 0x6f, 0x62, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x72, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x22,
	0x3f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xd9, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x45, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x40, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x77, 0x43,
	0x6f, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_proto_rawDescOnce sync.Once
	file_job_proto_rawDescData = file_job_proto_rawDesc
)

func file_job_proto_rawDescGZIP() []byte {
	file_job_proto_rawDescOnce.Do(func() {
		file_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_proto_rawDescData)
	})
	return file_job_proto_rawDescData
}

var file_job_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_job_proto_goTypes = []interface{}{
	(*JobReply)(nil),           // 0: api.JobReply
	(*JobScheduleReply)(nil),   // 1: api.JobScheduleReply
	(*JobPageReply)(nil),       // 2: api.JobPageReply
	(*JobRequest)(nil),         // 3: api.JobRequest
	(*JobScheduleRequest)(nil), // 4: api.JobScheduleRequest
	(*ListJobsRequest)(nil),    // 5: api.ListJobsRequest
	(*GetJobRequest)(nil),      // 6: api.GetJobRequest
}
var file_job_proto_depIdxs = []int32{
	1, // 0: api.JobReply.schedules:type_name -> api.JobScheduleReply
	0, // 1: api.JobPageReply.items:type_name -> api.JobReply
	4, // 2: api.JobRequest.schedules:type_name -> api.JobScheduleRequest
	5, // 3: api.JobService.ListJobs:input_type -> api.ListJobsRequest
	6, // 4: api.JobService.GetJob:input_type -> api.GetJobRequest
	3, // 5: api.JobService.CreateJob:input_type -> api.JobRequest
	2, // 6: api.JobService.ListJobs:output_type -> api.JobPageReply
	0, // 7: api.JobService.GetJob:output_type -> api.JobReply
	0, // 8: api.JobService.CreateJob:output_type -> api.JobReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_job_proto_init() }
func file_job_proto_init() {
	if File_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobReply); i {
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
		file_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobScheduleReply); i {
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
		file_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobPageReply); i {
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
		file_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobRequest); i {
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
		file_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobScheduleRequest); i {
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
		file_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
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
		file_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
			RawDescriptor: file_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_proto_goTypes,
		DependencyIndexes: file_job_proto_depIdxs,
		MessageInfos:      file_job_proto_msgTypes,
	}.Build()
	File_job_proto = out.File
	file_job_proto_rawDesc = nil
	file_job_proto_goTypes = nil
	file_job_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package heartbeat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HeartbeatMessage struct {
	// Cluster details
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Cluster metrics
	AMResourceLimitMB                              int32   `protobuf:"varint,2,opt,name=AMResourceLimitMB,proto3" json:"AMResourceLimitMB,omitempty"`
	AMResourceLimitVCores                          int32   `protobuf:"varint,3,opt,name=AMResourceLimitVCores,proto3" json:"AMResourceLimitVCores,omitempty"`
	UsedAMResourceMB                               int32   `protobuf:"varint,4,opt,name=UsedAMResourceMB,proto3" json:"UsedAMResourceMB,omitempty"`
	UsedAMResourceVCores                           int32   `protobuf:"varint,5,opt,name=UsedAMResourceVCores,proto3" json:"UsedAMResourceVCores,omitempty"`
	AppsSubmitted                                  int32   `protobuf:"varint,6,opt,name=AppsSubmitted,proto3" json:"AppsSubmitted,omitempty"`
	AppsRunning                                    int32   `protobuf:"varint,7,opt,name=AppsRunning,proto3" json:"AppsRunning,omitempty"`
	AppsPending                                    int32   `protobuf:"varint,8,opt,name=AppsPending,proto3" json:"AppsPending,omitempty"`
	AppsCompleted                                  int32   `protobuf:"varint,9,opt,name=AppsCompleted,proto3" json:"AppsCompleted,omitempty"`
	AppsKilled                                     int32   `protobuf:"varint,10,opt,name=AppsKilled,proto3" json:"AppsKilled,omitempty"`
	AppsFailed                                     int32   `protobuf:"varint,11,opt,name=AppsFailed,proto3" json:"AppsFailed,omitempty"`
	AggregateContainersPreempted                   int32   `protobuf:"varint,12,opt,name=AggregateContainersPreempted,proto3" json:"AggregateContainersPreempted,omitempty"`
	ActiveApplications                             int32   `protobuf:"varint,13,opt,name=ActiveApplications,proto3" json:"ActiveApplications,omitempty"`
	AppAttemptFirstContainerAllocationDelayNumOps  int32   `protobuf:"varint,14,opt,name=AppAttemptFirstContainerAllocationDelayNumOps,proto3" json:"AppAttemptFirstContainerAllocationDelayNumOps,omitempty"`
	AppAttemptFirstContainerAllocationDelayAvgTime float32 `protobuf:"fixed32,15,opt,name=AppAttemptFirstContainerAllocationDelayAvgTime,proto3" json:"AppAttemptFirstContainerAllocationDelayAvgTime,omitempty"`
	AllocatedMB                                    int32   `protobuf:"varint,16,opt,name=AllocatedMB,proto3" json:"AllocatedMB,omitempty"`
	AllocatedVCores                                int32   `protobuf:"varint,17,opt,name=AllocatedVCores,proto3" json:"AllocatedVCores,omitempty"`
	AllocatedContainers                            int32   `protobuf:"varint,18,opt,name=AllocatedContainers,proto3" json:"AllocatedContainers,omitempty"`
	AggregateContainersAllocated                   int32   `protobuf:"varint,19,opt,name=AggregateContainersAllocated,proto3" json:"AggregateContainersAllocated,omitempty"`
	AggregateContainersReleased                    int32   `protobuf:"varint,20,opt,name=AggregateContainersReleased,proto3" json:"AggregateContainersReleased,omitempty"`
	AvailableMB                                    int32   `protobuf:"varint,21,opt,name=AvailableMB,proto3" json:"AvailableMB,omitempty"`
	AvailableVCores                                int32   `protobuf:"varint,22,opt,name=AvailableVCores,proto3" json:"AvailableVCores,omitempty"`
	PendingMB                                      int32   `protobuf:"varint,23,opt,name=PendingMB,proto3" json:"PendingMB,omitempty"`
	PendingVCores                                  int32   `protobuf:"varint,24,opt,name=PendingVCores,proto3" json:"PendingVCores,omitempty"`
	PendingContainers                              int32   `protobuf:"varint,25,opt,name=PendingContainers,proto3" json:"PendingContainers,omitempty"`
	// Platform type
	ServiceType          string   `protobuf:"bytes,26,opt,name=ServiceType,proto3" json:"ServiceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatMessage) Reset()         { *m = HeartbeatMessage{} }
func (m *HeartbeatMessage) String() string { return proto.CompactTextString(m) }
func (*HeartbeatMessage) ProtoMessage()    {}
func (*HeartbeatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8aabfa9ad56473fd, []int{0}
}
func (m *HeartbeatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatMessage.Unmarshal(m, b)
}
func (m *HeartbeatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatMessage.Marshal(b, m, deterministic)
}
func (dst *HeartbeatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatMessage.Merge(dst, src)
}
func (m *HeartbeatMessage) XXX_Size() int {
	return xxx_messageInfo_HeartbeatMessage.Size(m)
}
func (m *HeartbeatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatMessage proto.InternalMessageInfo

func (m *HeartbeatMessage) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *HeartbeatMessage) GetAMResourceLimitMB() int32 {
	if m != nil {
		return m.AMResourceLimitMB
	}
	return 0
}

func (m *HeartbeatMessage) GetAMResourceLimitVCores() int32 {
	if m != nil {
		return m.AMResourceLimitVCores
	}
	return 0
}

func (m *HeartbeatMessage) GetUsedAMResourceMB() int32 {
	if m != nil {
		return m.UsedAMResourceMB
	}
	return 0
}

func (m *HeartbeatMessage) GetUsedAMResourceVCores() int32 {
	if m != nil {
		return m.UsedAMResourceVCores
	}
	return 0
}

func (m *HeartbeatMessage) GetAppsSubmitted() int32 {
	if m != nil {
		return m.AppsSubmitted
	}
	return 0
}

func (m *HeartbeatMessage) GetAppsRunning() int32 {
	if m != nil {
		return m.AppsRunning
	}
	return 0
}

func (m *HeartbeatMessage) GetAppsPending() int32 {
	if m != nil {
		return m.AppsPending
	}
	return 0
}

func (m *HeartbeatMessage) GetAppsCompleted() int32 {
	if m != nil {
		return m.AppsCompleted
	}
	return 0
}

func (m *HeartbeatMessage) GetAppsKilled() int32 {
	if m != nil {
		return m.AppsKilled
	}
	return 0
}

func (m *HeartbeatMessage) GetAppsFailed() int32 {
	if m != nil {
		return m.AppsFailed
	}
	return 0
}

func (m *HeartbeatMessage) GetAggregateContainersPreempted() int32 {
	if m != nil {
		return m.AggregateContainersPreempted
	}
	return 0
}

func (m *HeartbeatMessage) GetActiveApplications() int32 {
	if m != nil {
		return m.ActiveApplications
	}
	return 0
}

func (m *HeartbeatMessage) GetAppAttemptFirstContainerAllocationDelayNumOps() int32 {
	if m != nil {
		return m.AppAttemptFirstContainerAllocationDelayNumOps
	}
	return 0
}

func (m *HeartbeatMessage) GetAppAttemptFirstContainerAllocationDelayAvgTime() float32 {
	if m != nil {
		return m.AppAttemptFirstContainerAllocationDelayAvgTime
	}
	return 0
}

func (m *HeartbeatMessage) GetAllocatedMB() int32 {
	if m != nil {
		return m.AllocatedMB
	}
	return 0
}

func (m *HeartbeatMessage) GetAllocatedVCores() int32 {
	if m != nil {
		return m.AllocatedVCores
	}
	return 0
}

func (m *HeartbeatMessage) GetAllocatedContainers() int32 {
	if m != nil {
		return m.AllocatedContainers
	}
	return 0
}

func (m *HeartbeatMessage) GetAggregateContainersAllocated() int32 {
	if m != nil {
		return m.AggregateContainersAllocated
	}
	return 0
}

func (m *HeartbeatMessage) GetAggregateContainersReleased() int32 {
	if m != nil {
		return m.AggregateContainersReleased
	}
	return 0
}

func (m *HeartbeatMessage) GetAvailableMB() int32 {
	if m != nil {
		return m.AvailableMB
	}
	return 0
}

func (m *HeartbeatMessage) GetAvailableVCores() int32 {
	if m != nil {
		return m.AvailableVCores
	}
	return 0
}

func (m *HeartbeatMessage) GetPendingMB() int32 {
	if m != nil {
		return m.PendingMB
	}
	return 0
}

func (m *HeartbeatMessage) GetPendingVCores() int32 {
	if m != nil {
		return m.PendingVCores
	}
	return 0
}

func (m *HeartbeatMessage) GetPendingContainers() int32 {
	if m != nil {
		return m.PendingContainers
	}
	return 0
}

func (m *HeartbeatMessage) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartbeatMessage)(nil), "heartbeat.HeartbeatMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_8aabfa9ad56473fd) }

var fileDescriptor_message_8aabfa9ad56473fd = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xd5, 0xc1, 0x06, 0x3d, 0x5d, 0x59, 0xe7, 0x6d, 0x60, 0x60, 0x42, 0x05, 0x71, 0x51,
	0x21, 0xa8, 0x10, 0xf0, 0x00, 0xa4, 0x45, 0x13, 0x12, 0x64, 0x4c, 0xd9, 0xe0, 0x16, 0xb9, 0xc9,
	0x21, 0x58, 0x72, 0x12, 0xcb, 0x76, 0x2a, 0xed, 0xa9, 0x79, 0x05, 0x64, 0xc7, 0x49, 0xda, 0x35,
	0x9b, 0xb6, 0x4b, 0x7f, 0xff, 0x77, 0x7c, 0xa2, 0xe3, 0xd8, 0x30, 0xcc, 0x50, 0x6b, 0x96, 0xe2,
	0x54, 0xaa, 0xc2, 0x14, 0xa4, 0xff, 0x17, 0x99, 0x32, 0x0b, 0x64, 0xe6, 0xd5, 0xbf, 0x3e, 0x8c,
	0xbe, 0xd6, 0xab, 0xb0, 0xb2, 0xc8, 0x4b, 0xd8, 0x8d, 0x45, 0xa9, 0x0d, 0xaa, 0xdf, 0x39, 0xcb,
	0x90, 0xf6, 0xc6, 0xbd, 0x49, 0x3f, 0x1a, 0x78, 0x76, 0xca, 0x32, 0x24, 0x6f, 0x61, 0x3f, 0x08,
	0x23, 0xd4, 0x45, 0xa9, 0x62, 0xfc, 0xce, 0x33, 0x6e, 0xc2, 0x19, 0xdd, 0x1a, 0xf7, 0x26, 0xdb,
	0xd1, 0x66, 0x40, 0x3e, 0xc1, 0xd1, 0x15, 0xf8, 0x6b, 0x5e, 0x28, 0xd4, 0xf4, 0x9e, 0xab, 0xe8,
	0x0e, 0xc9, 0x1b, 0x18, 0xfd, 0xd4, 0x98, 0xb4, 0x61, 0x38, 0xa3, 0xf7, 0x5d, 0xc1, 0x06, 0x27,
	0x1f, 0xe0, 0x70, 0x9d, 0xf9, 0x06, 0xdb, 0xce, 0xef, 0xcc, 0xc8, 0x6b, 0x18, 0x06, 0x52, 0xea,
	0xf3, 0x72, 0x91, 0x71, 0x63, 0x30, 0xa1, 0x3b, 0x4e, 0x5e, 0x87, 0x64, 0x0c, 0x03, 0x0b, 0xa2,
	0x32, 0xcf, 0x79, 0x9e, 0xd2, 0x07, 0xce, 0x59, 0x45, 0xb5, 0x71, 0x86, 0x79, 0x62, 0x8d, 0x87,
	0xad, 0xe1, 0x51, 0xdd, 0x69, 0x5e, 0x64, 0x52, 0xa0, 0xed, 0xd4, 0x6f, 0x3b, 0x35, 0x90, 0xbc,
	0x00, 0xb0, 0xe0, 0x1b, 0x17, 0x02, 0x13, 0x0a, 0x4e, 0x59, 0x21, 0x75, 0x7e, 0xc2, 0xb8, 0xcd,
	0x07, 0x6d, 0x5e, 0x11, 0x32, 0x83, 0xe3, 0x20, 0x4d, 0x15, 0xa6, 0xcc, 0xe0, 0xbc, 0xc8, 0x0d,
	0xe3, 0x39, 0x2a, 0x7d, 0xa6, 0x10, 0x33, 0x69, 0x9b, 0xee, 0xba, 0x8a, 0x1b, 0x1d, 0x32, 0x05,
	0x12, 0xc4, 0x86, 0x2f, 0x31, 0x90, 0x52, 0xf0, 0x98, 0x19, 0x5e, 0xe4, 0x9a, 0x0e, 0x5d, 0x65,
	0x47, 0x42, 0x12, 0x78, 0x17, 0x48, 0x19, 0x18, 0x63, 0xeb, 0x4f, 0xb8, 0xd2, 0xa6, 0xd9, 0x35,
	0x10, 0xa2, 0xa8, 0xb4, 0x2f, 0x28, 0xd8, 0xe5, 0x69, 0x99, 0xfd, 0x90, 0x9a, 0x3e, 0x72, 0x5b,
	0xdd, 0xad, 0x88, 0xfc, 0x81, 0xe9, 0x2d, 0x0b, 0x82, 0x65, 0x7a, 0xc1, 0x33, 0xa4, 0x7b, 0xe3,
	0xde, 0x64, 0x2b, 0xba, 0x63, 0x95, 0x3b, 0xc9, 0x2a, 0xc1, 0x24, 0x9c, 0xd1, 0x91, 0x3f, 0xc9,
	0x16, 0x91, 0x09, 0xec, 0x35, 0x4b, 0xff, 0x8b, 0xed, 0x3b, 0xeb, 0x2a, 0x26, 0xef, 0xe1, 0xa0,
	0x41, 0xed, 0xa4, 0x29, 0x71, 0x76, 0x57, 0x74, 0xcd, 0xf9, 0x35, 0x26, 0x3d, 0xb8, 0xf6, 0xfc,
	0x1a, 0x87, 0x7c, 0x86, 0xe7, 0x1d, 0x79, 0x84, 0x02, 0x99, 0xc6, 0x84, 0x1e, 0xba, 0x2d, 0x6e,
	0x52, 0xdc, 0x0c, 0x96, 0x8c, 0x0b, 0xb6, 0x10, 0xf6, 0xc2, 0x1d, 0xf9, 0x19, 0xb4, 0xc8, 0xcd,
	0xa0, 0x5e, 0xfa, 0x19, 0x3c, 0xf6, 0x33, 0x58, 0xc7, 0xe4, 0x18, 0xfa, 0xfe, 0x0a, 0x84, 0x33,
	0xfa, 0xc4, 0x39, 0x2d, 0xb0, 0xb7, 0xc2, 0x2f, 0xfc, 0x2e, 0xb4, 0xba, 0x15, 0x6b, 0xd0, 0xbe,
	0x34, 0x1e, 0xac, 0x4c, 0xf1, 0x69, 0xf5, 0xd2, 0x6c, 0x04, 0xf6, 0xeb, 0xcf, 0x51, 0x2d, 0x79,
	0x8c, 0x17, 0x97, 0x12, 0xe9, 0xb3, 0xea, 0xe5, 0x5a, 0x41, 0x8b, 0x1d, 0xf7, 0x06, 0x7e, 0xfc,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xc6, 0xad, 0x84, 0x14, 0x05, 0x00, 0x00,
}
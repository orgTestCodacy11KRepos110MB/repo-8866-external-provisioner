// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/billing.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Billing related configuration of the service.
//
// The following example shows how to configure monitored resources and metrics
// for billing:
//
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     metrics:
//     - name: library.googleapis.com/book/borrowed_count
//       metric_kind: DELTA
//       value_type: INT64
//     billing:
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/borrowed_count
type Billing struct {
	// Billing configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations per service, each one must have
	// a different monitored resource type. A metric can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Billing_BillingDestination `protobuf:"bytes,8,rep,name=consumer_destinations,json=consumerDestinations,proto3" json:"consumer_destinations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Billing) Reset()         { *m = Billing{} }
func (m *Billing) String() string { return proto.CompactTextString(m) }
func (*Billing) ProtoMessage()    {}
func (*Billing) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_36fc4fe99841ce3d, []int{0}
}
func (m *Billing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Billing.Unmarshal(m, b)
}
func (m *Billing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Billing.Marshal(b, m, deterministic)
}
func (dst *Billing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Billing.Merge(dst, src)
}
func (m *Billing) XXX_Size() int {
	return xxx_messageInfo_Billing.Size(m)
}
func (m *Billing) XXX_DiscardUnknown() {
	xxx_messageInfo_Billing.DiscardUnknown(m)
}

var xxx_messageInfo_Billing proto.InternalMessageInfo

func (m *Billing) GetConsumerDestinations() []*Billing_BillingDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific billing destination (Currently only support
// bill against consumer project).
type Billing_BillingDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource,proto3" json:"monitored_resource,omitempty"`
	// Names of the metrics to report to this billing destination.
	// Each name must be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics              []string `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Billing_BillingDestination) Reset()         { *m = Billing_BillingDestination{} }
func (m *Billing_BillingDestination) String() string { return proto.CompactTextString(m) }
func (*Billing_BillingDestination) ProtoMessage()    {}
func (*Billing_BillingDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_36fc4fe99841ce3d, []int{0, 0}
}
func (m *Billing_BillingDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Billing_BillingDestination.Unmarshal(m, b)
}
func (m *Billing_BillingDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Billing_BillingDestination.Marshal(b, m, deterministic)
}
func (dst *Billing_BillingDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Billing_BillingDestination.Merge(dst, src)
}
func (m *Billing_BillingDestination) XXX_Size() int {
	return xxx_messageInfo_Billing_BillingDestination.Size(m)
}
func (m *Billing_BillingDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_Billing_BillingDestination.DiscardUnknown(m)
}

var xxx_messageInfo_Billing_BillingDestination proto.InternalMessageInfo

func (m *Billing_BillingDestination) GetMonitoredResource() string {
	if m != nil {
		return m.MonitoredResource
	}
	return ""
}

func (m *Billing_BillingDestination) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Billing)(nil), "google.api.Billing")
	proto.RegisterType((*Billing_BillingDestination)(nil), "google.api.Billing.BillingDestination")
}

func init() { proto.RegisterFile("google/api/billing.proto", fileDescriptor_billing_36fc4fe99841ce3d) }

var fileDescriptor_billing_36fc4fe99841ce3d = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x10, 0xc7, 0xe9, 0xee, 0xc7, 0xb7, 0x6e, 0x14, 0xc1, 0xa0, 0x50, 0x16, 0x0f, 0xc5, 0x83, 0xf4,
	0x62, 0x0a, 0x7a, 0xf4, 0x64, 0x51, 0xc4, 0x5b, 0xe9, 0x51, 0x91, 0x25, 0x9b, 0x1d, 0xc3, 0x40,
	0x3b, 0x53, 0x92, 0xac, 0x0f, 0xe4, 0xbb, 0xf8, 0x5e, 0x62, 0xd3, 0xba, 0x15, 0x4f, 0x61, 0xf2,
	0xfb, 0xfd, 0x27, 0x99, 0x11, 0xa9, 0x65, 0xb6, 0x0d, 0x14, 0xba, 0xc3, 0x62, 0x83, 0x4d, 0x83,
	0x64, 0x55, 0xe7, 0x38, 0xb0, 0x14, 0x91, 0x28, 0xdd, 0xe1, 0xea, 0x7c, 0x62, 0x69, 0x22, 0x0e,
	0x3a, 0x20, 0x93, 0x8f, 0xe6, 0xc5, 0x67, 0x22, 0x16, 0x65, 0xcc, 0xca, 0x17, 0x71, 0x66, 0x98,
	0xfc, 0xae, 0x05, 0xb7, 0xde, 0x82, 0x0f, 0x48, 0x51, 0x4d, 0x0f, 0xb2, 0x79, 0x7e, 0x78, 0x7d,
	0xa9, 0xf6, 0x5d, 0xd5, 0x90, 0x19, 0xcf, 0xfb, 0xbd, 0x5e, 0x9f, 0x8e, 0x4d, 0x26, 0x97, 0x7e,
	0xf5, 0x2a, 0xe4, 0x5f, 0x57, 0x5e, 0x09, 0xd9, 0x32, 0x61, 0x60, 0x07, 0xdb, 0xb5, 0x03, 0xcf,
	0x3b, 0x67, 0x20, 0x4d, 0xb2, 0x24, 0x5f, 0xd6, 0x27, 0x3f, 0xa4, 0x1e, 0x80, 0x4c, 0xc5, 0xa2,
	0x85, 0xe0, 0xd0, 0xf8, 0x74, 0x96, 0xcd, 0xf3, 0x65, 0x3d, 0x96, 0x25, 0x89, 0x63, 0xc3, 0xed,
	0xe4, 0x87, 0xe5, 0xd1, 0xf0, 0x5c, 0xf5, 0x3d, 0x67, 0x95, 0x3c, 0x3f, 0x0c, 0xcc, 0x72, 0xa3,
	0xc9, 0x2a, 0x76, 0xb6, 0xb0, 0x40, 0xfd, 0x16, 0x8a, 0x88, 0x74, 0x87, 0xbe, 0x5f, 0x93, 0x07,
	0xf7, 0x8e, 0x06, 0x0c, 0xd3, 0x1b, 0xda, 0xdb, 0x5f, 0xd5, 0xc7, 0xec, 0xdf, 0xe3, 0x5d, 0xf5,
	0xb4, 0xf9, 0xdf, 0x07, 0x6f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0x2a, 0x74, 0xa4, 0x84,
	0x01, 0x00, 0x00,
}

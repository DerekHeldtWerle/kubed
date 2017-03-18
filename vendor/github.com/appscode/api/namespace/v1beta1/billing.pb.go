// Code generated by protoc-gen-go.
// source: billing.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	billing.proto
	team.proto

It has these top-level messages:
	GetSubscriptionResponse
	GetQuotaResponse
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	IsAvailableRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSubscriptionResponse struct {
	Status     *appscode_dtypes.Status          `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Product    *GetSubscriptionResponse_Product `protobuf:"bytes,2,opt,name=product" json:"product,omitempty"`
	AutoExtend bool                             `protobuf:"varint,3,opt,name=auto_extend,json=autoExtend" json:"auto_extend,omitempty"`
	Quota      string                           `protobuf:"bytes,4,opt,name=quota" json:"quota,omitempty"`
	DateStart  int64                            `protobuf:"varint,5,opt,name=date_start,json=dateStart" json:"date_start,omitempty"`
	DateEnd    int64                            `protobuf:"varint,6,opt,name=date_end,json=dateEnd" json:"date_end,omitempty"`
}

func (m *GetSubscriptionResponse) Reset()                    { *m = GetSubscriptionResponse{} }
func (m *GetSubscriptionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSubscriptionResponse) ProtoMessage()               {}
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSubscriptionResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetSubscriptionResponse) GetProduct() *GetSubscriptionResponse_Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *GetSubscriptionResponse) GetAutoExtend() bool {
	if m != nil {
		return m.AutoExtend
	}
	return false
}

func (m *GetSubscriptionResponse) GetQuota() string {
	if m != nil {
		return m.Quota
	}
	return ""
}

func (m *GetSubscriptionResponse) GetDateStart() int64 {
	if m != nil {
		return m.DateStart
	}
	return 0
}

func (m *GetSubscriptionResponse) GetDateEnd() int64 {
	if m != nil {
		return m.DateEnd
	}
	return 0
}

type GetSubscriptionResponse_Product struct {
	Sku          string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Type         string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	DisplayName  string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	PricingModel string `protobuf:"bytes,4,opt,name=pricing_model,json=pricingModel" json:"pricing_model,omitempty"`
}

func (m *GetSubscriptionResponse_Product) Reset()         { *m = GetSubscriptionResponse_Product{} }
func (m *GetSubscriptionResponse_Product) String() string { return proto.CompactTextString(m) }
func (*GetSubscriptionResponse_Product) ProtoMessage()    {}
func (*GetSubscriptionResponse_Product) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *GetSubscriptionResponse_Product) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *GetSubscriptionResponse_Product) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetSubscriptionResponse_Product) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *GetSubscriptionResponse_Product) GetPricingModel() string {
	if m != nil {
		return m.PricingModel
	}
	return ""
}

type GetQuotaResponse struct {
	Status   *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Assigned map[string]int64        `protobuf:"bytes,2,rep,name=assigned" json:"assigned,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Used     map[string]int64        `protobuf:"bytes,3,rep,name=used" json:"used,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *GetQuotaResponse) Reset()                    { *m = GetQuotaResponse{} }
func (m *GetQuotaResponse) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaResponse) ProtoMessage()               {}
func (*GetQuotaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetQuotaResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetQuotaResponse) GetAssigned() map[string]int64 {
	if m != nil {
		return m.Assigned
	}
	return nil
}

func (m *GetQuotaResponse) GetUsed() map[string]int64 {
	if m != nil {
		return m.Used
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSubscriptionResponse)(nil), "appscode.namespace.v1beta1.GetSubscriptionResponse")
	proto.RegisterType((*GetSubscriptionResponse_Product)(nil), "appscode.namespace.v1beta1.GetSubscriptionResponse.Product")
	proto.RegisterType((*GetQuotaResponse)(nil), "appscode.namespace.v1beta1.GetQuotaResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Billing service

type BillingClient interface {
	// Checks current subscription of a namespace
	GetSubscription(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	GetQuota(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*GetQuotaResponse, error)
	CheckPaymentMethod(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type billingClient struct {
	cc *grpc.ClientConn
}

func NewBillingClient(cc *grpc.ClientConn) BillingClient {
	return &billingClient{cc}
}

func (c *billingClient) GetSubscription(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	out := new(GetSubscriptionResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Billing/GetSubscription", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) GetQuota(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*GetQuotaResponse, error) {
	out := new(GetQuotaResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Billing/GetQuota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) CheckPaymentMethod(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Billing/CheckPaymentMethod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Billing service

type BillingServer interface {
	// Checks current subscription of a namespace
	GetSubscription(context.Context, *appscode_dtypes.VoidRequest) (*GetSubscriptionResponse, error)
	GetQuota(context.Context, *appscode_dtypes.VoidRequest) (*GetQuotaResponse, error)
	CheckPaymentMethod(context.Context, *appscode_dtypes.VoidRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterBillingServer(s *grpc.Server, srv BillingServer) {
	s.RegisterService(&_Billing_serviceDesc, srv)
}

func _Billing_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Billing/GetSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetSubscription(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_GetQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).GetQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Billing/GetQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).GetQuota(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_CheckPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).CheckPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Billing/CheckPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).CheckPaymentMethod(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Billing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.namespace.v1beta1.Billing",
	HandlerType: (*BillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscription",
			Handler:    _Billing_GetSubscription_Handler,
		},
		{
			MethodName: "GetQuota",
			Handler:    _Billing_GetQuota_Handler,
		},
		{
			MethodName: "CheckPaymentMethod",
			Handler:    _Billing_CheckPaymentMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billing.proto",
}

func init() { proto.RegisterFile("billing.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0x86, 0x35, 0x99, 0xb6, 0x49, 0x4e, 0x5a, 0x7d, 0x95, 0xf5, 0x49, 0x1d, 0x46, 0x2d, 0x0d,
	0x61, 0x93, 0x05, 0x9a, 0x51, 0x5b, 0x54, 0x50, 0x7f, 0x90, 0x28, 0xaa, 0x2a, 0x2a, 0x15, 0x85,
	0xa9, 0xda, 0x05, 0x9b, 0xc8, 0x99, 0x39, 0x4a, 0x87, 0x24, 0xb6, 0x1b, 0x7b, 0x0a, 0xb3, 0xed,
	0x96, 0x25, 0x6c, 0xb9, 0x06, 0x56, 0x5c, 0x09, 0x6b, 0x76, 0x5c, 0x01, 0x57, 0x80, 0xec, 0x71,
	0x42, 0x43, 0x15, 0xf5, 0x67, 0x33, 0x1a, 0x9f, 0xe3, 0xf3, 0xfa, 0xf1, 0x7b, 0x6c, 0xc3, 0x42,
	0x27, 0xed, 0xf7, 0x53, 0xd6, 0x0d, 0xc4, 0x90, 0x2b, 0x4e, 0x7c, 0x2a, 0x84, 0x8c, 0x79, 0x82,
	0x01, 0xa3, 0x03, 0x94, 0x82, 0xc6, 0x18, 0x5c, 0xac, 0x75, 0x50, 0xd1, 0x35, 0x7f, 0xb9, 0xcb,
	0x79, 0xb7, 0x8f, 0x21, 0x15, 0x69, 0x48, 0x19, 0xe3, 0x8a, 0xaa, 0x94, 0x33, 0x59, 0x54, 0xfa,
	0x0f, 0x47, 0x95, 0x53, 0xf2, 0xab, 0x13, 0xf9, 0x44, 0xe5, 0x02, 0x65, 0x68, 0xbe, 0xc5, 0x84,
	0xc6, 0x27, 0x17, 0x96, 0x0e, 0x50, 0x1d, 0x67, 0x1d, 0x19, 0x0f, 0x53, 0xa1, 0x6b, 0x23, 0x94,
	0x82, 0x33, 0x89, 0x24, 0x84, 0x39, 0xa9, 0xa8, 0xca, 0xa4, 0xe7, 0xd4, 0x9d, 0x66, 0x6d, 0x7d,
	0x29, 0x18, 0x73, 0x16, 0x4a, 0xc1, 0xb1, 0x49, 0x47, 0x76, 0x1a, 0x39, 0x81, 0xb2, 0x18, 0xf2,
	0x24, 0x8b, 0x95, 0x57, 0x32, 0x15, 0xdb, 0xc1, 0xf4, 0x9d, 0x05, 0x53, 0x96, 0x0d, 0x5a, 0x85,
	0x44, 0x34, 0xd2, 0x22, 0xab, 0x50, 0xa3, 0x99, 0xe2, 0x6d, 0xfc, 0xa8, 0x90, 0x25, 0x9e, 0x5b,
	0x77, 0x9a, 0x95, 0x08, 0x74, 0x68, 0xdf, 0x44, 0xc8, 0xff, 0x30, 0x7b, 0x9e, 0x71, 0x45, 0xbd,
	0x99, 0xba, 0xd3, 0xac, 0x46, 0xc5, 0x80, 0xac, 0x00, 0x24, 0x54, 0x61, 0x5b, 0x2a, 0x3a, 0x54,
	0xde, 0x6c, 0xdd, 0x69, 0xba, 0x51, 0x55, 0x47, 0x8e, 0x75, 0x80, 0x3c, 0x80, 0x8a, 0x49, 0x6b,
	0xc9, 0x39, 0x93, 0x2c, 0xeb, 0xf1, 0x3e, 0x4b, 0xfc, 0x0f, 0x50, 0xb6, 0x10, 0x64, 0x11, 0x5c,
	0xd9, 0xcb, 0x8c, 0x01, 0xd5, 0x48, 0xff, 0x12, 0x02, 0x33, 0x7a, 0xf3, 0x66, 0x87, 0xd5, 0xc8,
	0xfc, 0x93, 0x47, 0x30, 0x9f, 0xa4, 0x52, 0xf4, 0x69, 0xde, 0xd6, 0xfb, 0x34, 0x88, 0xd5, 0xa8,
	0x66, 0x63, 0x6f, 0xe8, 0x00, 0xc9, 0x63, 0x58, 0x10, 0xc3, 0x34, 0x4e, 0x59, 0xb7, 0x3d, 0xe0,
	0x09, 0xf6, 0x2d, 0xeb, 0xbc, 0x0d, 0x1e, 0xe9, 0x58, 0xe3, 0x67, 0x09, 0x16, 0x0f, 0x50, 0xbd,
	0xd5, 0xfc, 0xf7, 0x6f, 0xc3, 0x29, 0x54, 0xa8, 0x94, 0x69, 0x97, 0x61, 0xe2, 0x95, 0xea, 0x6e,
	0xb3, 0xb6, 0xbe, 0x75, 0x43, 0x1f, 0x26, 0x16, 0x0c, 0x5e, 0xda, 0xe2, 0x7d, 0xa6, 0x86, 0x79,
	0x34, 0xd6, 0x22, 0x87, 0x30, 0x93, 0x49, 0xd4, 0x0d, 0xd0, 0x9a, 0x9b, 0x77, 0xd2, 0x3c, 0x91,
	0x23, 0x3d, 0xa3, 0xe1, 0x6f, 0xc3, 0xc2, 0xc4, 0x32, 0xda, 0xe8, 0x1e, 0xe6, 0x23, 0xa3, 0x7b,
	0x98, 0xeb, 0xae, 0x5e, 0xd0, 0x7e, 0x56, 0x38, 0xed, 0x46, 0xc5, 0x60, 0xab, 0xf4, 0xdc, 0xf1,
	0x9f, 0x41, 0x75, 0xac, 0x77, 0x97, 0xc2, 0xf5, 0xdf, 0x2e, 0x94, 0xf7, 0x8a, 0xab, 0x47, 0xbe,
	0x39, 0xf0, 0xdf, 0x3f, 0x47, 0x90, 0x2c, 0x5f, 0xb3, 0xf6, 0x94, 0xa7, 0x49, 0x84, 0xe7, 0x19,
	0x4a, 0xe5, 0x6f, 0xdc, 0xe3, 0x34, 0x37, 0x5e, 0x5f, 0x7e, 0xf7, 0x4a, 0x15, 0xe7, 0xf2, 0xc7,
	0xaf, 0xcf, 0xa5, 0x5d, 0xb2, 0x1d, 0xb6, 0x27, 0x6e, 0xe4, 0x58, 0x27, 0xb4, 0x3a, 0xa1, 0x7d,
	0x19, 0x42, 0x79, 0x45, 0x4c, 0x86, 0xef, 0x25, 0x67, 0xe4, 0x8b, 0x03, 0x95, 0x91, 0xb1, 0x37,
	0xa0, 0x3e, 0xb9, 0x4b, 0x73, 0x1a, 0x3b, 0x86, 0x6e, 0x93, 0x3c, 0xbd, 0x35, 0x9d, 0xb9, 0x61,
	0x16, 0xeb, 0xab, 0x03, 0xe4, 0xd5, 0x19, 0xc6, 0xbd, 0x16, 0xcd, 0x07, 0xc8, 0xd4, 0x11, 0xaa,
	0x33, 0x9e, 0xdc, 0x00, 0xb8, 0x32, 0x25, 0x6b, 0x89, 0x0e, 0xaf, 0xb8, 0xf6, 0x82, 0xec, 0xdc,
	0x9a, 0x4b, 0x14, 0x08, 0x03, 0x83, 0x50, 0xf0, 0xed, 0xed, 0x42, 0x23, 0xe6, 0x83, 0xbf, 0xeb,
	0x51, 0x91, 0x5e, 0x37, 0x65, 0x6f, 0xde, 0x9e, 0x8b, 0x96, 0x7e, 0x16, 0x5b, 0xce, 0xbb, 0xb2,
	0x4d, 0x74, 0xe6, 0xcc, 0x43, 0xb9, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xb4, 0xb9, 0xa9,
	0xb4, 0x05, 0x00, 0x00,
}
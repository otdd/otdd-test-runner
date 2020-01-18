// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testrunner.proto

package testrunner

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FetchTestCaseReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Mac                  string   `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchTestCaseReq) Reset()         { *m = FetchTestCaseReq{} }
func (m *FetchTestCaseReq) String() string { return proto.CompactTextString(m) }
func (*FetchTestCaseReq) ProtoMessage()    {}
func (*FetchTestCaseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{0}
}

func (m *FetchTestCaseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchTestCaseReq.Unmarshal(m, b)
}
func (m *FetchTestCaseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchTestCaseReq.Marshal(b, m, deterministic)
}
func (m *FetchTestCaseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchTestCaseReq.Merge(m, src)
}
func (m *FetchTestCaseReq) XXX_Size() int {
	return xxx_messageInfo_FetchTestCaseReq.Size(m)
}
func (m *FetchTestCaseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchTestCaseReq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchTestCaseReq proto.InternalMessageInfo

func (m *FetchTestCaseReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FetchTestCaseReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *FetchTestCaseReq) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type TestCase struct {
	TestId                 string   `protobuf:"bytes,1,opt,name=testId,proto3" json:"testId,omitempty"`
	RunId                  string   `protobuf:"bytes,2,opt,name=runId,proto3" json:"runId,omitempty"`
	Port                   int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	PassthroughConnections []string `protobuf:"bytes,4,rep,name=passthroughConnections,proto3" json:"passthroughConnections,omitempty"`
	InboundRequest         []byte   `protobuf:"bytes,5,opt,name=inboundRequest,proto3" json:"inboundRequest,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TestCase) Reset()         { *m = TestCase{} }
func (m *TestCase) String() string { return proto.CompactTextString(m) }
func (*TestCase) ProtoMessage()    {}
func (*TestCase) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{1}
}

func (m *TestCase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestCase.Unmarshal(m, b)
}
func (m *TestCase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestCase.Marshal(b, m, deterministic)
}
func (m *TestCase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestCase.Merge(m, src)
}
func (m *TestCase) XXX_Size() int {
	return xxx_messageInfo_TestCase.Size(m)
}
func (m *TestCase) XXX_DiscardUnknown() {
	xxx_messageInfo_TestCase.DiscardUnknown(m)
}

var xxx_messageInfo_TestCase proto.InternalMessageInfo

func (m *TestCase) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *TestCase) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *TestCase) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TestCase) GetPassthroughConnections() []string {
	if m != nil {
		return m.PassthroughConnections
	}
	return nil
}

func (m *TestCase) GetInboundRequest() []byte {
	if m != nil {
		return m.InboundRequest
	}
	return nil
}

type FetchOutboundRespReq struct {
	TestId               string   `protobuf:"bytes,1,opt,name=testId,proto3" json:"testId,omitempty"`
	RunId                string   `protobuf:"bytes,2,opt,name=runId,proto3" json:"runId,omitempty"`
	OutboundReq          []byte   `protobuf:"bytes,3,opt,name=outboundReq,proto3" json:"outboundReq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchOutboundRespReq) Reset()         { *m = FetchOutboundRespReq{} }
func (m *FetchOutboundRespReq) String() string { return proto.CompactTextString(m) }
func (*FetchOutboundRespReq) ProtoMessage()    {}
func (*FetchOutboundRespReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{2}
}

func (m *FetchOutboundRespReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchOutboundRespReq.Unmarshal(m, b)
}
func (m *FetchOutboundRespReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchOutboundRespReq.Marshal(b, m, deterministic)
}
func (m *FetchOutboundRespReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchOutboundRespReq.Merge(m, src)
}
func (m *FetchOutboundRespReq) XXX_Size() int {
	return xxx_messageInfo_FetchOutboundRespReq.Size(m)
}
func (m *FetchOutboundRespReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchOutboundRespReq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchOutboundRespReq proto.InternalMessageInfo

func (m *FetchOutboundRespReq) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *FetchOutboundRespReq) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *FetchOutboundRespReq) GetOutboundReq() []byte {
	if m != nil {
		return m.OutboundReq
	}
	return nil
}

type FetchOutboundRespResp struct {
	OutboundResp         []byte   `protobuf:"bytes,1,opt,name=outboundResp,proto3" json:"outboundResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchOutboundRespResp) Reset()         { *m = FetchOutboundRespResp{} }
func (m *FetchOutboundRespResp) String() string { return proto.CompactTextString(m) }
func (*FetchOutboundRespResp) ProtoMessage()    {}
func (*FetchOutboundRespResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{3}
}

func (m *FetchOutboundRespResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchOutboundRespResp.Unmarshal(m, b)
}
func (m *FetchOutboundRespResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchOutboundRespResp.Marshal(b, m, deterministic)
}
func (m *FetchOutboundRespResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchOutboundRespResp.Merge(m, src)
}
func (m *FetchOutboundRespResp) XXX_Size() int {
	return xxx_messageInfo_FetchOutboundRespResp.Size(m)
}
func (m *FetchOutboundRespResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchOutboundRespResp.DiscardUnknown(m)
}

var xxx_messageInfo_FetchOutboundRespResp proto.InternalMessageInfo

func (m *FetchOutboundRespResp) GetOutboundResp() []byte {
	if m != nil {
		return m.OutboundResp
	}
	return nil
}

type TestResult struct {
	TestId               string   `protobuf:"bytes,1,opt,name=testId,proto3" json:"testId,omitempty"`
	RunId                string   `protobuf:"bytes,2,opt,name=runId,proto3" json:"runId,omitempty"`
	InboundRequestErr    string   `protobuf:"bytes,3,opt,name=inboundRequestErr,proto3" json:"inboundRequestErr,omitempty"`
	InboundResponse      []byte   `protobuf:"bytes,4,opt,name=inboundResponse,proto3" json:"inboundResponse,omitempty"`
	InboundResponseErr   string   `protobuf:"bytes,5,opt,name=inboundResponseErr,proto3" json:"inboundResponseErr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResult) Reset()         { *m = TestResult{} }
func (m *TestResult) String() string { return proto.CompactTextString(m) }
func (*TestResult) ProtoMessage()    {}
func (*TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{4}
}

func (m *TestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResult.Unmarshal(m, b)
}
func (m *TestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResult.Marshal(b, m, deterministic)
}
func (m *TestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResult.Merge(m, src)
}
func (m *TestResult) XXX_Size() int {
	return xxx_messageInfo_TestResult.Size(m)
}
func (m *TestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResult.DiscardUnknown(m)
}

var xxx_messageInfo_TestResult proto.InternalMessageInfo

func (m *TestResult) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *TestResult) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *TestResult) GetInboundRequestErr() string {
	if m != nil {
		return m.InboundRequestErr
	}
	return ""
}

func (m *TestResult) GetInboundResponse() []byte {
	if m != nil {
		return m.InboundResponse
	}
	return nil
}

func (m *TestResult) GetInboundResponseErr() string {
	if m != nil {
		return m.InboundResponseErr
	}
	return ""
}

type ReportTestResultResp struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportTestResultResp) Reset()         { *m = ReportTestResultResp{} }
func (m *ReportTestResultResp) String() string { return proto.CompactTextString(m) }
func (*ReportTestResultResp) ProtoMessage()    {}
func (*ReportTestResultResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{5}
}

func (m *ReportTestResultResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportTestResultResp.Unmarshal(m, b)
}
func (m *ReportTestResultResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportTestResultResp.Marshal(b, m, deterministic)
}
func (m *ReportTestResultResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportTestResultResp.Merge(m, src)
}
func (m *ReportTestResultResp) XXX_Size() int {
	return xxx_messageInfo_ReportTestResultResp.Size(m)
}
func (m *ReportTestResultResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportTestResultResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReportTestResultResp proto.InternalMessageInfo

func (m *ReportTestResultResp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type LogReq struct {
	RunId                string   `protobuf:"bytes,1,opt,name=runId,proto3" json:"runId,omitempty"`
	TestId               string   `protobuf:"bytes,2,opt,name=testId,proto3" json:"testId,omitempty"`
	Log                  string   `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogReq) Reset()         { *m = LogReq{} }
func (m *LogReq) String() string { return proto.CompactTextString(m) }
func (*LogReq) ProtoMessage()    {}
func (*LogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_496d3b0f237f00f4, []int{6}
}

func (m *LogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogReq.Unmarshal(m, b)
}
func (m *LogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogReq.Marshal(b, m, deterministic)
}
func (m *LogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogReq.Merge(m, src)
}
func (m *LogReq) XXX_Size() int {
	return xxx_messageInfo_LogReq.Size(m)
}
func (m *LogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogReq proto.InternalMessageInfo

func (m *LogReq) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *LogReq) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *LogReq) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *LogReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchTestCaseReq)(nil), "testrunner.FetchTestCaseReq")
	proto.RegisterType((*TestCase)(nil), "testrunner.TestCase")
	proto.RegisterType((*FetchOutboundRespReq)(nil), "testrunner.FetchOutboundRespReq")
	proto.RegisterType((*FetchOutboundRespResp)(nil), "testrunner.FetchOutboundRespResp")
	proto.RegisterType((*TestResult)(nil), "testrunner.TestResult")
	proto.RegisterType((*ReportTestResultResp)(nil), "testrunner.ReportTestResultResp")
	proto.RegisterType((*LogReq)(nil), "testrunner.LogReq")
}

func init() { proto.RegisterFile("testrunner.proto", fileDescriptor_496d3b0f237f00f4) }

var fileDescriptor_496d3b0f237f00f4 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x25, 0x04, 0x10, 0x4c, 0x69, 0x17, 0x46, 0x14, 0x45, 0xe9, 0x1e, 0xa8, 0x0f, 0x2d, 0x87,
	0x2a, 0x2b, 0x75, 0xa5, 0x5e, 0x7a, 0xdb, 0x15, 0x95, 0x56, 0x5a, 0xb5, 0x2b, 0xb7, 0xa7, 0xde,
	0x02, 0x0c, 0x01, 0x89, 0xc4, 0xc1, 0x76, 0x56, 0xea, 0x3f, 0xf5, 0x0f, 0x7a, 0xeb, 0x97, 0x55,
	0x76, 0x02, 0x24, 0x81, 0x6a, 0xb5, 0x97, 0x28, 0xf3, 0xc6, 0x7e, 0x9e, 0xf7, 0xc6, 0x63, 0x18,
	0x68, 0x52, 0x5a, 0x66, 0x49, 0x42, 0x32, 0x48, 0xa5, 0xd0, 0x02, 0xe1, 0x88, 0xf8, 0x6f, 0x22,
	0x21, 0xa2, 0x2d, 0x5d, 0xd9, 0xcc, 0x3c, 0x5b, 0x5d, 0x51, 0x9c, 0xea, 0x5f, 0xf9, 0x42, 0xc6,
	0x61, 0xf0, 0x85, 0xf4, 0x62, 0xfd, 0x83, 0x94, 0xbe, 0x0d, 0x15, 0x71, 0xda, 0xa1, 0x0f, 0xdd,
	0x4c, 0x91, 0x4c, 0xc2, 0x98, 0x3c, 0x67, 0xe2, 0x4c, 0x7b, 0xfc, 0x10, 0xe3, 0x00, 0x5c, 0x1d,
	0x46, 0x5e, 0xd3, 0xc2, 0xe6, 0xd7, 0x20, 0x71, 0xb8, 0xf0, 0xdc, 0x1c, 0x89, 0xc3, 0x05, 0xfb,
	0xed, 0x40, 0x77, 0xcf, 0x87, 0x63, 0xe8, 0x98, 0x5a, 0xee, 0x96, 0x05, 0x55, 0x11, 0xe1, 0x08,
	0xda, 0x32, 0x4b, 0xee, 0x96, 0x05, 0x55, 0x1e, 0x20, 0x42, 0x2b, 0x15, 0x52, 0x5b, 0xb6, 0x36,
	0xb7, 0xff, 0xf8, 0x09, 0xc6, 0x69, 0xa8, 0x94, 0x5e, 0x4b, 0x91, 0x45, 0xeb, 0x5b, 0x91, 0x24,
	0xb4, 0xd0, 0x1b, 0x91, 0x28, 0xaf, 0x35, 0x71, 0xa7, 0x3d, 0xfe, 0x9f, 0x2c, 0xbe, 0x83, 0x57,
	0x9b, 0x64, 0x2e, 0xb2, 0x64, 0xc9, 0x69, 0x97, 0x91, 0xd2, 0x5e, 0x7b, 0xe2, 0x4c, 0xfb, 0xbc,
	0x86, 0xb2, 0x15, 0x8c, 0xac, 0x05, 0xdf, 0x32, 0x5d, 0xe0, 0x2a, 0x35, 0x36, 0x3c, 0xaf, 0xf2,
	0x09, 0xbc, 0x10, 0x07, 0x82, 0x9d, 0x15, 0xd0, 0xe7, 0x65, 0x88, 0x7d, 0x86, 0xd7, 0x67, 0xce,
	0x51, 0x29, 0x32, 0xe8, 0x8b, 0x12, 0x66, 0x8f, 0xeb, 0xf3, 0x0a, 0xc6, 0xfe, 0x3a, 0x00, 0xc6,
	0x53, 0x4e, 0x2a, 0xdb, 0xea, 0x67, 0xd6, 0xf6, 0x01, 0x86, 0x55, 0xcd, 0x33, 0x29, 0x8b, 0x86,
	0x9d, 0x26, 0x70, 0x0a, 0x17, 0x07, 0x50, 0xa5, 0x22, 0x51, 0xe4, 0xb5, 0x6c, 0x45, 0x75, 0x18,
	0x03, 0xc0, 0x1a, 0x64, 0x88, 0xdb, 0x96, 0xf8, 0x4c, 0x86, 0x05, 0x30, 0xe2, 0x64, 0x7a, 0x7a,
	0x54, 0x62, 0x0d, 0x18, 0x43, 0x47, 0xda, 0xc8, 0xaa, 0xe9, 0xf2, 0x22, 0x62, 0x2b, 0xe8, 0xdc,
	0x8b, 0xc8, 0xf4, 0xe2, 0xa0, 0xcb, 0x29, 0xeb, 0x3a, 0xba, 0xd0, 0xac, 0xb8, 0x30, 0x00, 0x77,
	0x2b, 0xa2, 0xfd, 0x95, 0xdc, 0x8a, 0x08, 0x2f, 0xa1, 0xa7, 0x37, 0x31, 0x29, 0x1d, 0xc6, 0xa9,
	0x55, 0xe3, 0xf2, 0x23, 0xf0, 0xf1, 0x4f, 0x13, 0x86, 0xb6, 0x24, 0x3b, 0x30, 0xdf, 0x49, 0x3e,
	0x6e, 0x16, 0x84, 0x33, 0x78, 0x59, 0x19, 0x0d, 0xbc, 0x0c, 0x4a, 0x73, 0x56, 0x9f, 0x1a, 0x7f,
	0x54, 0xce, 0xee, 0x13, 0xac, 0x81, 0x3f, 0x61, 0x78, 0xd2, 0x76, 0x9c, 0x9c, 0x50, 0xd5, 0x6e,
	0x9f, 0xff, 0xf6, 0x89, 0x15, 0x2a, 0x65, 0x0d, 0xfc, 0x0a, 0x83, 0xba, 0xa1, 0x38, 0xae, 0xd7,
	0x91, 0xe3, 0x7e, 0xe5, 0xc8, 0x73, 0x6d, 0x60, 0x0d, 0xbc, 0x06, 0xf7, 0x5e, 0x44, 0x88, 0xe5,
	0xa5, 0x79, 0x07, 0xfc, 0x71, 0x90, 0x3f, 0x23, 0xc1, 0xfe, 0x19, 0x09, 0x66, 0xe6, 0x19, 0x61,
	0x8d, 0x9b, 0xf7, 0x80, 0x1b, 0x11, 0x08, 0xbd, 0x5c, 0xda, 0x8f, 0x22, 0xf9, 0x48, 0xf2, 0xe6,
	0xe2, 0x68, 0xe8, 0x83, 0xd9, 0xf1, 0xe0, 0xcc, 0x3b, 0x76, 0xeb, 0xf5, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x52, 0xf9, 0xe2, 0xe6, 0xaf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestRunnerServiceClient is the client API for TestRunnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestRunnerServiceClient interface {
	FetchTestCase(ctx context.Context, in *FetchTestCaseReq, opts ...grpc.CallOption) (*TestCase, error)
	FetchOutboundResp(ctx context.Context, in *FetchOutboundRespReq, opts ...grpc.CallOption) (*FetchOutboundRespResp, error)
	ReportTestResult(ctx context.Context, in *TestResult, opts ...grpc.CallOption) (*ReportTestResultResp, error)
	Log(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type testRunnerServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestRunnerServiceClient(cc *grpc.ClientConn) TestRunnerServiceClient {
	return &testRunnerServiceClient{cc}
}

func (c *testRunnerServiceClient) FetchTestCase(ctx context.Context, in *FetchTestCaseReq, opts ...grpc.CallOption) (*TestCase, error) {
	out := new(TestCase)
	err := c.cc.Invoke(ctx, "/testrunner.TestRunnerService/FetchTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunnerServiceClient) FetchOutboundResp(ctx context.Context, in *FetchOutboundRespReq, opts ...grpc.CallOption) (*FetchOutboundRespResp, error) {
	out := new(FetchOutboundRespResp)
	err := c.cc.Invoke(ctx, "/testrunner.TestRunnerService/FetchOutboundResp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunnerServiceClient) ReportTestResult(ctx context.Context, in *TestResult, opts ...grpc.CallOption) (*ReportTestResultResp, error) {
	out := new(ReportTestResultResp)
	err := c.cc.Invoke(ctx, "/testrunner.TestRunnerService/ReportTestResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunnerServiceClient) Log(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/testrunner.TestRunnerService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestRunnerServiceServer is the server API for TestRunnerService service.
type TestRunnerServiceServer interface {
	FetchTestCase(context.Context, *FetchTestCaseReq) (*TestCase, error)
	FetchOutboundResp(context.Context, *FetchOutboundRespReq) (*FetchOutboundRespResp, error)
	ReportTestResult(context.Context, *TestResult) (*ReportTestResultResp, error)
	Log(context.Context, *LogReq) (*empty.Empty, error)
}

// UnimplementedTestRunnerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestRunnerServiceServer struct {
}

func (*UnimplementedTestRunnerServiceServer) FetchTestCase(ctx context.Context, req *FetchTestCaseReq) (*TestCase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTestCase not implemented")
}
func (*UnimplementedTestRunnerServiceServer) FetchOutboundResp(ctx context.Context, req *FetchOutboundRespReq) (*FetchOutboundRespResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOutboundResp not implemented")
}
func (*UnimplementedTestRunnerServiceServer) ReportTestResult(ctx context.Context, req *TestResult) (*ReportTestResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTestResult not implemented")
}
func (*UnimplementedTestRunnerServiceServer) Log(ctx context.Context, req *LogReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func RegisterTestRunnerServiceServer(s *grpc.Server, srv TestRunnerServiceServer) {
	s.RegisterService(&_TestRunnerService_serviceDesc, srv)
}

func _TestRunnerService_FetchTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchTestCaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRunnerServiceServer).FetchTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testrunner.TestRunnerService/FetchTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRunnerServiceServer).FetchTestCase(ctx, req.(*FetchTestCaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestRunnerService_FetchOutboundResp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchOutboundRespReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRunnerServiceServer).FetchOutboundResp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testrunner.TestRunnerService/FetchOutboundResp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRunnerServiceServer).FetchOutboundResp(ctx, req.(*FetchOutboundRespReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestRunnerService_ReportTestResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRunnerServiceServer).ReportTestResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testrunner.TestRunnerService/ReportTestResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRunnerServiceServer).ReportTestResult(ctx, req.(*TestResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestRunnerService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRunnerServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testrunner.TestRunnerService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRunnerServiceServer).Log(ctx, req.(*LogReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestRunnerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testrunner.TestRunnerService",
	HandlerType: (*TestRunnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTestCase",
			Handler:    _TestRunnerService_FetchTestCase_Handler,
		},
		{
			MethodName: "FetchOutboundResp",
			Handler:    _TestRunnerService_FetchOutboundResp_Handler,
		},
		{
			MethodName: "ReportTestResult",
			Handler:    _TestRunnerService_ReportTestResult_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _TestRunnerService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testrunner.proto",
}

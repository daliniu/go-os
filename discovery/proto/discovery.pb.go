// Code generated by protoc-gen-go.
// source: github.com/micro/go-platform/discovery/proto/discovery.proto
// DO NOT EDIT!

/*
Package go_micro_platform_discovery is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-platform/discovery/proto/discovery.proto

It has these top-level messages:
	Heartbeat
	Service
	Result
	Node
	Endpoint
	Value
*/
package go_micro_platform_discovery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Heartbeat struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Service   *Service `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Timestamp int64    `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Interval  int64    `protobuf:"varint,4,opt,name=interval" json:"interval,omitempty"`
	Ttl       int64    `protobuf:"varint,5,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Heartbeat) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Service struct {
	Name      string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version   string            `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,3,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Endpoints []*Endpoint       `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints,omitempty"`
	Nodes     []*Node           `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Result struct {
	Action    string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Service   *Service `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Timestamp int64    `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Result) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Node struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address  string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Port     int64             `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Endpoint struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Request  *Value            `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	Response *Value            `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Endpoint) GetRequest() *Value {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Endpoint) GetResponse() *Value {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Value struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type   string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Values []*Value `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Value) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Heartbeat)(nil), "go.micro.platform.discovery.Heartbeat")
	proto.RegisterType((*Service)(nil), "go.micro.platform.discovery.Service")
	proto.RegisterType((*Result)(nil), "go.micro.platform.discovery.Result")
	proto.RegisterType((*Node)(nil), "go.micro.platform.discovery.Node")
	proto.RegisterType((*Endpoint)(nil), "go.micro.platform.discovery.Endpoint")
	proto.RegisterType((*Value)(nil), "go.micro.platform.discovery.Value")
}

var fileDescriptor0 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x4b, 0x8f, 0x95, 0x40,
	0x10, 0x85, 0xc3, 0x7d, 0xf0, 0xa8, 0x71, 0x7c, 0xf4, 0x8a, 0xe8, 0x46, 0x89, 0x26, 0x6e, 0x6c,
	0xcc, 0xbd, 0x9a, 0x4c, 0x8c, 0x1b, 0x13, 0x6f, 0xe2, 0x46, 0x17, 0x9a, 0xe8, 0xba, 0x81, 0xf6,
	0xda, 0x11, 0x68, 0xec, 0x2e, 0x26, 0x61, 0xe7, 0xde, 0xdf, 0xe1, 0xaf, 0x74, 0x63, 0x01, 0x3d,
	0xa3, 0xc4, 0x04, 0x31, 0xce, 0xb2, 0x9a, 0x3a, 0xc5, 0xf9, 0x4e, 0x35, 0xc0, 0xf3, 0xa3, 0xc2,
	0x4f, 0x6d, 0xc6, 0x73, 0x5d, 0xa5, 0x95, 0xca, 0x8d, 0x4e, 0x8f, 0xfa, 0x51, 0x53, 0x0a, 0xfc,
	0xa8, 0x4d, 0x95, 0x16, 0xca, 0xe6, 0xfa, 0x5c, 0x9a, 0x2e, 0x6d, 0x8c, 0x46, 0xfd, 0xab, 0xe6,
	0x43, 0xcd, 0xee, 0x1c, 0x35, 0x1f, 0x54, 0xfc, 0x42, 0xc2, 0x2f, 0x5b, 0x92, 0xaf, 0x1e, 0x44,
	0xaf, 0xa4, 0x30, 0x98, 0x49, 0x81, 0x0c, 0x60, 0xa5, 0x8a, 0xd8, 0xbb, 0xeb, 0x3d, 0x8c, 0xd8,
	0x53, 0x08, 0xac, 0x34, 0xe7, 0x2a, 0x97, 0xf1, 0x8a, 0x0e, 0x4e, 0x76, 0xf7, 0xf9, 0xcc, 0x20,
	0xfe, 0x6e, 0xec, 0x65, 0xb7, 0x20, 0x42, 0x55, 0x49, 0x8b, 0xa2, 0x6a, 0xe2, 0x35, 0x09, 0xd7,
	0xec, 0x26, 0x84, 0xaa, 0x46, 0x7a, 0x2e, 0xca, 0x78, 0x33, 0x9c, 0x9c, 0xc0, 0x1a, 0xb1, 0x8c,
	0xb7, 0x7d, 0x91, 0x7c, 0x5b, 0x41, 0x70, 0xa1, 0xbe, 0x06, 0x9b, 0x5a, 0x54, 0xd2, 0x59, 0xb8,
	0x01, 0x01, 0xcd, 0xb6, 0x4a, 0xd7, 0x83, 0x85, 0x88, 0xbd, 0x84, 0xb0, 0x92, 0x28, 0x0a, 0x81,
	0x82, 0x66, 0xaf, 0xc9, 0xd4, 0x6e, 0x89, 0x29, 0xfe, 0xda, 0x89, 0x0e, 0x35, 0x9a, 0x8e, 0x9d,
	0x41, 0x24, 0xeb, 0xa2, 0xd1, 0x64, 0xca, 0x92, 0xa1, 0x7e, 0xcc, 0x83, 0xd9, 0x31, 0x07, 0xd7,
	0xcd, 0x1e, 0xc3, 0xb6, 0xd6, 0x85, 0xb4, 0xe4, 0xbc, 0x57, 0xdd, 0x9b, 0x55, 0xbd, 0xa1, 0xce,
	0xdb, 0x29, 0x9c, 0x4e, 0x5f, 0x4e, 0xe8, 0x9f, 0x65, 0xe7, 0x00, 0x4f, 0x61, 0x4b, 0xa1, 0xb4,
	0x63, 0xc2, 0xd1, 0xb3, 0xd5, 0x99, 0x97, 0x64, 0xe0, 0xbf, 0x95, 0xb6, 0x2d, 0x91, 0x5d, 0x07,
	0x5f, 0xe4, 0xd8, 0xc3, 0x5f, 0xf1, 0x42, 0x92, 0xef, 0x1e, 0x6c, 0x7a, 0x77, 0x93, 0x7d, 0x53,
	0xd8, 0xa2, 0x28, 0x8c, 0xb4, 0xd6, 0x85, 0x4d, 0xbb, 0x68, 0xb4, 0x41, 0xb7, 0xc4, 0x17, 0xbf,
	0x45, 0x3f, 0x66, 0x96, 0xfe, 0x95, 0x7e, 0x9a, 0xfb, 0xbf, 0x67, 0xf1, 0xc3, 0x83, 0xf0, 0x32,
	0xfb, 0xe9, 0xd5, 0xd8, 0x43, 0x60, 0xe4, 0x97, 0x96, 0xb0, 0x5c, 0x18, 0xc9, 0xac, 0x9b, 0xf7,
	0xfd, 0x64, 0xf6, 0x04, 0x42, 0xe2, 0x6b, 0x74, 0x6d, 0xe5, 0x40, 0xb5, 0x4c, 0x75, 0xf8, 0x83,
	0x7c, 0xbf, 0xe8, 0xb6, 0xfc, 0x2f, 0xfd, 0x07, 0xd8, 0x8e, 0x06, 0xa6, 0xe4, 0x54, 0x61, 0xd7,
	0xb8, 0x46, 0xb6, 0x03, 0x7f, 0xd0, 0x59, 0xf7, 0x3d, 0x2c, 0x00, 0xca, 0xfc, 0xe1, 0xbf, 0xb0,
	0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x82, 0xa2, 0xa8, 0x57, 0x04, 0x00, 0x00,
}

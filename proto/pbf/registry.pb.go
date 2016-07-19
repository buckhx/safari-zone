// Code generated by protoc-gen-gogo.
// source: registry.proto
// DO NOT EDIT!

package pbf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Trainer_Gender int32

const (
	BOY  Trainer_Gender = 0
	GIRL Trainer_Gender = 1
)

var Trainer_Gender_name = map[int32]string{
	0: "BOY",
	1: "GIRL",
}
var Trainer_Gender_value = map[string]int32{
	"BOY":  0,
	"GIRL": 1,
}

func (Trainer_Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0, 0} }

type Trainer struct {
	Uid      string              `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name     string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string              `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Age      int32               `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Gender   Trainer_Gender      `protobuf:"varint,5,opt,name=gender,proto3,enum=buckhx.safari.registry.Trainer_Gender" json:"gender,omitempty"`
	Start    *Timestamp          `protobuf:"bytes,6,opt,name=start" json:"start,omitempty"`
	Pc       *Pokemon_Collection `protobuf:"bytes,7,opt,name=pc" json:"pc,omitempty"`
	Scope    []string            `protobuf:"bytes,8,rep,name=scope" json:"scope,omitempty"`
}

func (m *Trainer) Reset()                    { *m = Trainer{} }
func (*Trainer) ProtoMessage()               {}
func (*Trainer) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{0} }

func (m *Trainer) GetStart() *Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Trainer) GetPc() *Pokemon_Collection {
	if m != nil {
		return m.Pc
	}
	return nil
}

type Token struct {
	Access string   `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Key    string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Secret string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Scope  []string `protobuf:"bytes,4,rep,name=scope" json:"scope,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{1} }

type Cert struct {
	Jwk []byte `protobuf:"bytes,1,opt,name=jwk,proto3" json:"jwk,omitempty"`
}

func (m *Cert) Reset()                    { *m = Cert{} }
func (*Cert) ProtoMessage()               {}
func (*Cert) Descriptor() ([]byte, []int) { return fileDescriptorRegistry, []int{2} }

func init() {
	proto.RegisterType((*Trainer)(nil), "buckhx.safari.registry.Trainer")
	proto.RegisterType((*Token)(nil), "buckhx.safari.registry.Token")
	proto.RegisterType((*Cert)(nil), "buckhx.safari.registry.Cert")
	proto.RegisterEnum("buckhx.safari.registry.Trainer_Gender", Trainer_Gender_name, Trainer_Gender_value)
}
func (x Trainer_Gender) String() string {
	s, ok := Trainer_Gender_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Trainer) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Trainer)
	if !ok {
		that2, ok := that.(Trainer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Uid != that1.Uid {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Password != that1.Password {
		return false
	}
	if this.Age != that1.Age {
		return false
	}
	if this.Gender != that1.Gender {
		return false
	}
	if !this.Start.Equal(that1.Start) {
		return false
	}
	if !this.Pc.Equal(that1.Pc) {
		return false
	}
	if len(this.Scope) != len(that1.Scope) {
		return false
	}
	for i := range this.Scope {
		if this.Scope[i] != that1.Scope[i] {
			return false
		}
	}
	return true
}
func (this *Token) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Token)
	if !ok {
		that2, ok := that.(Token)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Access != that1.Access {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Secret != that1.Secret {
		return false
	}
	if len(this.Scope) != len(that1.Scope) {
		return false
	}
	for i := range this.Scope {
		if this.Scope[i] != that1.Scope[i] {
			return false
		}
	}
	return true
}
func (this *Cert) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Cert)
	if !ok {
		that2, ok := that.(Cert)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Jwk, that1.Jwk) {
		return false
	}
	return true
}
func (this *Trainer) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&pbf.Trainer{")
	s = append(s, "Uid: "+fmt.Sprintf("%#v", this.Uid)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Password: "+fmt.Sprintf("%#v", this.Password)+",\n")
	s = append(s, "Age: "+fmt.Sprintf("%#v", this.Age)+",\n")
	s = append(s, "Gender: "+fmt.Sprintf("%#v", this.Gender)+",\n")
	if this.Start != nil {
		s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	}
	if this.Pc != nil {
		s = append(s, "Pc: "+fmt.Sprintf("%#v", this.Pc)+",\n")
	}
	s = append(s, "Scope: "+fmt.Sprintf("%#v", this.Scope)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Token) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pbf.Token{")
	s = append(s, "Access: "+fmt.Sprintf("%#v", this.Access)+",\n")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Secret: "+fmt.Sprintf("%#v", this.Secret)+",\n")
	s = append(s, "Scope: "+fmt.Sprintf("%#v", this.Scope)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Cert) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pbf.Cert{")
	s = append(s, "Jwk: "+fmt.Sprintf("%#v", this.Jwk)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRegistry(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRegistry(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Registry service

type RegistryClient interface {
	// Register makes a creates a new trainer in the safari
	//
	// Trainer name, password, age & gender are required.
	// Any other supplied fields will be ignored
	Register(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Response, error)
	// Get fetchs a trainer
	//
	// The populated fields will depend on the auth scope of the token
	GetTrainer(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Trainer, error)
	// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
	//
	// HTTPS required w/ HTTP basic access authentication via a header
	// Authorization: Basic BASE64({user:pass})
	Enter(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Token, error)
	// Access retrieves an access token from a token's key/secret combo
	//
	// HTTPS required. Only key/secret will be honored
	// Authorization: Basic BASE64({key:secret})
	Access(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	// Certificate returns the cert used to verify token signatures
	//
	// The cert is in JWK form as described in https://tools.ietf.org/html/rfc7517
	Certificate(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Cert, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) Register(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetTrainer(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Trainer, error) {
	out := new(Trainer)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/GetTrainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Enter(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Enter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Access(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Access", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) Certificate(ctx context.Context, in *Trainer, opts ...grpc.CallOption) (*Cert, error) {
	out := new(Cert)
	err := grpc.Invoke(ctx, "/buckhx.safari.registry.Registry/Certificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registry service

type RegistryServer interface {
	// Register makes a creates a new trainer in the safari
	//
	// Trainer name, password, age & gender are required.
	// Any other supplied fields will be ignored
	Register(context.Context, *Trainer) (*Response, error)
	// Get fetchs a trainer
	//
	// The populated fields will depend on the auth scope of the token
	GetTrainer(context.Context, *Trainer) (*Trainer, error)
	// Enter authenticates a user to retrieve a an access token to authorize requests for a safari
	//
	// HTTPS required w/ HTTP basic access authentication via a header
	// Authorization: Basic BASE64({user:pass})
	Enter(context.Context, *Trainer) (*Token, error)
	// Access retrieves an access token from a token's key/secret combo
	//
	// HTTPS required. Only key/secret will be honored
	// Authorization: Basic BASE64({key:secret})
	Access(context.Context, *Token) (*Token, error)
	// Certificate returns the cert used to verify token signatures
	//
	// The cert is in JWK form as described in https://tools.ietf.org/html/rfc7517
	Certificate(context.Context, *Trainer) (*Cert, error)
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Register(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetTrainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetTrainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/GetTrainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetTrainer(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Enter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Enter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Enter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Enter(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Access_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Access(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Access",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Access(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buckhx.safari.registry.Registry/Certificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).Certificate(ctx, req.(*Trainer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buckhx.safari.registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registry_Register_Handler,
		},
		{
			MethodName: "GetTrainer",
			Handler:    _Registry_GetTrainer_Handler,
		},
		{
			MethodName: "Enter",
			Handler:    _Registry_Enter_Handler,
		},
		{
			MethodName: "Access",
			Handler:    _Registry_Access_Handler,
		},
		{
			MethodName: "Certificate",
			Handler:    _Registry_Certificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorRegistry,
}

func (m *Trainer) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Trainer) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uid) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Uid)))
		i += copy(data[i:], m.Uid)
	}
	if len(m.Name) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.Password) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Password)))
		i += copy(data[i:], m.Password)
	}
	if m.Age != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Age))
	}
	if m.Gender != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Gender))
	}
	if m.Start != nil {
		data[i] = 0x32
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Start.Size()))
		n1, err := m.Start.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Pc != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintRegistry(data, i, uint64(m.Pc.Size()))
		n2, err := m.Pc.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Scope) > 0 {
		for _, s := range m.Scope {
			data[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *Token) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Token) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Access) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Access)))
		i += copy(data[i:], m.Access)
	}
	if len(m.Key) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if len(m.Secret) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Secret)))
		i += copy(data[i:], m.Secret)
	}
	if len(m.Scope) > 0 {
		for _, s := range m.Scope {
			data[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *Cert) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Cert) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Jwk) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRegistry(data, i, uint64(len(m.Jwk)))
		i += copy(data[i:], m.Jwk)
	}
	return i, nil
}

func encodeFixed64Registry(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Registry(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRegistry(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Trainer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovRegistry(uint64(m.Age))
	}
	if m.Gender != 0 {
		n += 1 + sovRegistry(uint64(m.Gender))
	}
	if m.Start != nil {
		l = m.Start.Size()
		n += 1 + l + sovRegistry(uint64(l))
	}
	if m.Pc != nil {
		l = m.Pc.Size()
		n += 1 + l + sovRegistry(uint64(l))
	}
	if len(m.Scope) > 0 {
		for _, s := range m.Scope {
			l = len(s)
			n += 1 + l + sovRegistry(uint64(l))
		}
	}
	return n
}

func (m *Token) Size() (n int) {
	var l int
	_ = l
	l = len(m.Access)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	if len(m.Scope) > 0 {
		for _, s := range m.Scope {
			l = len(s)
			n += 1 + l + sovRegistry(uint64(l))
		}
	}
	return n
}

func (m *Cert) Size() (n int) {
	var l int
	_ = l
	l = len(m.Jwk)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	return n
}

func sovRegistry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRegistry(x uint64) (n int) {
	return sovRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Trainer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Trainer{`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Password:` + fmt.Sprintf("%v", this.Password) + `,`,
		`Age:` + fmt.Sprintf("%v", this.Age) + `,`,
		`Gender:` + fmt.Sprintf("%v", this.Gender) + `,`,
		`Start:` + strings.Replace(fmt.Sprintf("%v", this.Start), "Timestamp", "Timestamp", 1) + `,`,
		`Pc:` + strings.Replace(fmt.Sprintf("%v", this.Pc), "Pokemon_Collection", "Pokemon_Collection", 1) + `,`,
		`Scope:` + fmt.Sprintf("%v", this.Scope) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Token) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Token{`,
		`Access:` + fmt.Sprintf("%v", this.Access) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Secret:` + fmt.Sprintf("%v", this.Secret) + `,`,
		`Scope:` + fmt.Sprintf("%v", this.Scope) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Cert) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Cert{`,
		`Jwk:` + fmt.Sprintf("%v", this.Jwk) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRegistry(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Trainer) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Trainer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trainer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Age |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gender", wireType)
			}
			m.Gender = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Gender |= (Trainer_Gender(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Start == nil {
				m.Start = &Timestamp{}
			}
			if err := m.Start.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pc == nil {
				m.Pc = &Pokemon_Collection{}
			}
			if err := m.Pc.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scope = append(m.Scope, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Token) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Access", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Access = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scope = append(m.Scope, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cert) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jwk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jwk = append(m.Jwk[:0], data[iNdEx:postIndex]...)
			if m.Jwk == nil {
				m.Jwk = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegistry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRegistry(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegistry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRegistry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRegistry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRegistry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRegistry(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRegistry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegistry   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorRegistry = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0xe3, 0xd8, 0x49, 0xb7, 0x4d, 0x14, 0x6d, 0x21, 0xb2, 0x92, 0x10, 0x2a, 0x1f, 0x10,
	0x70, 0xb0, 0x45, 0x38, 0xc1, 0x01, 0x89, 0x56, 0x28, 0x42, 0x42, 0x02, 0x59, 0xb9, 0x50, 0x84,
	0xd0, 0xc6, 0x99, 0xb8, 0x26, 0x89, 0xd7, 0xf2, 0x6e, 0xd4, 0x56, 0x08, 0x09, 0xf1, 0x04, 0x48,
	0xbc, 0x04, 0x6f, 0xc1, 0x95, 0x63, 0x25, 0x2e, 0x1c, 0x69, 0xe1, 0xc0, 0xb1, 0x8f, 0xc0, 0xec,
	0xda, 0xe1, 0xa7, 0x6a, 0x09, 0x87, 0x91, 0x67, 0x66, 0xbf, 0x99, 0x6f, 0xe7, 0x9b, 0x35, 0x69,
	0x64, 0x10, 0xc5, 0x42, 0x66, 0x87, 0x5e, 0x9a, 0x71, 0xc9, 0x69, 0x6b, 0xb4, 0x08, 0xa7, 0x7b,
	0x07, 0x9e, 0x60, 0x13, 0x96, 0xc5, 0xde, 0xf2, 0xb4, 0xdd, 0x8d, 0x38, 0x8f, 0x66, 0xe0, 0xb3,
	0x34, 0xf6, 0x59, 0x92, 0x70, 0xc9, 0x64, 0xcc, 0x13, 0x91, 0x57, 0xb5, 0xeb, 0x29, 0x9f, 0xc2,
	0x18, 0x0e, 0x8a, 0x70, 0x6d, 0x2e, 0xa2, 0xdc, 0x75, 0x3f, 0x96, 0x49, 0x75, 0x98, 0xb1, 0x38,
	0x81, 0x8c, 0x36, 0x89, 0xb9, 0x88, 0xc7, 0x8e, 0xb1, 0x65, 0x5c, 0x5f, 0x0b, 0x94, 0x4b, 0x29,
	0xa9, 0x24, 0x6c, 0x0e, 0x4e, 0x59, 0xa7, 0xb4, 0x4f, 0xdb, 0xa4, 0x96, 0x32, 0x21, 0xf6, 0x79,
	0x36, 0x76, 0x4c, 0x9d, 0xff, 0x15, 0xab, 0x0e, 0x2c, 0x02, 0xa7, 0x82, 0x69, 0x2b, 0x50, 0x2e,
	0xbd, 0x47, 0xec, 0x08, 0x92, 0x31, 0x64, 0x8e, 0x85, 0xc9, 0x46, 0xff, 0x9a, 0x77, 0xfe, 0x00,
	0x5e, 0x71, 0x09, 0x6f, 0xa0, 0xd1, 0x41, 0x51, 0x45, 0xfb, 0xc4, 0x12, 0x92, 0x65, 0xd2, 0xb1,
	0xb1, 0x7c, 0xbd, 0xdf, 0x3d, 0x53, 0xae, 0x06, 0x19, 0xc6, 0x73, 0x40, 0xcc, 0x3c, 0x0d, 0x72,
	0x28, 0xbd, 0x43, 0xca, 0x69, 0xe8, 0x54, 0x75, 0xc1, 0x8d, 0x33, 0x05, 0x4b, 0x21, 0x9e, 0xe0,
	0x77, 0xce, 0x13, 0x6f, 0x87, 0xcf, 0x66, 0x10, 0x2a, 0xad, 0x02, 0x2c, 0xa2, 0x97, 0x90, 0x2e,
	0xe4, 0x29, 0x38, 0xb5, 0x2d, 0x13, 0x27, 0xcb, 0x03, 0xb7, 0x43, 0xec, 0xfc, 0x5a, 0xb4, 0x4a,
	0xcc, 0xed, 0xc7, 0x4f, 0x9b, 0x25, 0x5a, 0x23, 0x95, 0xc1, 0xc3, 0xe0, 0x51, 0xd3, 0x70, 0x5f,
	0x10, 0x6b, 0x88, 0xcd, 0x12, 0xda, 0x22, 0x36, 0x0b, 0x43, 0x10, 0xa2, 0x50, 0xb0, 0x88, 0x94,
	0x28, 0x53, 0x38, 0x2c, 0x34, 0x54, 0xae, 0x42, 0x0a, 0x08, 0x33, 0x90, 0x85, 0x80, 0x45, 0xf4,
	0x9b, 0xbd, 0xf2, 0x27, 0xbb, 0x43, 0x2a, 0x3b, 0x80, 0x63, 0x61, 0x9f, 0x97, 0xfb, 0x53, 0xdd,
	0x7c, 0x23, 0x50, 0x6e, 0xff, 0xd4, 0x24, 0xb5, 0xa0, 0x10, 0x90, 0x3e, 0x5f, 0xfa, 0x78, 0xcd,
	0xab, 0x2b, 0x54, 0x6e, 0x77, 0xce, 0xd1, 0x31, 0x00, 0x91, 0xe2, 0x93, 0x01, 0x77, 0xf3, 0xed,
	0xe7, 0xef, 0xef, 0xcb, 0x75, 0xb7, 0xe6, 0xcb, 0x1c, 0x7e, 0xd7, 0xb8, 0x49, 0x27, 0x84, 0x0c,
	0x40, 0x2e, 0x9f, 0xca, 0x4a, 0x82, 0x55, 0x00, 0xb7, 0xa5, 0x49, 0x9a, 0xb4, 0xb1, 0x24, 0xf1,
	0x5f, 0xe1, 0x8b, 0x7b, 0x4d, 0xc7, 0xc4, 0x7a, 0x90, 0xfc, 0xd7, 0x0c, 0x57, 0x2e, 0x04, 0xa8,
	0x75, 0xb8, 0x1d, 0x4d, 0x70, 0x99, 0x6e, 0xfe, 0x4d, 0xe0, 0xb3, 0x85, 0xdc, 0xa3, 0xbb, 0xc4,
	0xbe, 0x9f, 0x6f, 0xe7, 0xdf, 0x5d, 0x56, 0x91, 0x50, 0x4d, 0xb2, 0xe1, 0x56, 0xfd, 0x7c, 0xd9,
	0x4a, 0xa9, 0x67, 0x64, 0x5d, 0xed, 0x2b, 0x9e, 0xc4, 0x21, 0x93, 0xb0, 0x7a, 0x8e, 0xee, 0x45,
	0x00, 0xd5, 0xc5, 0xad, 0x6b, 0x86, 0x2a, 0xb5, 0xfc, 0x10, 0xc3, 0xed, 0x5b, 0x47, 0xc7, 0xbd,
	0xd2, 0x17, 0xb4, 0xd3, 0xe3, 0x9e, 0xf1, 0xe6, 0xa4, 0x67, 0x7c, 0x40, 0xfb, 0x84, 0x76, 0x84,
	0xf6, 0x15, 0xed, 0xc7, 0x09, 0x9e, 0xe1, 0xf7, 0xdd, 0xb7, 0x5e, 0x69, 0xd7, 0x4c, 0x47, 0x93,
	0x91, 0xad, 0xff, 0xf4, 0xdb, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xe8, 0x9a, 0x31, 0x4b,
	0x04, 0x00, 0x00,
}

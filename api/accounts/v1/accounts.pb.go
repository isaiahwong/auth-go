// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts/v1/accounts.proto

package accounts

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8726aced901bdecf, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type AuthenticateResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8726aced901bdecf, []int{1}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type SignUpRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,3,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	CaptchaToken         string   `protobuf:"bytes,4,opt,name=captcha_token,json=captchaToken,proto3" json:"captcha_token,omitempty"`
	Ip                   string   `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8726aced901bdecf, []int{2}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *SignUpRequest) GetCaptchaToken() string {
	if m != nil {
		return m.CaptchaToken
	}
	return ""
}

func (m *SignUpRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	CaptchaToken         string   `protobuf:"bytes,3,opt,name=captcha_token,json=captchaToken,proto3" json:"captcha_token,omitempty"`
	Challenge            string   `protobuf:"bytes,4,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8726aced901bdecf, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetCaptchaToken() string {
	if m != nil {
		return m.CaptchaToken
	}
	return ""
}

func (m *LoginRequest) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8726aced901bdecf, []int{4}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.v1.accounts.Empty")
	proto.RegisterType((*AuthenticateResponse)(nil), "api.v1.accounts.AuthenticateResponse")
	proto.RegisterType((*SignUpRequest)(nil), "api.v1.accounts.SignUpRequest")
	proto.RegisterType((*LoginRequest)(nil), "api.v1.accounts.LoginRequest")
	proto.RegisterType((*UserResponse)(nil), "api.v1.accounts.UserResponse")
}

func init() { proto.RegisterFile("accounts/v1/accounts.proto", fileDescriptor_8726aced901bdecf) }

var fileDescriptor_8726aced901bdecf = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x95, 0xb4, 0x19, 0xd2, 0xc7, 0xb4, 0x19, 0xac, 0x32, 0x44, 0x51, 0x41, 0xc8, 0x08,
	0x89, 0x6e, 0x12, 0xb5, 0xac, 0x58, 0x16, 0x89, 0x05, 0x12, 0x0b, 0x94, 0xd2, 0x0d, 0x12, 0x2a,
	0x1e, 0xcf, 0x23, 0xb1, 0x48, 0x6c, 0x13, 0x3b, 0x83, 0xd8, 0xc2, 0x11, 0x38, 0x01, 0x97, 0xe0,
	0x22, 0x5c, 0x81, 0x83, 0xa0, 0xf1, 0x24, 0x61, 0x20, 0x23, 0x21, 0xb1, 0xcb, 0xfb, 0xdf, 0xef,
	0xcf, 0xbf, 0xf2, 0x1b, 0x12, 0xc6, 0xb9, 0x6a, 0xa5, 0x35, 0xd9, 0xea, 0x2c, 0xeb, 0xbf, 0x53,
	0xdd, 0x28, 0xab, 0x48, 0xc4, 0xb4, 0x48, 0x57, 0x67, 0x69, 0x2f, 0x27, 0x27, 0x85, 0x52, 0x45,
	0x85, 0x19, 0xd3, 0x22, 0x63, 0x52, 0x2a, 0xcb, 0xac, 0x50, 0xb2, 0xb3, 0x27, 0xf1, 0x36, 0xca,
	0xf0, 0x12, 0x6b, 0xb6, 0xd9, 0xd0, 0x1b, 0x10, 0x3c, 0xab, 0xb5, 0xfd, 0x44, 0x53, 0x38, 0xbe,
	0x68, 0x6d, 0x89, 0xd2, 0x0a, 0xce, 0x2c, 0xe6, 0x68, 0xb4, 0x92, 0x06, 0xc9, 0x1c, 0x26, 0xc6,
	0x32, 0xdb, 0x9a, 0xd8, 0xbb, 0xef, 0x3d, 0x0a, 0xf3, 0x6e, 0xa2, 0xdf, 0x3c, 0x38, 0xbc, 0x14,
	0x85, 0xbc, 0xd2, 0x39, 0x7e, 0x68, 0xd1, 0x58, 0x72, 0x0c, 0x01, 0xd6, 0x4c, 0x54, 0xce, 0x78,
	0x90, 0x6f, 0x06, 0x92, 0x40, 0xa8, 0x99, 0x31, 0x1f, 0x55, 0xb3, 0x8c, 0x7d, 0xb7, 0x18, 0x66,
	0x72, 0x0a, 0x33, 0xae, 0xe4, 0x3b, 0xd1, 0xd4, 0xd7, 0x83, 0x67, 0xcf, 0x79, 0xa2, 0x4e, 0x7f,
	0xd9, 0x5b, 0x1f, 0xc0, 0x21, 0x67, 0xda, 0xf2, 0x92, 0x5d, 0x5b, 0xf5, 0x1e, 0x65, 0xbc, 0xef,
	0x7c, 0xd3, 0x4e, 0x7c, 0xb5, 0xd6, 0xc8, 0x11, 0xf8, 0x42, 0xc7, 0x81, 0xdb, 0xf8, 0x42, 0xd3,
	0x2f, 0x1e, 0x4c, 0x5f, 0xa8, 0x42, 0xc8, 0xff, 0x8f, 0x38, 0xba, 0x77, 0x6f, 0xc7, 0xbd, 0x27,
	0x70, 0xc0, 0x4b, 0x56, 0x55, 0x28, 0x0b, 0xec, 0x82, 0xfd, 0x16, 0xe8, 0x13, 0x98, 0x5e, 0x19,
	0x6c, 0x86, 0x3f, 0x7a, 0x0a, 0xfb, 0xad, 0xc1, 0xc6, 0x65, 0xb8, 0x79, 0x7e, 0x3b, 0xfd, 0xab,
	0xca, 0xd4, 0x99, 0x9d, 0xe5, 0xfc, 0xbb, 0x0f, 0xd1, 0x45, 0xa7, 0x5f, 0x62, 0xb3, 0x12, 0x1c,
	0xc9, 0x02, 0xa2, 0xe7, 0x66, 0xbb, 0xaa, 0x25, 0x99, 0x8f, 0x18, 0xae, 0xd3, 0xe4, 0xe1, 0x48,
	0xdf, 0x55, 0x31, 0x9d, 0x7d, 0xfe, 0xf1, 0xf3, 0xab, 0x0f, 0x34, 0x74, 0x0f, 0xad, 0xb5, 0x25,
	0x79, 0x0b, 0x93, 0x4d, 0xb7, 0xe4, 0xde, 0x08, 0xf1, 0x47, 0xe9, 0xc9, 0xdd, 0xdd, 0xf1, 0x7b,
	0xf4, 0x1d, 0x87, 0xbe, 0x45, 0xa3, 0x1e, 0x9d, 0x19, 0x51, 0xc8, 0x56, 0x93, 0x37, 0x10, 0xb8,
	0x66, 0xc8, 0x18, 0xb0, 0xdd, 0xd8, 0xbf, 0xf8, 0x73, 0xc7, 0x9f, 0xd1, 0xa3, 0x81, 0x5f, 0xad,
	0x4f, 0x3f, 0x85, 0xd7, 0x61, 0x7f, 0x60, 0x31, 0x71, 0x2f, 0xfd, 0xf1, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x37, 0x3f, 0x67, 0x50, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsServiceClient is the client API for AccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsServiceClient interface {
	IsAuthenticated(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type accountsServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountsServiceClient(cc *grpc.ClientConn) AccountsServiceClient {
	return &accountsServiceClient{cc}
}

func (c *accountsServiceClient) IsAuthenticated(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/api.v1.accounts.AccountsService/IsAuthenticated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/api.v1.accounts.AccountsService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/api.v1.accounts.AccountsService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServiceServer is the server API for AccountsService service.
type AccountsServiceServer interface {
	IsAuthenticated(context.Context, *Empty) (*AuthenticateResponse, error)
	SignUp(context.Context, *SignUpRequest) (*UserResponse, error)
	Login(context.Context, *LoginRequest) (*UserResponse, error)
}

// UnimplementedAccountsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountsServiceServer struct {
}

func (*UnimplementedAccountsServiceServer) IsAuthenticated(ctx context.Context, req *Empty) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthenticated not implemented")
}
func (*UnimplementedAccountsServiceServer) SignUp(ctx context.Context, req *SignUpRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAccountsServiceServer) Login(ctx context.Context, req *LoginRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAccountsServiceServer(s *grpc.Server, srv AccountsServiceServer) {
	s.RegisterService(&_AccountsService_serviceDesc, srv)
}

func _AccountsService_IsAuthenticated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).IsAuthenticated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.accounts.AccountsService/IsAuthenticated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).IsAuthenticated(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.accounts.AccountsService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.accounts.AccountsService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.accounts.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthenticated",
			Handler:    _AccountsService_IsAuthenticated_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _AccountsService_SignUp_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AccountsService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/v1/accounts.proto",
}
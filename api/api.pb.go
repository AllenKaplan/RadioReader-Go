// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package RadioReader

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Song struct {
	Artist               string   `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (m *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(m, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *Song) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Song                 *Song    `protobuf:"bytes,2,opt,name=song,proto3" json:"song,omitempty"`
	Songs                []*Song  `protobuf:"bytes,3,rep,name=songs,proto3" json:"songs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *Response) GetSongs() []*Song {
	if m != nil {
		return m.Songs
	}
	return nil
}

type SongCountResponse struct {
	Song                 *Song    `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	Plays                int32    `protobuf:"varint,2,opt,name=plays,proto3" json:"plays,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SongCountResponse) Reset()         { *m = SongCountResponse{} }
func (m *SongCountResponse) String() string { return proto.CompactTextString(m) }
func (*SongCountResponse) ProtoMessage()    {}
func (*SongCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{3}
}

func (m *SongCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongCountResponse.Unmarshal(m, b)
}
func (m *SongCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongCountResponse.Marshal(b, m, deterministic)
}
func (m *SongCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongCountResponse.Merge(m, src)
}
func (m *SongCountResponse) XXX_Size() int {
	return xxx_messageInfo_SongCountResponse.Size(m)
}
func (m *SongCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SongCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SongCountResponse proto.InternalMessageInfo

func (m *SongCountResponse) GetSong() *Song {
	if m != nil {
		return m.Song
	}
	return nil
}

func (m *SongCountResponse) GetPlays() int32 {
	if m != nil {
		return m.Plays
	}
	return 0
}

type GetSongsByArtistRequest struct {
	Artist               string   `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSongsByArtistRequest) Reset()         { *m = GetSongsByArtistRequest{} }
func (m *GetSongsByArtistRequest) String() string { return proto.CompactTextString(m) }
func (*GetSongsByArtistRequest) ProtoMessage()    {}
func (*GetSongsByArtistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{4}
}

func (m *GetSongsByArtistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSongsByArtistRequest.Unmarshal(m, b)
}
func (m *GetSongsByArtistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSongsByArtistRequest.Marshal(b, m, deterministic)
}
func (m *GetSongsByArtistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSongsByArtistRequest.Merge(m, src)
}
func (m *GetSongsByArtistRequest) XXX_Size() int {
	return xxx_messageInfo_GetSongsByArtistRequest.Size(m)
}
func (m *GetSongsByArtistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSongsByArtistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSongsByArtistRequest proto.InternalMessageInfo

func (m *GetSongsByArtistRequest) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

type GetSongsByTitleRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSongsByTitleRequest) Reset()         { *m = GetSongsByTitleRequest{} }
func (m *GetSongsByTitleRequest) String() string { return proto.CompactTextString(m) }
func (*GetSongsByTitleRequest) ProtoMessage()    {}
func (*GetSongsByTitleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{5}
}

func (m *GetSongsByTitleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSongsByTitleRequest.Unmarshal(m, b)
}
func (m *GetSongsByTitleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSongsByTitleRequest.Marshal(b, m, deterministic)
}
func (m *GetSongsByTitleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSongsByTitleRequest.Merge(m, src)
}
func (m *GetSongsByTitleRequest) XXX_Size() int {
	return xxx_messageInfo_GetSongsByTitleRequest.Size(m)
}
func (m *GetSongsByTitleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSongsByTitleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSongsByTitleRequest proto.InternalMessageInfo

func (m *GetSongsByTitleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "RadioReader-Go.GetRequest")
	proto.RegisterType((*Song)(nil), "RadioReader-Go.Song")
	proto.RegisterType((*Response)(nil), "RadioReader-Go.Response")
	proto.RegisterType((*SongCountResponse)(nil), "RadioReader-Go.SongCountResponse")
	proto.RegisterType((*GetSongsByArtistRequest)(nil), "RadioReader-Go.GetSongsByArtistRequest")
	proto.RegisterType((*GetSongsByTitleRequest)(nil), "RadioReader-Go.GetSongsByTitleRequest")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0x53, 0xa0, 0xc0, 0x3b, 0xf0, 0x46, 0xd9, 0x20, 0x34, 0x1e, 0x0c, 0xa9, 0x1a, 0x3d,
	0xd5, 0x08, 0x5e, 0x3d, 0x54, 0x0f, 0x9c, 0x8c, 0x64, 0xf1, 0x0b, 0xac, 0x74, 0x42, 0x36, 0x69,
	0xba, 0x75, 0x77, 0x30, 0xe1, 0x33, 0xfb, 0x25, 0xcc, 0x2e, 0xff, 0xa1, 0xf5, 0xd4, 0xcc, 0xec,
	0x3c, 0xfb, 0x3c, 0xf3, 0xeb, 0xc2, 0x7f, 0x91, 0xcb, 0x07, 0x91, 0xcb, 0x28, 0xd7, 0x8a, 0x14,
	0x6b, 0x71, 0x91, 0x48, 0xc5, 0x51, 0x24, 0xa8, 0xc3, 0x36, 0xc0, 0x18, 0x89, 0xe3, 0xd7, 0x02,
	0x0d, 0x85, 0x4f, 0x50, 0x9b, 0xaa, 0x6c, 0xce, 0x7a, 0x50, 0x17, 0x9a, 0xa4, 0xa1, 0xc0, 0x1b,
	0x78, 0xf7, 0xff, 0xf8, 0xba, 0x62, 0x5d, 0xf0, 0x49, 0x52, 0x8a, 0x41, 0xc5, 0xb5, 0x57, 0x45,
	0x48, 0xd0, 0xe4, 0x68, 0x72, 0x95, 0x19, 0x64, 0x01, 0x34, 0x66, 0x1a, 0x05, 0x61, 0xe2, 0xa4,
	0x4d, 0xbe, 0x29, 0xd9, 0x2d, 0xd4, 0x8c, 0xca, 0xe6, 0x4e, 0xda, 0x1a, 0x76, 0xa2, 0xbd, 0x14,
	0x91, 0x35, 0xe5, 0xee, 0x98, 0xdd, 0x81, 0x6f, 0xbf, 0x26, 0xa8, 0x0e, 0xaa, 0xc5, 0x73, 0xab,
	0xf3, 0x70, 0x02, 0x1d, 0x5b, 0xbe, 0xaa, 0x45, 0x46, 0x5b, 0xfb, 0x8d, 0x89, 0xf7, 0xb7, 0x49,
	0x17, 0xfc, 0x3c, 0x15, 0x4b, 0xe3, 0xc2, 0xf8, 0x7c, 0x55, 0x84, 0x8f, 0xd0, 0x1f, 0x23, 0xd9,
	0x31, 0xf3, 0xb2, 0x8c, 0xdd, 0xc6, 0x6b, 0x30, 0x65, 0x40, 0xc2, 0x08, 0x7a, 0x3b, 0xc9, 0x87,
	0xa5, 0xb1, 0x51, 0x6c, 0x51, 0x79, 0x7b, 0xa8, 0x86, 0x3f, 0x15, 0x60, 0x7a, 0x97, 0x69, 0x8a,
	0xfa, 0x5b, 0xce, 0x90, 0x8d, 0xa0, 0x11, 0x27, 0x89, 0x43, 0x7f, 0x9a, 0xf9, 0xf2, 0xe2, 0xa0,
	0xb5, 0xdd, 0xf5, 0x19, 0x5a, 0x63, 0xa4, 0x38, 0x4d, 0x9d, 0x3d, 0xeb, 0x1f, 0x4c, 0xed, 0x7e,
	0x6a, 0x99, 0x3c, 0x86, 0xf6, 0x3a, 0xfa, 0xc4, 0x6e, 0x5f, 0x64, 0x7c, 0x75, 0xd2, 0x3a, 0xa4,
	0xfd, 0x0e, 0xe7, 0xc7, 0xc0, 0xd8, 0xcd, 0x71, 0x8c, 0x22, 0x9e, 0x65, 0x99, 0xde, 0xe0, 0xec,
	0x08, 0x27, 0xbb, 0x2e, 0xb9, 0x6f, 0x1f, 0x76, 0xc9, 0x75, 0x9f, 0x75, 0xf7, 0xe0, 0x47, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x53, 0xbc, 0x75, 0x9a, 0x01, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RadioReaderServiceClient is the client API for RadioReaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RadioReaderServiceClient interface {
	AddSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
	GetAllSongs(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	GetSongPlays(ctx context.Context, in *Song, opts ...grpc.CallOption) (*SongCountResponse, error)
	GetSongsByArtist(ctx context.Context, in *GetSongsByArtistRequest, opts ...grpc.CallOption) (*Response, error)
	GetSongsByTitle(ctx context.Context, in *GetSongsByTitleRequest, opts ...grpc.CallOption) (*Response, error)
}

type radioReaderServiceClient struct {
	cc *grpc.ClientConn
}

func NewRadioReaderServiceClient(cc *grpc.ClientConn) RadioReaderServiceClient {
	return &radioReaderServiceClient{cc}
}

func (c *radioReaderServiceClient) AddSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RadioReader-Go.radioReaderService/AddSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radioReaderServiceClient) GetAllSongs(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RadioReader-Go.radioReaderService/GetAllSongs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radioReaderServiceClient) GetSongPlays(ctx context.Context, in *Song, opts ...grpc.CallOption) (*SongCountResponse, error) {
	out := new(SongCountResponse)
	err := c.cc.Invoke(ctx, "/RadioReader-Go.radioReaderService/GetSongPlays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radioReaderServiceClient) GetSongsByArtist(ctx context.Context, in *GetSongsByArtistRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RadioReader-Go.radioReaderService/GetSongsByArtist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radioReaderServiceClient) GetSongsByTitle(ctx context.Context, in *GetSongsByTitleRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RadioReader-Go.radioReaderService/GetSongsByTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RadioReaderServiceServer is the server API for RadioReaderService service.
type RadioReaderServiceServer interface {
	AddSong(context.Context, *Song) (*Response, error)
	GetAllSongs(context.Context, *GetRequest) (*Response, error)
	GetSongPlays(context.Context, *Song) (*SongCountResponse, error)
	GetSongsByArtist(context.Context, *GetSongsByArtistRequest) (*Response, error)
	GetSongsByTitle(context.Context, *GetSongsByTitleRequest) (*Response, error)
}

// UnimplementedRadioReaderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRadioReaderServiceServer struct {
}

func (*UnimplementedRadioReaderServiceServer) AddSong(ctx context.Context, req *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (*UnimplementedRadioReaderServiceServer) GetAllSongs(ctx context.Context, req *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSongs not implemented")
}
func (*UnimplementedRadioReaderServiceServer) GetSongPlays(ctx context.Context, req *Song) (*SongCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongPlays not implemented")
}
func (*UnimplementedRadioReaderServiceServer) GetSongsByArtist(ctx context.Context, req *GetSongsByArtistRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongsByArtist not implemented")
}
func (*UnimplementedRadioReaderServiceServer) GetSongsByTitle(ctx context.Context, req *GetSongsByTitleRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongsByTitle not implemented")
}

func RegisterRadioReaderServiceServer(s *grpc.Server, srv RadioReaderServiceServer) {
	s.RegisterService(&_RadioReaderService_serviceDesc, srv)
}

func _RadioReaderService_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioReaderServiceServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RadioReader-Go.radioReaderService/AddSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioReaderServiceServer).AddSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _RadioReaderService_GetAllSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioReaderServiceServer).GetAllSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RadioReader-Go.radioReaderService/GetAllSongs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioReaderServiceServer).GetAllSongs(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RadioReaderService_GetSongPlays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioReaderServiceServer).GetSongPlays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RadioReader-Go.radioReaderService/GetSongPlays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioReaderServiceServer).GetSongPlays(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _RadioReaderService_GetSongsByArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongsByArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioReaderServiceServer).GetSongsByArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RadioReader-Go.radioReaderService/GetSongsByArtist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioReaderServiceServer).GetSongsByArtist(ctx, req.(*GetSongsByArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RadioReaderService_GetSongsByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongsByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioReaderServiceServer).GetSongsByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RadioReader-Go.radioReaderService/GetSongsByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioReaderServiceServer).GetSongsByTitle(ctx, req.(*GetSongsByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RadioReaderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RadioReader-Go.radioReaderService",
	HandlerType: (*RadioReaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSong",
			Handler:    _RadioReaderService_AddSong_Handler,
		},
		{
			MethodName: "GetAllSongs",
			Handler:    _RadioReaderService_GetAllSongs_Handler,
		},
		{
			MethodName: "GetSongPlays",
			Handler:    _RadioReaderService_GetSongPlays_Handler,
		},
		{
			MethodName: "GetSongsByArtist",
			Handler:    _RadioReaderService_GetSongsByArtist_Handler,
		},
		{
			MethodName: "GetSongsByTitle",
			Handler:    _RadioReaderService_GetSongsByTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

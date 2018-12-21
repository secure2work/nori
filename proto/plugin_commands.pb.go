// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin_commands.proto

package commands

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

type ErrorReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorReply) Reset()         { *m = ErrorReply{} }
func (m *ErrorReply) String() string { return proto.CompactTextString(m) }
func (*ErrorReply) ProtoMessage()    {}
func (*ErrorReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{0}
}
func (m *ErrorReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorReply.Unmarshal(m, b)
}
func (m *ErrorReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorReply.Marshal(b, m, deterministic)
}
func (dst *ErrorReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorReply.Merge(dst, src)
}
func (m *ErrorReply) XXX_Size() int {
	return xxx_messageInfo_ErrorReply.Size(m)
}
func (m *ErrorReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorReply.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorReply proto.InternalMessageInfo

func (m *ErrorReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ErrorReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// cmd: plugin list
type PluginListRequest struct {
	Installed            bool     `protobuf:"varint,1,opt,name=installed,proto3" json:"installed,omitempty"`
	Running              bool     `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginListRequest) Reset()         { *m = PluginListRequest{} }
func (m *PluginListRequest) String() string { return proto.CompactTextString(m) }
func (*PluginListRequest) ProtoMessage()    {}
func (*PluginListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{1}
}
func (m *PluginListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginListRequest.Unmarshal(m, b)
}
func (m *PluginListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginListRequest.Marshal(b, m, deterministic)
}
func (dst *PluginListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginListRequest.Merge(dst, src)
}
func (m *PluginListRequest) XXX_Size() int {
	return xxx_messageInfo_PluginListRequest.Size(m)
}
func (m *PluginListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginListRequest proto.InternalMessageInfo

func (m *PluginListRequest) GetInstalled() bool {
	if m != nil {
		return m.Installed
	}
	return false
}

func (m *PluginListRequest) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

type PluginDependency struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginDependency) Reset()         { *m = PluginDependency{} }
func (m *PluginDependency) String() string { return proto.CompactTextString(m) }
func (*PluginDependency) ProtoMessage()    {}
func (*PluginDependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{2}
}
func (m *PluginDependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginDependency.Unmarshal(m, b)
}
func (m *PluginDependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginDependency.Marshal(b, m, deterministic)
}
func (dst *PluginDependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginDependency.Merge(dst, src)
}
func (m *PluginDependency) XXX_Size() int {
	return xxx_messageInfo_PluginDependency.Size(m)
}
func (m *PluginDependency) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginDependency.DiscardUnknown(m)
}

var xxx_messageInfo_PluginDependency proto.InternalMessageInfo

func (m *PluginDependency) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PluginDependency) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PluginList struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author               string              `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Dependencies         []*PluginDependency `protobuf:"bytes,4,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Installed            bool                `protobuf:"varint,5,opt,name=installed,proto3" json:"installed,omitempty"`
	Running              bool                `protobuf:"varint,6,opt,name=running,proto3" json:"running,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PluginList) Reset()         { *m = PluginList{} }
func (m *PluginList) String() string { return proto.CompactTextString(m) }
func (*PluginList) ProtoMessage()    {}
func (*PluginList) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{3}
}
func (m *PluginList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginList.Unmarshal(m, b)
}
func (m *PluginList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginList.Marshal(b, m, deterministic)
}
func (dst *PluginList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginList.Merge(dst, src)
}
func (m *PluginList) XXX_Size() int {
	return xxx_messageInfo_PluginList.Size(m)
}
func (m *PluginList) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginList.DiscardUnknown(m)
}

var xxx_messageInfo_PluginList proto.InternalMessageInfo

func (m *PluginList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PluginList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginList) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *PluginList) GetDependencies() []*PluginDependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *PluginList) GetInstalled() bool {
	if m != nil {
		return m.Installed
	}
	return false
}

func (m *PluginList) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

type PluginListReply struct {
	Data                 []*PluginList `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Error                *ErrorReply   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PluginListReply) Reset()         { *m = PluginListReply{} }
func (m *PluginListReply) String() string { return proto.CompactTextString(m) }
func (*PluginListReply) ProtoMessage()    {}
func (*PluginListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{4}
}
func (m *PluginListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginListReply.Unmarshal(m, b)
}
func (m *PluginListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginListReply.Marshal(b, m, deterministic)
}
func (dst *PluginListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginListReply.Merge(dst, src)
}
func (m *PluginListReply) XXX_Size() int {
	return xxx_messageInfo_PluginListReply.Size(m)
}
func (m *PluginListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginListReply.DiscardUnknown(m)
}

var xxx_messageInfo_PluginListReply proto.InternalMessageInfo

func (m *PluginListReply) GetData() []*PluginList {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PluginListReply) GetError() *ErrorReply {
	if m != nil {
		return m.Error
	}
	return nil
}

// cmd: plugin get
type PluginGetRequest struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	InstallDependencies  bool     `protobuf:"varint,2,opt,name=install_dependencies,json=installDependencies,proto3" json:"install_dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginGetRequest) Reset()         { *m = PluginGetRequest{} }
func (m *PluginGetRequest) String() string { return proto.CompactTextString(m) }
func (*PluginGetRequest) ProtoMessage()    {}
func (*PluginGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{5}
}
func (m *PluginGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginGetRequest.Unmarshal(m, b)
}
func (m *PluginGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginGetRequest.Marshal(b, m, deterministic)
}
func (dst *PluginGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginGetRequest.Merge(dst, src)
}
func (m *PluginGetRequest) XXX_Size() int {
	return xxx_messageInfo_PluginGetRequest.Size(m)
}
func (m *PluginGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginGetRequest proto.InternalMessageInfo

func (m *PluginGetRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *PluginGetRequest) GetInstallDependencies() bool {
	if m != nil {
		return m.InstallDependencies
	}
	return false
}

// cmd: plugin remove
type PluginRemoveRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginRemoveRequest) Reset()         { *m = PluginRemoveRequest{} }
func (m *PluginRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*PluginRemoveRequest) ProtoMessage()    {}
func (*PluginRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{6}
}
func (m *PluginRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginRemoveRequest.Unmarshal(m, b)
}
func (m *PluginRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginRemoveRequest.Marshal(b, m, deterministic)
}
func (dst *PluginRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginRemoveRequest.Merge(dst, src)
}
func (m *PluginRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_PluginRemoveRequest.Size(m)
}
func (m *PluginRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginRemoveRequest proto.InternalMessageInfo

func (m *PluginRemoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// cmd: plugin install
type PluginInstallRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInstallRequest) Reset()         { *m = PluginInstallRequest{} }
func (m *PluginInstallRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInstallRequest) ProtoMessage()    {}
func (*PluginInstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{7}
}
func (m *PluginInstallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInstallRequest.Unmarshal(m, b)
}
func (m *PluginInstallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInstallRequest.Marshal(b, m, deterministic)
}
func (dst *PluginInstallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInstallRequest.Merge(dst, src)
}
func (m *PluginInstallRequest) XXX_Size() int {
	return xxx_messageInfo_PluginInstallRequest.Size(m)
}
func (m *PluginInstallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInstallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInstallRequest proto.InternalMessageInfo

func (m *PluginInstallRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// cmd: plugin uninstall
type PluginUninstallRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginUninstallRequest) Reset()         { *m = PluginUninstallRequest{} }
func (m *PluginUninstallRequest) String() string { return proto.CompactTextString(m) }
func (*PluginUninstallRequest) ProtoMessage()    {}
func (*PluginUninstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{8}
}
func (m *PluginUninstallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginUninstallRequest.Unmarshal(m, b)
}
func (m *PluginUninstallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginUninstallRequest.Marshal(b, m, deterministic)
}
func (dst *PluginUninstallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginUninstallRequest.Merge(dst, src)
}
func (m *PluginUninstallRequest) XXX_Size() int {
	return xxx_messageInfo_PluginUninstallRequest.Size(m)
}
func (m *PluginUninstallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginUninstallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginUninstallRequest proto.InternalMessageInfo

func (m *PluginUninstallRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// cmd: plugin upload
type PluginUploadRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	So                   []byte   `protobuf:"bytes,2,opt,name=so,proto3" json:"so,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginUploadRequest) Reset()         { *m = PluginUploadRequest{} }
func (m *PluginUploadRequest) String() string { return proto.CompactTextString(m) }
func (*PluginUploadRequest) ProtoMessage()    {}
func (*PluginUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{9}
}
func (m *PluginUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginUploadRequest.Unmarshal(m, b)
}
func (m *PluginUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginUploadRequest.Marshal(b, m, deterministic)
}
func (dst *PluginUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginUploadRequest.Merge(dst, src)
}
func (m *PluginUploadRequest) XXX_Size() int {
	return xxx_messageInfo_PluginUploadRequest.Size(m)
}
func (m *PluginUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginUploadRequest proto.InternalMessageInfo

func (m *PluginUploadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginUploadRequest) GetSo() []byte {
	if m != nil {
		return m.So
	}
	return nil
}

// cmd: plugin meta
type PluginMetaRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Dependencies         bool     `protobuf:"varint,2,opt,name=dependencies,proto3" json:"dependencies,omitempty"`
	DependenciesStatus   bool     `protobuf:"varint,3,opt,name=dependencies_status,json=dependenciesStatus,proto3" json:"dependencies_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginMetaRequest) Reset()         { *m = PluginMetaRequest{} }
func (m *PluginMetaRequest) String() string { return proto.CompactTextString(m) }
func (*PluginMetaRequest) ProtoMessage()    {}
func (*PluginMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{10}
}
func (m *PluginMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginMetaRequest.Unmarshal(m, b)
}
func (m *PluginMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginMetaRequest.Marshal(b, m, deterministic)
}
func (dst *PluginMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginMetaRequest.Merge(dst, src)
}
func (m *PluginMetaRequest) XXX_Size() int {
	return xxx_messageInfo_PluginMetaRequest.Size(m)
}
func (m *PluginMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginMetaRequest proto.InternalMessageInfo

func (m *PluginMetaRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PluginMetaRequest) GetDependencies() bool {
	if m != nil {
		return m.Dependencies
	}
	return false
}

func (m *PluginMetaRequest) GetDependenciesStatus() bool {
	if m != nil {
		return m.DependenciesStatus
	}
	return false
}

type PluginMetaReply struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginMetaReply) Reset()         { *m = PluginMetaReply{} }
func (m *PluginMetaReply) String() string { return proto.CompactTextString(m) }
func (*PluginMetaReply) ProtoMessage()    {}
func (*PluginMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_commands_2735544750f7d704, []int{11}
}
func (m *PluginMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginMetaReply.Unmarshal(m, b)
}
func (m *PluginMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginMetaReply.Marshal(b, m, deterministic)
}
func (dst *PluginMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginMetaReply.Merge(dst, src)
}
func (m *PluginMetaReply) XXX_Size() int {
	return xxx_messageInfo_PluginMetaReply.Size(m)
}
func (m *PluginMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_PluginMetaReply proto.InternalMessageInfo

func (m *PluginMetaReply) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorReply)(nil), "commands.ErrorReply")
	proto.RegisterType((*PluginListRequest)(nil), "commands.PluginListRequest")
	proto.RegisterType((*PluginDependency)(nil), "commands.PluginDependency")
	proto.RegisterType((*PluginList)(nil), "commands.PluginList")
	proto.RegisterType((*PluginListReply)(nil), "commands.PluginListReply")
	proto.RegisterType((*PluginGetRequest)(nil), "commands.PluginGetRequest")
	proto.RegisterType((*PluginRemoveRequest)(nil), "commands.PluginRemoveRequest")
	proto.RegisterType((*PluginInstallRequest)(nil), "commands.PluginInstallRequest")
	proto.RegisterType((*PluginUninstallRequest)(nil), "commands.PluginUninstallRequest")
	proto.RegisterType((*PluginUploadRequest)(nil), "commands.PluginUploadRequest")
	proto.RegisterType((*PluginMetaRequest)(nil), "commands.PluginMetaRequest")
	proto.RegisterType((*PluginMetaReply)(nil), "commands.PluginMetaReply")
}

func init() {
	proto.RegisterFile("plugin_commands.proto", fileDescriptor_plugin_commands_2735544750f7d704)
}

var fileDescriptor_plugin_commands_2735544750f7d704 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5f, 0xab, 0xd3, 0x30,
	0x1c, 0xa5, 0x7f, 0xee, 0xdc, 0x7e, 0x5e, 0xae, 0xd7, 0xac, 0x5e, 0x8a, 0xf8, 0x30, 0x02, 0x57,
	0x8a, 0x0f, 0x57, 0xd4, 0x27, 0x45, 0x7c, 0x9a, 0x88, 0xa8, 0x20, 0x91, 0xe1, 0xe3, 0x88, 0x6b,
	0x98, 0x91, 0x36, 0xa9, 0x49, 0x3a, 0xdc, 0xb7, 0xf3, 0xa3, 0x49, 0xf3, 0x67, 0x6b, 0x2b, 0xdb,
	0x5b, 0x4e, 0x7f, 0x27, 0xa7, 0x27, 0xe7, 0x24, 0xf0, 0xa8, 0xa9, 0xda, 0x2d, 0x17, 0xeb, 0x8d,
	0xac, 0x6b, 0x2a, 0x4a, 0x7d, 0xd7, 0x28, 0x69, 0x24, 0x9a, 0x06, 0x8c, 0xdf, 0x00, 0xbc, 0x57,
	0x4a, 0x2a, 0xc2, 0x9a, 0x6a, 0x8f, 0x6e, 0x60, 0xa2, 0x0d, 0x35, 0xad, 0xce, 0xa3, 0x45, 0x54,
	0x4c, 0x89, 0x47, 0x28, 0x83, 0x0b, 0xd6, 0xb1, 0xf2, 0x78, 0x11, 0x15, 0x33, 0xe2, 0x00, 0xfe,
	0x04, 0x0f, 0xbf, 0x5a, 0xf9, 0xcf, 0x5c, 0x1b, 0xc2, 0x7e, 0xb7, 0x4c, 0x1b, 0xf4, 0x04, 0x66,
	0x5c, 0x68, 0x43, 0xab, 0x8a, 0x95, 0x5e, 0xe5, 0xf8, 0x01, 0xe5, 0x70, 0x4f, 0xb5, 0x42, 0x70,
	0xb1, 0xb5, 0x52, 0x53, 0x12, 0x20, 0x7e, 0x0b, 0xd7, 0x4e, 0x6c, 0xc9, 0x1a, 0x26, 0x4a, 0x26,
	0x36, 0x7b, 0x74, 0x05, 0x31, 0x77, 0x22, 0x33, 0x12, 0x73, 0xbb, 0x7b, 0xc7, 0x94, 0xe6, 0x52,
	0x78, 0x23, 0x01, 0xe2, 0xbf, 0x11, 0xc0, 0xd1, 0xcb, 0x7f, 0x1b, 0x11, 0xa4, 0x82, 0xd6, 0xcc,
	0xef, 0xb2, 0xeb, 0xee, 0xac, 0xb4, 0x35, 0x3f, 0xa5, 0xca, 0x13, 0xfb, 0xd5, 0x23, 0xf4, 0x0e,
	0x2e, 0xcb, 0x60, 0x81, 0x33, 0x9d, 0xa7, 0x8b, 0xa4, 0xb8, 0xff, 0xf2, 0xf1, 0xdd, 0x21, 0xc2,
	0xb1, 0x4d, 0x32, 0xe0, 0x0f, 0x03, 0xb8, 0x38, 0x13, 0xc0, 0x64, 0x18, 0xc0, 0x16, 0x1e, 0xf4,
	0xd3, 0xec, 0xea, 0x28, 0x20, 0x2d, 0xa9, 0xa1, 0x79, 0x64, 0x2d, 0x64, 0x63, 0x0b, 0x96, 0x68,
	0x19, 0xe8, 0x59, 0xbf, 0xa0, 0x01, 0xf5, 0xd8, 0x6e, 0xa8, 0xed, 0x7b, 0x48, 0xfa, 0x03, 0x3b,
	0xb4, 0x76, 0x0d, 0x49, 0xab, 0xb8, 0x4f, 0xac, 0x5b, 0xa2, 0x17, 0x90, 0x79, 0xd7, 0xeb, 0x41,
	0x1c, 0xae, 0xb6, 0xb9, 0x9f, 0x2d, 0x7b, 0x23, 0x7c, 0x0b, 0x73, 0x27, 0x4c, 0x58, 0x2d, 0x77,
	0x2c, 0x68, 0x8f, 0xca, 0xc0, 0x4f, 0x21, 0x73, 0xb4, 0x8f, 0x4e, 0xe3, 0x14, 0xaf, 0x80, 0x1b,
	0xc7, 0x5b, 0x09, 0x7e, 0x9e, 0xf9, 0x3a, 0xfc, 0x78, 0xd5, 0x54, 0x92, 0x96, 0x81, 0x16, 0x5a,
	0x8f, 0x7a, 0xad, 0x5f, 0x41, 0xac, 0xa5, 0x3d, 0xc4, 0x25, 0x89, 0xb5, 0xc4, 0x7f, 0xc2, 0x1d,
	0xfe, 0xc2, 0x0c, 0x3d, 0xa1, 0x8f, 0xf0, 0xe8, 0x4a, 0xb8, 0x0c, 0x86, 0xb5, 0x3f, 0x87, 0x79,
	0x1f, 0xaf, 0xfd, 0x3b, 0x4a, 0x2c, 0x15, 0xf5, 0x47, 0xdf, 0xec, 0x04, 0xdf, 0x86, 0xbe, 0xdd,
	0x9f, 0xbb, 0xbe, 0x11, 0xa4, 0xbf, 0xb4, 0x14, 0xc1, 0x70, 0xb7, 0xfe, 0x31, 0xb1, 0x2f, 0xf6,
	0xd5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0xb0, 0xca, 0x3c, 0xca, 0x03, 0x00, 0x00,
}

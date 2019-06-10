// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/janna_api.proto

package apiv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Storage type. If 'cluster' type was chosen then Janna try to get DRS recommendation to choose proper 'datastore'.
type VMDeployRequest_Datastore_Type int32

const (
	// zero value
	VMDeployRequest_Datastore_TYPE_INVALID VMDeployRequest_Datastore_Type = 0
	// cluster. DRS will be applied.
	VMDeployRequest_Datastore_TYPE_CLUSTER VMDeployRequest_Datastore_Type = 1
	// single datastore.
	VMDeployRequest_Datastore_TYPE_DATASTORE VMDeployRequest_Datastore_Type = 2
)

var VMDeployRequest_Datastore_Type_name = map[int32]string{
	0: "TYPE_INVALID",
	1: "TYPE_CLUSTER",
	2: "TYPE_DATASTORE",
}

var VMDeployRequest_Datastore_Type_value = map[string]int32{
	"TYPE_INVALID":   0,
	"TYPE_CLUSTER":   1,
	"TYPE_DATASTORE": 2,
}

func (x VMDeployRequest_Datastore_Type) String() string {
	return proto.EnumName(VMDeployRequest_Datastore_Type_name, int32(x))
}

func (VMDeployRequest_Datastore_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{4, 1, 0}
}

// Computer Resource Type.
type VMDeployRequest_ComputerResources_Type int32

const (
	// zero value
	VMDeployRequest_ComputerResources_TYPE_INVALID VMDeployRequest_ComputerResources_Type = 0
	// host
	VMDeployRequest_ComputerResources_TYPE_HOST VMDeployRequest_ComputerResources_Type = 1
	// cluster
	VMDeployRequest_ComputerResources_TYPE_CLUSTER VMDeployRequest_ComputerResources_Type = 2
	// resource pool
	VMDeployRequest_ComputerResources_TYPE_RP VMDeployRequest_ComputerResources_Type = 3
)

var VMDeployRequest_ComputerResources_Type_name = map[int32]string{
	0: "TYPE_INVALID",
	1: "TYPE_HOST",
	2: "TYPE_CLUSTER",
	3: "TYPE_RP",
}

var VMDeployRequest_ComputerResources_Type_value = map[string]int32{
	"TYPE_INVALID": 0,
	"TYPE_HOST":    1,
	"TYPE_CLUSTER": 2,
	"TYPE_RP":      3,
}

func (x VMDeployRequest_ComputerResources_Type) String() string {
	return proto.EnumName(VMDeployRequest_ComputerResources_Type_name, int32(x))
}

func (VMDeployRequest_ComputerResources_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{4, 2, 0}
}

type AppInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfoRequest) Reset()         { *m = AppInfoRequest{} }
func (m *AppInfoRequest) String() string { return proto.CompactTextString(m) }
func (*AppInfoRequest) ProtoMessage()    {}
func (*AppInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{0}
}

func (m *AppInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfoRequest.Unmarshal(m, b)
}
func (m *AppInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfoRequest.Marshal(b, m, deterministic)
}
func (m *AppInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfoRequest.Merge(m, src)
}
func (m *AppInfoRequest) XXX_Size() int {
	return xxx_messageInfo_AppInfoRequest.Size(m)
}
func (m *AppInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfoRequest proto.InternalMessageInfo

type AppInfoResponse struct {
	Commit               string   `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	BuildTime            string   `protobuf:"bytes,2,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfoResponse) Reset()         { *m = AppInfoResponse{} }
func (m *AppInfoResponse) String() string { return proto.CompactTextString(m) }
func (*AppInfoResponse) ProtoMessage()    {}
func (*AppInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{1}
}

func (m *AppInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfoResponse.Unmarshal(m, b)
}
func (m *AppInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfoResponse.Marshal(b, m, deterministic)
}
func (m *AppInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfoResponse.Merge(m, src)
}
func (m *AppInfoResponse) XXX_Size() int {
	return xxx_messageInfo_AppInfoResponse.Size(m)
}
func (m *AppInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfoResponse proto.InternalMessageInfo

func (m *AppInfoResponse) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *AppInfoResponse) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

type VMInfoRequest struct {
	VmUuid               string   `protobuf:"bytes,1,opt,name=vm_uuid,json=vmUuid,proto3" json:"vm_uuid,omitempty"`
	Datacenter           string   `protobuf:"bytes,2,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMInfoRequest) Reset()         { *m = VMInfoRequest{} }
func (m *VMInfoRequest) String() string { return proto.CompactTextString(m) }
func (*VMInfoRequest) ProtoMessage()    {}
func (*VMInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{2}
}

func (m *VMInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMInfoRequest.Unmarshal(m, b)
}
func (m *VMInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMInfoRequest.Marshal(b, m, deterministic)
}
func (m *VMInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMInfoRequest.Merge(m, src)
}
func (m *VMInfoRequest) XXX_Size() int {
	return xxx_messageInfo_VMInfoRequest.Size(m)
}
func (m *VMInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VMInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VMInfoRequest proto.InternalMessageInfo

func (m *VMInfoRequest) GetVmUuid() string {
	if m != nil {
		return m.VmUuid
	}
	return ""
}

func (m *VMInfoRequest) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

type VMInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	GuestId              string   `protobuf:"bytes,3,opt,name=guest_id,json=guestId,proto3" json:"guest_id,omitempty"`
	Annotation           string   `protobuf:"bytes,4,opt,name=annotation,proto3" json:"annotation,omitempty"`
	PowerState           string   `protobuf:"bytes,5,opt,name=power_state,json=powerState,proto3" json:"power_state,omitempty"`
	NumCpu               uint32   `protobuf:"varint,6,opt,name=num_cpu,json=numCpu,proto3" json:"num_cpu,omitempty"`
	NumEthernetCards     uint32   `protobuf:"varint,7,opt,name=num_ethernet_cards,json=numEthernetCards,proto3" json:"num_ethernet_cards,omitempty"`
	NumVirtualDisks      uint32   `protobuf:"varint,8,opt,name=num_virtual_disks,json=numVirtualDisks,proto3" json:"num_virtual_disks,omitempty"`
	Template             bool     `protobuf:"varint,9,opt,name=template,proto3" json:"template,omitempty"`
	IpAddresses          []string `protobuf:"bytes,20,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMInfoResponse) Reset()         { *m = VMInfoResponse{} }
func (m *VMInfoResponse) String() string { return proto.CompactTextString(m) }
func (*VMInfoResponse) ProtoMessage()    {}
func (*VMInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{3}
}

func (m *VMInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMInfoResponse.Unmarshal(m, b)
}
func (m *VMInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMInfoResponse.Marshal(b, m, deterministic)
}
func (m *VMInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMInfoResponse.Merge(m, src)
}
func (m *VMInfoResponse) XXX_Size() int {
	return xxx_messageInfo_VMInfoResponse.Size(m)
}
func (m *VMInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VMInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VMInfoResponse proto.InternalMessageInfo

func (m *VMInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VMInfoResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *VMInfoResponse) GetGuestId() string {
	if m != nil {
		return m.GuestId
	}
	return ""
}

func (m *VMInfoResponse) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

func (m *VMInfoResponse) GetPowerState() string {
	if m != nil {
		return m.PowerState
	}
	return ""
}

func (m *VMInfoResponse) GetNumCpu() uint32 {
	if m != nil {
		return m.NumCpu
	}
	return 0
}

func (m *VMInfoResponse) GetNumEthernetCards() uint32 {
	if m != nil {
		return m.NumEthernetCards
	}
	return 0
}

func (m *VMInfoResponse) GetNumVirtualDisks() uint32 {
	if m != nil {
		return m.NumVirtualDisks
	}
	return 0
}

func (m *VMInfoResponse) GetTemplate() bool {
	if m != nil {
		return m.Template
	}
	return false
}

func (m *VMInfoResponse) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

type VMDeployRequest struct {
	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OvaUrl     string `protobuf:"bytes,2,opt,name=ova_url,json=ovaUrl,proto3" json:"ova_url,omitempty"`
	Datacenter string `protobuf:"bytes,3,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	Folder     string `protobuf:"bytes,4,opt,name=folder,proto3" json:"folder,omitempty"`
	Annotation string `protobuf:"bytes,5,opt,name=annotation,proto3" json:"annotation,omitempty"`
	// "OVF network": "ESXi network"
	Networks             map[string]string                  `protobuf:"bytes,10,rep,name=networks,proto3" json:"networks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Datastores           *VMDeployRequest_Datastore         `protobuf:"bytes,11,opt,name=datastores,proto3" json:"datastores,omitempty"`
	ComputerResources    *VMDeployRequest_ComputerResources `protobuf:"bytes,12,opt,name=computer_resources,json=computerResources,proto3" json:"computer_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *VMDeployRequest) Reset()         { *m = VMDeployRequest{} }
func (m *VMDeployRequest) String() string { return proto.CompactTextString(m) }
func (*VMDeployRequest) ProtoMessage()    {}
func (*VMDeployRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{4}
}

func (m *VMDeployRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDeployRequest.Unmarshal(m, b)
}
func (m *VMDeployRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDeployRequest.Marshal(b, m, deterministic)
}
func (m *VMDeployRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDeployRequest.Merge(m, src)
}
func (m *VMDeployRequest) XXX_Size() int {
	return xxx_messageInfo_VMDeployRequest.Size(m)
}
func (m *VMDeployRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDeployRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VMDeployRequest proto.InternalMessageInfo

func (m *VMDeployRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VMDeployRequest) GetOvaUrl() string {
	if m != nil {
		return m.OvaUrl
	}
	return ""
}

func (m *VMDeployRequest) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

func (m *VMDeployRequest) GetFolder() string {
	if m != nil {
		return m.Folder
	}
	return ""
}

func (m *VMDeployRequest) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

func (m *VMDeployRequest) GetNetworks() map[string]string {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *VMDeployRequest) GetDatastores() *VMDeployRequest_Datastore {
	if m != nil {
		return m.Datastores
	}
	return nil
}

func (m *VMDeployRequest) GetComputerResources() *VMDeployRequest_ComputerResources {
	if m != nil {
		return m.ComputerResources
	}
	return nil
}

// Describe a datastore.
type VMDeployRequest_Datastore struct {
	Type                 VMDeployRequest_Datastore_Type `protobuf:"varint,1,opt,name=type,proto3,enum=api.v1.VMDeployRequest_Datastore_Type" json:"type,omitempty"`
	Names                []string                       `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *VMDeployRequest_Datastore) Reset()         { *m = VMDeployRequest_Datastore{} }
func (m *VMDeployRequest_Datastore) String() string { return proto.CompactTextString(m) }
func (*VMDeployRequest_Datastore) ProtoMessage()    {}
func (*VMDeployRequest_Datastore) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{4, 1}
}

func (m *VMDeployRequest_Datastore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDeployRequest_Datastore.Unmarshal(m, b)
}
func (m *VMDeployRequest_Datastore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDeployRequest_Datastore.Marshal(b, m, deterministic)
}
func (m *VMDeployRequest_Datastore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDeployRequest_Datastore.Merge(m, src)
}
func (m *VMDeployRequest_Datastore) XXX_Size() int {
	return xxx_messageInfo_VMDeployRequest_Datastore.Size(m)
}
func (m *VMDeployRequest_Datastore) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDeployRequest_Datastore.DiscardUnknown(m)
}

var xxx_messageInfo_VMDeployRequest_Datastore proto.InternalMessageInfo

func (m *VMDeployRequest_Datastore) GetType() VMDeployRequest_Datastore_Type {
	if m != nil {
		return m.Type
	}
	return VMDeployRequest_Datastore_TYPE_INVALID
}

func (m *VMDeployRequest_Datastore) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// Describes a computer resource. Such as single host or cluster.
type VMDeployRequest_ComputerResources struct {
	Type                 VMDeployRequest_ComputerResources_Type `protobuf:"varint,1,opt,name=type,proto3,enum=api.v1.VMDeployRequest_ComputerResources_Type" json:"type,omitempty"`
	Path                 string                                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *VMDeployRequest_ComputerResources) Reset()         { *m = VMDeployRequest_ComputerResources{} }
func (m *VMDeployRequest_ComputerResources) String() string { return proto.CompactTextString(m) }
func (*VMDeployRequest_ComputerResources) ProtoMessage()    {}
func (*VMDeployRequest_ComputerResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{4, 2}
}

func (m *VMDeployRequest_ComputerResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDeployRequest_ComputerResources.Unmarshal(m, b)
}
func (m *VMDeployRequest_ComputerResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDeployRequest_ComputerResources.Marshal(b, m, deterministic)
}
func (m *VMDeployRequest_ComputerResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDeployRequest_ComputerResources.Merge(m, src)
}
func (m *VMDeployRequest_ComputerResources) XXX_Size() int {
	return xxx_messageInfo_VMDeployRequest_ComputerResources.Size(m)
}
func (m *VMDeployRequest_ComputerResources) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDeployRequest_ComputerResources.DiscardUnknown(m)
}

var xxx_messageInfo_VMDeployRequest_ComputerResources proto.InternalMessageInfo

func (m *VMDeployRequest_ComputerResources) GetType() VMDeployRequest_ComputerResources_Type {
	if m != nil {
		return m.Type
	}
	return VMDeployRequest_ComputerResources_TYPE_INVALID
}

func (m *VMDeployRequest_ComputerResources) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type VMDeployResponse struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMDeployResponse) Reset()         { *m = VMDeployResponse{} }
func (m *VMDeployResponse) String() string { return proto.CompactTextString(m) }
func (*VMDeployResponse) ProtoMessage()    {}
func (*VMDeployResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c929af3cc935ee9, []int{5}
}

func (m *VMDeployResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDeployResponse.Unmarshal(m, b)
}
func (m *VMDeployResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDeployResponse.Marshal(b, m, deterministic)
}
func (m *VMDeployResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDeployResponse.Merge(m, src)
}
func (m *VMDeployResponse) XXX_Size() int {
	return xxx_messageInfo_VMDeployResponse.Size(m)
}
func (m *VMDeployResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDeployResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VMDeployResponse proto.InternalMessageInfo

func (m *VMDeployResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.v1.VMDeployRequest_Datastore_Type", VMDeployRequest_Datastore_Type_name, VMDeployRequest_Datastore_Type_value)
	proto.RegisterEnum("api.v1.VMDeployRequest_ComputerResources_Type", VMDeployRequest_ComputerResources_Type_name, VMDeployRequest_ComputerResources_Type_value)
	proto.RegisterType((*AppInfoRequest)(nil), "api.v1.AppInfoRequest")
	proto.RegisterType((*AppInfoResponse)(nil), "api.v1.AppInfoResponse")
	proto.RegisterType((*VMInfoRequest)(nil), "api.v1.VMInfoRequest")
	proto.RegisterType((*VMInfoResponse)(nil), "api.v1.VMInfoResponse")
	proto.RegisterType((*VMDeployRequest)(nil), "api.v1.VMDeployRequest")
	proto.RegisterMapType((map[string]string)(nil), "api.v1.VMDeployRequest.NetworksEntry")
	proto.RegisterType((*VMDeployRequest_Datastore)(nil), "api.v1.VMDeployRequest.Datastore")
	proto.RegisterType((*VMDeployRequest_ComputerResources)(nil), "api.v1.VMDeployRequest.ComputerResources")
	proto.RegisterType((*VMDeployResponse)(nil), "api.v1.VMDeployResponse")
}

func init() { proto.RegisterFile("v1/janna_api.proto", fileDescriptor_6c929af3cc935ee9) }

var fileDescriptor_6c929af3cc935ee9 = []byte{
	// 1017 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0x49, 0x9a, 0x8f, 0x37, 0x6d, 0xd7, 0x9d, 0x5d, 0xb6, 0xde, 0x88, 0x0f, 0xaf, 0x25,
	0x50, 0x28, 0xbb, 0xc9, 0x26, 0xf4, 0x00, 0x5d, 0x09, 0xc9, 0xfd, 0x40, 0x0d, 0xea, 0x6e, 0x8b,
	0x9b, 0x46, 0x05, 0x21, 0x59, 0x53, 0x7b, 0x36, 0xf1, 0x36, 0x9e, 0x19, 0x3c, 0x63, 0x97, 0x08,
	0x71, 0xe1, 0xc6, 0x15, 0x6e, 0x48, 0x5c, 0x38, 0x72, 0x45, 0xfc, 0x05, 0x0e, 0x5c, 0xf9, 0x0b,
	0xfc, 0x05, 0xce, 0x20, 0x8f, 0xed, 0x7e, 0x24, 0x54, 0x70, 0xf2, 0xfb, 0x3e, 0xf3, 0xce, 0x33,
	0xcf, 0x3c, 0xe3, 0x79, 0x07, 0x50, 0xd2, 0xeb, 0xbe, 0xc4, 0x94, 0x62, 0x17, 0xf3, 0xa0, 0xc3,
	0x23, 0x26, 0x19, 0xaa, 0xa6, 0x61, 0xd2, 0x6b, 0xbd, 0x36, 0x66, 0x6c, 0x3c, 0x25, 0x5d, 0xcc,
	0x83, 0x2e, 0xa6, 0x94, 0x49, 0x2c, 0x03, 0x46, 0x45, 0x56, 0xd5, 0x7a, 0xa4, 0x3e, 0xde, 0xe3,
	0x31, 0xa1, 0x8f, 0xc5, 0x05, 0x1e, 0x8f, 0x49, 0xd4, 0x65, 0x5c, 0x55, 0x2c, 0x56, 0x5b, 0x3a,
	0xac, 0xda, 0x9c, 0x0f, 0xe8, 0x0b, 0xe6, 0x90, 0x2f, 0x62, 0x22, 0xa4, 0xb5, 0x0f, 0x77, 0x2e,
	0x11, 0xc1, 0x19, 0x15, 0x04, 0xdd, 0x87, 0xaa, 0xc7, 0xc2, 0x30, 0x90, 0x86, 0x66, 0x6a, 0xed,
	0x86, 0x93, 0x67, 0xe8, 0x75, 0x80, 0xb3, 0x38, 0x98, 0xfa, 0xae, 0x0c, 0x42, 0x62, 0x94, 0xd4,
	0x58, 0x43, 0x21, 0xc3, 0x20, 0x24, 0xd6, 0x3e, 0xac, 0x8c, 0x9e, 0x5d, 0xa3, 0x46, 0xeb, 0x50,
	0x4b, 0x42, 0x37, 0x8e, 0x03, 0xbf, 0x20, 0x4a, 0xc2, 0x93, 0x38, 0xf0, 0xd1, 0x1b, 0x00, 0x3e,
	0x96, 0xd8, 0x23, 0x54, 0x92, 0x28, 0x27, 0xba, 0x86, 0x58, 0xbf, 0x95, 0x60, 0xb5, 0xa0, 0xca,
	0x35, 0x21, 0xa8, 0x50, 0x1c, 0x92, 0x9c, 0x48, 0xc5, 0x29, 0xa6, 0xc8, 0x33, 0x02, 0x15, 0xa3,
	0x07, 0x50, 0x1f, 0xa7, 0x8b, 0xbb, 0x81, 0x6f, 0x94, 0x15, 0x5e, 0x53, 0xf9, 0x40, 0xad, 0x7a,
	0x65, 0x88, 0x51, 0xc9, 0x56, 0xbd, 0x42, 0xd0, 0x9b, 0xd0, 0xe4, 0xec, 0x82, 0x44, 0xae, 0x90,
	0x58, 0x12, 0x63, 0x29, 0x2b, 0x50, 0xd0, 0x71, 0x8a, 0xa4, 0xfb, 0xa1, 0x71, 0xe8, 0x7a, 0x3c,
	0x36, 0xaa, 0xa6, 0xd6, 0x5e, 0x71, 0xaa, 0x34, 0x0e, 0x77, 0x78, 0x8c, 0x1e, 0x01, 0x4a, 0x07,
	0x88, 0x9c, 0x90, 0x88, 0x12, 0xe9, 0x7a, 0x38, 0xf2, 0x85, 0x51, 0x53, 0x35, 0x3a, 0x8d, 0xc3,
	0xbd, 0x7c, 0x60, 0x27, 0xc5, 0xd1, 0x06, 0xac, 0xa5, 0xd5, 0x49, 0x10, 0xc9, 0x18, 0x4f, 0x5d,
	0x3f, 0x10, 0xe7, 0xc2, 0xa8, 0xab, 0xe2, 0x3b, 0x34, 0x0e, 0x47, 0x19, 0xbe, 0x9b, 0xc2, 0xa8,
	0x05, 0x75, 0x49, 0x42, 0x3e, 0x4d, 0x05, 0x35, 0x4c, 0xad, 0x5d, 0x77, 0x2e, 0x73, 0xf4, 0x10,
	0x96, 0x03, 0xee, 0x62, 0xdf, 0x8f, 0x88, 0x10, 0x44, 0x18, 0xf7, 0xcc, 0x72, 0xbb, 0xe1, 0x34,
	0x03, 0x6e, 0x17, 0x90, 0xf5, 0x6d, 0x15, 0xee, 0x8c, 0x9e, 0xed, 0x12, 0x3e, 0x65, 0xb3, 0xe2,
	0x54, 0xfe, 0xcd, 0xc9, 0x75, 0xa8, 0xb1, 0x04, 0xbb, 0x71, 0x34, 0xcd, 0xcd, 0xac, 0xb2, 0x04,
	0x9f, 0x44, 0xd3, 0xb9, 0x93, 0x2a, 0xcf, 0x9f, 0x54, 0xfa, 0xab, 0xbc, 0x60, 0x53, 0x9f, 0x44,
	0xb9, 0x9f, 0x79, 0x36, 0xe7, 0xf5, 0xd2, 0x82, 0xd7, 0x36, 0xd4, 0x29, 0x91, 0x17, 0x2c, 0x3a,
	0x17, 0x06, 0x98, 0xe5, 0x76, 0xb3, 0xff, 0x56, 0x27, 0xfb, 0xdd, 0x3b, 0x73, 0x7a, 0x3b, 0xcf,
	0xf3, 0xba, 0x3d, 0x2a, 0xa3, 0x99, 0x73, 0x39, 0x0d, 0xd9, 0x99, 0x34, 0x21, 0x59, 0x44, 0x84,
	0xd1, 0x34, 0xb5, 0x76, 0xb3, 0xff, 0xf0, 0x36, 0x92, 0xdd, 0xa2, 0xd2, 0xb9, 0x36, 0x09, 0x9d,
	0x02, 0xf2, 0x58, 0xc8, 0x63, 0x49, 0x22, 0x37, 0x22, 0x82, 0xc5, 0x91, 0x47, 0x84, 0xb1, 0xac,
	0xa8, 0xde, 0xb9, 0x8d, 0x6a, 0x27, 0x9f, 0xe1, 0x14, 0x13, 0x9c, 0x35, 0x6f, 0x1e, 0x6a, 0x3d,
	0x85, 0x95, 0x1b, 0xba, 0x91, 0x0e, 0xe5, 0x73, 0x32, 0xcb, 0x4d, 0x4f, 0x43, 0x74, 0x0f, 0x96,
	0x12, 0x3c, 0x8d, 0x8b, 0x8b, 0x94, 0x25, 0x5b, 0xa5, 0xf7, 0xb5, 0xd6, 0x8f, 0x1a, 0x34, 0x2e,
	0x05, 0xa3, 0x2d, 0xa8, 0xc8, 0x19, 0xcf, 0xce, 0x6b, 0xb5, 0xff, 0xf6, 0x7f, 0xee, 0xb0, 0x33,
	0x9c, 0x71, 0xe2, 0xa8, 0x39, 0xe9, 0x1a, 0xe9, 0xf9, 0x0a, 0xa3, 0xa4, 0xfe, 0x8d, 0x2c, 0xb1,
	0x3e, 0x84, 0x4a, 0x5a, 0x83, 0x74, 0x58, 0x1e, 0x7e, 0x7a, 0xb4, 0xe7, 0x0e, 0x9e, 0x8f, 0xec,
	0x83, 0xc1, 0xae, 0xfe, 0xca, 0x25, 0xb2, 0x73, 0x70, 0x72, 0x3c, 0xdc, 0x73, 0x74, 0x0d, 0x21,
	0x58, 0x55, 0xc8, 0xae, 0x3d, 0xb4, 0x8f, 0x87, 0x87, 0xce, 0x9e, 0x5e, 0x6a, 0xfd, 0xa2, 0xc1,
	0xda, 0x82, 0x0b, 0x68, 0xfb, 0x86, 0xce, 0xce, 0xff, 0xb6, 0xef, 0xba, 0x5e, 0x04, 0x15, 0x8e,
	0xe5, 0xa4, 0xb8, 0xd1, 0x69, 0x6c, 0x7d, 0x74, 0xab, 0xda, 0x15, 0x68, 0x28, 0x64, 0xff, 0xf0,
	0x78, 0xa8, 0x6b, 0x0b, 0xe2, 0x4b, 0xa8, 0x09, 0x35, 0x85, 0x38, 0x47, 0x7a, 0xd9, 0x7a, 0x17,
	0xf4, 0x2b, 0x2d, 0x79, 0x57, 0x59, 0x87, 0x9a, 0xc4, 0xe2, 0xdc, 0xbd, 0xea, 0x50, 0x69, 0x3a,
	0xf0, 0xfb, 0x7f, 0x6b, 0x50, 0xff, 0x38, 0xed, 0xc7, 0xf6, 0xd1, 0x00, 0x1d, 0x40, 0x2d, 0x6f,
	0x91, 0xe8, 0x7e, 0xb1, 0xad, 0x9b, 0x5d, 0xb4, 0xb5, 0xbe, 0x80, 0x67, 0x2b, 0x58, 0xfa, 0x37,
	0x7f, 0xfc, 0xf9, 0x7d, 0x09, 0x50, 0xbd, 0x9b, 0xf4, 0xba, 0x41, 0x4a, 0xe1, 0x40, 0x35, 0xeb,
	0x6d, 0xe8, 0xd5, 0x2b, 0x8f, 0xae, 0x73, 0xdd, 0x9f, 0x87, 0x73, 0xaa, 0x07, 0x8a, 0xea, 0x2e,
	0x5a, 0x4b, 0xa9, 0x92, 0x50, 0x74, 0xbf, 0xca, 0xbb, 0xeb, 0xd7, 0xe8, 0x13, 0xa8, 0x17, 0x7b,
	0x43, 0xeb, 0xb7, 0x38, 0xdf, 0x32, 0x16, 0x07, 0x72, 0x66, 0xa4, 0x98, 0x97, 0xad, 0x5a, 0xce,
	0xbc, 0xa5, 0x6d, 0x6c, 0xff, 0xa5, 0x7d, 0x67, 0xff, 0xaa, 0xa1, 0x43, 0x58, 0x51, 0x3e, 0x98,
	0x82, 0x44, 0x49, 0xe0, 0x11, 0xeb, 0x03, 0xb8, 0x3b, 0x9a, 0x61, 0x6f, 0x42, 0xc4, 0x14, 0x27,
	0xe6, 0x90, 0x44, 0x7e, 0x4c, 0x59, 0x82, 0xac, 0x89, 0x94, 0x5c, 0x6c, 0x75, 0xbb, 0xe3, 0x40,
	0x4e, 0xe2, 0xb3, 0x8e, 0xc7, 0xc2, 0x6e, 0x22, 0xf3, 0xd1, 0xec, 0x69, 0xeb, 0x97, 0x7b, 0x9d,
	0x27, 0x1b, 0x9a, 0xd6, 0xd7, 0x31, 0xe7, 0xd3, 0xc0, 0x53, 0x9d, 0xa0, 0xfb, 0x52, 0x30, 0xba,
	0xb5, 0x80, 0x38, 0x4f, 0xa1, 0xbc, 0xf9, 0x64, 0x13, 0x6d, 0xc2, 0x86, 0x43, 0x64, 0x1c, 0x51,
	0xe2, 0x9b, 0x17, 0x13, 0x42, 0x4d, 0x39, 0x21, 0x66, 0x71, 0x55, 0x4d, 0x9f, 0x11, 0x61, 0x52,
	0x26, 0x4d, 0xf2, 0x65, 0x20, 0x64, 0x07, 0x55, 0xa1, 0xf2, 0x43, 0x49, 0xab, 0x01, 0x78, 0x2c,
	0xcc, 0x37, 0xbb, 0x9d, 0xc9, 0xb7, 0x79, 0x70, 0x94, 0x3e, 0x80, 0x47, 0xda, 0x67, 0x4b, 0x98,
	0x07, 0x49, 0xef, 0xa7, 0x52, 0xd9, 0x3e, 0x3d, 0xfd, 0xb9, 0x54, 0xb5, 0x79, 0xd0, 0x19, 0xf5,
	0x7e, 0x57, 0xc1, 0xe7, 0xa3, 0xde, 0x59, 0x55, 0x3d, 0x94, 0xef, 0xfd, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x92, 0x0f, 0xae, 0x47, 0x92, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JannaAPIClient is the client API for JannaAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JannaAPIClient interface {
	// Get the application build information.
	AppInfo(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*AppInfoResponse, error)
	// Get a Virtual Machine information by its UUID.
	VMInfo(ctx context.Context, in *VMInfoRequest, opts ...grpc.CallOption) (*VMInfoResponse, error)
	// Deploy Virtual Machine.
	VMDeploy(ctx context.Context, in *VMDeployRequest, opts ...grpc.CallOption) (*VMDeployResponse, error)
}

type jannaAPIClient struct {
	cc *grpc.ClientConn
}

func NewJannaAPIClient(cc *grpc.ClientConn) JannaAPIClient {
	return &jannaAPIClient{cc}
}

func (c *jannaAPIClient) AppInfo(ctx context.Context, in *AppInfoRequest, opts ...grpc.CallOption) (*AppInfoResponse, error) {
	out := new(AppInfoResponse)
	err := c.cc.Invoke(ctx, "/api.v1.JannaAPI/AppInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jannaAPIClient) VMInfo(ctx context.Context, in *VMInfoRequest, opts ...grpc.CallOption) (*VMInfoResponse, error) {
	out := new(VMInfoResponse)
	err := c.cc.Invoke(ctx, "/api.v1.JannaAPI/VMInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jannaAPIClient) VMDeploy(ctx context.Context, in *VMDeployRequest, opts ...grpc.CallOption) (*VMDeployResponse, error) {
	out := new(VMDeployResponse)
	err := c.cc.Invoke(ctx, "/api.v1.JannaAPI/VMDeploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JannaAPIServer is the server API for JannaAPI service.
type JannaAPIServer interface {
	// Get the application build information.
	AppInfo(context.Context, *AppInfoRequest) (*AppInfoResponse, error)
	// Get a Virtual Machine information by its UUID.
	VMInfo(context.Context, *VMInfoRequest) (*VMInfoResponse, error)
	// Deploy Virtual Machine.
	VMDeploy(context.Context, *VMDeployRequest) (*VMDeployResponse, error)
}

func RegisterJannaAPIServer(s *grpc.Server, srv JannaAPIServer) {
	s.RegisterService(&_JannaAPI_serviceDesc, srv)
}

func _JannaAPI_AppInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JannaAPIServer).AppInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.JannaAPI/AppInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JannaAPIServer).AppInfo(ctx, req.(*AppInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JannaAPI_VMInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JannaAPIServer).VMInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.JannaAPI/VMInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JannaAPIServer).VMInfo(ctx, req.(*VMInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JannaAPI_VMDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JannaAPIServer).VMDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.JannaAPI/VMDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JannaAPIServer).VMDeploy(ctx, req.(*VMDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JannaAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.JannaAPI",
	HandlerType: (*JannaAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppInfo",
			Handler:    _JannaAPI_AppInfo_Handler,
		},
		{
			MethodName: "VMInfo",
			Handler:    _JannaAPI_VMInfo_Handler,
		},
		{
			MethodName: "VMDeploy",
			Handler:    _JannaAPI_VMDeploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/janna_api.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/public/uac/Collaborator.proto

package uac

import (
	context "context"
	fmt "fmt"
	common "github.com/VertaAI/modeldb/protos/gen/go/protos/public/common"
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

type ShareViaEnum int32

const (
	ShareViaEnum_USER_ID  ShareViaEnum = 0
	ShareViaEnum_EMAIL_ID ShareViaEnum = 1
	ShareViaEnum_USERNAME ShareViaEnum = 2
)

var ShareViaEnum_name = map[int32]string{
	0: "USER_ID",
	1: "EMAIL_ID",
	2: "USERNAME",
}

var ShareViaEnum_value = map[string]int32{
	"USER_ID":  0,
	"EMAIL_ID": 1,
	"USERNAME": 2,
}

func (x ShareViaEnum) String() string {
	return proto.EnumName(ShareViaEnum_name, int32(x))
}

func (ShareViaEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{0}
}

type AddCollaboratorRequest struct {
	EntityIds            []string                                     `protobuf:"bytes,1,rep,name=entity_ids,json=entityIds,proto3" json:"entity_ids,omitempty"`
	ShareWith            string                                       `protobuf:"bytes,2,opt,name=share_with,json=shareWith,proto3" json:"share_with,omitempty"`
	CollaboratorType     common.CollaboratorTypeEnum_CollaboratorType `protobuf:"varint,3,opt,name=collaborator_type,json=collaboratorType,proto3,enum=ai.verta.common.CollaboratorTypeEnum_CollaboratorType" json:"collaborator_type,omitempty"`
	Message              string                                       `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	DateCreated          uint64                                       `protobuf:"varint,5,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated          uint64                                       `protobuf:"varint,6,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	CanDeploy            common.TernaryEnum_Ternary                   `protobuf:"varint,7,opt,name=can_deploy,json=canDeploy,proto3,enum=ai.verta.common.TernaryEnum_Ternary" json:"can_deploy,omitempty"`
	AuthzEntityType      EntitiesEnum_EntitiesTypes                   `protobuf:"varint,8,opt,name=authz_entity_type,json=authzEntityType,proto3,enum=ai.verta.uac.EntitiesEnum_EntitiesTypes" json:"authz_entity_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *AddCollaboratorRequest) Reset()         { *m = AddCollaboratorRequest{} }
func (m *AddCollaboratorRequest) String() string { return proto.CompactTextString(m) }
func (*AddCollaboratorRequest) ProtoMessage()    {}
func (*AddCollaboratorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{0}
}

func (m *AddCollaboratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCollaboratorRequest.Unmarshal(m, b)
}
func (m *AddCollaboratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCollaboratorRequest.Marshal(b, m, deterministic)
}
func (m *AddCollaboratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCollaboratorRequest.Merge(m, src)
}
func (m *AddCollaboratorRequest) XXX_Size() int {
	return xxx_messageInfo_AddCollaboratorRequest.Size(m)
}
func (m *AddCollaboratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCollaboratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddCollaboratorRequest proto.InternalMessageInfo

func (m *AddCollaboratorRequest) GetEntityIds() []string {
	if m != nil {
		return m.EntityIds
	}
	return nil
}

func (m *AddCollaboratorRequest) GetShareWith() string {
	if m != nil {
		return m.ShareWith
	}
	return ""
}

func (m *AddCollaboratorRequest) GetCollaboratorType() common.CollaboratorTypeEnum_CollaboratorType {
	if m != nil {
		return m.CollaboratorType
	}
	return common.CollaboratorTypeEnum_READ_ONLY
}

func (m *AddCollaboratorRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddCollaboratorRequest) GetDateCreated() uint64 {
	if m != nil {
		return m.DateCreated
	}
	return 0
}

func (m *AddCollaboratorRequest) GetDateUpdated() uint64 {
	if m != nil {
		return m.DateUpdated
	}
	return 0
}

func (m *AddCollaboratorRequest) GetCanDeploy() common.TernaryEnum_Ternary {
	if m != nil {
		return m.CanDeploy
	}
	return common.TernaryEnum_UNKNOWN
}

func (m *AddCollaboratorRequest) GetAuthzEntityType() EntitiesEnum_EntitiesTypes {
	if m != nil {
		return m.AuthzEntityType
	}
	return EntitiesEnum_UNKNOWN
}

type AddCollaboratorRequest_Response struct {
	SelfAllowedActions []*Action `protobuf:"bytes,5,rep,name=self_allowed_actions,json=selfAllowedActions,proto3" json:"self_allowed_actions,omitempty"`
	Status             bool      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are valid to be assigned to Collaborator:
	//	*AddCollaboratorRequest_Response_CollaboratorUserInfo
	//	*AddCollaboratorRequest_Response_CollaboratorOrganization
	//	*AddCollaboratorRequest_Response_CollaboratorTeam
	Collaborator         isAddCollaboratorRequest_Response_Collaborator `protobuf_oneof:"collaborator"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *AddCollaboratorRequest_Response) Reset()         { *m = AddCollaboratorRequest_Response{} }
func (m *AddCollaboratorRequest_Response) String() string { return proto.CompactTextString(m) }
func (*AddCollaboratorRequest_Response) ProtoMessage()    {}
func (*AddCollaboratorRequest_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{0, 0}
}

func (m *AddCollaboratorRequest_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCollaboratorRequest_Response.Unmarshal(m, b)
}
func (m *AddCollaboratorRequest_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCollaboratorRequest_Response.Marshal(b, m, deterministic)
}
func (m *AddCollaboratorRequest_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCollaboratorRequest_Response.Merge(m, src)
}
func (m *AddCollaboratorRequest_Response) XXX_Size() int {
	return xxx_messageInfo_AddCollaboratorRequest_Response.Size(m)
}
func (m *AddCollaboratorRequest_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCollaboratorRequest_Response.DiscardUnknown(m)
}

var xxx_messageInfo_AddCollaboratorRequest_Response proto.InternalMessageInfo

func (m *AddCollaboratorRequest_Response) GetSelfAllowedActions() []*Action {
	if m != nil {
		return m.SelfAllowedActions
	}
	return nil
}

func (m *AddCollaboratorRequest_Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type isAddCollaboratorRequest_Response_Collaborator interface {
	isAddCollaboratorRequest_Response_Collaborator()
}

type AddCollaboratorRequest_Response_CollaboratorUserInfo struct {
	CollaboratorUserInfo *UserInfo `protobuf:"bytes,2,opt,name=collaborator_user_info,json=collaboratorUserInfo,proto3,oneof"`
}

type AddCollaboratorRequest_Response_CollaboratorOrganization struct {
	CollaboratorOrganization *Organization `protobuf:"bytes,3,opt,name=collaborator_organization,json=collaboratorOrganization,proto3,oneof"`
}

type AddCollaboratorRequest_Response_CollaboratorTeam struct {
	CollaboratorTeam *Team `protobuf:"bytes,4,opt,name=collaborator_team,json=collaboratorTeam,proto3,oneof"`
}

func (*AddCollaboratorRequest_Response_CollaboratorUserInfo) isAddCollaboratorRequest_Response_Collaborator() {
}

func (*AddCollaboratorRequest_Response_CollaboratorOrganization) isAddCollaboratorRequest_Response_Collaborator() {
}

func (*AddCollaboratorRequest_Response_CollaboratorTeam) isAddCollaboratorRequest_Response_Collaborator() {
}

func (m *AddCollaboratorRequest_Response) GetCollaborator() isAddCollaboratorRequest_Response_Collaborator {
	if m != nil {
		return m.Collaborator
	}
	return nil
}

func (m *AddCollaboratorRequest_Response) GetCollaboratorUserInfo() *UserInfo {
	if x, ok := m.GetCollaborator().(*AddCollaboratorRequest_Response_CollaboratorUserInfo); ok {
		return x.CollaboratorUserInfo
	}
	return nil
}

func (m *AddCollaboratorRequest_Response) GetCollaboratorOrganization() *Organization {
	if x, ok := m.GetCollaborator().(*AddCollaboratorRequest_Response_CollaboratorOrganization); ok {
		return x.CollaboratorOrganization
	}
	return nil
}

func (m *AddCollaboratorRequest_Response) GetCollaboratorTeam() *Team {
	if x, ok := m.GetCollaborator().(*AddCollaboratorRequest_Response_CollaboratorTeam); ok {
		return x.CollaboratorTeam
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AddCollaboratorRequest_Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AddCollaboratorRequest_Response_CollaboratorUserInfo)(nil),
		(*AddCollaboratorRequest_Response_CollaboratorOrganization)(nil),
		(*AddCollaboratorRequest_Response_CollaboratorTeam)(nil),
	}
}

type RemoveCollaborator struct {
	EntityId             string                     `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	ShareWith            string                     `protobuf:"bytes,2,opt,name=share_with,json=shareWith,proto3" json:"share_with,omitempty"`
	DateDeleted          uint64                     `protobuf:"varint,3,opt,name=date_deleted,json=dateDeleted,proto3" json:"date_deleted,omitempty"`
	AuthzEntityType      EntitiesEnum_EntitiesTypes `protobuf:"varint,4,opt,name=authz_entity_type,json=authzEntityType,proto3,enum=ai.verta.uac.EntitiesEnum_EntitiesTypes" json:"authz_entity_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RemoveCollaborator) Reset()         { *m = RemoveCollaborator{} }
func (m *RemoveCollaborator) String() string { return proto.CompactTextString(m) }
func (*RemoveCollaborator) ProtoMessage()    {}
func (*RemoveCollaborator) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{1}
}

func (m *RemoveCollaborator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveCollaborator.Unmarshal(m, b)
}
func (m *RemoveCollaborator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveCollaborator.Marshal(b, m, deterministic)
}
func (m *RemoveCollaborator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveCollaborator.Merge(m, src)
}
func (m *RemoveCollaborator) XXX_Size() int {
	return xxx_messageInfo_RemoveCollaborator.Size(m)
}
func (m *RemoveCollaborator) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveCollaborator.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveCollaborator proto.InternalMessageInfo

func (m *RemoveCollaborator) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *RemoveCollaborator) GetShareWith() string {
	if m != nil {
		return m.ShareWith
	}
	return ""
}

func (m *RemoveCollaborator) GetDateDeleted() uint64 {
	if m != nil {
		return m.DateDeleted
	}
	return 0
}

func (m *RemoveCollaborator) GetAuthzEntityType() EntitiesEnum_EntitiesTypes {
	if m != nil {
		return m.AuthzEntityType
	}
	return EntitiesEnum_UNKNOWN
}

type RemoveCollaborator_Response struct {
	Status               bool      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	SelfAllowedActions   []*Action `protobuf:"bytes,5,rep,name=self_allowed_actions,json=selfAllowedActions,proto3" json:"self_allowed_actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RemoveCollaborator_Response) Reset()         { *m = RemoveCollaborator_Response{} }
func (m *RemoveCollaborator_Response) String() string { return proto.CompactTextString(m) }
func (*RemoveCollaborator_Response) ProtoMessage()    {}
func (*RemoveCollaborator_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{1, 0}
}

func (m *RemoveCollaborator_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveCollaborator_Response.Unmarshal(m, b)
}
func (m *RemoveCollaborator_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveCollaborator_Response.Marshal(b, m, deterministic)
}
func (m *RemoveCollaborator_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveCollaborator_Response.Merge(m, src)
}
func (m *RemoveCollaborator_Response) XXX_Size() int {
	return xxx_messageInfo_RemoveCollaborator_Response.Size(m)
}
func (m *RemoveCollaborator_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveCollaborator_Response.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveCollaborator_Response proto.InternalMessageInfo

func (m *RemoveCollaborator_Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *RemoveCollaborator_Response) GetSelfAllowedActions() []*Action {
	if m != nil {
		return m.SelfAllowedActions
	}
	return nil
}

type GetCollaboratorResponse struct {
	UserId               string                                       `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // Deprecated: Do not use.
	CollaboratorType     common.CollaboratorTypeEnum_CollaboratorType `protobuf:"varint,2,opt,name=collaborator_type,json=collaboratorType,proto3,enum=ai.verta.common.CollaboratorTypeEnum_CollaboratorType" json:"collaborator_type,omitempty"`
	ShareViaType         ShareViaEnum                                 `protobuf:"varint,3,opt,name=share_via_type,json=shareViaType,proto3,enum=ai.verta.uac.ShareViaEnum" json:"share_via_type,omitempty"`
	VertaId              string                                       `protobuf:"bytes,4,opt,name=verta_id,json=vertaId,proto3" json:"verta_id,omitempty"`
	CanDeploy            common.TernaryEnum_Ternary                   `protobuf:"varint,5,opt,name=can_deploy,json=canDeploy,proto3,enum=ai.verta.common.TernaryEnum_Ternary" json:"can_deploy,omitempty"`
	AuthzEntityType      EntitiesEnum_EntitiesTypes                   `protobuf:"varint,6,opt,name=authz_entity_type,json=authzEntityType,proto3,enum=ai.verta.uac.EntitiesEnum_EntitiesTypes" json:"authz_entity_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *GetCollaboratorResponse) Reset()         { *m = GetCollaboratorResponse{} }
func (m *GetCollaboratorResponse) String() string { return proto.CompactTextString(m) }
func (*GetCollaboratorResponse) ProtoMessage()    {}
func (*GetCollaboratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{2}
}

func (m *GetCollaboratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollaboratorResponse.Unmarshal(m, b)
}
func (m *GetCollaboratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollaboratorResponse.Marshal(b, m, deterministic)
}
func (m *GetCollaboratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollaboratorResponse.Merge(m, src)
}
func (m *GetCollaboratorResponse) XXX_Size() int {
	return xxx_messageInfo_GetCollaboratorResponse.Size(m)
}
func (m *GetCollaboratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollaboratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollaboratorResponse proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *GetCollaboratorResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GetCollaboratorResponse) GetCollaboratorType() common.CollaboratorTypeEnum_CollaboratorType {
	if m != nil {
		return m.CollaboratorType
	}
	return common.CollaboratorTypeEnum_READ_ONLY
}

func (m *GetCollaboratorResponse) GetShareViaType() ShareViaEnum {
	if m != nil {
		return m.ShareViaType
	}
	return ShareViaEnum_USER_ID
}

func (m *GetCollaboratorResponse) GetVertaId() string {
	if m != nil {
		return m.VertaId
	}
	return ""
}

func (m *GetCollaboratorResponse) GetCanDeploy() common.TernaryEnum_Ternary {
	if m != nil {
		return m.CanDeploy
	}
	return common.TernaryEnum_UNKNOWN
}

func (m *GetCollaboratorResponse) GetAuthzEntityType() EntitiesEnum_EntitiesTypes {
	if m != nil {
		return m.AuthzEntityType
	}
	return EntitiesEnum_UNKNOWN
}

type GetCollaborator struct {
	EntityId             string   `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCollaborator) Reset()         { *m = GetCollaborator{} }
func (m *GetCollaborator) String() string { return proto.CompactTextString(m) }
func (*GetCollaborator) ProtoMessage()    {}
func (*GetCollaborator) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{3}
}

func (m *GetCollaborator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollaborator.Unmarshal(m, b)
}
func (m *GetCollaborator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollaborator.Marshal(b, m, deterministic)
}
func (m *GetCollaborator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollaborator.Merge(m, src)
}
func (m *GetCollaborator) XXX_Size() int {
	return xxx_messageInfo_GetCollaborator.Size(m)
}
func (m *GetCollaborator) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollaborator.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollaborator proto.InternalMessageInfo

func (m *GetCollaborator) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

type GetCollaborator_Response struct {
	SharedUsers          []*GetCollaboratorResponse `protobuf:"bytes,1,rep,name=shared_users,json=sharedUsers,proto3" json:"shared_users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetCollaborator_Response) Reset()         { *m = GetCollaborator_Response{} }
func (m *GetCollaborator_Response) String() string { return proto.CompactTextString(m) }
func (*GetCollaborator_Response) ProtoMessage()    {}
func (*GetCollaborator_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_30132867311f1000, []int{3, 0}
}

func (m *GetCollaborator_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCollaborator_Response.Unmarshal(m, b)
}
func (m *GetCollaborator_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCollaborator_Response.Marshal(b, m, deterministic)
}
func (m *GetCollaborator_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCollaborator_Response.Merge(m, src)
}
func (m *GetCollaborator_Response) XXX_Size() int {
	return xxx_messageInfo_GetCollaborator_Response.Size(m)
}
func (m *GetCollaborator_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCollaborator_Response.DiscardUnknown(m)
}

var xxx_messageInfo_GetCollaborator_Response proto.InternalMessageInfo

func (m *GetCollaborator_Response) GetSharedUsers() []*GetCollaboratorResponse {
	if m != nil {
		return m.SharedUsers
	}
	return nil
}

func init() {
	proto.RegisterEnum("ai.verta.uac.ShareViaEnum", ShareViaEnum_name, ShareViaEnum_value)
	proto.RegisterType((*AddCollaboratorRequest)(nil), "ai.verta.uac.AddCollaboratorRequest")
	proto.RegisterType((*AddCollaboratorRequest_Response)(nil), "ai.verta.uac.AddCollaboratorRequest.Response")
	proto.RegisterType((*RemoveCollaborator)(nil), "ai.verta.uac.RemoveCollaborator")
	proto.RegisterType((*RemoveCollaborator_Response)(nil), "ai.verta.uac.RemoveCollaborator.Response")
	proto.RegisterType((*GetCollaboratorResponse)(nil), "ai.verta.uac.GetCollaboratorResponse")
	proto.RegisterType((*GetCollaborator)(nil), "ai.verta.uac.GetCollaborator")
	proto.RegisterType((*GetCollaborator_Response)(nil), "ai.verta.uac.GetCollaborator.Response")
}

func init() {
	proto.RegisterFile("protos/public/uac/Collaborator.proto", fileDescriptor_30132867311f1000)
}

var fileDescriptor_30132867311f1000 = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x93, 0x36, 0x3f, 0x93, 0xa8, 0xdb, 0x1d, 0xaa, 0xae, 0x9b, 0xdd, 0x45, 0x21, 0x14,
	0x14, 0x22, 0x88, 0x21, 0x20, 0x90, 0x7a, 0x81, 0x48, 0xd3, 0x40, 0x23, 0xb1, 0x3f, 0x9a, 0xa6,
	0x8b, 0xe0, 0xc6, 0x9a, 0xd8, 0xa7, 0x89, 0x57, 0xb6, 0x27, 0x78, 0xc6, 0x59, 0x65, 0x2f, 0x11,
	0x4f, 0xc0, 0xde, 0x72, 0xc3, 0x23, 0x70, 0xc1, 0x93, 0xf0, 0x0a, 0x20, 0xf1, 0x18, 0x68, 0xc6,
	0x4e, 0x6b, 0xc7, 0x6e, 0x17, 0x69, 0xab, 0x8a, 0xab, 0x76, 0xbe, 0x73, 0xe6, 0x9c, 0x6f, 0xe6,
	0xfb, 0xe6, 0x38, 0xe8, 0x60, 0x1e, 0x30, 0xc1, 0xb8, 0x31, 0x0f, 0x27, 0xae, 0x63, 0x19, 0x21,
	0xb5, 0x8c, 0x01, 0x73, 0x5d, 0x3a, 0x61, 0x01, 0x15, 0x2c, 0xe8, 0xaa, 0x30, 0xae, 0x53, 0xa7,
	0xbb, 0x80, 0x40, 0xd0, 0x6e, 0x48, 0xad, 0xc6, 0x83, 0x29, 0x63, 0x53, 0x17, 0x0c, 0x3a, 0x77,
	0x0c, 0xea, 0xfb, 0x4c, 0x50, 0xe1, 0x30, 0x9f, 0x47, 0xb9, 0x8d, 0x56, 0xb6, 0xe2, 0x59, 0x7f,
	0x70, 0x0a, 0xc1, 0xc2, 0xb1, 0x20, 0xce, 0xc9, 0xe9, 0xfa, 0x24, 0x98, 0x52, 0xdf, 0x79, 0xa9,
	0x4a, 0xc5, 0x59, 0x0f, 0xb2, 0x59, 0x63, 0xa0, 0x5e, 0x1c, 0x7d, 0x37, 0x1b, 0x25, 0xcc, 0x85,
	0x74, 0xa3, 0x76, 0x3a, 0xc9, 0x62, 0x9e, 0xc7, 0x7c, 0x63, 0xa0, 0xfe, 0xa4, 0x32, 0x5b, 0xbf,
	0x95, 0xd0, 0x5e, 0xdf, 0xb6, 0x93, 0x87, 0x27, 0xf0, 0x63, 0x08, 0x5c, 0xe0, 0x87, 0x08, 0x81,
	0x2f, 0x1c, 0xb1, 0x34, 0x1d, 0x9b, 0xeb, 0x5a, 0xb3, 0xd8, 0xae, 0x92, 0x6a, 0x84, 0x8c, 0x6c,
	0x2e, 0xc3, 0x7c, 0x46, 0x03, 0x30, 0x5f, 0x38, 0x62, 0xa6, 0x17, 0x9a, 0x9a, 0x0c, 0x2b, 0xe4,
	0x3b, 0x47, 0xcc, 0xb0, 0x85, 0xee, 0x5a, 0x89, 0xa2, 0xa6, 0x58, 0xce, 0x41, 0x2f, 0x36, 0xb5,
	0xf6, 0x76, 0xef, 0xf3, 0xee, 0xc5, 0xbd, 0x46, 0xcc, 0xba, 0xc9, 0xf6, 0xe3, 0xe5, 0x1c, 0x86,
	0x7e, 0xe8, 0x65, 0x40, 0xb2, 0x63, 0xad, 0x21, 0x58, 0x47, 0x65, 0x0f, 0x38, 0xa7, 0x53, 0xd0,
	0x37, 0x15, 0x81, 0xd5, 0x12, 0xbf, 0x83, 0xea, 0x36, 0x15, 0x60, 0x5a, 0x01, 0x50, 0x01, 0xb6,
	0xbe, 0xd5, 0xd4, 0xda, 0x9b, 0xa4, 0x26, 0xb1, 0x41, 0x04, 0x5d, 0xa4, 0x84, 0x73, 0x5b, 0xa5,
	0x94, 0x2e, 0x53, 0xce, 0x22, 0x08, 0x0f, 0x10, 0xb2, 0xa8, 0x6f, 0xda, 0x30, 0x77, 0xd9, 0x52,
	0x2f, 0x2b, 0xf6, 0x07, 0x19, 0xf6, 0x63, 0x08, 0x7c, 0x1a, 0x2c, 0x15, 0xe9, 0xf8, 0x7f, 0x52,
	0xb5, 0xa8, 0x7f, 0xac, 0xb6, 0xe1, 0x31, 0xba, 0x4b, 0x43, 0x31, 0x7b, 0x69, 0xc6, 0xb7, 0xa9,
	0x6e, 0xa2, 0xa2, 0x6a, 0xb5, 0xbb, 0x49, 0x87, 0x75, 0x87, 0x32, 0xc1, 0x01, 0xae, 0x2a, 0xad,
	0x16, 0xf2, 0xa0, 0x9c, 0xdc, 0x51, 0x25, 0x14, 0xb6, 0x94, 0x48, 0xe3, 0x9f, 0x02, 0xaa, 0x10,
	0xe0, 0x73, 0xe6, 0x73, 0xc0, 0x5f, 0xa3, 0x5d, 0x0e, 0xee, 0xb9, 0x49, 0x5d, 0x97, 0xbd, 0x00,
	0xdb, 0xa4, 0x96, 0xb2, 0xa6, 0xbe, 0xd5, 0x2c, 0xb6, 0x6b, 0xbd, 0xdd, 0x74, 0x97, 0xbe, 0x0a,
	0x12, 0x2c, 0x77, 0xf4, 0xa3, 0x0d, 0x11, 0xc4, 0xf1, 0x1e, 0x2a, 0x71, 0x41, 0x45, 0x28, 0xe5,
	0xd6, 0xda, 0x15, 0x12, 0xaf, 0xf0, 0x63, 0xb4, 0x97, 0x12, 0x33, 0xe4, 0x10, 0x98, 0x8e, 0x7f,
	0xce, 0x94, 0xee, 0xb5, 0xde, 0x5e, 0xba, 0xc3, 0x19, 0x87, 0x60, 0xe4, 0x9f, 0xb3, 0x93, 0x0d,
	0xb2, 0x9b, 0xdc, 0xb7, 0xc2, 0xf1, 0xf7, 0x68, 0x3f, 0x55, 0x8f, 0x25, 0x5e, 0x81, 0x32, 0x49,
	0xad, 0xd7, 0x48, 0x97, 0x4c, 0xbe, 0x93, 0x93, 0x0d, 0xa2, 0x27, 0xb7, 0x27, 0x63, 0xb8, 0xbf,
	0xee, 0x3b, 0xa0, 0x9e, 0x32, 0x47, 0xad, 0x87, 0xd3, 0x25, 0xe5, 0xa3, 0x3a, 0xd9, 0x58, 0x73,
	0x15, 0x50, 0xef, 0x68, 0x1b, 0xd5, 0x93, 0x58, 0xeb, 0x8f, 0x02, 0xc2, 0x04, 0x3c, 0xb6, 0x80,
	0xa4, 0x25, 0xf1, 0x7d, 0x54, 0xbd, 0x78, 0x1f, 0xea, 0xbe, 0xaa, 0xa4, 0xb2, 0x7a, 0x1e, 0xaf,
	0x7b, 0x1d, 0x2b, 0xef, 0xd9, 0xe0, 0x82, 0xf4, 0x5e, 0xf1, 0xd2, 0x7b, 0xc7, 0x11, 0x94, 0x6f,
	0x9b, 0xcd, 0x37, 0xb5, 0xcd, 0xf3, 0x84, 0x6b, 0xae, 0x52, 0xfb, 0x86, 0xdc, 0xd4, 0x7a, 0x55,
	0x44, 0xf7, 0xbe, 0x01, 0x91, 0x9e, 0x2d, 0x71, 0xef, 0xfb, 0xa8, 0x1c, 0x99, 0x28, 0xbe, 0xba,
	0xa3, 0x82, 0xae, 0x91, 0x92, 0x84, 0x46, 0x76, 0xfe, 0xec, 0x28, 0xdc, 0xf0, 0xec, 0xf8, 0x0a,
	0x6d, 0x47, 0x0a, 0x2d, 0x1c, 0x9a, 0x9c, 0x4e, 0x6b, 0xc6, 0x3b, 0x95, 0x39, 0xcf, 0x1c, 0x2a,
	0xcb, 0x92, 0x3a, 0x8f, 0x57, 0xaa, 0xc2, 0x3e, 0xaa, 0xa8, 0x3c, 0x79, 0x88, 0x78, 0xfc, 0xa8,
	0xf5, 0x68, 0x7d, 0x70, 0x6c, 0xdd, 0xe0, 0xe0, 0x28, 0xbd, 0xa1, 0x03, 0x5a, 0x3f, 0x6b, 0xe8,
	0xce, 0x9a, 0x2a, 0xd7, 0x5a, 0xb9, 0x31, 0x4e, 0x58, 0xe6, 0x04, 0x45, 0x57, 0x60, 0xab, 0x11,
	0x10, 0x7d, 0x15, 0x6a, 0xbd, 0xf7, 0xd2, 0x6c, 0xae, 0xd0, 0x9c, 0xd4, 0xa2, 0xad, 0x72, 0x0a,
	0xf0, 0xce, 0x17, 0xa8, 0x9e, 0xbc, 0x5a, 0x5c, 0x43, 0xe5, 0xb3, 0xd3, 0x21, 0x31, 0x47, 0xc7,
	0x3b, 0x1b, 0xb8, 0x8e, 0x2a, 0xc3, 0x47, 0xfd, 0xd1, 0xb7, 0x72, 0xa5, 0xc9, 0x95, 0x0c, 0x3d,
	0xee, 0x3f, 0x1a, 0xee, 0x14, 0x7a, 0x7f, 0x97, 0xd1, 0x5b, 0xc9, 0xf2, 0xf1, 0xf7, 0x0c, 0xff,
	0xae, 0xa1, 0xb7, 0xa9, 0x6d, 0x3f, 0x09, 0xa2, 0xe1, 0xfd, 0x34, 0x60, 0xcf, 0xc1, 0x4a, 0x1f,
	0xf3, 0x60, 0xcd, 0xba, 0xb9, 0xdf, 0xbd, 0xc6, 0x47, 0xff, 0x25, 0xab, 0xbb, 0x3a, 0x55, 0xeb,
	0xf0, 0xa7, 0x3f, 0xff, 0x7a, 0x55, 0xf8, 0xac, 0x65, 0x18, 0x8b, 0x4f, 0x8c, 0xa4, 0xcb, 0x8c,
	0xeb, 0xd9, 0x1c, 0x6a, 0x1d, 0xfc, 0xab, 0x86, 0xf6, 0x03, 0x35, 0x59, 0xf2, 0xe8, 0x36, 0xd3,
	0x44, 0xb2, 0x23, 0xa8, 0xf1, 0xc1, 0xeb, 0x32, 0x2e, 0x69, 0xf6, 0x14, 0xcd, 0x0f, 0x3b, 0x9d,
	0x0c, 0xcd, 0xab, 0x09, 0xfc, 0xa2, 0xa1, 0x7b, 0x53, 0x10, 0x39, 0x21, 0x8e, 0x1f, 0x5e, 0xab,
	0x79, 0xe3, 0xfd, 0x6b, 0xc3, 0x97, 0xb4, 0x3e, 0x56, 0xb4, 0x3a, 0xb8, 0x9d, 0xa1, 0x75, 0x55,
	0xe3, 0x35, 0x9d, 0x8f, 0xa9, 0xa0, 0x1c, 0xfe, 0x2f, 0x3a, 0xe7, 0xb0, 0x49, 0xeb, 0x9c, 0x47,
	0xf7, 0x56, 0x75, 0xce, 0x23, 0x10, 0xeb, 0x9c, 0x13, 0xba, 0x1d, 0x9d, 0xf3, 0x1a, 0x1f, 0x7d,
	0xf9, 0x54, 0xfb, 0xe1, 0x70, 0xea, 0x88, 0x59, 0x38, 0x91, 0xf3, 0xd2, 0x78, 0x26, 0x9b, 0xf4,
	0x47, 0x86, 0xc7, 0x6c, 0x70, 0xed, 0x89, 0x11, 0xff, 0xc4, 0x9d, 0x82, 0x6f, 0x4c, 0x99, 0x91,
	0xf9, 0x55, 0x3c, 0x29, 0x29, 0xe8, 0xd3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab, 0x46, 0x12,
	0xe5, 0xeb, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollaboratorServiceClient is the client API for CollaboratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollaboratorServiceClient interface {
	AddOrUpdateProjectCollaborator(ctx context.Context, in *AddCollaboratorRequest, opts ...grpc.CallOption) (*AddCollaboratorRequest_Response, error)
	RemoveProjectCollaborator(ctx context.Context, in *RemoveCollaborator, opts ...grpc.CallOption) (*RemoveCollaborator_Response, error)
	GetProjectCollaborators(ctx context.Context, in *GetCollaborator, opts ...grpc.CallOption) (*GetCollaborator_Response, error)
	AddOrUpdateDatasetCollaborator(ctx context.Context, in *AddCollaboratorRequest, opts ...grpc.CallOption) (*AddCollaboratorRequest_Response, error)
	RemoveDatasetCollaborator(ctx context.Context, in *RemoveCollaborator, opts ...grpc.CallOption) (*RemoveCollaborator_Response, error)
	GetDatasetCollaborators(ctx context.Context, in *GetCollaborator, opts ...grpc.CallOption) (*GetCollaborator_Response, error)
}

type collaboratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollaboratorServiceClient(cc grpc.ClientConnInterface) CollaboratorServiceClient {
	return &collaboratorServiceClient{cc}
}

func (c *collaboratorServiceClient) AddOrUpdateProjectCollaborator(ctx context.Context, in *AddCollaboratorRequest, opts ...grpc.CallOption) (*AddCollaboratorRequest_Response, error) {
	out := new(AddCollaboratorRequest_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.CollaboratorService/addOrUpdateProjectCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaboratorServiceClient) RemoveProjectCollaborator(ctx context.Context, in *RemoveCollaborator, opts ...grpc.CallOption) (*RemoveCollaborator_Response, error) {
	out := new(RemoveCollaborator_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.CollaboratorService/removeProjectCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaboratorServiceClient) GetProjectCollaborators(ctx context.Context, in *GetCollaborator, opts ...grpc.CallOption) (*GetCollaborator_Response, error) {
	out := new(GetCollaborator_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.CollaboratorService/getProjectCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaboratorServiceClient) AddOrUpdateDatasetCollaborator(ctx context.Context, in *AddCollaboratorRequest, opts ...grpc.CallOption) (*AddCollaboratorRequest_Response, error) {
	out := new(AddCollaboratorRequest_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.CollaboratorService/addOrUpdateDatasetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaboratorServiceClient) RemoveDatasetCollaborator(ctx context.Context, in *RemoveCollaborator, opts ...grpc.CallOption) (*RemoveCollaborator_Response, error) {
	out := new(RemoveCollaborator_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.CollaboratorService/removeDatasetCollaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaboratorServiceClient) GetDatasetCollaborators(ctx context.Context, in *GetCollaborator, opts ...grpc.CallOption) (*GetCollaborator_Response, error) {
	out := new(GetCollaborator_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.CollaboratorService/getDatasetCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollaboratorServiceServer is the server API for CollaboratorService service.
type CollaboratorServiceServer interface {
	AddOrUpdateProjectCollaborator(context.Context, *AddCollaboratorRequest) (*AddCollaboratorRequest_Response, error)
	RemoveProjectCollaborator(context.Context, *RemoveCollaborator) (*RemoveCollaborator_Response, error)
	GetProjectCollaborators(context.Context, *GetCollaborator) (*GetCollaborator_Response, error)
	AddOrUpdateDatasetCollaborator(context.Context, *AddCollaboratorRequest) (*AddCollaboratorRequest_Response, error)
	RemoveDatasetCollaborator(context.Context, *RemoveCollaborator) (*RemoveCollaborator_Response, error)
	GetDatasetCollaborators(context.Context, *GetCollaborator) (*GetCollaborator_Response, error)
}

// UnimplementedCollaboratorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCollaboratorServiceServer struct {
}

func (*UnimplementedCollaboratorServiceServer) AddOrUpdateProjectCollaborator(ctx context.Context, req *AddCollaboratorRequest) (*AddCollaboratorRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateProjectCollaborator not implemented")
}
func (*UnimplementedCollaboratorServiceServer) RemoveProjectCollaborator(ctx context.Context, req *RemoveCollaborator) (*RemoveCollaborator_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProjectCollaborator not implemented")
}
func (*UnimplementedCollaboratorServiceServer) GetProjectCollaborators(ctx context.Context, req *GetCollaborator) (*GetCollaborator_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectCollaborators not implemented")
}
func (*UnimplementedCollaboratorServiceServer) AddOrUpdateDatasetCollaborator(ctx context.Context, req *AddCollaboratorRequest) (*AddCollaboratorRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateDatasetCollaborator not implemented")
}
func (*UnimplementedCollaboratorServiceServer) RemoveDatasetCollaborator(ctx context.Context, req *RemoveCollaborator) (*RemoveCollaborator_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDatasetCollaborator not implemented")
}
func (*UnimplementedCollaboratorServiceServer) GetDatasetCollaborators(ctx context.Context, req *GetCollaborator) (*GetCollaborator_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetCollaborators not implemented")
}

func RegisterCollaboratorServiceServer(s *grpc.Server, srv CollaboratorServiceServer) {
	s.RegisterService(&_CollaboratorService_serviceDesc, srv)
}

func _CollaboratorService_AddOrUpdateProjectCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaboratorServiceServer).AddOrUpdateProjectCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.CollaboratorService/AddOrUpdateProjectCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaboratorServiceServer).AddOrUpdateProjectCollaborator(ctx, req.(*AddCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaboratorService_RemoveProjectCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCollaborator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaboratorServiceServer).RemoveProjectCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.CollaboratorService/RemoveProjectCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaboratorServiceServer).RemoveProjectCollaborator(ctx, req.(*RemoveCollaborator))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaboratorService_GetProjectCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollaborator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaboratorServiceServer).GetProjectCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.CollaboratorService/GetProjectCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaboratorServiceServer).GetProjectCollaborators(ctx, req.(*GetCollaborator))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaboratorService_AddOrUpdateDatasetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaboratorServiceServer).AddOrUpdateDatasetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.CollaboratorService/AddOrUpdateDatasetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaboratorServiceServer).AddOrUpdateDatasetCollaborator(ctx, req.(*AddCollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaboratorService_RemoveDatasetCollaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCollaborator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaboratorServiceServer).RemoveDatasetCollaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.CollaboratorService/RemoveDatasetCollaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaboratorServiceServer).RemoveDatasetCollaborator(ctx, req.(*RemoveCollaborator))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaboratorService_GetDatasetCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollaborator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaboratorServiceServer).GetDatasetCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.CollaboratorService/GetDatasetCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaboratorServiceServer).GetDatasetCollaborators(ctx, req.(*GetCollaborator))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollaboratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.verta.uac.CollaboratorService",
	HandlerType: (*CollaboratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addOrUpdateProjectCollaborator",
			Handler:    _CollaboratorService_AddOrUpdateProjectCollaborator_Handler,
		},
		{
			MethodName: "removeProjectCollaborator",
			Handler:    _CollaboratorService_RemoveProjectCollaborator_Handler,
		},
		{
			MethodName: "getProjectCollaborators",
			Handler:    _CollaboratorService_GetProjectCollaborators_Handler,
		},
		{
			MethodName: "addOrUpdateDatasetCollaborator",
			Handler:    _CollaboratorService_AddOrUpdateDatasetCollaborator_Handler,
		},
		{
			MethodName: "removeDatasetCollaborator",
			Handler:    _CollaboratorService_RemoveDatasetCollaborator_Handler,
		},
		{
			MethodName: "getDatasetCollaborators",
			Handler:    _CollaboratorService_GetDatasetCollaborators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/public/uac/Collaborator.proto",
}

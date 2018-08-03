// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/user/service.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import track "user-api/rpc/track"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                     string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username               string               `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email                  string               `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	DisplayName            string               `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	FullName               string               `protobuf:"bytes,5,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	FirstName              string               `protobuf:"bytes,6,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName               string               `protobuf:"bytes,7,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Member                 bool                 `protobuf:"varint,8,opt,name=member" json:"member,omitempty"`
	Avatar                 []byte               `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	NewsletterNotification bool                 `protobuf:"varint,10,opt,name=newsletter_notification,json=newsletterNotification" json:"newsletter_notification,omitempty"`
	FavoriteTracks         []string             `protobuf:"bytes,11,rep,name=favorite_tracks,json=favoriteTracks" json:"favorite_tracks,omitempty"`
	FollowedGroups         []string             `protobuf:"bytes,12,rep,name=followed_groups,json=followedGroups" json:"followed_groups,omitempty"`
	OwnerOfGroups          []*UserGroupResponse `protobuf:"bytes,13,rep,name=owner_of_groups,json=ownerOfGroups" json:"owner_of_groups,omitempty"`
	ResidenceAddress       *StreetAddress       `protobuf:"bytes,14,opt,name=residence_address,json=residenceAddress" json:"residence_address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_661e9044d01207fe, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetMember() bool {
	if m != nil {
		return m.Member
	}
	return false
}

func (m *User) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *User) GetNewsletterNotification() bool {
	if m != nil {
		return m.NewsletterNotification
	}
	return false
}

func (m *User) GetFavoriteTracks() []string {
	if m != nil {
		return m.FavoriteTracks
	}
	return nil
}

func (m *User) GetFollowedGroups() []string {
	if m != nil {
		return m.FollowedGroups
	}
	return nil
}

func (m *User) GetOwnerOfGroups() []*UserGroupResponse {
	if m != nil {
		return m.OwnerOfGroups
	}
	return nil
}

func (m *User) GetResidenceAddress() *StreetAddress {
	if m != nil {
		return m.ResidenceAddress
	}
	return nil
}

type Playlists struct {
	Playlists            []*track.RelatedTrackGroup `protobuf:"bytes,1,rep,name=playlists" json:"playlists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Playlists) Reset()         { *m = Playlists{} }
func (m *Playlists) String() string { return proto.CompactTextString(m) }
func (*Playlists) ProtoMessage()    {}
func (*Playlists) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_661e9044d01207fe, []int{1}
}
func (m *Playlists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Playlists.Unmarshal(m, b)
}
func (m *Playlists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Playlists.Marshal(b, m, deterministic)
}
func (dst *Playlists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Playlists.Merge(dst, src)
}
func (m *Playlists) XXX_Size() int {
	return xxx_messageInfo_Playlists.Size(m)
}
func (m *Playlists) XXX_DiscardUnknown() {
	xxx_messageInfo_Playlists.DiscardUnknown(m)
}

var xxx_messageInfo_Playlists proto.InternalMessageInfo

func (m *Playlists) GetPlaylists() []*track.RelatedTrackGroup {
	if m != nil {
		return m.Playlists
	}
	return nil
}

type StreetAddress struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PersonalData         bool              `protobuf:"varint,3,opt,name=personal_data,json=personalData" json:"personal_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreetAddress) Reset()         { *m = StreetAddress{} }
func (m *StreetAddress) String() string { return proto.CompactTextString(m) }
func (*StreetAddress) ProtoMessage()    {}
func (*StreetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_661e9044d01207fe, []int{2}
}
func (m *StreetAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreetAddress.Unmarshal(m, b)
}
func (m *StreetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreetAddress.Marshal(b, m, deterministic)
}
func (dst *StreetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreetAddress.Merge(dst, src)
}
func (m *StreetAddress) XXX_Size() int {
	return xxx_messageInfo_StreetAddress.Size(m)
}
func (m *StreetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_StreetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_StreetAddress proto.InternalMessageInfo

func (m *StreetAddress) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreetAddress) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StreetAddress) GetPersonalData() bool {
	if m != nil {
		return m.PersonalData
	}
	return false
}

type UserGroupResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Avatar               []byte   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGroupResponse) Reset()         { *m = UserGroupResponse{} }
func (m *UserGroupResponse) String() string { return proto.CompactTextString(m) }
func (*UserGroupResponse) ProtoMessage()    {}
func (*UserGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_661e9044d01207fe, []int{3}
}
func (m *UserGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroupResponse.Unmarshal(m, b)
}
func (m *UserGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroupResponse.Marshal(b, m, deterministic)
}
func (dst *UserGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupResponse.Merge(dst, src)
}
func (m *UserGroupResponse) XXX_Size() int {
	return xxx_messageInfo_UserGroupResponse.Size(m)
}
func (m *UserGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupResponse proto.InternalMessageInfo

func (m *UserGroupResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserGroupResponse) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *UserGroupResponse) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

// Used for:
// - connecting to user group
// - follow/unfollow group
type UserToUserGroup struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserGroupId          string   `protobuf:"bytes,2,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToUserGroup) Reset()         { *m = UserToUserGroup{} }
func (m *UserToUserGroup) String() string { return proto.CompactTextString(m) }
func (*UserToUserGroup) ProtoMessage()    {}
func (*UserToUserGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_661e9044d01207fe, []int{4}
}
func (m *UserToUserGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToUserGroup.Unmarshal(m, b)
}
func (m *UserToUserGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToUserGroup.Marshal(b, m, deterministic)
}
func (dst *UserToUserGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToUserGroup.Merge(dst, src)
}
func (m *UserToUserGroup) XXX_Size() int {
	return xxx_messageInfo_UserToUserGroup.Size(m)
}
func (m *UserToUserGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToUserGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UserToUserGroup proto.InternalMessageInfo

func (m *UserToUserGroup) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserToUserGroup) GetUserGroupId() string {
	if m != nil {
		return m.UserGroupId
	}
	return ""
}

// Used for:
// - add/remove favorite tracks
type UserToTrack struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	TrackId              string   `protobuf:"bytes,2,opt,name=track_id,json=trackId" json:"track_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToTrack) Reset()         { *m = UserToTrack{} }
func (m *UserToTrack) String() string { return proto.CompactTextString(m) }
func (*UserToTrack) ProtoMessage()    {}
func (*UserToTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_661e9044d01207fe, []int{5}
}
func (m *UserToTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToTrack.Unmarshal(m, b)
}
func (m *UserToTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToTrack.Marshal(b, m, deterministic)
}
func (dst *UserToTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToTrack.Merge(dst, src)
}
func (m *UserToTrack) XXX_Size() int {
	return xxx_messageInfo_UserToTrack.Size(m)
}
func (m *UserToTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToTrack.DiscardUnknown(m)
}

var xxx_messageInfo_UserToTrack proto.InternalMessageInfo

func (m *UserToTrack) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserToTrack) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "resonate.api.user.User")
	proto.RegisterType((*Playlists)(nil), "resonate.api.user.Playlists")
	proto.RegisterType((*StreetAddress)(nil), "resonate.api.user.StreetAddress")
	proto.RegisterMapType((map[string]string)(nil), "resonate.api.user.StreetAddress.DataEntry")
	proto.RegisterType((*UserGroupResponse)(nil), "resonate.api.user.UserGroupResponse")
	proto.RegisterType((*UserToUserGroup)(nil), "resonate.api.user.UserToUserGroup")
	proto.RegisterType((*UserToTrack)(nil), "resonate.api.user.UserToTrack")
}

func init() { proto.RegisterFile("rpc/user/service.proto", fileDescriptor_service_661e9044d01207fe) }

var fileDescriptor_service_661e9044d01207fe = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6a, 0xdb, 0x4a,
	0x10, 0x46, 0xb6, 0x63, 0x5b, 0x63, 0x3b, 0x3f, 0x7b, 0x0e, 0x89, 0x8e, 0xcf, 0x0f, 0x3a, 0x3a,
	0x07, 0x6a, 0x0a, 0x75, 0x20, 0xbd, 0x48, 0x69, 0x21, 0x25, 0x69, 0x7e, 0x08, 0xcd, 0x4f, 0x51,
	0x92, 0x9b, 0x5e, 0x54, 0x6c, 0xbc, 0xe3, 0x22, 0x22, 0x4b, 0x62, 0x77, 0xed, 0xe0, 0x97, 0xe8,
	0x13, 0xf4, 0x55, 0xfa, 0x6e, 0x65, 0x47, 0x92, 0xed, 0xd4, 0x71, 0x02, 0x21, 0x37, 0xc6, 0xf3,
	0x7d, 0xf3, 0x7d, 0xb3, 0xbb, 0x33, 0xbb, 0x82, 0x75, 0x99, 0xf6, 0x36, 0x87, 0x0a, 0xe5, 0xa6,
	0x42, 0x39, 0x0a, 0x7b, 0xd8, 0x4d, 0x65, 0xa2, 0x13, 0xb6, 0x26, 0x51, 0x25, 0x31, 0xd7, 0xd8,
	0xe5, 0x69, 0xd8, 0x35, 0x09, 0x6d, 0xd7, 0xfc, 0xbe, 0xe2, 0x69, 0xb8, 0x69, 0x34, 0x5a, 0xf2,
	0xde, 0xcd, 0x5d, 0x91, 0xf7, 0xbd, 0x02, 0x95, 0x2b, 0x85, 0x92, 0x2d, 0x43, 0x29, 0x14, 0x8e,
	0xe5, 0x5a, 0x1d, 0xdb, 0x2f, 0x85, 0x82, 0xb5, 0xa1, 0x6e, 0xc4, 0x31, 0x1f, 0xa0, 0x53, 0x22,
	0x74, 0x12, 0xb3, 0xdf, 0x61, 0x09, 0x07, 0x3c, 0x8c, 0x9c, 0x32, 0x11, 0x59, 0xc0, 0xfe, 0x85,
	0xa6, 0x08, 0x55, 0x1a, 0xf1, 0x71, 0x40, 0xaa, 0x0a, 0x91, 0x8d, 0x1c, 0x3b, 0x33, 0xc2, 0x3f,
	0xc1, 0xee, 0x0f, 0xa3, 0x28, 0xe3, 0x97, 0x32, 0x57, 0x03, 0x10, 0xf9, 0x37, 0x40, 0x3f, 0x94,
	0x4a, 0x67, 0x6c, 0x95, 0x58, 0x9b, 0x90, 0x42, 0x1b, 0xf1, 0x82, 0xad, 0x65, 0x5a, 0x03, 0x10,
	0xb9, 0x0e, 0xd5, 0x01, 0x0e, 0xae, 0x51, 0x3a, 0x75, 0xd7, 0xea, 0xd4, 0xfd, 0x3c, 0x32, 0x38,
	0x1f, 0x71, 0xcd, 0xa5, 0x63, 0xbb, 0x56, 0xa7, 0xe9, 0xe7, 0x11, 0xdb, 0x86, 0x8d, 0x18, 0x6f,
	0x55, 0x84, 0x5a, 0xa3, 0x0c, 0xe2, 0x44, 0x87, 0xfd, 0xb0, 0xc7, 0x75, 0x98, 0xc4, 0x0e, 0x90,
	0xc1, 0xfa, 0x94, 0x3e, 0x9b, 0x61, 0xd9, 0x0b, 0x58, 0xe9, 0xf3, 0x51, 0x22, 0x43, 0x8d, 0x01,
	0x9d, 0xa7, 0x72, 0x1a, 0x6e, 0xb9, 0x63, 0xfb, 0xcb, 0x05, 0x7c, 0x49, 0x28, 0x25, 0x26, 0x51,
	0x94, 0xdc, 0xa2, 0x08, 0xbe, 0xca, 0x64, 0x98, 0x2a, 0xa7, 0x99, 0x27, 0xe6, 0xf0, 0x11, 0xa1,
	0xec, 0x04, 0x56, 0x92, 0xdb, 0x18, 0x65, 0x90, 0xf4, 0x8b, 0xc4, 0x96, 0x5b, 0xee, 0x34, 0xb6,
	0xfe, 0xef, 0xce, 0x35, 0xb4, 0x6b, 0x5a, 0x45, 0x3a, 0x1f, 0x55, 0x9a, 0xc4, 0x0a, 0xfd, 0x16,
	0x89, 0xcf, 0xfb, 0xb9, 0xdb, 0x29, 0x98, 0x31, 0x08, 0x05, 0xc6, 0x3d, 0x0c, 0xb8, 0x10, 0x12,
	0x95, 0x72, 0x96, 0x5d, 0xab, 0xd3, 0xd8, 0x72, 0xef, 0xf1, 0xbb, 0xd0, 0x12, 0x51, 0xef, 0x66,
	0x79, 0xfe, 0xea, 0x44, 0x9a, 0x23, 0xde, 0x39, 0xd8, 0x9f, 0x22, 0x3e, 0x8e, 0x42, 0xa5, 0x15,
	0xdb, 0x03, 0x3b, 0x2d, 0x02, 0xc7, 0x5a, 0xb8, 0x46, 0x1f, 0x23, 0xae, 0x51, 0xd0, 0x39, 0x64,
	0x6b, 0x9d, 0xca, 0xbc, 0x1f, 0x16, 0xb4, 0xee, 0x14, 0x9d, 0x1b, 0xbc, 0x1d, 0xa8, 0x08, 0xae,
	0xb9, 0x53, 0xa2, 0x02, 0x2f, 0x1f, 0x5b, 0x74, 0x77, 0x9f, 0x6b, 0x7e, 0x10, 0x6b, 0x39, 0xf6,
	0x49, 0xc7, 0xfe, 0x83, 0x56, 0x8a, 0xd2, 0x68, 0xa2, 0x80, 0x8c, 0xca, 0xd4, 0xd0, 0x66, 0x01,
	0x9a, 0xfc, 0xf6, 0x36, 0xd8, 0x13, 0x1d, 0x5b, 0x85, 0xf2, 0x0d, 0x8e, 0xf3, 0x25, 0x98, 0xbf,
	0x66, 0xc0, 0x47, 0x3c, 0x1a, 0x16, 0x93, 0x9f, 0x05, 0x6f, 0x4b, 0x6f, 0x2c, 0xef, 0x0b, 0xac,
	0xcd, 0xf5, 0x60, 0x6e, 0x0b, 0xbf, 0xde, 0x84, 0xd2, 0xfc, 0x4d, 0x98, 0x0e, 0x66, 0x79, 0x76,
	0x30, 0xbd, 0x33, 0x58, 0x31, 0xfe, 0x97, 0xc9, 0xa4, 0x0a, 0xdb, 0x80, 0x9a, 0xd9, 0x76, 0x30,
	0x29, 0x51, 0x35, 0xe1, 0xb1, 0x60, 0x1e, 0xb4, 0x88, 0xa0, 0xa9, 0x31, 0x74, 0x5e, 0x67, 0x58,
	0x48, 0x8f, 0x85, 0xb7, 0x0b, 0x8d, 0xcc, 0x8f, 0xda, 0xb1, 0xd8, 0xeb, 0x0f, 0xa8, 0xd3, 0x38,
	0x4f, 0x6d, 0x6a, 0x14, 0x1f, 0x8b, 0xad, 0x6f, 0x4b, 0x99, 0xc7, 0x45, 0xf6, 0x70, 0xb0, 0x1d,
	0x80, 0x0f, 0x12, 0xb9, 0x46, 0x7a, 0x37, 0x36, 0x16, 0x4c, 0x69, 0x7b, 0x11, 0xc1, 0xde, 0x41,
	0xed, 0x08, 0xf5, 0x13, 0xc5, 0xef, 0x01, 0xae, 0x52, 0xf1, 0x68, 0x71, 0xe7, 0x1e, 0xe2, 0x60,
	0x90, 0xea, 0xb1, 0x31, 0xd8, 0xc7, 0x08, 0x9f, 0x6e, 0xf0, 0x11, 0x1a, 0x87, 0x74, 0x83, 0xb3,
	0xee, 0x78, 0x0b, 0x1c, 0x66, 0x3a, 0xf8, 0x80, 0xd9, 0x29, 0xb4, 0xae, 0xe2, 0xfe, 0xb3, 0xd9,
	0x9d, 0xc0, 0xea, 0xae, 0x10, 0x87, 0xb3, 0x2f, 0x11, 0xfb, 0x67, 0xa1, 0x23, 0xf1, 0x0f, 0xb8,
	0x9d, 0xc3, 0x6f, 0x3e, 0x0e, 0x92, 0x11, 0x3e, 0x97, 0xe1, 0x01, 0x34, 0x8f, 0x50, 0x4f, 0x1f,
	0x94, 0x85, 0xa7, 0xff, 0xd7, 0x3d, 0xc4, 0x44, 0xb6, 0x57, 0xfd, 0x5c, 0x31, 0xc8, 0x75, 0x95,
	0x3e, 0x61, 0xaf, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x66, 0xaa, 0x45, 0x53, 0x11, 0x07, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_notifications.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PushNotificationRegistration_TokenType int32

const (
	PushNotificationRegistration_UNKNOWN_TOKEN_TYPE PushNotificationRegistration_TokenType = 0
	PushNotificationRegistration_APN_TOKEN          PushNotificationRegistration_TokenType = 1
	PushNotificationRegistration_FIREBASE_TOKEN     PushNotificationRegistration_TokenType = 2
)

var PushNotificationRegistration_TokenType_name = map[int32]string{
	0: "UNKNOWN_TOKEN_TYPE",
	1: "APN_TOKEN",
	2: "FIREBASE_TOKEN",
}

var PushNotificationRegistration_TokenType_value = map[string]int32{
	"UNKNOWN_TOKEN_TYPE": 0,
	"APN_TOKEN":          1,
	"FIREBASE_TOKEN":     2,
}

func (x PushNotificationRegistration_TokenType) String() string {
	return proto.EnumName(PushNotificationRegistration_TokenType_name, int32(x))
}

func (PushNotificationRegistration_TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{0, 0}
}

type PushNotificationRegistrationResponse_ErrorType int32

const (
	PushNotificationRegistrationResponse_UNKNOWN_ERROR_TYPE     PushNotificationRegistrationResponse_ErrorType = 0
	PushNotificationRegistrationResponse_MALFORMED_MESSAGE      PushNotificationRegistrationResponse_ErrorType = 1
	PushNotificationRegistrationResponse_VERSION_MISMATCH       PushNotificationRegistrationResponse_ErrorType = 2
	PushNotificationRegistrationResponse_UNSUPPORTED_TOKEN_TYPE PushNotificationRegistrationResponse_ErrorType = 3
	PushNotificationRegistrationResponse_INTERNAL_ERROR         PushNotificationRegistrationResponse_ErrorType = 4
)

var PushNotificationRegistrationResponse_ErrorType_name = map[int32]string{
	0: "UNKNOWN_ERROR_TYPE",
	1: "MALFORMED_MESSAGE",
	2: "VERSION_MISMATCH",
	3: "UNSUPPORTED_TOKEN_TYPE",
	4: "INTERNAL_ERROR",
}

var PushNotificationRegistrationResponse_ErrorType_value = map[string]int32{
	"UNKNOWN_ERROR_TYPE":     0,
	"MALFORMED_MESSAGE":      1,
	"VERSION_MISMATCH":       2,
	"UNSUPPORTED_TOKEN_TYPE": 3,
	"INTERNAL_ERROR":         4,
}

func (x PushNotificationRegistrationResponse_ErrorType) String() string {
	return proto.EnumName(PushNotificationRegistrationResponse_ErrorType_name, int32(x))
}

func (PushNotificationRegistrationResponse_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{1, 0}
}

type PushNotificationReport_ErrorType int32

const (
	PushNotificationReport_UNKNOWN_ERROR_TYPE PushNotificationReport_ErrorType = 0
	PushNotificationReport_WRONG_TOKEN        PushNotificationReport_ErrorType = 1
	PushNotificationReport_INTERNAL_ERROR     PushNotificationReport_ErrorType = 2
	PushNotificationReport_NOT_REGISTERED     PushNotificationReport_ErrorType = 3
)

var PushNotificationReport_ErrorType_name = map[int32]string{
	0: "UNKNOWN_ERROR_TYPE",
	1: "WRONG_TOKEN",
	2: "INTERNAL_ERROR",
	3: "NOT_REGISTERED",
}

var PushNotificationReport_ErrorType_value = map[string]int32{
	"UNKNOWN_ERROR_TYPE": 0,
	"WRONG_TOKEN":        1,
	"INTERNAL_ERROR":     2,
	"NOT_REGISTERED":     3,
}

func (x PushNotificationReport_ErrorType) String() string {
	return proto.EnumName(PushNotificationReport_ErrorType_name, int32(x))
}

func (PushNotificationReport_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{9, 0}
}

type PushNotificationRegistration struct {
	TokenType            PushNotificationRegistration_TokenType `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=protobuf.PushNotificationRegistration_TokenType" json:"token_type,omitempty"`
	Token                string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	InstallationId       string                                 `protobuf:"bytes,3,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	AccessToken          string                                 `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Enabled              bool                                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Version              uint64                                 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	AllowedUserList      [][]byte                               `protobuf:"bytes,7,rep,name=allowed_user_list,json=allowedUserList,proto3" json:"allowed_user_list,omitempty"`
	BlockedChatList      [][]byte                               `protobuf:"bytes,8,rep,name=blocked_chat_list,json=blockedChatList,proto3" json:"blocked_chat_list,omitempty"`
	Unregister           bool                                   `protobuf:"varint,9,opt,name=unregister,proto3" json:"unregister,omitempty"`
	Grant                []byte                                 `protobuf:"bytes,10,opt,name=grant,proto3" json:"grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *PushNotificationRegistration) Reset()         { *m = PushNotificationRegistration{} }
func (m *PushNotificationRegistration) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRegistration) ProtoMessage()    {}
func (*PushNotificationRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{0}
}

func (m *PushNotificationRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRegistration.Unmarshal(m, b)
}
func (m *PushNotificationRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRegistration.Marshal(b, m, deterministic)
}
func (m *PushNotificationRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRegistration.Merge(m, src)
}
func (m *PushNotificationRegistration) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRegistration.Size(m)
}
func (m *PushNotificationRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRegistration proto.InternalMessageInfo

func (m *PushNotificationRegistration) GetTokenType() PushNotificationRegistration_TokenType {
	if m != nil {
		return m.TokenType
	}
	return PushNotificationRegistration_UNKNOWN_TOKEN_TYPE
}

func (m *PushNotificationRegistration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PushNotificationRegistration) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotificationRegistration) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationRegistration) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PushNotificationRegistration) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushNotificationRegistration) GetAllowedUserList() [][]byte {
	if m != nil {
		return m.AllowedUserList
	}
	return nil
}

func (m *PushNotificationRegistration) GetBlockedChatList() [][]byte {
	if m != nil {
		return m.BlockedChatList
	}
	return nil
}

func (m *PushNotificationRegistration) GetUnregister() bool {
	if m != nil {
		return m.Unregister
	}
	return false
}

func (m *PushNotificationRegistration) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

type PushNotificationRegistrationResponse struct {
	Success              bool                                           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                PushNotificationRegistrationResponse_ErrorType `protobuf:"varint,2,opt,name=error,proto3,enum=protobuf.PushNotificationRegistrationResponse_ErrorType" json:"error,omitempty"`
	RequestId            []byte                                         `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *PushNotificationRegistrationResponse) Reset()         { *m = PushNotificationRegistrationResponse{} }
func (m *PushNotificationRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRegistrationResponse) ProtoMessage()    {}
func (*PushNotificationRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{1}
}

func (m *PushNotificationRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Unmarshal(m, b)
}
func (m *PushNotificationRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRegistrationResponse.Merge(m, src)
}
func (m *PushNotificationRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Size(m)
}
func (m *PushNotificationRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRegistrationResponse proto.InternalMessageInfo

func (m *PushNotificationRegistrationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PushNotificationRegistrationResponse) GetError() PushNotificationRegistrationResponse_ErrorType {
	if m != nil {
		return m.Error
	}
	return PushNotificationRegistrationResponse_UNKNOWN_ERROR_TYPE
}

func (m *PushNotificationRegistrationResponse) GetRequestId() []byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type PushNotificationAdvertisementInfo struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	InstallationId       string   `protobuf:"bytes,3,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationAdvertisementInfo) Reset()         { *m = PushNotificationAdvertisementInfo{} }
func (m *PushNotificationAdvertisementInfo) String() string { return proto.CompactTextString(m) }
func (*PushNotificationAdvertisementInfo) ProtoMessage()    {}
func (*PushNotificationAdvertisementInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{2}
}

func (m *PushNotificationAdvertisementInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationAdvertisementInfo.Unmarshal(m, b)
}
func (m *PushNotificationAdvertisementInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationAdvertisementInfo.Marshal(b, m, deterministic)
}
func (m *PushNotificationAdvertisementInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationAdvertisementInfo.Merge(m, src)
}
func (m *PushNotificationAdvertisementInfo) XXX_Size() int {
	return xxx_messageInfo_PushNotificationAdvertisementInfo.Size(m)
}
func (m *PushNotificationAdvertisementInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationAdvertisementInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationAdvertisementInfo proto.InternalMessageInfo

func (m *PushNotificationAdvertisementInfo) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationAdvertisementInfo) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationAdvertisementInfo) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

type ContactCodeAdvertisement struct {
	PushNotificationInfo []*PushNotificationAdvertisementInfo `protobuf:"bytes,1,rep,name=push_notification_info,json=pushNotificationInfo,proto3" json:"push_notification_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ContactCodeAdvertisement) Reset()         { *m = ContactCodeAdvertisement{} }
func (m *ContactCodeAdvertisement) String() string { return proto.CompactTextString(m) }
func (*ContactCodeAdvertisement) ProtoMessage()    {}
func (*ContactCodeAdvertisement) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{3}
}

func (m *ContactCodeAdvertisement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactCodeAdvertisement.Unmarshal(m, b)
}
func (m *ContactCodeAdvertisement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactCodeAdvertisement.Marshal(b, m, deterministic)
}
func (m *ContactCodeAdvertisement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactCodeAdvertisement.Merge(m, src)
}
func (m *ContactCodeAdvertisement) XXX_Size() int {
	return xxx_messageInfo_ContactCodeAdvertisement.Size(m)
}
func (m *ContactCodeAdvertisement) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactCodeAdvertisement.DiscardUnknown(m)
}

var xxx_messageInfo_ContactCodeAdvertisement proto.InternalMessageInfo

func (m *ContactCodeAdvertisement) GetPushNotificationInfo() []*PushNotificationAdvertisementInfo {
	if m != nil {
		return m.PushNotificationInfo
	}
	return nil
}

type PushNotificationQuery struct {
	PublicKeys           [][]byte `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationQuery) Reset()         { *m = PushNotificationQuery{} }
func (m *PushNotificationQuery) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQuery) ProtoMessage()    {}
func (*PushNotificationQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{4}
}

func (m *PushNotificationQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQuery.Unmarshal(m, b)
}
func (m *PushNotificationQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQuery.Marshal(b, m, deterministic)
}
func (m *PushNotificationQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQuery.Merge(m, src)
}
func (m *PushNotificationQuery) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQuery.Size(m)
}
func (m *PushNotificationQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQuery proto.InternalMessageInfo

func (m *PushNotificationQuery) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type PushNotificationQueryInfo struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	InstallationId       string   `protobuf:"bytes,2,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AllowedUserList      [][]byte `protobuf:"bytes,4,rep,name=allowed_user_list,json=allowedUserList,proto3" json:"allowed_user_list,omitempty"`
	Grant                []byte   `protobuf:"bytes,5,opt,name=grant,proto3" json:"grant,omitempty"`
	Version              uint64   `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationQueryInfo) Reset()         { *m = PushNotificationQueryInfo{} }
func (m *PushNotificationQueryInfo) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQueryInfo) ProtoMessage()    {}
func (*PushNotificationQueryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{5}
}

func (m *PushNotificationQueryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQueryInfo.Unmarshal(m, b)
}
func (m *PushNotificationQueryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQueryInfo.Marshal(b, m, deterministic)
}
func (m *PushNotificationQueryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQueryInfo.Merge(m, src)
}
func (m *PushNotificationQueryInfo) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQueryInfo.Size(m)
}
func (m *PushNotificationQueryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQueryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQueryInfo proto.InternalMessageInfo

func (m *PushNotificationQueryInfo) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationQueryInfo) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotificationQueryInfo) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetAllowedUserList() [][]byte {
	if m != nil {
		return m.AllowedUserList
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type PushNotificationQueryResponse struct {
	Info                 []*PushNotificationQueryInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	MessageId            []byte                       `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Success              bool                         `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PushNotificationQueryResponse) Reset()         { *m = PushNotificationQueryResponse{} }
func (m *PushNotificationQueryResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQueryResponse) ProtoMessage()    {}
func (*PushNotificationQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{6}
}

func (m *PushNotificationQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQueryResponse.Unmarshal(m, b)
}
func (m *PushNotificationQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQueryResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQueryResponse.Merge(m, src)
}
func (m *PushNotificationQueryResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQueryResponse.Size(m)
}
func (m *PushNotificationQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQueryResponse proto.InternalMessageInfo

func (m *PushNotificationQueryResponse) GetInfo() []*PushNotificationQueryInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PushNotificationQueryResponse) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *PushNotificationQueryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PushNotification struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ChatId               string   `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	InstallationId       string   `protobuf:"bytes,4,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	Message              []byte   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotification) Reset()         { *m = PushNotification{} }
func (m *PushNotification) String() string { return proto.CompactTextString(m) }
func (*PushNotification) ProtoMessage()    {}
func (*PushNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{7}
}

func (m *PushNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotification.Unmarshal(m, b)
}
func (m *PushNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotification.Marshal(b, m, deterministic)
}
func (m *PushNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotification.Merge(m, src)
}
func (m *PushNotification) XXX_Size() int {
	return xxx_messageInfo_PushNotification.Size(m)
}
func (m *PushNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotification proto.InternalMessageInfo

func (m *PushNotification) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotification) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *PushNotification) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotification) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotification) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type PushNotificationRequest struct {
	Requests             []*PushNotification `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	MessageId            []byte              `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PushNotificationRequest) Reset()         { *m = PushNotificationRequest{} }
func (m *PushNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRequest) ProtoMessage()    {}
func (*PushNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{8}
}

func (m *PushNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRequest.Unmarshal(m, b)
}
func (m *PushNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRequest.Marshal(b, m, deterministic)
}
func (m *PushNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRequest.Merge(m, src)
}
func (m *PushNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRequest.Size(m)
}
func (m *PushNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRequest proto.InternalMessageInfo

func (m *PushNotificationRequest) GetRequests() []*PushNotification {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *PushNotificationRequest) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

type PushNotificationReport struct {
	Success              bool                             `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                PushNotificationReport_ErrorType `protobuf:"varint,2,opt,name=error,proto3,enum=protobuf.PushNotificationReport_ErrorType" json:"error,omitempty"`
	PublicKey            []byte                           `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	InstallationId       string                           `protobuf:"bytes,4,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PushNotificationReport) Reset()         { *m = PushNotificationReport{} }
func (m *PushNotificationReport) String() string { return proto.CompactTextString(m) }
func (*PushNotificationReport) ProtoMessage()    {}
func (*PushNotificationReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{9}
}

func (m *PushNotificationReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationReport.Unmarshal(m, b)
}
func (m *PushNotificationReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationReport.Marshal(b, m, deterministic)
}
func (m *PushNotificationReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationReport.Merge(m, src)
}
func (m *PushNotificationReport) XXX_Size() int {
	return xxx_messageInfo_PushNotificationReport.Size(m)
}
func (m *PushNotificationReport) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationReport.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationReport proto.InternalMessageInfo

func (m *PushNotificationReport) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PushNotificationReport) GetError() PushNotificationReport_ErrorType {
	if m != nil {
		return m.Error
	}
	return PushNotificationReport_UNKNOWN_ERROR_TYPE
}

func (m *PushNotificationReport) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationReport) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

type PushNotificationResponse struct {
	MessageId            []byte                    `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Reports              []*PushNotificationReport `protobuf:"bytes,2,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PushNotificationResponse) Reset()         { *m = PushNotificationResponse{} }
func (m *PushNotificationResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationResponse) ProtoMessage()    {}
func (*PushNotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{10}
}

func (m *PushNotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationResponse.Unmarshal(m, b)
}
func (m *PushNotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationResponse.Merge(m, src)
}
func (m *PushNotificationResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationResponse.Size(m)
}
func (m *PushNotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationResponse proto.InternalMessageInfo

func (m *PushNotificationResponse) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *PushNotificationResponse) GetReports() []*PushNotificationReport {
	if m != nil {
		return m.Reports
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.PushNotificationRegistration_TokenType", PushNotificationRegistration_TokenType_name, PushNotificationRegistration_TokenType_value)
	proto.RegisterEnum("protobuf.PushNotificationRegistrationResponse_ErrorType", PushNotificationRegistrationResponse_ErrorType_name, PushNotificationRegistrationResponse_ErrorType_value)
	proto.RegisterEnum("protobuf.PushNotificationReport_ErrorType", PushNotificationReport_ErrorType_name, PushNotificationReport_ErrorType_value)
	proto.RegisterType((*PushNotificationRegistration)(nil), "protobuf.PushNotificationRegistration")
	proto.RegisterType((*PushNotificationRegistrationResponse)(nil), "protobuf.PushNotificationRegistrationResponse")
	proto.RegisterType((*PushNotificationAdvertisementInfo)(nil), "protobuf.PushNotificationAdvertisementInfo")
	proto.RegisterType((*ContactCodeAdvertisement)(nil), "protobuf.ContactCodeAdvertisement")
	proto.RegisterType((*PushNotificationQuery)(nil), "protobuf.PushNotificationQuery")
	proto.RegisterType((*PushNotificationQueryInfo)(nil), "protobuf.PushNotificationQueryInfo")
	proto.RegisterType((*PushNotificationQueryResponse)(nil), "protobuf.PushNotificationQueryResponse")
	proto.RegisterType((*PushNotification)(nil), "protobuf.PushNotification")
	proto.RegisterType((*PushNotificationRequest)(nil), "protobuf.PushNotificationRequest")
	proto.RegisterType((*PushNotificationReport)(nil), "protobuf.PushNotificationReport")
	proto.RegisterType((*PushNotificationResponse)(nil), "protobuf.PushNotificationResponse")
}

func init() { proto.RegisterFile("push_notifications.proto", fileDescriptor_200acd86044eaa5d) }

var fileDescriptor_200acd86044eaa5d = []byte{
	// 858 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0x66, 0x9d, 0xb4, 0x89, 0xa7, 0xa1, 0xcd, 0xad, 0x7a, 0x3d, 0x73, 0xe2, 0x20, 0x67, 0x90,
	0x88, 0x0e, 0x29, 0x42, 0x45, 0x82, 0x13, 0x4f, 0x84, 0xd6, 0x2d, 0x56, 0x1b, 0x3b, 0x6c, 0x5c,
	0x4e, 0x48, 0x48, 0x96, 0x13, 0x6f, 0x5b, 0xab, 0x39, 0xdb, 0xec, 0xae, 0x8b, 0xf2, 0x80, 0xc4,
	0x2f, 0x40, 0xe2, 0x95, 0x5f, 0x81, 0xf8, 0x3b, 0xbc, 0xf1, 0x4b, 0x90, 0xd7, 0x76, 0x70, 0x1c,
	0x37, 0xcd, 0x03, 0x4f, 0xf1, 0xcc, 0xce, 0xcc, 0xee, 0x7c, 0xdf, 0xcc, 0x17, 0xd0, 0xe2, 0x84,
	0xdf, 0xba, 0x61, 0x24, 0x82, 0xeb, 0x60, 0xe6, 0x89, 0x20, 0x0a, 0xf9, 0x20, 0x66, 0x91, 0x88,
	0x70, 0x5b, 0xfe, 0x4c, 0x93, 0x6b, 0xfd, 0xef, 0x06, 0xbc, 0x3f, 0x4e, 0xf8, 0xad, 0x55, 0x8a,
	0x22, 0xf4, 0x26, 0xe0, 0x82, 0xc9, 0x6f, 0x6c, 0x03, 0x88, 0xe8, 0x8e, 0x86, 0xae, 0x58, 0xc4,
	0x54, 0x43, 0x3d, 0xd4, 0xdf, 0x3f, 0xfe, 0x6c, 0x50, 0xe4, 0x0f, 0x36, 0xe5, 0x0e, 0x9c, 0x34,
	0xd1, 0x59, 0xc4, 0x94, 0xa8, 0xa2, 0xf8, 0xc4, 0x87, 0xb0, 0x23, 0x0d, 0x4d, 0xe9, 0xa1, 0xbe,
	0x4a, 0x32, 0x03, 0x7f, 0x02, 0x07, 0x41, 0xc8, 0x85, 0x37, 0x9f, 0xcb, 0x54, 0x37, 0xf0, 0xb5,
	0x86, 0x3c, 0xdf, 0x2f, 0xbb, 0x4d, 0x1f, 0xbf, 0x84, 0x8e, 0x37, 0x9b, 0x51, 0xce, 0xdd, 0xac,
	0x4a, 0x53, 0x46, 0xed, 0x65, 0x3e, 0x79, 0x21, 0xd6, 0xa0, 0x45, 0x43, 0x6f, 0x3a, 0xa7, 0xbe,
	0xb6, 0xd3, 0x43, 0xfd, 0x36, 0x29, 0xcc, 0xf4, 0xe4, 0x9e, 0x32, 0x1e, 0x44, 0xa1, 0xb6, 0xdb,
	0x43, 0xfd, 0x26, 0x29, 0x4c, 0xfc, 0x0a, 0x9e, 0x78, 0xf3, 0x79, 0xf4, 0x33, 0xf5, 0xdd, 0x84,
	0x53, 0xe6, 0xce, 0x03, 0x2e, 0xb4, 0x56, 0xaf, 0xd1, 0xef, 0x90, 0x83, 0xfc, 0xe0, 0x8a, 0x53,
	0x76, 0x19, 0x70, 0x91, 0xc6, 0x4e, 0xe7, 0xd1, 0xec, 0x8e, 0xfa, 0xee, 0xec, 0xd6, 0x13, 0x59,
	0x6c, 0x3b, 0x8b, 0xcd, 0x0f, 0x4e, 0x6e, 0x3d, 0x21, 0x63, 0x3f, 0x00, 0x48, 0x42, 0x26, 0x41,
	0xa1, 0x4c, 0x53, 0xe5, 0x73, 0x4a, 0x9e, 0x14, 0x8d, 0x1b, 0xe6, 0x85, 0x42, 0x83, 0x1e, 0xea,
	0x77, 0x48, 0x66, 0xe8, 0x67, 0xa0, 0x2e, 0xb1, 0xc3, 0x47, 0x80, 0xaf, 0xac, 0x0b, 0xcb, 0x7e,
	0x63, 0xb9, 0x8e, 0x7d, 0x61, 0x58, 0xae, 0xf3, 0xc3, 0xd8, 0xe8, 0xbe, 0x83, 0xdf, 0x05, 0x75,
	0x38, 0xce, 0x7d, 0x5d, 0x84, 0x31, 0xec, 0x9f, 0x99, 0xc4, 0xf8, 0x66, 0x38, 0x31, 0x72, 0x9f,
	0xa2, 0xff, 0xa5, 0xc0, 0xc7, 0x9b, 0x18, 0x22, 0x94, 0xc7, 0x51, 0xc8, 0x69, 0x0a, 0x0c, 0x4f,
	0x24, 0x84, 0x92, 0xe2, 0x36, 0x29, 0x4c, 0x6c, 0xc1, 0x0e, 0x65, 0x2c, 0x62, 0x92, 0xae, 0xfd,
	0xe3, 0xd7, 0xdb, 0x51, 0x5f, 0x14, 0x1e, 0x18, 0x69, 0xae, 0x1c, 0x81, 0xac, 0x0c, 0x7e, 0x01,
	0xc0, 0xe8, 0x4f, 0x09, 0xe5, 0xa2, 0xe0, 0xb8, 0x43, 0xd4, 0xdc, 0x63, 0xfa, 0xfa, 0xaf, 0x08,
	0xd4, 0x65, 0x4e, 0xb9, 0x75, 0x83, 0x10, 0x9b, 0x14, 0xad, 0x3f, 0x85, 0x27, 0xa3, 0xe1, 0xe5,
	0x99, 0x4d, 0x46, 0xc6, 0xa9, 0x3b, 0x32, 0x26, 0x93, 0xe1, 0xb9, 0xd1, 0x45, 0xf8, 0x10, 0xba,
	0xdf, 0x1b, 0x64, 0x62, 0xda, 0x96, 0x3b, 0x32, 0x27, 0xa3, 0xa1, 0x73, 0xf2, 0x6d, 0x57, 0xc1,
	0xcf, 0xe1, 0xe8, 0xca, 0x9a, 0x5c, 0x8d, 0xc7, 0x36, 0x71, 0x8c, 0xd3, 0x32, 0x86, 0x8d, 0x14,
	0x34, 0xd3, 0x72, 0x0c, 0x62, 0x0d, 0x2f, 0xb3, 0x1b, 0xba, 0x4d, 0xfd, 0x37, 0x04, 0x2f, 0xab,
	0xbd, 0x0d, 0xfd, 0x7b, 0xca, 0x44, 0xc0, 0xe9, 0x5b, 0x1a, 0x0a, 0x33, 0xbc, 0x8e, 0xd2, 0x3e,
	0xe2, 0x64, 0x3a, 0x0f, 0x66, 0xee, 0x1d, 0x5d, 0x48, 0xd0, 0x3a, 0x44, 0xcd, 0x3c, 0x17, 0x74,
	0xb1, 0x36, 0xa6, 0xca, 0xfa, 0x98, 0x6e, 0x3b, 0xf2, 0xfa, 0x2f, 0xa0, 0x9d, 0x44, 0xa1, 0xf0,
	0x66, 0xe2, 0x24, 0xf2, 0xe9, 0xca, 0x53, 0xb0, 0x07, 0x47, 0x6b, 0x5b, 0xee, 0x06, 0xe1, 0x75,
	0xa4, 0xa1, 0x5e, 0xa3, 0xbf, 0x77, 0xfc, 0xe9, 0xc3, 0x7c, 0xad, 0xf5, 0x44, 0x0e, 0xe3, 0x4a,
	0x48, 0xea, 0xd5, 0x5f, 0xc3, 0xd3, 0x6a, 0xea, 0x77, 0x09, 0x65, 0x0b, 0xfc, 0x21, 0xec, 0xfd,
	0x07, 0x01, 0x97, 0x17, 0x76, 0x08, 0x2c, 0x31, 0xe0, 0xfa, 0x3f, 0x08, 0xde, 0xab, 0x4d, 0x95,
	0x08, 0x56, 0x21, 0x42, 0x5b, 0x41, 0xa4, 0xd4, 0xaa, 0xc2, 0x2a, 0x1b, 0x8d, 0x2a, 0x1b, 0xb5,
	0xdb, 0xdd, 0xac, 0xdf, 0xee, 0xe5, 0x46, 0xee, 0x94, 0x36, 0xf2, 0x61, 0xe5, 0xd0, 0x7f, 0x47,
	0xf0, 0xa2, 0xb6, 0xc9, 0xe5, 0x72, 0x7d, 0x09, 0xcd, 0x12, 0x23, 0x1f, 0x3d, 0xcc, 0xc8, 0x12,
	0x1b, 0x22, 0x13, 0xd2, 0xae, 0xde, 0x52, 0xce, 0xbd, 0x1b, 0x5a, 0x74, 0xde, 0x21, 0x6a, 0xee,
	0x31, 0xfd, 0xf2, 0xd2, 0x36, 0x56, 0x96, 0x56, 0xff, 0x13, 0x41, 0xb7, 0x5a, 0x7c, 0x1b, 0xbc,
	0x9f, 0x41, 0x4b, 0x2a, 0xda, 0x12, 0xe7, 0xdd, 0xd4, 0x7c, 0x1c, 0xdf, 0x1a, 0x9e, 0x9a, 0xb5,
	0x3c, 0x69, 0xd0, 0xca, 0xdf, 0x9f, 0xc3, 0x5b, 0x98, 0x7a, 0x0c, 0xcf, 0xd6, 0x05, 0x45, 0xaa,
	0x02, 0xfe, 0x02, 0xda, 0xb9, 0x40, 0xf0, 0x1c, 0xc3, 0xe7, 0x1b, 0x54, 0x68, 0x19, 0xfb, 0x08,
	0x7c, 0xfa, 0x1f, 0x0a, 0x1c, 0xad, 0x5f, 0x19, 0x47, 0x4c, 0x6c, 0x90, 0xc3, 0xaf, 0x57, 0xe5,
	0xf0, 0xd5, 0x26, 0x39, 0x4c, 0x4b, 0xd5, 0x0a, 0xe0, 0xff, 0x01, 0xa5, 0xfe, 0xe3, 0x36, 0x42,
	0x79, 0x00, 0x7b, 0x6f, 0x88, 0x6d, 0x9d, 0x97, 0xff, 0x25, 0x2a, 0x82, 0xa7, 0xa4, 0x3e, 0xcb,
	0x76, 0x5c, 0x62, 0x9c, 0x9b, 0x13, 0xc7, 0x20, 0xc6, 0x69, 0xb7, 0xa1, 0x27, 0xa0, 0xad, 0x37,
	0x94, 0xcf, 0xf3, 0x2a, 0xae, 0xa8, 0x3a, 0x96, 0x5f, 0x41, 0x8b, 0xc9, 0xde, 0xb9, 0xa6, 0x48,
	0xb6, 0x7a, 0x8f, 0x81, 0x44, 0x8a, 0x84, 0xe9, 0xae, 0x8c, 0xfc, 0xfc, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x04, 0xbd, 0xc3, 0x50, 0xbb, 0x08, 0x00, 0x00,
}

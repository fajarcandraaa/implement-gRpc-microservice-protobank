// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/proto/user_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// =================
// Metadata
// =================
type UserMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total   int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *UserMeta) Reset() {
	*x = UserMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeta) ProtoMessage() {}

func (x *UserMeta) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeta.ProtoReflect.Descriptor instead.
func (*UserMeta) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserMeta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UserMeta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *UserMeta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// =================
// Parameter Request
// =================
type UserParameterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage         int32  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	OrderBy         string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderMethod     string `protobuf:"bytes,4,opt,name=order_method,json=orderMethod,proto3" json:"order_method,omitempty"`
	SearchCondition string `protobuf:"bytes,5,opt,name=search_condition,json=searchCondition,proto3" json:"search_condition,omitempty"`
	Equal           string `protobuf:"bytes,6,opt,name=equal,proto3" json:"equal,omitempty"`
	Not             string `protobuf:"bytes,7,opt,name=not,proto3" json:"not,omitempty"`
	Like            string `protobuf:"bytes,8,opt,name=like,proto3" json:"like,omitempty"`
	DateRangeBy     string `protobuf:"bytes,9,opt,name=date_range_by,json=dateRangeBy,proto3" json:"date_range_by,omitempty"`
	DateStart       string `protobuf:"bytes,10,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd         string `protobuf:"bytes,11,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
}

func (x *UserParameterRequest) Reset() {
	*x = UserParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserParameterRequest) ProtoMessage() {}

func (x *UserParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserParameterRequest.ProtoReflect.Descriptor instead.
func (*UserParameterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserParameterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UserParameterRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *UserParameterRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *UserParameterRequest) GetOrderMethod() string {
	if x != nil {
		return x.OrderMethod
	}
	return ""
}

func (x *UserParameterRequest) GetSearchCondition() string {
	if x != nil {
		return x.SearchCondition
	}
	return ""
}

func (x *UserParameterRequest) GetEqual() string {
	if x != nil {
		return x.Equal
	}
	return ""
}

func (x *UserParameterRequest) GetNot() string {
	if x != nil {
		return x.Not
	}
	return ""
}

func (x *UserParameterRequest) GetLike() string {
	if x != nil {
		return x.Like
	}
	return ""
}

func (x *UserParameterRequest) GetDateRangeBy() string {
	if x != nil {
		return x.DateRangeBy
	}
	return ""
}

func (x *UserParameterRequest) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *UserParameterRequest) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UserResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Meta *UserMeta       `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *UsersResponse) Reset() {
	*x = UsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersResponse) ProtoMessage() {}

func (x *UsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersResponse.ProtoReflect.Descriptor instead.
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *UsersResponse) GetData() []*UserResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UsersResponse) GetMeta() *UserMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type UserStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *UserResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserStatusResponse) Reset() {
	*x = UserStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatusResponse) ProtoMessage() {}

func (x *UserStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatusResponse.ProtoReflect.Descriptor instead.
func (*UserStatusResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserStatusResponse) GetData() *UserResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FindUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindUserByIdRequest) Reset() {
	*x = FindUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByIdRequest) ProtoMessage() {}

func (x *FindUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByIdRequest.ProtoReflect.Descriptor instead.
func (*FindUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *FindUserByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_grpc_proto_user_service_proto protoreflect.FileDescriptor

var file_grpc_proto_user_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xc8, 0x02,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x71, 0x75, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x69, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x22, 0x75, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x75, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbf, 0x01, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x13, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x58, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_user_service_proto_rawDescOnce sync.Once
	file_grpc_proto_user_service_proto_rawDescData = file_grpc_proto_user_service_proto_rawDesc
)

func file_grpc_proto_user_service_proto_rawDescGZIP() []byte {
	file_grpc_proto_user_service_proto_rawDescOnce.Do(func() {
		file_grpc_proto_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_user_service_proto_rawDescData)
	})
	return file_grpc_proto_user_service_proto_rawDescData
}

var file_grpc_proto_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_proto_user_service_proto_goTypes = []interface{}{
	(*UserMeta)(nil),             // 0: userservice.UserMeta
	(*UserParameterRequest)(nil), // 1: userservice.UserParameterRequest
	(*UserResponse)(nil),         // 2: userservice.UserResponse
	(*UsersResponse)(nil),        // 3: userservice.UsersResponse
	(*UserStatusResponse)(nil),   // 4: userservice.UserStatusResponse
	(*CreateUserRequest)(nil),    // 5: userservice.CreateUserRequest
	(*FindUserByIdRequest)(nil),  // 6: userservice.FindUserByIdRequest
}
var file_grpc_proto_user_service_proto_depIdxs = []int32{
	2, // 0: userservice.UsersResponse.data:type_name -> userservice.UserResponse
	0, // 1: userservice.UsersResponse.meta:type_name -> userservice.UserMeta
	2, // 2: userservice.UserStatusResponse.data:type_name -> userservice.UserResponse
	5, // 3: userservice.UserService.ServiceRegisterUser:input_type -> userservice.CreateUserRequest
	6, // 4: userservice.UserService.ServiceFindUserById:input_type -> userservice.FindUserByIdRequest
	4, // 5: userservice.UserService.ServiceRegisterUser:output_type -> userservice.UserStatusResponse
	4, // 6: userservice.UserService.ServiceFindUserById:output_type -> userservice.UserStatusResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_proto_user_service_proto_init() }
func file_grpc_proto_user_service_proto_init() {
	if File_grpc_proto_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserParameterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_user_service_proto_goTypes,
		DependencyIndexes: file_grpc_proto_user_service_proto_depIdxs,
		MessageInfos:      file_grpc_proto_user_service_proto_msgTypes,
	}.Build()
	File_grpc_proto_user_service_proto = out.File
	file_grpc_proto_user_service_proto_rawDesc = nil
	file_grpc_proto_user_service_proto_goTypes = nil
	file_grpc_proto_user_service_proto_depIdxs = nil
}

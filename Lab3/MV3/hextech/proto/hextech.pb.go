// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: hextech.proto

package proto

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

type AddProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product  string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddProductRequest) Reset() {
	*x = AddProductRequest{}
	mi := &file_hextech_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductRequest) ProtoMessage() {}

func (x *AddProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductRequest.ProtoReflect.Descriptor instead.
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{0}
}

func (x *AddProductRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AddProductRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *AddProductRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_hextech_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteProductRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DeleteProductRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product  string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_hextech_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProductRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpdateProductRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *UpdateProductRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RenameProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region     string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	OldProduct string `protobuf:"bytes,2,opt,name=old_product,json=oldProduct,proto3" json:"old_product,omitempty"`
	NewProduct string `protobuf:"bytes,3,opt,name=new_product,json=newProduct,proto3" json:"new_product,omitempty"`
}

func (x *RenameProductRequest) Reset() {
	*x = RenameProductRequest{}
	mi := &file_hextech_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RenameProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameProductRequest) ProtoMessage() {}

func (x *RenameProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameProductRequest.ProtoReflect.Descriptor instead.
func (*RenameProductRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{3}
}

func (x *RenameProductRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RenameProductRequest) GetOldProduct() string {
	if x != nil {
		return x.OldProduct
	}
	return ""
}

func (x *RenameProductRequest) GetNewProduct() string {
	if x != nil {
		return x.NewProduct
	}
	return ""
}

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_hextech_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GetProductRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

type AddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressResponse) Reset() {
	*x = AddressResponse{}
	mi := &file_hextech_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressResponse) ProtoMessage() {}

func (x *AddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressResponse.ProtoReflect.Descriptor instead.
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{5}
}

func (x *AddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ClockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []int32 `protobuf:"varint,1,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *ClockResponse) Reset() {
	*x = ClockResponse{}
	mi := &file_hextech_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockResponse) ProtoMessage() {}

func (x *ClockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockResponse.ProtoReflect.Descriptor instead.
func (*ClockResponse) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{6}
}

func (x *ClockResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity    int32   `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	VectorClock []int32 `protobuf:"varint,3,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_hextech_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{7}
}

func (x *ProductResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProductResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type ErrorMergeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region      string  `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	VectorClock []int32 `protobuf:"varint,2,rep,packed,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *ErrorMergeRequest) Reset() {
	*x = ErrorMergeRequest{}
	mi := &file_hextech_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorMergeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMergeRequest) ProtoMessage() {}

func (x *ErrorMergeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMergeRequest.ProtoReflect.Descriptor instead.
func (*ErrorMergeRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{8}
}

func (x *ErrorMergeRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ErrorMergeRequest) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type ConfirmationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmation string `protobuf:"bytes,1,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *ConfirmationError) Reset() {
	*x = ConfirmationError{}
	mi := &file_hextech_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationError) ProtoMessage() {}

func (x *ConfirmationError) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationError.ProtoReflect.Descriptor instead.
func (*ConfirmationError) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{9}
}

func (x *ConfirmationError) GetConfirmation() string {
	if x != nil {
		return x.Confirmation
	}
	return ""
}

//-------------------------------------- Solo Servidores -------------------//
type PropagationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region      string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	ChangeLog   []string `protobuf:"bytes,2,rep,name=changeLog,proto3" json:"changeLog,omitempty"`
	VectorClock []int32  `protobuf:"varint,3,rep,packed,name=vectorClock,proto3" json:"vectorClock,omitempty"`
}

func (x *PropagationRequest) Reset() {
	*x = PropagationRequest{}
	mi := &file_hextech_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PropagationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagationRequest) ProtoMessage() {}

func (x *PropagationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagationRequest.ProtoReflect.Descriptor instead.
func (*PropagationRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{10}
}

func (x *PropagationRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *PropagationRequest) GetChangeLog() []string {
	if x != nil {
		return x.ChangeLog
	}
	return nil
}

func (x *PropagationRequest) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type PropagationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // "success" o "error"
}

func (x *PropagationResponse) Reset() {
	*x = PropagationResponse{}
	mi := &file_hextech_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PropagationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagationResponse) ProtoMessage() {}

func (x *PropagationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagationResponse.ProtoReflect.Descriptor instead.
func (*PropagationResponse) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{11}
}

func (x *PropagationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MergeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *MergeRequest) Reset() {
	*x = MergeRequest{}
	mi := &file_hextech_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MergeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeRequest) ProtoMessage() {}

func (x *MergeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeRequest.ProtoReflect.Descriptor instead.
func (*MergeRequest) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{12}
}

func (x *MergeRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type MergeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangeLog   []string `protobuf:"bytes,1,rep,name=changeLog,proto3" json:"changeLog,omitempty"`
	VectorClock []int32  `protobuf:"varint,2,rep,packed,name=vectorClock,proto3" json:"vectorClock,omitempty"`
}

func (x *MergeResponse) Reset() {
	*x = MergeResponse{}
	mi := &file_hextech_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MergeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeResponse) ProtoMessage() {}

func (x *MergeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hextech_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeResponse.ProtoReflect.Descriptor instead.
func (*MergeResponse) Descriptor() ([]byte, []int) {
	return file_hextech_proto_rawDescGZIP(), []int{13}
}

func (x *MergeResponse) GetChangeLog() []string {
	if x != nil {
		return x.ChangeLog
	}
	return nil
}

func (x *MergeResponse) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

var File_hextech_proto protoreflect.FileDescriptor

var file_hextech_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x22, 0x61, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x64, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x70, 0x0a, 0x14, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x45, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x32, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x50, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x4e, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x6c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x2d,
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a,
	0x0c, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4c, 0x6f, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0xe0, 0x04, 0x0a, 0x0e, 0x48, 0x65, 0x78, 0x74, 0x65,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x78, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65,
	0x63, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a,
	0x13, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x65,
	0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x10, 0x50,
	0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12,
	0x1b, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68,
	0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x78,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hextech_proto_rawDescOnce sync.Once
	file_hextech_proto_rawDescData = file_hextech_proto_rawDesc
)

func file_hextech_proto_rawDescGZIP() []byte {
	file_hextech_proto_rawDescOnce.Do(func() {
		file_hextech_proto_rawDescData = protoimpl.X.CompressGZIP(file_hextech_proto_rawDescData)
	})
	return file_hextech_proto_rawDescData
}

var file_hextech_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_hextech_proto_goTypes = []any{
	(*AddProductRequest)(nil),    // 0: hextech.AddProductRequest
	(*DeleteProductRequest)(nil), // 1: hextech.DeleteProductRequest
	(*UpdateProductRequest)(nil), // 2: hextech.UpdateProductRequest
	(*RenameProductRequest)(nil), // 3: hextech.RenameProductRequest
	(*GetProductRequest)(nil),    // 4: hextech.GetProductRequest
	(*AddressResponse)(nil),      // 5: hextech.AddressResponse
	(*ClockResponse)(nil),        // 6: hextech.ClockResponse
	(*ProductResponse)(nil),      // 7: hextech.ProductResponse
	(*ErrorMergeRequest)(nil),    // 8: hextech.ErrorMergeRequest
	(*ConfirmationError)(nil),    // 9: hextech.ConfirmationError
	(*PropagationRequest)(nil),   // 10: hextech.PropagationRequest
	(*PropagationResponse)(nil),  // 11: hextech.PropagationResponse
	(*MergeRequest)(nil),         // 12: hextech.MergeRequest
	(*MergeResponse)(nil),        // 13: hextech.MergeResponse
}
var file_hextech_proto_depIdxs = []int32{
	0,  // 0: hextech.HextechService.AddProductServer:input_type -> hextech.AddProductRequest
	1,  // 1: hextech.HextechService.DeleteProductServer:input_type -> hextech.DeleteProductRequest
	2,  // 2: hextech.HextechService.UpdateProductServer:input_type -> hextech.UpdateProductRequest
	3,  // 3: hextech.HextechService.RenameProductServer:input_type -> hextech.RenameProductRequest
	4,  // 4: hextech.HextechService.GetProductServer:input_type -> hextech.GetProductRequest
	8,  // 5: hextech.HextechService.ForceMerge:input_type -> hextech.ErrorMergeRequest
	10, // 6: hextech.HextechService.PropagateChanges:input_type -> hextech.PropagationRequest
	12, // 7: hextech.HextechService.RequestMerge:input_type -> hextech.MergeRequest
	6,  // 8: hextech.HextechService.AddProductServer:output_type -> hextech.ClockResponse
	6,  // 9: hextech.HextechService.DeleteProductServer:output_type -> hextech.ClockResponse
	6,  // 10: hextech.HextechService.UpdateProductServer:output_type -> hextech.ClockResponse
	6,  // 11: hextech.HextechService.RenameProductServer:output_type -> hextech.ClockResponse
	7,  // 12: hextech.HextechService.GetProductServer:output_type -> hextech.ProductResponse
	9,  // 13: hextech.HextechService.ForceMerge:output_type -> hextech.ConfirmationError
	11, // 14: hextech.HextechService.PropagateChanges:output_type -> hextech.PropagationResponse
	13, // 15: hextech.HextechService.RequestMerge:output_type -> hextech.MergeResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_hextech_proto_init() }
func file_hextech_proto_init() {
	if File_hextech_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hextech_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hextech_proto_goTypes,
		DependencyIndexes: file_hextech_proto_depIdxs,
		MessageInfos:      file_hextech_proto_msgTypes,
	}.Build()
	File_hextech_proto = out.File
	file_hextech_proto_rawDesc = nil
	file_hextech_proto_goTypes = nil
	file_hextech_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: shop/v1/shop.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	mi := &file_shop_v1_shop_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{0}
}

// InfoResponse представляет информацию о монетах, инвентаре и истории транзакций
type InfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Coins         int32                  `protobuf:"varint,1,opt,name=coins,proto3" json:"coins,omitempty"`
	Inventory     []*InventoryItem       `protobuf:"bytes,2,rep,name=inventory,proto3" json:"inventory,omitempty"`
	CoinHistory   *CoinHistoryDetails    `protobuf:"bytes,3,opt,name=coin_history,json=coinHistory,proto3" json:"coin_history,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	mi := &file_shop_v1_shop_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetCoins() int32 {
	if x != nil {
		return x.Coins
	}
	return 0
}

func (x *InfoResponse) GetInventory() []*InventoryItem {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *InfoResponse) GetCoinHistory() *CoinHistoryDetails {
	if x != nil {
		return x.CoinHistory
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_shop_v1_shop_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Предмет в инвентаре
type InventoryItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryItem) Reset() {
	*x = InventoryItem{}
	mi := &file_shop_v1_shop_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItem) ProtoMessage() {}

func (x *InventoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItem.ProtoReflect.Descriptor instead.
func (*InventoryItem) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{3}
}

func (x *InventoryItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InventoryItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// История транзакций с монетами
type CoinHistoryDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Received      []*ReceivedTransaction `protobuf:"bytes,1,rep,name=received,proto3" json:"received,omitempty"`
	Sent          []*SentTransaction     `protobuf:"bytes,2,rep,name=sent,proto3" json:"sent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CoinHistoryDetails) Reset() {
	*x = CoinHistoryDetails{}
	mi := &file_shop_v1_shop_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CoinHistoryDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinHistoryDetails) ProtoMessage() {}

func (x *CoinHistoryDetails) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinHistoryDetails.ProtoReflect.Descriptor instead.
func (*CoinHistoryDetails) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{4}
}

func (x *CoinHistoryDetails) GetReceived() []*ReceivedTransaction {
	if x != nil {
		return x.Received
	}
	return nil
}

func (x *CoinHistoryDetails) GetSent() []*SentTransaction {
	if x != nil {
		return x.Sent
	}
	return nil
}

// Полученная транзакция
type ReceivedTransaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromUser      string                 `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceivedTransaction) Reset() {
	*x = ReceivedTransaction{}
	mi := &file_shop_v1_shop_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceivedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedTransaction) ProtoMessage() {}

func (x *ReceivedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedTransaction.ProtoReflect.Descriptor instead.
func (*ReceivedTransaction) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{5}
}

func (x *ReceivedTransaction) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *ReceivedTransaction) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Отправленная транзакция
type SentTransaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ToUser        string                 `protobuf:"bytes,1,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SentTransaction) Reset() {
	*x = SentTransaction{}
	mi := &file_shop_v1_shop_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SentTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentTransaction) ProtoMessage() {}

func (x *SentTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentTransaction.ProtoReflect.Descriptor instead.
func (*SentTransaction) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{6}
}

func (x *SentTransaction) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *SentTransaction) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Запрос на аутентификацию
type AuthRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Имя пользователя для аутентификации
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Пароль для аутентификации
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	mi := &file_shop_v1_shop_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{7}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Ответ на запрос аутентификации
type AuthResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*AuthResponse_Token
	//	*AuthResponse_Error
	Data          isAuthResponse_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_shop_v1_shop_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{8}
}

func (x *AuthResponse) GetData() isAuthResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		if x, ok := x.Data.(*AuthResponse_Token); ok {
			return x.Token
		}
	}
	return ""
}

func (x *AuthResponse) GetError() string {
	if x != nil {
		if x, ok := x.Data.(*AuthResponse_Error); ok {
			return x.Error
		}
	}
	return ""
}

type isAuthResponse_Data interface {
	isAuthResponse_Data()
}

type AuthResponse_Token struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3,oneof"`
}

type AuthResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*AuthResponse_Token) isAuthResponse_Data() {}

func (*AuthResponse_Error) isAuthResponse_Data() {}

type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         string                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_shop_v1_shop_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{9}
}

func (x *BaseResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_shop_v1_shop_proto protoreflect.FileDescriptor

var file_shop_v1_shop_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e,
	0x73, 0x12, 0x34, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x6f, 0x69, 0x6e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x1a, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x7c, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x73, 0x65,
	0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42,
	0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x45, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x24, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xba, 0x02, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x46, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x43, 0x6f, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x15,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a,
	0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x48, 0x0a, 0x07, 0x42, 0x75, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0d, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x75, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x49, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x42, 0x15, 0x5a, 0x13, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_shop_v1_shop_proto_rawDescOnce sync.Once
	file_shop_v1_shop_proto_rawDescData []byte
)

func file_shop_v1_shop_proto_rawDescGZIP() []byte {
	file_shop_v1_shop_proto_rawDescOnce.Do(func() {
		file_shop_v1_shop_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shop_v1_shop_proto_rawDesc), len(file_shop_v1_shop_proto_rawDesc)))
	})
	return file_shop_v1_shop_proto_rawDescData
}

var file_shop_v1_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_shop_v1_shop_proto_goTypes = []any{
	(*InfoRequest)(nil),         // 0: shop.v1.InfoRequest
	(*InfoResponse)(nil),        // 1: shop.v1.InfoResponse
	(*Item)(nil),                // 2: shop.v1.Item
	(*InventoryItem)(nil),       // 3: shop.v1.InventoryItem
	(*CoinHistoryDetails)(nil),  // 4: shop.v1.CoinHistoryDetails
	(*ReceivedTransaction)(nil), // 5: shop.v1.ReceivedTransaction
	(*SentTransaction)(nil),     // 6: shop.v1.SentTransaction
	(*AuthRequest)(nil),         // 7: shop.v1.AuthRequest
	(*AuthResponse)(nil),        // 8: shop.v1.AuthResponse
	(*BaseResponse)(nil),        // 9: shop.v1.BaseResponse
}
var file_shop_v1_shop_proto_depIdxs = []int32{
	3, // 0: shop.v1.InfoResponse.inventory:type_name -> shop.v1.InventoryItem
	4, // 1: shop.v1.InfoResponse.coin_history:type_name -> shop.v1.CoinHistoryDetails
	5, // 2: shop.v1.CoinHistoryDetails.received:type_name -> shop.v1.ReceivedTransaction
	6, // 3: shop.v1.CoinHistoryDetails.sent:type_name -> shop.v1.SentTransaction
	0, // 4: shop.v1.Shop.Info:input_type -> shop.v1.InfoRequest
	6, // 5: shop.v1.Shop.SendCoin:input_type -> shop.v1.SentTransaction
	2, // 6: shop.v1.Shop.BuyItem:input_type -> shop.v1.Item
	7, // 7: shop.v1.Shop.Auth:input_type -> shop.v1.AuthRequest
	1, // 8: shop.v1.Shop.Info:output_type -> shop.v1.InfoResponse
	9, // 9: shop.v1.Shop.SendCoin:output_type -> shop.v1.BaseResponse
	9, // 10: shop.v1.Shop.BuyItem:output_type -> shop.v1.BaseResponse
	8, // 11: shop.v1.Shop.Auth:output_type -> shop.v1.AuthResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shop_v1_shop_proto_init() }
func file_shop_v1_shop_proto_init() {
	if File_shop_v1_shop_proto != nil {
		return
	}
	file_shop_v1_shop_proto_msgTypes[8].OneofWrappers = []any{
		(*AuthResponse_Token)(nil),
		(*AuthResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shop_v1_shop_proto_rawDesc), len(file_shop_v1_shop_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_v1_shop_proto_goTypes,
		DependencyIndexes: file_shop_v1_shop_proto_depIdxs,
		MessageInfos:      file_shop_v1_shop_proto_msgTypes,
	}.Build()
	File_shop_v1_shop_proto = out.File
	file_shop_v1_shop_proto_goTypes = nil
	file_shop_v1_shop_proto_depIdxs = nil
}

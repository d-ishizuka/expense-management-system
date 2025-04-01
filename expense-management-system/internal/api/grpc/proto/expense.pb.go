// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: internal/api/grpc/proto/expense.proto

package proto

import (
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

type CreateExpenseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount        float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Date          string                 `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateExpenseRequest) Reset() {
	*x = CreateExpenseRequest{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExpenseRequest) ProtoMessage() {}

func (x *CreateExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExpenseRequest.ProtoReflect.Descriptor instead.
func (*CreateExpenseRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExpenseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateExpenseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateExpenseRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateExpenseRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateExpenseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpenseId     string                 `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateExpenseResponse) Reset() {
	*x = CreateExpenseResponse{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExpenseResponse) ProtoMessage() {}

func (x *CreateExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExpenseResponse.ProtoReflect.Descriptor instead.
func (*CreateExpenseResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExpenseResponse) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

type GetExpenseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpenseId     string                 `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetExpenseRequest) Reset() {
	*x = GetExpenseRequest{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExpenseRequest) ProtoMessage() {}

func (x *GetExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExpenseRequest.ProtoReflect.Descriptor instead.
func (*GetExpenseRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{2}
}

func (x *GetExpenseRequest) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

type GetExpenseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpenseId     string                 `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount        float64                `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Date          string                 `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetExpenseResponse) Reset() {
	*x = GetExpenseResponse{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExpenseResponse) ProtoMessage() {}

func (x *GetExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExpenseResponse.ProtoReflect.Descriptor instead.
func (*GetExpenseResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{3}
}

func (x *GetExpenseResponse) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

func (x *GetExpenseResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetExpenseResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetExpenseResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetExpenseResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type UpdateExpenseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpenseId     string                 `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount        float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Date          string                 `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateExpenseRequest) Reset() {
	*x = UpdateExpenseRequest{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExpenseRequest) ProtoMessage() {}

func (x *UpdateExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExpenseRequest.ProtoReflect.Descriptor instead.
func (*UpdateExpenseRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateExpenseRequest) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

func (x *UpdateExpenseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateExpenseRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateExpenseRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type UpdateExpenseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateExpenseResponse) Reset() {
	*x = UpdateExpenseResponse{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExpenseResponse) ProtoMessage() {}

func (x *UpdateExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExpenseResponse.ProtoReflect.Descriptor instead.
func (*UpdateExpenseResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateExpenseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteExpenseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExpenseId     string                 `protobuf:"bytes,1,opt,name=expense_id,json=expenseId,proto3" json:"expense_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteExpenseRequest) Reset() {
	*x = DeleteExpenseRequest{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExpenseRequest) ProtoMessage() {}

func (x *DeleteExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExpenseRequest.ProtoReflect.Descriptor instead.
func (*DeleteExpenseRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteExpenseRequest) GetExpenseId() string {
	if x != nil {
		return x.ExpenseId
	}
	return ""
}

type DeleteExpenseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteExpenseResponse) Reset() {
	*x = DeleteExpenseResponse{}
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExpenseResponse) ProtoMessage() {}

func (x *DeleteExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_grpc_proto_expense_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExpenseResponse.ProtoReflect.Descriptor instead.
func (*DeleteExpenseResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_grpc_proto_expense_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteExpenseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_internal_api_grpc_proto_expense_proto protoreflect.FileDescriptor

const file_internal_api_grpc_proto_expense_proto_rawDesc = "" +
	"\n" +
	"%internal/api/grpc/proto/expense.proto\x12\aexpense\"}\n" +
	"\x14CreateExpenseRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x12\n" +
	"\x04date\x18\x04 \x01(\tR\x04date\"6\n" +
	"\x15CreateExpenseResponse\x12\x1d\n" +
	"\n" +
	"expense_id\x18\x01 \x01(\tR\texpenseId\"2\n" +
	"\x11GetExpenseRequest\x12\x1d\n" +
	"\n" +
	"expense_id\x18\x01 \x01(\tR\texpenseId\"\x9a\x01\n" +
	"\x12GetExpenseResponse\x12\x1d\n" +
	"\n" +
	"expense_id\x18\x01 \x01(\tR\texpenseId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\x01R\x06amount\x12\x12\n" +
	"\x04date\x18\x05 \x01(\tR\x04date\"\x83\x01\n" +
	"\x14UpdateExpenseRequest\x12\x1d\n" +
	"\n" +
	"expense_id\x18\x01 \x01(\tR\texpenseId\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x12\n" +
	"\x04date\x18\x04 \x01(\tR\x04date\"1\n" +
	"\x15UpdateExpenseResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"5\n" +
	"\x14DeleteExpenseRequest\x12\x1d\n" +
	"\n" +
	"expense_id\x18\x01 \x01(\tR\texpenseId\"1\n" +
	"\x15DeleteExpenseResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xc7\x02\n" +
	"\x0eExpenseService\x12N\n" +
	"\rCreateExpense\x12\x1d.expense.CreateExpenseRequest\x1a\x1e.expense.CreateExpenseResponse\x12E\n" +
	"\n" +
	"GetExpense\x12\x1a.expense.GetExpenseRequest\x1a\x1b.expense.GetExpenseResponse\x12N\n" +
	"\rUpdateExpense\x12\x1d.expense.UpdateExpenseRequest\x1a\x1e.expense.UpdateExpenseResponse\x12N\n" +
	"\rDeleteExpense\x12\x1d.expense.DeleteExpenseRequest\x1a\x1e.expense.DeleteExpenseResponseB3Z1expense-management-system/internal/api/grpc/protob\x06proto3"

var (
	file_internal_api_grpc_proto_expense_proto_rawDescOnce sync.Once
	file_internal_api_grpc_proto_expense_proto_rawDescData []byte
)

func file_internal_api_grpc_proto_expense_proto_rawDescGZIP() []byte {
	file_internal_api_grpc_proto_expense_proto_rawDescOnce.Do(func() {
		file_internal_api_grpc_proto_expense_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_api_grpc_proto_expense_proto_rawDesc), len(file_internal_api_grpc_proto_expense_proto_rawDesc)))
	})
	return file_internal_api_grpc_proto_expense_proto_rawDescData
}

var file_internal_api_grpc_proto_expense_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_api_grpc_proto_expense_proto_goTypes = []any{
	(*CreateExpenseRequest)(nil),  // 0: expense.CreateExpenseRequest
	(*CreateExpenseResponse)(nil), // 1: expense.CreateExpenseResponse
	(*GetExpenseRequest)(nil),     // 2: expense.GetExpenseRequest
	(*GetExpenseResponse)(nil),    // 3: expense.GetExpenseResponse
	(*UpdateExpenseRequest)(nil),  // 4: expense.UpdateExpenseRequest
	(*UpdateExpenseResponse)(nil), // 5: expense.UpdateExpenseResponse
	(*DeleteExpenseRequest)(nil),  // 6: expense.DeleteExpenseRequest
	(*DeleteExpenseResponse)(nil), // 7: expense.DeleteExpenseResponse
}
var file_internal_api_grpc_proto_expense_proto_depIdxs = []int32{
	0, // 0: expense.ExpenseService.CreateExpense:input_type -> expense.CreateExpenseRequest
	2, // 1: expense.ExpenseService.GetExpense:input_type -> expense.GetExpenseRequest
	4, // 2: expense.ExpenseService.UpdateExpense:input_type -> expense.UpdateExpenseRequest
	6, // 3: expense.ExpenseService.DeleteExpense:input_type -> expense.DeleteExpenseRequest
	1, // 4: expense.ExpenseService.CreateExpense:output_type -> expense.CreateExpenseResponse
	3, // 5: expense.ExpenseService.GetExpense:output_type -> expense.GetExpenseResponse
	5, // 6: expense.ExpenseService.UpdateExpense:output_type -> expense.UpdateExpenseResponse
	7, // 7: expense.ExpenseService.DeleteExpense:output_type -> expense.DeleteExpenseResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_api_grpc_proto_expense_proto_init() }
func file_internal_api_grpc_proto_expense_proto_init() {
	if File_internal_api_grpc_proto_expense_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_api_grpc_proto_expense_proto_rawDesc), len(file_internal_api_grpc_proto_expense_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_grpc_proto_expense_proto_goTypes,
		DependencyIndexes: file_internal_api_grpc_proto_expense_proto_depIdxs,
		MessageInfos:      file_internal_api_grpc_proto_expense_proto_msgTypes,
	}.Build()
	File_internal_api_grpc_proto_expense_proto = out.File
	file_internal_api_grpc_proto_expense_proto_goTypes = nil
	file_internal_api_grpc_proto_expense_proto_depIdxs = nil
}

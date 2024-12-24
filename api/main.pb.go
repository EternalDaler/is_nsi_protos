// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto/main.proto

package api

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_proto_main_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
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
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type    string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Symbols []string `protobuf:"bytes,4,rep,name=symbols,proto3" json:"symbols,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	mi := &file_proto_main_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

func (x *Rule) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Rule) GetSymbols() []string {
	if x != nil {
		return x.Symbols
	}
	return nil
}

type RuleResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId      uint32 `protobuf:"varint,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	Failed      bool   `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RuleResult) Reset() {
	*x = RuleResult{}
	mi := &file_proto_main_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleResult) ProtoMessage() {}

func (x *RuleResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleResult.ProtoReflect.Descriptor instead.
func (*RuleResult) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

func (x *RuleResult) GetRuleId() uint32 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

func (x *RuleResult) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

func (x *RuleResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RuleSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleId      uint32 `protobuf:"varint,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	FailedCount bool   `protobuf:"varint,2,opt,name=failed_count,json=failedCount,proto3" json:"failed_count,omitempty"`
}

func (x *RuleSummary) Reset() {
	*x = RuleSummary{}
	mi := &file_proto_main_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleSummary) ProtoMessage() {}

func (x *RuleSummary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleSummary.ProtoReflect.Descriptor instead.
func (*RuleSummary) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{3}
}

func (x *RuleSummary) GetRuleId() uint32 {
	if x != nil {
		return x.RuleId
	}
	return 0
}

func (x *RuleSummary) GetFailedCount() bool {
	if x != nil {
		return x.FailedCount
	}
	return false
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OriginalName string        `protobuf:"bytes,2,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	CleanedName  string        `protobuf:"bytes,3,opt,name=cleaned_name,json=cleanedName,proto3" json:"cleaned_name,omitempty"`
	RulesResults []*RuleResult `protobuf:"bytes,4,rep,name=rules_results,json=rulesResults,proto3" json:"rules_results,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_proto_main_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *Record) GetCleanedName() string {
	if x != nil {
		return x.CleanedName
	}
	return ""
}

func (x *Record) GetRulesResults() []*RuleResult {
	if x != nil {
		return x.RulesResults
	}
	return nil
}

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalNames uint32         `protobuf:"varint,1,opt,name=total_names,json=totalNames,proto3" json:"total_names,omitempty"`
	Rules      []*RuleSummary `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	mi := &file_proto_main_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{5}
}

func (x *Summary) GetTotalNames() uint32 {
	if x != nil {
		return x.TotalNames
	}
	return 0
}

func (x *Summary) GetRules() []*RuleSummary {
	if x != nil {
		return x.Rules
	}
	return nil
}

type NormalizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Rules []*Rule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *NormalizeRequest) Reset() {
	*x = NormalizeRequest{}
	mi := &file_proto_main_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NormalizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizeRequest) ProtoMessage() {}

func (x *NormalizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizeRequest.ProtoReflect.Descriptor instead.
func (*NormalizeRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{6}
}

func (x *NormalizeRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *NormalizeRequest) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type NormalizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	Summary *Summary  `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *NormalizeResponse) Reset() {
	*x = NormalizeResponse{}
	mi := &file_proto_main_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NormalizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizeResponse) ProtoMessage() {}

func (x *NormalizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizeResponse.ProtoReflect.Descriptor instead.
func (*NormalizeResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{7}
}

func (x *NormalizeResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *NormalizeResponse) GetSummary() *Summary {
	if x != nil {
		return x.Summary
	}
	return nil
}

type DuplicateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DuplicateRequest) Reset() {
	*x = DuplicateRequest{}
	mi := &file_proto_main_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DuplicateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuplicateRequest) ProtoMessage() {}

func (x *DuplicateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuplicateRequest.ProtoReflect.Descriptor instead.
func (*DuplicateRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{8}
}

func (x *DuplicateRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type DuplicateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TotalItems int32    `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Duplicates []string `protobuf:"bytes,3,rep,name=duplicates,proto3" json:"duplicates,omitempty"`
}

func (x *DuplicateResponse) Reset() {
	*x = DuplicateResponse{}
	mi := &file_proto_main_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DuplicateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuplicateResponse) ProtoMessage() {}

func (x *DuplicateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuplicateResponse.ProtoReflect.Descriptor instead.
func (*DuplicateResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{9}
}

func (x *DuplicateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DuplicateResponse) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *DuplicateResponse) GetDuplicates() []string {
	if x != nil {
		return x.Duplicates
	}
	return nil
}

type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GramMorphError string `protobuf:"bytes,1,opt,name=gram_morph_error,json=gramMorphError,proto3" json:"gram_morph_error,omitempty"`
	FirstNounError string `protobuf:"bytes,2,opt,name=first_noun_error,json=firstNounError,proto3" json:"first_noun_error,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	mi := &file_proto_main_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

func (x *Description) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{10}
}

func (x *Description) GetGramMorphError() string {
	if x != nil {
		return x.GramMorphError
	}
	return ""
}

func (x *Description) GetFirstNounError() string {
	if x != nil {
		return x.FirstNounError
	}
	return ""
}

type FirstSingularNoun_And_NoGrammaticalErrorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrectedName string       `protobuf:"bytes,1,opt,name=corrected_name,json=correctedName,proto3" json:"corrected_name,omitempty"`
	GramMorphRule bool         `protobuf:"varint,2,opt,name=gram_morph_rule,json=gramMorphRule,proto3" json:"gram_morph_rule,omitempty"`
	FirstNounRule bool         `protobuf:"varint,3,opt,name=first_noun_rule,json=firstNounRule,proto3" json:"first_noun_rule,omitempty"`
	Description   *Description `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) Reset() {
	*x = FirstSingularNoun_And_NoGrammaticalErrorsResponse{}
	mi := &file_proto_main_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirstSingularNoun_And_NoGrammaticalErrorsResponse) ProtoMessage() {}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirstSingularNoun_And_NoGrammaticalErrorsResponse.ProtoReflect.Descriptor instead.
func (*FirstSingularNoun_And_NoGrammaticalErrorsResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{11}
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) GetCorrectedName() string {
	if x != nil {
		return x.CorrectedName
	}
	return ""
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) GetGramMorphRule() bool {
	if x != nil {
		return x.GramMorphRule
	}
	return false
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) GetFirstNounRule() bool {
	if x != nil {
		return x.FirstNounRule
	}
	return false
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsResponse) GetDescription() *Description {
	if x != nil {
		return x.Description
	}
	return nil
}

type FirstSingularNoun_And_NoGrammaticalErrorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsRequest) Reset() {
	*x = FirstSingularNoun_And_NoGrammaticalErrorsRequest{}
	mi := &file_proto_main_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FirstSingularNoun_And_NoGrammaticalErrorsRequest) ProtoMessage() {}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FirstSingularNoun_And_NoGrammaticalErrorsRequest.ProtoReflect.Descriptor instead.
func (*FirstSingularNoun_And_NoGrammaticalErrorsRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{12}
}

func (x *FirstSingularNoun_And_NoGrammaticalErrorsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_main_proto protoreflect.FileDescriptor

var file_proto_main_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x2a, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x22, 0x5f,
	0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x49, 0x0a, 0x0b, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x0d, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x10, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x20, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x22, 0x64, 0x0a, 0x11, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x27,
	0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x34, 0x0a, 0x10, 0x44, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x6e, 0x0a,
	0x11, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x22, 0x61, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10,
	0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x72, 0x61, 0x6d, 0x4d, 0x6f, 0x72, 0x70,
	0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x6f, 0x75, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x6f, 0x75, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0xdf, 0x01, 0x0a, 0x31, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x4e, 0x6f, 0x75, 0x6e, 0x5f, 0x41, 0x6e, 0x64, 0x5f, 0x4e, 0x6f, 0x47, 0x72, 0x61,
	0x6d, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x5f, 0x72, 0x75, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x67, 0x72, 0x61, 0x6d, 0x4d, 0x6f, 0x72, 0x70,
	0x68, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x6f, 0x75, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x6f, 0x75, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x33, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x30, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x4e, 0x6f, 0x75, 0x6e, 0x5f, 0x41, 0x6e, 0x64, 0x5f, 0x4e, 0x6f, 0x47, 0x72,
	0x61, 0x6d, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x50, 0x0a, 0x10, 0x4e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf5, 0x01, 0x0a,
	0x12, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x50, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x75, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x29, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53,
	0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4e, 0x6f, 0x75, 0x6e, 0x5f, 0x41, 0x6e, 0x64, 0x5f,
	0x4e, 0x6f, 0x47, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x36, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4e, 0x6f, 0x75, 0x6e, 0x5f, 0x41, 0x6e, 0x64,
	0x5f, 0x4e, 0x6f, 0x47, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x4e, 0x6f, 0x75, 0x6e, 0x5f, 0x41, 0x6e, 0x64, 0x5f, 0x4e, 0x6f, 0x47, 0x72, 0x61, 0x6d, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData = file_proto_main_proto_rawDesc
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_proto_rawDescData)
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_main_proto_goTypes = []any{
	(*Item)(nil),              // 0: main.Item
	(*Rule)(nil),              // 1: main.Rule
	(*RuleResult)(nil),        // 2: main.RuleResult
	(*RuleSummary)(nil),       // 3: main.RuleSummary
	(*Record)(nil),            // 4: main.Record
	(*Summary)(nil),           // 5: main.Summary
	(*NormalizeRequest)(nil),  // 6: main.NormalizeRequest
	(*NormalizeResponse)(nil), // 7: main.NormalizeResponse
	(*DuplicateRequest)(nil),  // 8: main.DuplicateRequest
	(*DuplicateResponse)(nil), // 9: main.DuplicateResponse
	(*Description)(nil),       // 10: main.Description
	(*FirstSingularNoun_And_NoGrammaticalErrorsResponse)(nil), // 11: main.FirstSingularNoun_And_NoGrammaticalErrorsResponse
	(*FirstSingularNoun_And_NoGrammaticalErrorsRequest)(nil),  // 12: main.FirstSingularNoun_And_NoGrammaticalErrorsRequest
}
var file_proto_main_proto_depIdxs = []int32{
	2,  // 0: main.Record.rules_results:type_name -> main.RuleResult
	3,  // 1: main.Summary.rules:type_name -> main.RuleSummary
	0,  // 2: main.NormalizeRequest.items:type_name -> main.Item
	1,  // 3: main.NormalizeRequest.rules:type_name -> main.Rule
	4,  // 4: main.NormalizeResponse.records:type_name -> main.Record
	5,  // 5: main.NormalizeResponse.summary:type_name -> main.Summary
	0,  // 6: main.DuplicateRequest.items:type_name -> main.Item
	10, // 7: main.FirstSingularNoun_And_NoGrammaticalErrorsResponse.description:type_name -> main.Description
	6,  // 8: main.NormalizeService.Normalize:input_type -> main.NormalizeRequest
	8,  // 9: main.NormalizePyService.GetDuplicates:input_type -> main.DuplicateRequest
	12, // 10: main.NormalizePyService.FirstSingularNoun_And_NoGrammaticalErrors:input_type -> main.FirstSingularNoun_And_NoGrammaticalErrorsRequest
	7,  // 11: main.NormalizeService.Normalize:output_type -> main.NormalizeResponse
	9,  // 12: main.NormalizePyService.GetDuplicates:output_type -> main.DuplicateResponse
	11, // 13: main.NormalizePyService.FirstSingularNoun_And_NoGrammaticalErrors:output_type -> main.FirstSingularNoun_And_NoGrammaticalErrorsResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_rawDesc = nil
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}

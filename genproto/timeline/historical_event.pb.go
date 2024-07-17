// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: submodule-for-timecapsule/timeline_service/historical_event.proto

package timeline

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

// HistoricalEvent represents a historical event associated with a user.
type HistoricalEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // Added user_id field
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"` // Use string for date (YYYY-MM-DD format)
	Category    string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	SourceUrl   string `protobuf:"bytes,7,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *HistoricalEvent) Reset() {
	*x = HistoricalEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricalEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalEvent) ProtoMessage() {}

func (x *HistoricalEvent) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricalEvent.ProtoReflect.Descriptor instead.
func (*HistoricalEvent) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP(), []int{0}
}

func (x *HistoricalEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HistoricalEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HistoricalEvent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *HistoricalEvent) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *HistoricalEvent) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *HistoricalEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HistoricalEvent) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *HistoricalEvent) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// GetHistoricalEventByIdRequest represents a request to retrieve a historical event by its ID.
type GetHistoricalEventByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHistoricalEventByIdRequest) Reset() {
	*x = GetHistoricalEventByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoricalEventByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoricalEventByIdRequest) ProtoMessage() {}

func (x *GetHistoricalEventByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoricalEventByIdRequest.ProtoReflect.Descriptor instead.
func (*GetHistoricalEventByIdRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP(), []int{1}
}

func (x *GetHistoricalEventByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteHistoricalEventRequest represents a request to delete a historical event by its ID.
type DeleteHistoricalEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHistoricalEventRequest) Reset() {
	*x = DeleteHistoricalEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHistoricalEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHistoricalEventRequest) ProtoMessage() {}

func (x *DeleteHistoricalEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHistoricalEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteHistoricalEventRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteHistoricalEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteHistoricalEventResponse represents a response to a historical event deletion request.
type DeleteHistoricalEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteHistoricalEventResponse) Reset() {
	*x = DeleteHistoricalEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHistoricalEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHistoricalEventResponse) ProtoMessage() {}

func (x *DeleteHistoricalEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHistoricalEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteHistoricalEventResponse) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteHistoricalEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// GetAllHistoricalEventsRequest represents a request to retrieve all historical events.
type GetAllHistoricalEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit       int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                          // Filter by title
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`              // Filter by description
	Category    string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`                    // Filter by category
	StartDate   string `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"` // Filter by start date (inclusive)
	EndDate     string `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`       // Filter by end date (inclusive)
}

func (x *GetAllHistoricalEventsRequest) Reset() {
	*x = GetAllHistoricalEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHistoricalEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHistoricalEventsRequest) ProtoMessage() {}

func (x *GetAllHistoricalEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHistoricalEventsRequest.ProtoReflect.Descriptor instead.
func (*GetAllHistoricalEventsRequest) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllHistoricalEventsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllHistoricalEventsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllHistoricalEventsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetAllHistoricalEventsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetAllHistoricalEventsRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GetAllHistoricalEventsRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetAllHistoricalEventsRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// GetAllHistoricalEventsResponse represents a response containing a list of historical events.
type GetAllHistoricalEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoricalEvents []*HistoricalEvent `protobuf:"bytes,1,rep,name=historical_events,json=historicalEvents,proto3" json:"historical_events,omitempty"`
	Count            int32              `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllHistoricalEventsResponse) Reset() {
	*x = GetAllHistoricalEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHistoricalEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHistoricalEventsResponse) ProtoMessage() {}

func (x *GetAllHistoricalEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHistoricalEventsResponse.ProtoReflect.Descriptor instead.
func (*GetAllHistoricalEventsResponse) Descriptor() ([]byte, []int) {
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllHistoricalEventsResponse) GetHistoricalEvents() []*HistoricalEvent {
	if x != nil {
		return x.HistoricalEvents
	}
	return nil
}

func (x *GetAllHistoricalEventsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_submodule_for_timecapsule_timeline_service_historical_event_proto protoreflect.FileDescriptor

var file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDesc = []byte{
	0x0a, 0x41, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x66, 0x6f, 0x72, 0x2d,
	0x74, 0x69, 0x6d, 0x65, 0x63, 0x61, 0x70, 0x73, 0x75, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xe0, 0x01,
	0x0a, 0x0f, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x2f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x39, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xd7, 0x01, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x7e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x10,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xcd, 0x02, 0x0a, 0x16, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x27, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x68, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescOnce sync.Once
	file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescData = file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDesc
)

func file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescGZIP() []byte {
	file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescOnce.Do(func() {
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescData)
	})
	return file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDescData
}

var file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_submodule_for_timecapsule_timeline_service_historical_event_proto_goTypes = []any{
	(*HistoricalEvent)(nil),                // 0: timeline.HistoricalEvent
	(*GetHistoricalEventByIdRequest)(nil),  // 1: timeline.GetHistoricalEventByIdRequest
	(*DeleteHistoricalEventRequest)(nil),   // 2: timeline.DeleteHistoricalEventRequest
	(*DeleteHistoricalEventResponse)(nil),  // 3: timeline.DeleteHistoricalEventResponse
	(*GetAllHistoricalEventsRequest)(nil),  // 4: timeline.GetAllHistoricalEventsRequest
	(*GetAllHistoricalEventsResponse)(nil), // 5: timeline.GetAllHistoricalEventsResponse
}
var file_submodule_for_timecapsule_timeline_service_historical_event_proto_depIdxs = []int32{
	0, // 0: timeline.GetAllHistoricalEventsResponse.historical_events:type_name -> timeline.HistoricalEvent
	1, // 1: timeline.HistoricalEventService.GetHistoricalEventById:input_type -> timeline.GetHistoricalEventByIdRequest
	2, // 2: timeline.HistoricalEventService.DeleteHistoricalEvent:input_type -> timeline.DeleteHistoricalEventRequest
	4, // 3: timeline.HistoricalEventService.GetAllHistoricalEvents:input_type -> timeline.GetAllHistoricalEventsRequest
	0, // 4: timeline.HistoricalEventService.GetHistoricalEventById:output_type -> timeline.HistoricalEvent
	3, // 5: timeline.HistoricalEventService.DeleteHistoricalEvent:output_type -> timeline.DeleteHistoricalEventResponse
	5, // 6: timeline.HistoricalEventService.GetAllHistoricalEvents:output_type -> timeline.GetAllHistoricalEventsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submodule_for_timecapsule_timeline_service_historical_event_proto_init() }
func file_submodule_for_timecapsule_timeline_service_historical_event_proto_init() {
	if File_submodule_for_timecapsule_timeline_service_historical_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HistoricalEvent); i {
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
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetHistoricalEventByIdRequest); i {
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
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHistoricalEventRequest); i {
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
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHistoricalEventResponse); i {
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
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllHistoricalEventsRequest); i {
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
		file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllHistoricalEventsResponse); i {
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
			RawDescriptor: file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submodule_for_timecapsule_timeline_service_historical_event_proto_goTypes,
		DependencyIndexes: file_submodule_for_timecapsule_timeline_service_historical_event_proto_depIdxs,
		MessageInfos:      file_submodule_for_timecapsule_timeline_service_historical_event_proto_msgTypes,
	}.Build()
	File_submodule_for_timecapsule_timeline_service_historical_event_proto = out.File
	file_submodule_for_timecapsule_timeline_service_historical_event_proto_rawDesc = nil
	file_submodule_for_timecapsule_timeline_service_historical_event_proto_goTypes = nil
	file_submodule_for_timecapsule_timeline_service_historical_event_proto_depIdxs = nil
}

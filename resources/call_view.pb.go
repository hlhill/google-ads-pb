// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: google/ads/googleads/v11/resources/call_view.proto

package resources

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A call view that includes data for call tracking of call-only ads or call
// extensions.
type CallView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the call view.
	// Call view resource names have the form:
	//
	// `customers/{customer_id}/callViews/{call_detail_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Country code of the caller.
	CallerCountryCode string `protobuf:"bytes,2,opt,name=caller_country_code,json=callerCountryCode,proto3" json:"caller_country_code,omitempty"`
	// Output only. Area code of the caller. Null if the call duration is shorter than 15
	// seconds.
	CallerAreaCode string `protobuf:"bytes,3,opt,name=caller_area_code,json=callerAreaCode,proto3" json:"caller_area_code,omitempty"`
	// Output only. The advertiser-provided call duration in seconds.
	CallDurationSeconds int64 `protobuf:"varint,4,opt,name=call_duration_seconds,json=callDurationSeconds,proto3" json:"call_duration_seconds,omitempty"`
	// Output only. The advertiser-provided call start date time.
	StartCallDateTime string `protobuf:"bytes,5,opt,name=start_call_date_time,json=startCallDateTime,proto3" json:"start_call_date_time,omitempty"`
	// Output only. The advertiser-provided call end date time.
	EndCallDateTime string `protobuf:"bytes,6,opt,name=end_call_date_time,json=endCallDateTime,proto3" json:"end_call_date_time,omitempty"`
	// Output only. The call tracking display location.
	CallTrackingDisplayLocation enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation `protobuf:"varint,7,opt,name=call_tracking_display_location,json=callTrackingDisplayLocation,proto3,enum=google.ads.googleads.v11.enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation" json:"call_tracking_display_location,omitempty"`
	// Output only. The type of the call.
	Type enums.CallTypeEnum_CallType `protobuf:"varint,8,opt,name=type,proto3,enum=google.ads.googleads.v11.enums.CallTypeEnum_CallType" json:"type,omitempty"`
	// Output only. The status of the call.
	CallStatus enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus `protobuf:"varint,9,opt,name=call_status,json=callStatus,proto3,enum=google.ads.googleads.v11.enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus" json:"call_status,omitempty"`
}

func (x *CallView) Reset() {
	*x = CallView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v11_resources_call_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallView) ProtoMessage() {}

func (x *CallView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v11_resources_call_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallView.ProtoReflect.Descriptor instead.
func (*CallView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v11_resources_call_view_proto_rawDescGZIP(), []int{0}
}

func (x *CallView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CallView) GetCallerCountryCode() string {
	if x != nil {
		return x.CallerCountryCode
	}
	return ""
}

func (x *CallView) GetCallerAreaCode() string {
	if x != nil {
		return x.CallerAreaCode
	}
	return ""
}

func (x *CallView) GetCallDurationSeconds() int64 {
	if x != nil {
		return x.CallDurationSeconds
	}
	return 0
}

func (x *CallView) GetStartCallDateTime() string {
	if x != nil {
		return x.StartCallDateTime
	}
	return ""
}

func (x *CallView) GetEndCallDateTime() string {
	if x != nil {
		return x.EndCallDateTime
	}
	return ""
}

func (x *CallView) GetCallTrackingDisplayLocation() enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation {
	if x != nil {
		return x.CallTrackingDisplayLocation
	}
	return enums.CallTrackingDisplayLocationEnum_UNSPECIFIED
}

func (x *CallView) GetType() enums.CallTypeEnum_CallType {
	if x != nil {
		return x.Type
	}
	return enums.CallTypeEnum_UNSPECIFIED
}

func (x *CallView) GetCallStatus() enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus {
	if x != nil {
		return x.CallStatus
	}
	return enums.GoogleVoiceCallStatusEnum_UNSPECIFIED
}

var File_google_ads_googleads_v11_resources_call_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v11_resources_call_view_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x06, 0x0a, 0x08, 0x43, 0x61, 0x6c,
	0x6c, 0x56, 0x69, 0x65, 0x77, 0x12, 0x4e, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41,
	0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x61, 0x6c, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x13, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x15, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x13, 0x63,
	0x61, 0x6c, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x34, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x12, 0x65, 0x6e, 0x64, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x43, 0x61,
	0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0xa5, 0x01, 0x0a, 0x1e, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1b, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x75, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x5a, 0xea, 0x41, 0x57, 0x0a, 0x21,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6c, 0x6c, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x32, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x56,
	0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0xff, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x31, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x31, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v11_resources_call_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v11_resources_call_view_proto_rawDescData = file_google_ads_googleads_v11_resources_call_view_proto_rawDesc
)

func file_google_ads_googleads_v11_resources_call_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v11_resources_call_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v11_resources_call_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v11_resources_call_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v11_resources_call_view_proto_rawDescData
}

var file_google_ads_googleads_v11_resources_call_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v11_resources_call_view_proto_goTypes = []interface{}{
	(*CallView)(nil), // 0: google.ads.googleads.v11.resources.CallView
	(enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation)(0), // 1: google.ads.googleads.v11.enums.CallTrackingDisplayLocationEnum.CallTrackingDisplayLocation
	(enums.CallTypeEnum_CallType)(0),                                       // 2: google.ads.googleads.v11.enums.CallTypeEnum.CallType
	(enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus)(0),             // 3: google.ads.googleads.v11.enums.GoogleVoiceCallStatusEnum.GoogleVoiceCallStatus
}
var file_google_ads_googleads_v11_resources_call_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v11.resources.CallView.call_tracking_display_location:type_name -> google.ads.googleads.v11.enums.CallTrackingDisplayLocationEnum.CallTrackingDisplayLocation
	2, // 1: google.ads.googleads.v11.resources.CallView.type:type_name -> google.ads.googleads.v11.enums.CallTypeEnum.CallType
	3, // 2: google.ads.googleads.v11.resources.CallView.call_status:type_name -> google.ads.googleads.v11.enums.GoogleVoiceCallStatusEnum.GoogleVoiceCallStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v11_resources_call_view_proto_init() }
func file_google_ads_googleads_v11_resources_call_view_proto_init() {
	if File_google_ads_googleads_v11_resources_call_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v11_resources_call_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallView); i {
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
			RawDescriptor: file_google_ads_googleads_v11_resources_call_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v11_resources_call_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v11_resources_call_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v11_resources_call_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v11_resources_call_view_proto = out.File
	file_google_ads_googleads_v11_resources_call_view_proto_rawDesc = nil
	file_google_ads_googleads_v11_resources_call_view_proto_goTypes = nil
	file_google_ads_googleads_v11_resources_call_view_proto_depIdxs = nil
}

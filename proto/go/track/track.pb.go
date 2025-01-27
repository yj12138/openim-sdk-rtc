// Copyright 2023 LiveKit, Inc.
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
// 	protoc-gen-go v1.36.1
// 	protoc        v5.27.1
// source: track.proto

package track

import (
	e2ee "github.com/openimsdk/openim-rtc/proto/go/e2ee"
	handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	stats "github.com/openimsdk/openim-rtc/proto/go/stats"
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

type TrackKind int32

const (
	TrackKind_KIND_UNKNOWN TrackKind = 0
	TrackKind_KIND_AUDIO   TrackKind = 1
	TrackKind_KIND_VIDEO   TrackKind = 2
)

// Enum value maps for TrackKind.
var (
	TrackKind_name = map[int32]string{
		0: "KIND_UNKNOWN",
		1: "KIND_AUDIO",
		2: "KIND_VIDEO",
	}
	TrackKind_value = map[string]int32{
		"KIND_UNKNOWN": 0,
		"KIND_AUDIO":   1,
		"KIND_VIDEO":   2,
	}
)

func (x TrackKind) Enum() *TrackKind {
	p := new(TrackKind)
	*p = x
	return p
}

func (x TrackKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackKind) Descriptor() protoreflect.EnumDescriptor {
	return file_track_proto_enumTypes[0].Descriptor()
}

func (TrackKind) Type() protoreflect.EnumType {
	return &file_track_proto_enumTypes[0]
}

func (x TrackKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackKind.Descriptor instead.
func (TrackKind) EnumDescriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{0}
}

type TrackSource int32

const (
	TrackSource_SOURCE_UNKNOWN           TrackSource = 0
	TrackSource_SOURCE_CAMERA            TrackSource = 1
	TrackSource_SOURCE_MICROPHONE        TrackSource = 2
	TrackSource_SOURCE_SCREENSHARE       TrackSource = 3
	TrackSource_SOURCE_SCREENSHARE_AUDIO TrackSource = 4
)

// Enum value maps for TrackSource.
var (
	TrackSource_name = map[int32]string{
		0: "SOURCE_UNKNOWN",
		1: "SOURCE_CAMERA",
		2: "SOURCE_MICROPHONE",
		3: "SOURCE_SCREENSHARE",
		4: "SOURCE_SCREENSHARE_AUDIO",
	}
	TrackSource_value = map[string]int32{
		"SOURCE_UNKNOWN":           0,
		"SOURCE_CAMERA":            1,
		"SOURCE_MICROPHONE":        2,
		"SOURCE_SCREENSHARE":       3,
		"SOURCE_SCREENSHARE_AUDIO": 4,
	}
)

func (x TrackSource) Enum() *TrackSource {
	p := new(TrackSource)
	*p = x
	return p
}

func (x TrackSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackSource) Descriptor() protoreflect.EnumDescriptor {
	return file_track_proto_enumTypes[1].Descriptor()
}

func (TrackSource) Type() protoreflect.EnumType {
	return &file_track_proto_enumTypes[1]
}

func (x TrackSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackSource.Descriptor instead.
func (TrackSource) EnumDescriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{1}
}

type StreamState int32

const (
	StreamState_STATE_UNKNOWN StreamState = 0
	StreamState_STATE_ACTIVE  StreamState = 1
	StreamState_STATE_PAUSED  StreamState = 2
)

// Enum value maps for StreamState.
var (
	StreamState_name = map[int32]string{
		0: "STATE_UNKNOWN",
		1: "STATE_ACTIVE",
		2: "STATE_PAUSED",
	}
	StreamState_value = map[string]int32{
		"STATE_UNKNOWN": 0,
		"STATE_ACTIVE":  1,
		"STATE_PAUSED":  2,
	}
)

func (x StreamState) Enum() *StreamState {
	p := new(StreamState)
	*p = x
	return p
}

func (x StreamState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamState) Descriptor() protoreflect.EnumDescriptor {
	return file_track_proto_enumTypes[2].Descriptor()
}

func (StreamState) Type() protoreflect.EnumType {
	return &file_track_proto_enumTypes[2]
}

func (x StreamState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamState.Descriptor instead.
func (StreamState) EnumDescriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{2}
}

// Create a new VideoTrack from a VideoSource
type CreateVideoTrackRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	SourceHandle  uint64                 `protobuf:"varint,2,opt,name=source_handle,json=sourceHandle,proto3" json:"source_handle"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVideoTrackRequest) Reset() {
	*x = CreateVideoTrackRequest{}
	mi := &file_track_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVideoTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoTrackRequest) ProtoMessage() {}

func (x *CreateVideoTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoTrackRequest.ProtoReflect.Descriptor instead.
func (*CreateVideoTrackRequest) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVideoTrackRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVideoTrackRequest) GetSourceHandle() uint64 {
	if x != nil {
		return x.SourceHandle
	}
	return 0
}

type CreateVideoTrackResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Track         *OwnedTrack            `protobuf:"bytes,1,opt,name=track,proto3" json:"track"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVideoTrackResponse) Reset() {
	*x = CreateVideoTrackResponse{}
	mi := &file_track_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVideoTrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoTrackResponse) ProtoMessage() {}

func (x *CreateVideoTrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoTrackResponse.ProtoReflect.Descriptor instead.
func (*CreateVideoTrackResponse) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVideoTrackResponse) GetTrack() *OwnedTrack {
	if x != nil {
		return x.Track
	}
	return nil
}

// Create a new AudioTrack from a AudioSource
type CreateAudioTrackRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	SourceHandle  uint64                 `protobuf:"varint,2,opt,name=source_handle,json=sourceHandle,proto3" json:"source_handle"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAudioTrackRequest) Reset() {
	*x = CreateAudioTrackRequest{}
	mi := &file_track_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAudioTrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAudioTrackRequest) ProtoMessage() {}

func (x *CreateAudioTrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAudioTrackRequest.ProtoReflect.Descriptor instead.
func (*CreateAudioTrackRequest) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAudioTrackRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAudioTrackRequest) GetSourceHandle() uint64 {
	if x != nil {
		return x.SourceHandle
	}
	return 0
}

type CreateAudioTrackResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Track         *OwnedTrack            `protobuf:"bytes,1,opt,name=track,proto3" json:"track"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAudioTrackResponse) Reset() {
	*x = CreateAudioTrackResponse{}
	mi := &file_track_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAudioTrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAudioTrackResponse) ProtoMessage() {}

func (x *CreateAudioTrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAudioTrackResponse.ProtoReflect.Descriptor instead.
func (*CreateAudioTrackResponse) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAudioTrackResponse) GetTrack() *OwnedTrack {
	if x != nil {
		return x.Track
	}
	return nil
}

type GetStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TrackHandle   uint64                 `protobuf:"varint,1,opt,name=track_handle,json=trackHandle,proto3" json:"track_handle"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	mi := &file_track_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{4}
}

func (x *GetStatsRequest) GetTrackHandle() uint64 {
	if x != nil {
		return x.TrackHandle
	}
	return 0
}

type GetStatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AsyncId       uint64                 `protobuf:"varint,1,opt,name=async_id,json=asyncId,proto3" json:"async_id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	mi := &file_track_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{5}
}

func (x *GetStatsResponse) GetAsyncId() uint64 {
	if x != nil {
		return x.AsyncId
	}
	return 0
}

type GetStatsCallback struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AsyncId       uint64                 `protobuf:"varint,1,opt,name=async_id,json=asyncId,proto3" json:"async_id"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
	Stats         []*stats.RtcStats      `protobuf:"bytes,3,rep,name=stats,proto3" json:"stats"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatsCallback) Reset() {
	*x = GetStatsCallback{}
	mi := &file_track_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsCallback) ProtoMessage() {}

func (x *GetStatsCallback) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsCallback.ProtoReflect.Descriptor instead.
func (*GetStatsCallback) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{6}
}

func (x *GetStatsCallback) GetAsyncId() uint64 {
	if x != nil {
		return x.AsyncId
	}
	return 0
}

func (x *GetStatsCallback) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetStatsCallback) GetStats() []*stats.RtcStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type TrackEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackEvent) Reset() {
	*x = TrackEvent{}
	mi := &file_track_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackEvent) ProtoMessage() {}

func (x *TrackEvent) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackEvent.ProtoReflect.Descriptor instead.
func (*TrackEvent) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{7}
}

type TrackPublicationInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Sid            string                 `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Kind           TrackKind              `protobuf:"varint,3,opt,name=kind,proto3,enum=livekit.proto.TrackKind" json:"kind"`
	Source         TrackSource            `protobuf:"varint,4,opt,name=source,proto3,enum=livekit.proto.TrackSource" json:"source"`
	Simulcasted    bool                   `protobuf:"varint,5,opt,name=simulcasted,proto3" json:"simulcasted"`
	Width          uint32                 `protobuf:"varint,6,opt,name=width,proto3" json:"width"`
	Height         uint32                 `protobuf:"varint,7,opt,name=height,proto3" json:"height"`
	MimeType       string                 `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type"`
	Muted          bool                   `protobuf:"varint,9,opt,name=muted,proto3" json:"muted"`
	Remote         bool                   `protobuf:"varint,10,opt,name=remote,proto3" json:"remote"`
	EncryptionType e2ee.EncryptionType    `protobuf:"varint,11,opt,name=encryption_type,json=encryptionType,proto3,enum=livekit.proto.EncryptionType" json:"encryption_type"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TrackPublicationInfo) Reset() {
	*x = TrackPublicationInfo{}
	mi := &file_track_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackPublicationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackPublicationInfo) ProtoMessage() {}

func (x *TrackPublicationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackPublicationInfo.ProtoReflect.Descriptor instead.
func (*TrackPublicationInfo) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{8}
}

func (x *TrackPublicationInfo) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *TrackPublicationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrackPublicationInfo) GetKind() TrackKind {
	if x != nil {
		return x.Kind
	}
	return TrackKind_KIND_UNKNOWN
}

func (x *TrackPublicationInfo) GetSource() TrackSource {
	if x != nil {
		return x.Source
	}
	return TrackSource_SOURCE_UNKNOWN
}

func (x *TrackPublicationInfo) GetSimulcasted() bool {
	if x != nil {
		return x.Simulcasted
	}
	return false
}

func (x *TrackPublicationInfo) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *TrackPublicationInfo) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TrackPublicationInfo) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *TrackPublicationInfo) GetMuted() bool {
	if x != nil {
		return x.Muted
	}
	return false
}

func (x *TrackPublicationInfo) GetRemote() bool {
	if x != nil {
		return x.Remote
	}
	return false
}

func (x *TrackPublicationInfo) GetEncryptionType() e2ee.EncryptionType {
	if x != nil {
		return x.EncryptionType
	}
	return e2ee.EncryptionType(0)
}

type OwnedTrackPublication struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Handle        *handle.FfiOwnedHandle `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle"`
	Info          *TrackPublicationInfo  `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OwnedTrackPublication) Reset() {
	*x = OwnedTrackPublication{}
	mi := &file_track_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OwnedTrackPublication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnedTrackPublication) ProtoMessage() {}

func (x *OwnedTrackPublication) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnedTrackPublication.ProtoReflect.Descriptor instead.
func (*OwnedTrackPublication) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{9}
}

func (x *OwnedTrackPublication) GetHandle() *handle.FfiOwnedHandle {
	if x != nil {
		return x.Handle
	}
	return nil
}

func (x *OwnedTrackPublication) GetInfo() *TrackPublicationInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type TrackInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sid           string                 `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Kind          TrackKind              `protobuf:"varint,3,opt,name=kind,proto3,enum=livekit.proto.TrackKind" json:"kind"`
	StreamState   StreamState            `protobuf:"varint,4,opt,name=stream_state,json=streamState,proto3,enum=livekit.proto.StreamState" json:"stream_state"`
	Muted         bool                   `protobuf:"varint,5,opt,name=muted,proto3" json:"muted"`
	Remote        bool                   `protobuf:"varint,6,opt,name=remote,proto3" json:"remote"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackInfo) Reset() {
	*x = TrackInfo{}
	mi := &file_track_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackInfo) ProtoMessage() {}

func (x *TrackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackInfo.ProtoReflect.Descriptor instead.
func (*TrackInfo) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{10}
}

func (x *TrackInfo) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *TrackInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrackInfo) GetKind() TrackKind {
	if x != nil {
		return x.Kind
	}
	return TrackKind_KIND_UNKNOWN
}

func (x *TrackInfo) GetStreamState() StreamState {
	if x != nil {
		return x.StreamState
	}
	return StreamState_STATE_UNKNOWN
}

func (x *TrackInfo) GetMuted() bool {
	if x != nil {
		return x.Muted
	}
	return false
}

func (x *TrackInfo) GetRemote() bool {
	if x != nil {
		return x.Remote
	}
	return false
}

type OwnedTrack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Handle        *handle.FfiOwnedHandle `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle"`
	Info          *TrackInfo             `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OwnedTrack) Reset() {
	*x = OwnedTrack{}
	mi := &file_track_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OwnedTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnedTrack) ProtoMessage() {}

func (x *OwnedTrack) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnedTrack.ProtoReflect.Descriptor instead.
func (*OwnedTrack) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{11}
}

func (x *OwnedTrack) GetHandle() *handle.FfiOwnedHandle {
	if x != nil {
		return x.Handle
	}
	return nil
}

func (x *OwnedTrack) GetInfo() *TrackInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_track_proto protoreflect.FileDescriptor

var file_track_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x65, 0x32,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x4b, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x22, 0x52, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x64, 0x69, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x4b, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x05,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x74, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x0c,
	0x0a, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x81, 0x03, 0x0a,
	0x14, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x63, 0x61, 0x73, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x63, 0x61, 0x73, 0x74, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75,
	0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x87, 0x01, 0x0a, 0x15, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x66, 0x69, 0x4f, 0x77,
	0x6e, 0x65, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x09, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x75, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x74, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x71, 0x0a, 0x0a, 0x4f, 0x77, 0x6e,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x66, 0x69, 0x4f, 0x77, 0x6e, 0x65, 0x64,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x2a, 0x3d, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x02, 0x2a, 0x81, 0x01, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x45, 0x52, 0x41,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4d, 0x49, 0x43,
	0x52, 0x4f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10,
	0x03, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x43, 0x52, 0x45,
	0x45, 0x4e, 0x53, 0x48, 0x41, 0x52, 0x45, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x04, 0x2a,
	0x44, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x55,
	0x53, 0x45, 0x44, 0x10, 0x02, 0x42, 0x40, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x73, 0x64, 0x6b, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x69, 0x6d, 0x2d, 0x72, 0x74, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0xaa, 0x02, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_track_proto_rawDescOnce sync.Once
	file_track_proto_rawDescData = file_track_proto_rawDesc
)

func file_track_proto_rawDescGZIP() []byte {
	file_track_proto_rawDescOnce.Do(func() {
		file_track_proto_rawDescData = protoimpl.X.CompressGZIP(file_track_proto_rawDescData)
	})
	return file_track_proto_rawDescData
}

var file_track_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_track_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_track_proto_goTypes = []any{
	(TrackKind)(0),                   // 0: livekit.proto.TrackKind
	(TrackSource)(0),                 // 1: livekit.proto.TrackSource
	(StreamState)(0),                 // 2: livekit.proto.StreamState
	(*CreateVideoTrackRequest)(nil),  // 3: livekit.proto.CreateVideoTrackRequest
	(*CreateVideoTrackResponse)(nil), // 4: livekit.proto.CreateVideoTrackResponse
	(*CreateAudioTrackRequest)(nil),  // 5: livekit.proto.CreateAudioTrackRequest
	(*CreateAudioTrackResponse)(nil), // 6: livekit.proto.CreateAudioTrackResponse
	(*GetStatsRequest)(nil),          // 7: livekit.proto.GetStatsRequest
	(*GetStatsResponse)(nil),         // 8: livekit.proto.GetStatsResponse
	(*GetStatsCallback)(nil),         // 9: livekit.proto.GetStatsCallback
	(*TrackEvent)(nil),               // 10: livekit.proto.TrackEvent
	(*TrackPublicationInfo)(nil),     // 11: livekit.proto.TrackPublicationInfo
	(*OwnedTrackPublication)(nil),    // 12: livekit.proto.OwnedTrackPublication
	(*TrackInfo)(nil),                // 13: livekit.proto.TrackInfo
	(*OwnedTrack)(nil),               // 14: livekit.proto.OwnedTrack
	(*stats.RtcStats)(nil),           // 15: livekit.proto.RtcStats
	(e2ee.EncryptionType)(0),         // 16: livekit.proto.EncryptionType
	(*handle.FfiOwnedHandle)(nil),    // 17: livekit.proto.FfiOwnedHandle
}
var file_track_proto_depIdxs = []int32{
	14, // 0: livekit.proto.CreateVideoTrackResponse.track:type_name -> livekit.proto.OwnedTrack
	14, // 1: livekit.proto.CreateAudioTrackResponse.track:type_name -> livekit.proto.OwnedTrack
	15, // 2: livekit.proto.GetStatsCallback.stats:type_name -> livekit.proto.RtcStats
	0,  // 3: livekit.proto.TrackPublicationInfo.kind:type_name -> livekit.proto.TrackKind
	1,  // 4: livekit.proto.TrackPublicationInfo.source:type_name -> livekit.proto.TrackSource
	16, // 5: livekit.proto.TrackPublicationInfo.encryption_type:type_name -> livekit.proto.EncryptionType
	17, // 6: livekit.proto.OwnedTrackPublication.handle:type_name -> livekit.proto.FfiOwnedHandle
	11, // 7: livekit.proto.OwnedTrackPublication.info:type_name -> livekit.proto.TrackPublicationInfo
	0,  // 8: livekit.proto.TrackInfo.kind:type_name -> livekit.proto.TrackKind
	2,  // 9: livekit.proto.TrackInfo.stream_state:type_name -> livekit.proto.StreamState
	17, // 10: livekit.proto.OwnedTrack.handle:type_name -> livekit.proto.FfiOwnedHandle
	13, // 11: livekit.proto.OwnedTrack.info:type_name -> livekit.proto.TrackInfo
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_track_proto_init() }
func file_track_proto_init() {
	if File_track_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_track_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_track_proto_goTypes,
		DependencyIndexes: file_track_proto_depIdxs,
		EnumInfos:         file_track_proto_enumTypes,
		MessageInfos:      file_track_proto_msgTypes,
	}.Build()
	File_track_proto = out.File
	file_track_proto_rawDesc = nil
	file_track_proto_goTypes = nil
	file_track_proto_depIdxs = nil
}

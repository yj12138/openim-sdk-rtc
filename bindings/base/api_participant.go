package base

import (
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/openimsdk/openim-rtc/proto/go/e2ee"
	pb_ffi "github.com/openimsdk/openim-rtc/proto/go/ffi"
	pb_handle "github.com/openimsdk/openim-rtc/proto/go/handle"
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
	"github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) PublishTrack(req *pb_room.PublishTrackRequest) (*pb_room.PublishTrackResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	track := api.getLocalTrack(req.TrackHandle)
	go func() {
		pub, err := participant.PublishTrack(track, &lksdk.TrackPublicationOptions{})
		if err != nil {
			dispatchEvent(&pb_ffi.FfiEvent{
				Message: &pb_ffi.FfiEvent_PublishTrack{
					PublishTrack: &pb_room.PublishTrackCallback{
						AsyncId: asyncId,
						Message: &pb_room.PublishTrackCallback_Error{
							Error: err.Error(),
						},
					},
				},
			})
		} else {
			track.Publication = pub
			dispatchEvent(&pb_ffi.FfiEvent{
				Message: &pb_ffi.FfiEvent_PublishTrack{
					PublishTrack: &pb_room.PublishTrackCallback{
						AsyncId: asyncId,
						Message: &pb_room.PublishTrackCallback_Publication{
							Publication: &pb_track.OwnedTrackPublication{
								Handle: &pb_handle.FfiOwnedHandle{Id: api.storeObj(pub)},
								Info: &pb_track.TrackPublicationInfo{
									Sid:            pub.SID(),
									Name:           pub.Name(),
									Kind:           sdk.ConvertTrackKind(pub.Kind()),
									Source:         sdk.ConvertTrackSource(pub.Source()),
									Simulcasted:    false,
									Width:          pub.TrackInfo().Width,
									Height:         pub.TrackInfo().Height,
									MimeType:       pub.TrackInfo().MimeType,
									Muted:          pub.IsMuted(),
									Remote:         false,
									EncryptionType: e2ee.EncryptionType_NONE,
								},
							},
						},
					},
				},
			})
		}

	}()

	res := &pb_room.PublishTrackResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) UnpublishTrack(req *pb_room.UnpublishTrackRequest) (*pb_room.UnpublishTrackResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.UnpublishTrack(req.TrackSid)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_UnpublishTrack{
				UnpublishTrack: &pb_room.UnpublishTrackCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	// TODO req.StopOnUnpublish
	res := &pb_room.UnpublishTrackResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) PublishData(req *pb_room.PublishDataRequest) (*pb_room.PublishDataResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		data := cPointerToGoByteSliceNoCopyFunc(req.DataPtr, req.DataLen)
		err := participant.SendData(req.Topic, data, req.Reliable, req.DestinationIdentities)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_PublishData{
				PublishData: &pb_room.PublishDataCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.PublishDataResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) PublishSipDtmf(req *pb_room.PublishSipDtmfRequest) (*pb_room.PublishSipDtmfResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendSipDtmf(req.Code, req.Digit)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_PublishSipDtmf{
				PublishSipDtmf: &pb_room.PublishSipDtmfCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.PublishSipDtmfResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SetLocalMetadata(req *pb_room.SetLocalMetadataRequest) (*pb_room.SetLocalMetadataResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		participant.SetMetadata(req.Metadata)
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_SetLocalMetadata{
				SetLocalMetadata: &pb_room.SetLocalMetadataCallback{
					AsyncId: asyncId,
					Error:   "",
				},
			},
		})
	}()
	res := &pb_room.SetLocalMetadataResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SetLocalName(req *pb_room.SetLocalNameRequest) (*pb_room.SetLocalNameResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		participant.SetMetadata(req.Name)
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_SetLocalName{
				SetLocalName: &pb_room.SetLocalNameCallback{
					AsyncId: asyncId,
					Error:   "",
				},
			},
		})
	}()
	res := &pb_room.SetLocalNameResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SetLocalAttributes(req *pb_room.SetLocalAttributesRequest) (*pb_room.SetLocalAttributesResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		attributes := make(map[string]string)
		for _, entry := range req.Attributes {
			attributes[entry.Key] = entry.Value
		}
		participant.SetAttributes(attributes)
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_SetLocalAttributes{
				SetLocalAttributes: &pb_room.SetLocalAttributesCallback{
					AsyncId: asyncId,
					Error:   "",
				},
			},
		})
	}()
	res := &pb_room.SetLocalAttributesResponse{
		AsyncId: asyncId,
	}
	return res, nil
}

func (api *API) PublishTranscription(req *pb_room.PublishTranscriptionRequest) (*pb_room.PublishTranscriptionResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendTranscription(req.ParticipantIdentity, req.TrackId, req.Segments)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_PublishTranscription{
				PublishTranscription: &pb_room.PublishTranscriptionCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.PublishTranscriptionResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SendChatMessage(req *pb_room.SendChatMessageRequest) (*pb_room.SendChatMessageResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendChatMessage(req.SenderIdentity, req.DestinationIdentities, req.Message)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		if err != nil {
			dispatchEvent(&pb_ffi.FfiEvent{
				Message: &pb_ffi.FfiEvent_ChatMessage{
					ChatMessage: &pb_room.SendChatMessageCallback{
						AsyncId: asyncId,
						Message: &pb_room.SendChatMessageCallback_Error{
							Error: errStr,
						},
					},
				},
			})
		} else {
			dispatchEvent(&pb_ffi.FfiEvent{
				Message: &pb_ffi.FfiEvent_ChatMessage{
					ChatMessage: &pb_room.SendChatMessageCallback{
						AsyncId: asyncId,
						Message: &pb_room.SendChatMessageCallback_ChatMessage{
							ChatMessage: &pb_room.ChatMessage{
								Id:            req.SenderIdentity,
								Timestamp:     0,
								Message:       req.Message,
								EditTimestamp: 0,
								Deleted:       false,
								Generated:     false,
							},
						},
					},
				},
			})
		}

	}()
	res := &pb_room.SendChatMessageResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) EditChatMessage(req *pb_room.EditChatMessageRequest) (*pb_room.SendChatMessageResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.EditChatMessage(req.SenderIdentity, req.DestinationIdentities, req.EditText, req.OriginalMessage)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		if err != nil {
			dispatchEvent(&pb_ffi.FfiEvent{
				Message: &pb_ffi.FfiEvent_ChatMessage{
					ChatMessage: &pb_room.SendChatMessageCallback{
						AsyncId: asyncId,
						Message: &pb_room.SendChatMessageCallback_Error{
							Error: errStr,
						},
					},
				},
			})
		} else {
			dispatchEvent(&pb_ffi.FfiEvent{
				Message: &pb_ffi.FfiEvent_ChatMessage{
					ChatMessage: &pb_room.SendChatMessageCallback{
						AsyncId: asyncId,
						Message: &pb_room.SendChatMessageCallback_ChatMessage{
							ChatMessage: &pb_room.ChatMessage{
								Id:            req.SenderIdentity,
								Timestamp:     0,
								Message:       req.EditText,
								EditTimestamp: 0,
								Deleted:       false,
								Generated:     false,
							},
						},
					},
				},
			})
		}
	}()
	res := &pb_room.SendChatMessageResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SendTranscription(req *pb_room.PublishTranscriptionRequest) (*pb_room.PublishTranscriptionResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendTranscription(req.ParticipantIdentity, req.TrackId, req.Segments)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_PublishTranscription{
				PublishTranscription: &pb_room.PublishTranscriptionCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.PublishTranscriptionResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SendStreamHeader(req *pb_room.SendStreamHeaderRequest) (*pb_room.SendStreamHeaderResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendStreamHeader(req.SenderIdentity, req.DestinationIdentities, req.Header)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_SendStreamHeader{
				SendStreamHeader: &pb_room.SendStreamHeaderCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.SendStreamHeaderResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SendStreamChunk(req *pb_room.SendStreamChunkRequest) (*pb_room.SendStreamChunkResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendStreamChunk(req.SenderIdentity, req.DestinationIdentities, req.Chunk)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_SendStreamChunk{
				SendStreamChunk: &pb_room.SendStreamChunkCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.SendStreamChunkResponse{
		AsyncId: asyncId,
	}
	return res, nil
}
func (api *API) SendStreamTrailer(req *pb_room.SendStreamTrailerRequest) (*pb_room.SendStreamTrailerResponse, error) {
	asyncId := api.nextAsyncId()
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	go func() {
		err := participant.SendStreamTrailer(req.SenderIdentity, req.DestinationIdentities, req.Trailer)
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		dispatchEvent(&pb_ffi.FfiEvent{
			Message: &pb_ffi.FfiEvent_SendStreamTrailer{
				SendStreamTrailer: &pb_room.SendStreamTrailerCallback{
					AsyncId: asyncId,
					Error:   errStr,
				},
			},
		})
	}()
	res := &pb_room.SendStreamTrailerResponse{
		AsyncId: asyncId,
	}
	return res, nil
}

func (api *API) SetTrackSubscriptionPermissions(req *pb_track.SetTrackSubscriptionPermissionsRequest) (*pb_track.SetTrackSubscriptionPermissionsResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SetSubscriptionPermissionWrap(req.AllParticipantsAllowed, req.Permissions)
	res := &pb_track.SetTrackSubscriptionPermissionsResponse{}
	return res, nil
}

func (api *API) SetDataChannelBufferedAmountLowThreshold(req *pb_room.SetDataChannelBufferedAmountLowThresholdRequest) (*pb_room.SetDataChannelBufferedAmountLowThresholdResponse, error) {
	// TODO
	// participant := api.getLocalParticipant(req.LocalParticipantHandle)
	res := &pb_room.SetDataChannelBufferedAmountLowThresholdResponse{}
	return res, nil
}

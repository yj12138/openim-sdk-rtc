package base

import (
	pb_room "github.com/openimsdk/openim-rtc/proto/go/room"
	pb_track "github.com/openimsdk/openim-rtc/proto/go/track"
)

func (api *API) PublishTrack(req *pb_room.PublishTrackRequest) (*pb_room.PublishTrackResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	track := api.getLocalTrack(req.TrackHandle)
	participant.PublicTrack(track)
	res := &pb_room.PublishTrackResponse{}
	return res, nil
}
func (api *API) UnpublishTrack(req *pb_room.UnpublishTrackRequest) (*pb_room.UnpublishTrackResponse, error) {
	// participant := api.getLocalParticipant(req.LocalParticipantHandle)
	// track := api.getLocalTrack(req.TrackHandle)
	// TODO req.StopOnUnpublish
	// participant.PublicTrack(track)
	res := &pb_room.UnpublishTrackResponse{}
	return res, nil
}
func (api *API) PublishData(req *pb_room.PublishDataRequest) (*pb_room.PublishDataResponse, error) {
	// participant := api.getLocalParticipant(req.LocalParticipantHandle)
	// participant.SendData(req.Topic, req.Data, req.Reliable, req.Identifies)
	res := &pb_room.PublishDataResponse{}
	return res, nil
}

func (api *API) SetLocalMetadata(req *pb_room.SetLocalMetadataRequest) (*pb_room.SetLocalMetadataResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SetMetadata(req.Metadata)
	res := &pb_room.SetLocalMetadataResponse{}
	return res, nil
}
func (api *API) SetLocalName(req *pb_room.SetLocalNameRequest) (*pb_room.SetLocalNameResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SetMetadata(req.Name)
	res := &pb_room.SetLocalNameResponse{}
	return res, nil
}
func (api *API) SetLocalAttributes(req *pb_room.SetLocalAttributesRequest) (*pb_room.SetLocalAttributesResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	attributes := make(map[string]string)
	for _, entry := range req.Attributes {
		attributes[entry.Key] = entry.Value
	}
	participant.SetAttributes(attributes)
	res := &pb_room.SetLocalAttributesResponse{}
	return res, nil
}

func (api *API) PublishTranscription(req *pb_room.PublishTranscriptionRequest) (*pb_room.PublishTranscriptionResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendTranscription(req.ParticipantIdentity, req.TrackId, req.Segments)
	res := &pb_room.PublishTranscriptionResponse{}
	return res, nil
}
func (api *API) SendChatMessage(req *pb_room.SendChatMessageRequest) (*pb_room.SendChatMessageResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendChatMessage(req.SenderIdentity, req.DestinationIdentities, req.Message)
	res := &pb_room.SendChatMessageResponse{}
	return res, nil
}
func (api *API) EditChatMessage(req *pb_room.EditChatMessageRequest) (*pb_room.EditChatMessageRequest, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.EditChatMessage(req.SenderIdentity, req.DestinationIdentities, req.EditText, req.OriginalMessage)
	return req, nil
}
func (api *API) SendTranscription(req *pb_room.PublishTranscriptionRequest) (*pb_room.PublishTranscriptionResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendTranscription(req.ParticipantIdentity, req.TrackId, req.Segments)
	res := &pb_room.PublishTranscriptionResponse{}
	return res, nil
}
func (api *API) SendStreamHeader(req *pb_room.SendStreamHeaderRequest) (*pb_room.SendStreamHeaderResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendStreamHeader(req.SenderIdentity, req.DestinationIdentities, req.Header)
	res := &pb_room.SendStreamHeaderResponse{}
	return res, nil
}
func (api *API) SendStreamChunk(req *pb_room.SendStreamChunkRequest) (*pb_room.SendStreamChunkResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendStreamChunk(req.SenderIdentity, req.DestinationIdentities, req.Chunk)
	res := &pb_room.SendStreamChunkResponse{}
	return res, nil
}
func (api *API) SendStreamTrailer(req *pb_room.SendStreamTrailerRequest) (*pb_room.SendStreamTrailerResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendStreamTrailer(req.SenderIdentity, req.DestinationIdentities, req.Trailer)
	res := &pb_room.SendStreamTrailerResponse{}
	return res, nil
}

func (api *API) SetTrackSubscriptionPermissions(req *pb_track.SetTrackSubscriptionPermissionsRequest) (*pb_track.SetTrackSubscriptionPermissionsResponse, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SetSubscriptionPermission(req.AllParticipantsAllowed, req.Permissions)
	res := &pb_track.SetTrackSubscriptionPermissionsResponse{}
	return res, nil
}

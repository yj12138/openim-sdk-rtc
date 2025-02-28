package base

import (
	pb_participant "github.com/openimsdk/openim-rtc/proto/go/participant"
	// "github.com/openimsdk/openim-rtc/sdk"
)

func (api *API) PublishTrack(req *pb_participant.PublishTrackReq) (*pb_participant.PublishTrackRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	track := api.getLocalTrack(req.TrackHandle)
	participant.PublicTrack(track)
	res := &pb_participant.PublishTrackRes{}
	return res, nil
}
func (api *API) UnpublishTrack(req *pb_participant.UnpublishTrackReq) (*pb_participant.UnpublishTrackRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	track := api.getLocalTrack(req.TrackHandle)
	// TODO req.StopOnUnpublish
	participant.PublicTrack(track)
	res := &pb_participant.UnpublishTrackRes{}
	return res, nil
}
func (api *API) PublishData(req *pb_participant.PublishDataReq) (*pb_participant.PublishDataRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendData(req.Topic, req.Data, req.Reliable, req.Identifies)
	res := &pb_participant.PublishDataRes{}
	return res, nil
}
func (api *API) SetSubscribed(req *pb_participant.SetSubscribedReq) (*pb_participant.SetSubscribedRes, error) {
	participant := api.getRemoteParticipant(req.RemoteParticipantHandle)
	participant.SetSubscribed(req.PublicationSid, req.Subscribe)
	return nil, nil
}
func (api *API) SetLocalMetadata(req *pb_participant.SetLocalMetadataReq) (*pb_participant.SetLocalMetadataRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SetMetadata(req.Metadata)
	res := &pb_participant.SetLocalMetadataRes{}
	return res, nil
}
func (api *API) SetLocalName(req *pb_participant.SetLocalNameReq) (*pb_participant.SetLocalNameRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SetMetadata(req.Name)
	res := &pb_participant.SetLocalNameRes{}
	return res, nil
}
func (api *API) SetLocalAttributes(req *pb_participant.SetLocalAttributesReq) (*pb_participant.SetLocalAttributesRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)

	attributes := make(map[string]string)
	for _, entry := range req.Attributes {
		attributes[entry.Key] = entry.Value
	}
	participant.SetAttributes(attributes)

	res := &pb_participant.SetLocalAttributesRes{}
	return res, nil
}
func (api *API) EnableRemoteTrackPublication(req *pb_participant.EnableRemoteTrackPublicationReq) (*pb_participant.EnableRemoteTrackPublicationRes, error) {

	return nil, nil
}

func (api *API) UpdateRemoteTrackPublicationDimension(req *pb_participant.UpdateRemoteTrackPublicationDimensionReq) (*pb_participant.UpdateRemoteTrackPublicationDimensionRes, error) {

	return nil, nil
}

func (api *API) PublishTranscription(req *pb_participant.PublishTranscriptionReq) (*pb_participant.PublishTranscriptionRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendTranscription(req.ParticipantIdentify, req.TrackId, req.Segments)
	res := &pb_participant.PublishTranscriptionRes{}
	return res, nil
}
func (api *API) SendChatMessage(req *pb_participant.SendChatMessageReq) (*pb_participant.SendChatMessageRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendChatMessage(req.SenderIdentity, req.DestinationIdentities, req.Message)
	res := &pb_participant.SendChatMessageRes{}
	return res, nil
}
func (api *API) EditChatMessage(req *pb_participant.EditChatMessageReq) (*pb_participant.EditChatMessageRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.EditChatMessage(req.SenderIdentity, req.DestinationIdentities, req.EditText, req.OriginalMessage)
	res := &pb_participant.EditChatMessageRes{}
	return res, nil
}
func (api *API) SendTranscription(req *pb_participant.PublishTranscriptionReq) (*pb_participant.PublishTranscriptionRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendTranscription(req.ParticipantIdentify, req.TrackId, req.Segments)
	res := &pb_participant.PublishTranscriptionRes{}
	return res, nil
}
func (api *API) SendStreamHeader(req *pb_participant.SendStreamHeaderReq) (*pb_participant.SendStreamHeaderRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendStreamHeader(req.SenderIdentity, req.DestinationIdentities, req.Header)
	res := &pb_participant.SendStreamHeaderRes{}
	return res, nil
}
func (api *API) SendStreamChunk(req *pb_participant.SendStreamChunkReq) (*pb_participant.SendStreamChunkRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendStreamChunk(req.SenderIdentity, req.DestinationIdentities, req.Chunk)
	res := &pb_participant.SendStreamChunkRes{}
	return res, nil
}
func (api *API) SendStreamTrailer(req *pb_participant.SendStreamTrailerReq) (*pb_participant.SendStreamTrailerRes, error) {
	participant := api.getLocalParticipant(req.LocalParticipantHandle)
	participant.SendStreamTrailer(req.SenderIdentity, req.DestinationIdentities, req.Trailer)
	res := &pb_participant.SendStreamTrailerRes{}
	return res, nil
}
func (api *API) SetDataChannelBufferedAmountLowThreshold(req *pb_participant.SetDataChannelBufferedAmountLowThresholdReq) (*pb_participant.SetDataChannelBufferedAmountLowThresholdRes, error) {
	return nil, nil
}

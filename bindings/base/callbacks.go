package base

type RoomListener struct {
}

func (l *RoomListener) OnDisconnected() {

}
func (l *RoomListener) OnDisconnectedWithReason(reason string) {

}
func (l *RoomListener) OnParticipantConnected() {

}
func (l *RoomListener) OnParticipantDisconnected() {

}
func (l *RoomListener) OnActiveSpeakersChanged() {

}
func (l *RoomListener) OnRoomMetadataChanged(metadata string) {

}
func (l *RoomListener) OnReconnecting() {

}
func (l *RoomListener) OnReconnected() {

}

func NewRoomListener() *RoomListener {
	return &RoomListener{}
}

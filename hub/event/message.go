package event

import (
	"github.com/sniper365/Go-gRPC-K8S/common/log"
	"github.com/sniper365/Go-gRPC-K8S/hub"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
)

func init() {
	RegisterEvent(&proto.HubMessage_EventMessage{}, eventMessage)
	RegisterEvent(&proto.HubMessage_EventMessage{}, eventMessage)
}

func eventMessage(conn *hub.Conn, message *proto.HubMessage) {
	log.Debugf("recvd message from %d: %s", conn.UserId, message.GetEventMessage().Message)
	conn.Parent.Broadcast(message)
}

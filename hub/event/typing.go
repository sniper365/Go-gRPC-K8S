package event

import (
	"github.com/sniper365/Go-gRPC-K8S/hub"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
)

func init() {
	RegisterEvent(&proto.HubMessage_EventTyping{}, eventTyping)
}

func eventTyping(conn *hub.Conn, message *proto.HubMessage) {

}

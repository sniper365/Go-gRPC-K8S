package event

import (
	"github.com/sniper365/Go-gRPC-K8S/hub"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
)

func init() {
	RegisterEvent(&proto.HubMessage_EventUserAdded{}, eventUserAdded)
}

func eventUserAdded(conn *hub.Conn, message *proto.HubMessage) {

}

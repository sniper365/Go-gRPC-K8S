package event

import (
	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
)

func init() {
	RegisterEvent(&proto.HubMessage_EventUserRemoved{}, eventUserRemoved)
}

func eventUserRemoved(conn *hub.Conn, message *proto.HubMessage) {

}

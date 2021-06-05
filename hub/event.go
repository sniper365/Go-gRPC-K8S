package hub

import (
	"reflect"

	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
)

type EventFunc func(conn *Conn, message *proto.HubMessage)
type Handler map[reflect.Type]EventFunc

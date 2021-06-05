package stream

import (
	"io"

	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
)

type Stream interface {
	Recv() (*proto.HubMessage, error)
	Send(msg *proto.HubMessage) error
	io.Closer
}

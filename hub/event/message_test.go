package event

import (
	"reflect"
	"testing"

	"time"

	"github.com/sniper365/Go-gRPC-K8S/common/test"
	"github.com/sniper365/Go-gRPC-K8S/hub"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
	"github.com/sniper365/Go-gRPC-K8S/hub/stream"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	msg := &proto.HubMessage{
		Event: &proto.HubMessage_EventMessage{
			EventMessage: &proto.EventMessage{
				Message: "hello",
			},
		},
	}

	h := hub.New(nil)
	go h.Start()
	m := stream.NewMock()
	c := hub.NewConn(h, m)
	h.Register(c)

	GetHandler()[reflect.TypeOf(&proto.HubMessage_EventMessage{})](c, msg)

	assert.True(t, test.RanWithinTimeout(100*time.Millisecond, func() {
		assert.Equal(t, msg, <-m.SendBuf)
	}))
}

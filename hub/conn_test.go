package hub

import (
	"fmt"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/sniper365/Go-gRPC-K8S/common/test"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
	"github.com/sniper365/Go-gRPC-K8S/hub/stream"
	"github.com/stretchr/testify/require"
)

func TestCloseStopsLoops(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)

	mock := stream.NewMock()
	c := NewConn(nil, mock)
	time.AfterFunc(100*time.Millisecond, func() {
		mock.Close()
		time.AfterFunc(500*time.Millisecond, func() {
			a.FailNow("failed to get msg")
		})
	})
	c.Loop()
}

func TestConnSendsToChannel(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)
	mock := stream.NewMock()

	c := NewConn(nil, mock)
	msg := &proto.HubMessage{}
	c.Send(msg)

	a.True(test.RanWithinTimeout(100*time.Millisecond, func() {
		a.Equal(msg, <-mock.SendBuf)
	}))
}

func TestConnErrorInSendEndsConn(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)
	mock := stream.NewMock()
	c := NewConn(nil, mock)

	mock.SendErr = fmt.Errorf("send error")
	a.Equal(mock.SendErr, c.Send(nil))

}

func TestConnRecvs(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)
	mock := stream.NewMock()
	c := NewConn(New(Handler{}), mock)

	msg := &proto.HubMessage{}
	mock.RecvBuf <- msg
	a.Len(mock.RecvBuf, 1)
	go c.Loop()

	a.True(test.EqualsWithinTimeout(50*time.Millisecond, func() bool {
		return len(mock.RecvBuf) == 0
	}))
	c.Close()
}

package service

import (
	"github.com/sniper365/Go-gRPC-K8S/common/log"
	"github.com/sniper365/Go-gRPC-K8S/hub"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
	"github.com/sniper365/Go-gRPC-K8S/hub/stream"
	"google.golang.org/grpc/peer"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/hub.proto

type HubService struct {
	hub *hub.Hub
}

func (h *HubService) Connect(s proto.Hub_ConnectServer) error {
	conn := hub.NewConn(h.hub, stream.NewGrpc(s))
	p, _ := peer.FromContext(s.Context())
	if p != nil {
		log.Debug("User connected to grpc server ", p.Addr.String())
	}

	h.hub.Register(conn)
	conn.Loop()
	h.hub.Deregister(conn)
	return nil
}

func New(hub *hub.Hub) proto.HubServer {
	return &HubService{
		hub: hub,
	}
}

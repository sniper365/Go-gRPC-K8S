package stream

import "github.com/sniper365/Go-gRPC-K8S/hub/proto"

type grpcStream struct {
	proto.Hub_ConnectServer
}

func NewGrpc(s proto.Hub_ConnectServer) Stream {
	return &grpcStream{s}
}

func (g *grpcStream) Close() error {
	return nil
}

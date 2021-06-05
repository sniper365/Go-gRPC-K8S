package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/sniper365/Go-gRPC-K8S/common/log"
	libservice "github.com/sniper365/Go-gRPC-K8S/common/service"
	"github.com/sniper365/Go-gRPC-K8S/hub"
	"github.com/sniper365/Go-gRPC-K8S/hub/event"
	"github.com/sniper365/Go-gRPC-K8S/hub/proto"
	"github.com/sniper365/Go-gRPC-K8S/hub/service"
	"github.com/sniper365/Go-gRPC-K8S/hub/version"
)

func init() {
	log.Infof("%s %s %s %s", version.Name, version.Version, version.BuildTime, version.Commit)
}

func main() {
	v := flag.Bool("v", false, version.Name+" version")
	laddr := flag.String("laddr", "0.0.0.0:8080", version.Name+" version")
	flag.Parse()

	if *v {
		// version is printed in init()
		return
	}

	parts := strings.Split(*laddr, ":")
	port, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	rpc := libservice.NewGrpcService(libservice.Options{
		ServiceName:    version.Name,
		ServiceVersion: version.Version,
		MetricsAddr:    "0.0.0.0:8181",
		IP:             parts[0],
		Port:           int(port),
	})

	h := hub.New(event.GetHandler())
	go h.Start()
	proto.RegisterHubServer(rpc.GrpcServer(), service.New(h))
	if err = rpc.Serve(); err != nil {
		log.Fatal(err)
	}
}

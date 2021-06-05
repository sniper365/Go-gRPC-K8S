package main

import (
	"flag"
	"strconv"
	"strings"

	"os"

	"github.com/sniper365/Go-gRPC-K8S/common/log"
	libservice "github.com/sniper365/Go-gRPC-K8S/common/service"
	proto "github.com/sniper365/Go-gRPC-K8S/users/proto"
	"github.com/sniper365/Go-gRPC-K8S/users/repo"
	"github.com/sniper365/Go-gRPC-K8S/users/service"
	"github.com/sniper365/Go-gRPC-K8S/users/version"
)

func init() {
	log.Infof("%s %s %s %s", version.Name, version.Version, version.BuildTime, version.Commit)
}

const (
	DbHost    = "DB_HOST"
	DnUser    = "DB_USER"
	DbPass    = "DB_PASS"
	DbName    = "DB_NAME"
	JwtSecret = "JWT_SECRET"
)

var (
	v     = flag.Bool("v", false, version.Name+" version")
	laddr = flag.String("laddr", ":8080", version.Name+" version")
)

func main() {
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
		MetricsAddr:    ":8181",
		IP:             parts[0],
		Port:           int(port),
	})

	tkn := service.NewTokenService(os.Getenv(JwtSecret))

	db, err := repo.CreateConnection(os.Getenv(DbHost), os.Getenv(DnUser), os.Getenv(DbPass), os.Getenv(DbName))
	if err != nil {
		log.Fatal("Unable to connect to db", db)
	}
	defer db.Close()

	srvc := service.New(repo.NewMySql(db), tkn)

	proto.RegisterUsersServer(rpc.GrpcServer(), srvc)
	log.Fatal(rpc.Serve())
}

package main

import (
	"fmt"
	"net"

	"github.com/cagedmantis/sabre/toothapi"
	envflag "github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	defaultPort = 8010
)

func main() {
	log.Info("starting tooth server...")

	var binDir string
	var port int
	envflag.StringVar(&binDir, "toothserver_bin_dir", "/data/toothserver", "path to bin directory")
	envflag.IntVar(&port, "toothserver_port", defaultPort, "port listening on")

	envflag.Parse()

	log.WithFields(log.Fields{
		"bin.directory": binDir,
	}).Info("config settings")

	//log.Info("exiting tooth server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc := grpc.NewServer()
	tapis := toothapi.NewToothAPI()
	toothapi.RegisterToothAPIServer(grpc, tapis)

	log.Fatal(grpc.Serve(lis))
}

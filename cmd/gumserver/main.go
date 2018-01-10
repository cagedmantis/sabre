package main

import (
	"fmt"
	"net"

	"github.com/cagedmantis/sabre/chunk/store"
	"github.com/cagedmantis/sabre/toothapi"

	envflag "github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	defaultPort = 8011
)

func main() {
	log.Info("starting gum server...")

	var (
		port int
	)

	envflag.IntVar(&port, "gumserver_port", defaultPort, "port listening on")
	envflag.Parse()

	// log.WithFields(log.Fields{
	// 	"bin.directory": binDir,
	// }).Info("config settings")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create chunk store
	// add chunk store to server

	grpc := grpc.NewServer()

	s, err := store.NewDiskStore(binDir)
	if err != nil {
		log.Fatalf("unable to create store: %s", err)
	}
	tapis := toothapi.NewToothAPI(s)

	toothapi.RegisterToothAPIServer(grpc, tapis)

	log.Fatal(grpc.Serve(lis))
}
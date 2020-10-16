package main

import (
	"net"
	"net/http"
	"os"

	contrl "github.com/EurosportDigital/gps-decengine/controllers"
	pb "github.com/EurosportDigital/gps-decengine/model"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type ctx struct {
	id   int
	name string
}

const (
	port = ":9090"
)

// These exist so they can be mocked in tests.
var getHostname func() (string, error) = os.Hostname
var startServer func(*http.Server) error = func(srv *http.Server) error {
	return srv.ListenAndServe()
}

func main() {
	// Setup gRPC server
	lis, errListen := net.Listen("tcp", port)
	if errListen != nil {
		log.Print(errListen)
	}
	s := grpc.NewServer()
	pb.RegisterDecisioningEngineServer(s, contrl.Server{}) // call function getServerStruct
	if err := s.Serve(lis); err != nil {
		log.Print(errListen)
	}
}

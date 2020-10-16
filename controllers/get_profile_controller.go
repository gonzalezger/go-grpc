package controllers

import (
	"context"

	pb "github.com/EurosportDigital/gps-decengine/model"
)

type Server struct {
	*server
}
type server struct {
	pb.UnimplementedDecisioningEngineServer
}

func (s *server) GetBestProfiles(ctx context.Context, in *pb.ClientInfo) (*pb.Profile, error) {
	return &pb.Profile{PrimaryProfile: "hls_h264", FallbackProfile: "hls_h264"}, nil
}

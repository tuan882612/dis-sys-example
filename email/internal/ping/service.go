package ping

import (
	"context"
	"fmt"

	"dissys/internal/proto/pb/pingpb"
)

type Service struct {
	pingpb.UnimplementedPingServiceServer
}

func New() pingpb.PingServiceServer {
	return &Service{}
}

func (s *Service) Ping(ctx context.Context, pingData *pingpb.PingData) (*pingpb.PingData, error) {
	return &pingpb.PingData{
		Message: fmt.Sprintf("Origin: %s - successfull ping", pingData.GetMessage()),
	}, nil
}

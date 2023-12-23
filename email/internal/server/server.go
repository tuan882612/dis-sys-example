package server

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"dissys/internal/deps"
	"dissys/internal/email"
	"dissys/internal/ping"
	"dissys/internal/proto/pb/authpb"
	"dissys/internal/proto/pb/pingpb"
)

type Server struct {
	svr  *grpc.Server
	deps *deps.Dependencies
	addr string
}

func New(d *deps.Dependencies) *Server {
	log.Info().Msg("initializing server...")
	return &Server{
		svr:  grpc.NewServer(),
		addr: d.Config.Server.Host + ":" + d.Config.Server.Port,
		deps: d,
	}
}

func (s *Server) setup() {
	log.Info().Msg("registering server services...")

	emailSvc := email.New(s.deps)
	authpb.RegisterAuthServiceServer(s.svr, emailSvc)

	pingSvc := ping.New()
	pingpb.RegisterPingServiceServer(s.svr, pingSvc)
}

func (s *Server) Start() error {
	s.setup()

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Error().Msgf("failed to listen: %v", err)
		return err
	}
	
	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := s.svr.Serve(lis); err != nil {
		log.Error().Msgf("failed to serve: %v", err)
		return err
	}
	
	return nil
}

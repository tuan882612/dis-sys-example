package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"dissys/internal/deps"
	"dissys/internal/email"
	"dissys/internal/ping"
	"dissys/internal/proto/pb/emailpb"
	"dissys/internal/proto/pb/pingpb"
)

type Server struct {
	svr  *grpc.Server
	deps *deps.Dependencies
	addr string
}

func New(d *deps.Dependencies) (*Server, error) {
	log.Info().Msg("initializing server...")

	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Error().Err(err).Msg("failed to read ca.crt")
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		err := errors.New("failed to append ca.crt to cert pool")
		log.Error().Err(err).Msg("failed server initialization")
		return nil, err
	}

	serverCert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Error().Err(err).Msg("failed to load server.crt and server.key")
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
	}

	return &Server{
		svr:  grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig))),
		addr: d.Config.Host + ":" + d.Config.Port,
		deps: d,
	}, nil
}

func (s *Server) setup() error {
	log.Info().Msg("registering server services...")

	emailSvc := email.NewService(s.deps)
	emailpb.RegisterEmailServiceServer(s.svr, emailSvc)

	pingSvc := ping.New()
	pingpb.RegisterPingServiceServer(s.svr, pingSvc)

	return nil
}

func (s *Server) Start() error {
	if err := s.setup(); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return err
	}

	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := s.svr.Serve(lis); err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return err
	}

	return nil
}

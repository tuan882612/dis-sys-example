package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"dissys/internal/deps"
	"dissys/internal/proto"
	"dissys/internal/proto/pb/pingpb"
	"dissys/internal/server/middlewares"
	"dissys/internal/server/routes"
)

type Server struct {
	svr *echo.Echo
	d   *deps.Dependencies
}

func New(d *deps.Dependencies) *Server {
	log.Info().Msg("initializing server...")
	return &Server{
		svr: echo.New(),
		d:   d,
	}
}

func (s *Server) setup() error {
	log.Info().Msg("setting up server...")

	// setup middlewares
	middlewares.SetCORS(s.svr)
	s.svr.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(15)))

	// setup routes
	g := s.svr.Group("/api/v1")
	routes.SetBaseRoutes(g)

	// setup auth routes
	authG := g.Group("/auth")

	if err := routes.SetTwoFARoutes(authG, s.d); err != nil {
		return err
	}

	if err := routes.SetOAuthRoutes(authG, s.d); err != nil {
		return err
	}

	return nil
}

func (s *Server) pingRPCServer() error {
	conn, err := proto.DialServer(s.d.Config.MailRpcHost + ":" + s.d.Config.MailRpcPort)
	if err != nil {
		return err
	}

	// create new client
	client := pingpb.NewPingServiceClient(conn)

	// request to ping server
	req := &pingpb.PingData{
		Message: "auth",
	}

	// ping server and get the response
	res, err := client.Ping(context.Background(), req)
	if err != nil {
		log.Error().Err(err).Msg("failed to ping server")
		return err
	}

	log.Info().Msg(res.Message)
	return nil
}

func (s *Server) Start() error {
	// setup server
	if err := s.setup(); err != nil {
		return err
	}

	// ping server
	if err := s.pingRPCServer(); err != nil {
		return nil
	}

	addr := s.d.Config.Host + ":" + s.d.Config.Port
	log.Info().Msg("starting server...")
	log.Info().Msgf("server listening at %v", addr)

	if err := s.svr.Start(addr); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
		return err
	}

	return nil
}

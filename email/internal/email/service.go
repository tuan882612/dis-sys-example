package email

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"

	"dissys/internal/deps"
	"dissys/internal/email/smtp"
	"dissys/internal/proto/pb/authpb"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	SmtpSvc *smtp.Service
}

func New(d *deps.Dependencies) authpb.AuthServiceServer {
	return &Service{
		SmtpSvc: smtp.New(d.Config.Email),
	}
}

func (s *Service) SendTwoFACode(ctx context.Context, payload *authpb.TwoFAPayload) (*empty.Empty, error) {
	if err := s.SmtpSvc.SendEmail(payload.GetEmail(), smtp.AuthEmail); err != nil {
		log.Error().Err(err).Msgf("%v: failed to send two factor code", payload.GetUserId())
		return nil, err
	}

	log.Info().Msgf("%v: verification email sent", payload.GetUserId())
	return &empty.Empty{}, nil
}

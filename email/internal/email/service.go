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
	smtpSvc *smtp.Service
	cache   *repository
}

func New(d *deps.Dependencies) authpb.AuthServiceServer {
	log.Info().Msg("initializing email service...")
	return &Service{
		smtpSvc: smtp.New(d.Config.Email),
		cache:   NewRepository(d.Database.Redis),
	}
}

func (s *Service) SendTwoFACode(ctx context.Context, payload *authpb.TwoFAPayload) (*empty.Empty, error) {
	// generate a 2FA code
	code, err := s.smtpSvc.Generate2FACode()
	if err != nil {
		log.Error().Err(err).Msgf("%v: failed to generate two factor code", payload.GetUserId())
		return nil, err
	}

	// store the 2FA data
	go func() {
		if err := s.cache.StoreTwoFAData(payload.GetUserId(), payload.GetUserStatus(), code); err != nil {
			return
		}

		log.Info().Msgf("%v: two factor data stored", payload.GetUserId())
	}()

	// create the email payload
	mailPayload := smtp.NewMailPayload(smtp.AuthEmail, payload.GetUserId(), code)
	if err := s.smtpSvc.SendEmail(payload.GetEmail(), mailPayload); err != nil {
		log.Error().Err(err).Msgf("%v: failed to send two factor code", payload.GetUserId())
		return nil, err
	}

	log.Info().Msgf("%v: verification email sent", payload.GetUserId())
	return &empty.Empty{}, nil
}

package email

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"

	"dissys/internal/deps"
	"dissys/internal/email/smtp"
	"dissys/internal/proto/pb/emailpb"
)

type Service struct {
	emailpb.UnimplementedEmailServiceServer
	smtpSvc *smtp.Service
	cache   *cache
}

func NewService(d *deps.Dependencies) emailpb.EmailServiceServer {
	log.Info().Msg("initializing email service...")
	return &Service{
		smtpSvc: smtp.New(d.Config.Email),
		cache:   NewCache(d.Database.Redis),
	}
}

func (s *Service) SendTwoFACode(ctx context.Context, payload *emailpb.TwoFAPayload) (*empty.Empty, error) {
	// generate a 2FA code
	code, err := s.smtpSvc.Generate2FACode()
	if err != nil {
		log.Error().Err(err).Msgf("%v: failed to generate two factor code", payload.GetUserId())
		return nil, err
	}

	// store the 2FA data
	go func() {
		if err := s.cache.StoreTwoFAData(payload.GetUserId(), payload.GetRole(), payload.GetStatus(), code); err != nil {
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

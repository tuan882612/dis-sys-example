package twofa

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"dissys/internal/auth"
	"dissys/internal/auth/jwt"
	"dissys/internal/deps"
	"dissys/internal/proto"
	"dissys/internal/proto/pb/emailpb"
)

type tfaService struct {
	emailSvc emailpb.EmailServiceClient
	authRepo *auth.Repository
	jwtProv  *jwt.Provider
}

func NewService(d *deps.Dependencies) (*tfaService, error) {
	conn, err := proto.DialServer(d.Config.MailRpcHost + ":" + d.Config.MailRpcPort)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial grpc server")
	}

	return &tfaService{
		emailSvc: emailpb.NewEmailServiceClient(conn),
		authRepo: d.AuthRepo,
		jwtProv:  d.JWTProv,
	}, nil
}

func (s *tfaService) Login(ctx context.Context, req *auth.Request) (uuid.UUID, error) {
	creds, err := s.authRepo.FetchCredentials(ctx, req.Email)
	if err != nil {
		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(creds.Password), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return uuid.Nil, echo.NewHTTPError(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		}

		log.Error().Err(err).Msg("failed to compare password")
		return uuid.Nil, err
	}

	// send mail
	go func() {
		// load data into
		req := &emailpb.TwoFAPayload{
			UserId: creds.UserID.String(),
			Email:  req.Email,
			Role:   creds.Role,
			Status: creds.Status,
		}

		// new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		// send request to email server via grpc
		_, err := s.emailSvc.SendTwoFACode(ctx, req)
		if err != nil {
			log.Error().Err(err).Msg("failed to send 2fa request to email server")
			return
		}

		log.Info().Msgf("%v: successfully sent 2fa request to email server", creds.UserID)
	}()

	return creds.UserID, nil
}

func (s *tfaService) Verify(ctx context.Context) {

}

func (s *tfaService) Resend(ctx context.Context) {

}

func (s *tfaService) Register(ctx context.Context, req *auth.Request) (uuid.UUID, error) {
	user, err := auth.NewUser(req)
	if err != nil {
		return uuid.Nil, err
	}

	// create transaction
	tx, err := s.authRepo.CreateTx(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	// insert user
	if err := s.authRepo.InsertUser(ctx, tx, user); err != nil {
		return uuid.Nil, err
	}

	// commit transaction
	if err := s.authRepo.CommitTx(ctx, tx); err != nil {
		return uuid.Nil, err
	}

	// send email
	go func() {
		// load data into
		req := &emailpb.TwoFAPayload{
			UserId: user.UserID.String(),
			Email:  req.Email,
			Role:   user.Role,
			Status: user.Status,
		}

		// new context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
		defer cancel()

		// send request to email server via grpc
		_, err := s.emailSvc.SendTwoFACode(ctx, req)
		if err != nil {
			log.Error().Err(err).Msg("failed to send 2fa request to email server")
			return
		}

		log.Info().Msgf("%v: successfully sent 2fa request to email server", user.UserID)
	}()

	return user.UserID, nil
}

func (s *tfaService) Reset(ctx context.Context) {

}

func (s *tfaService) ResetFinal(ctx context.Context) {

}

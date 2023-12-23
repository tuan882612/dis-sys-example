package smtp

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/smtp"
	"strings"

	"github.com/absentbird/loginauth"
	"github.com/rs/zerolog/log"

	"dissys/internal/config"
)

type Service struct {
	auth smtp.Auth
	cfg  *config.Email
}

func New(cfg *config.Email) *Service {
	log.Info().Msg("initializing smtp service...")
	return &Service{
		cfg:  cfg,
		auth: loginauth.Auth(cfg.Sender, cfg.Password),
	}
}

func (s *Service) SendEmail(to string, payload *MailPayload) error {
	// Set the headers
	hMap := make(map[string]string)
	hMap["From"] = s.cfg.Sender
	hMap["To"] = to
	hMap["Content-Type"] = "text/html; charset=\"utf-8\""

	var body string
	switch payload.Type {
	case AuthEmail:
		code, ok := payload.Data.(int)
		if !ok {
			err := errors.New("failed to assert data type (int)")
			log.Error().Err(err).Msgf("%v: failed to send auth email", payload.UserID)
			return err
		}

		hMap["Subject"] = fmt.Sprintf("2FA Code: %d", code)
		body = AuthTemplate
	}

	// Build the email message
	var sb strings.Builder
	for k, v := range hMap {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	sb.WriteString("\r\n")
	sb.WriteString(body)

	// Send the email
	err := smtp.SendMail(s.cfg.SmtpHost, s.auth, s.cfg.Sender, []string{to}, []byte(sb.String()))
	if err != nil {
		log.Error().Err(err).Msgf("%v: failed to send email", payload.UserID)
		return err
	}

	return nil
}

func (s *Service) Generate2FACode() (int, error) {
	// Define the range for a 6-digit code
	const min, max = 100000, 999999

	// Generate a random number in the range [min, max]
	numRange := big.NewInt(max - min + 1)
	num, err := rand.Int(rand.Reader, numRange)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate 2FA code")
		return 0, err
	}

	code := int(num.Int64()) + min
	return code, nil
}

package smtp

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"sync"

	"github.com/absentbird/loginauth"
	"github.com/rs/zerolog/log"

	"dissys/internal/config"
)

type Service struct {
	headers  map[string]string
	auth     smtp.Auth
	smtpHost string
	mu       sync.Mutex
}

func New(cfg *config.Email) *Service {
	log.Info().Msg("initializing smtp service...")

	preHeaders := make(map[string]string)
	preHeaders["From"] = cfg.Sender
	preHeaders["To"] = ""
	preHeaders["Subject"] = ""
	preHeaders["Content-Type"] = "text/html; charset=\"utf-8\""

	return &Service{
		headers:  preHeaders,
		smtpHost: cfg.SmtpHost,
		auth:     loginauth.Auth(cfg.Sender, cfg.Password),
	}
}

func (s *Service) SendEmail(to string, emailType EmailType) error {
	s.mu.Lock()
	s.headers["To"] = to

	var body string
	switch emailType {
	case AuthEmail:
		code, err := s.Generate2FACode()
		if err != nil {
			return err
		}

		s.headers["Subject"] = fmt.Sprintf("2 Factor Authentication Code: %d", code)
		body = AuthTemplate
	}

	s.mu.Unlock()

	msg := ""
	for k, v := range s.headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body

	err := smtp.SendMail(s.smtpHost, s.auth, s.headers["From"], []string{to}, []byte(msg))
	if err != nil {
		log.Error().Msgf("failed to send email: %v", err)
		return err
	}

	return nil
}

func (s *Service) Generate2FACode() (int, error) {
	// Define the range for a 6-digit code
	const min = 100000
	const max = 999999

	// Generate a random number in the range [min, max]
	numRange := big.NewInt(max - min + 1)
	num, err := rand.Int(rand.Reader, numRange)
	if err != nil {
		log.Error().Msgf("failed to generate 2FA code: %v", err)
		return 0, err
	}

	code := int(num.Int64()) + min
	return code, nil
}

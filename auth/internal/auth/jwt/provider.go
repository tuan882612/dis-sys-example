package jwt

import (
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"dissys/internal/auth"
	"dissys/internal/config"
)

// handles the generation and validation of JWT tokens
type Provider struct {
	secret []byte
}

// New returns a new JWT provider
func NewProvider(cfg *config.JWT) (*Provider, error) {
	// decode the secret from base64
	byteSecret, err := base64.StdEncoding.DecodeString(cfg.Secret)
	if err != nil {
		log.Error().Err(err).Msg("failed to decode JWT secret")
		return nil, err
	}

	return &Provider{
		secret: byteSecret,
	}, nil
}

// GenerateToken generates a JWT token for the given user ID with RSA256
func (s *Provider) GenerateToken(userInfo *auth.UserInfo) (string, error) {
	claims := &Claims{
		UserID: userInfo.UserID.String(),
		Role:   userInfo.Role,
		Status: userInfo.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "dissys",
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// sign the token with secret key
	signedToken, err := newToken.SignedString(s.secret)
	if err != nil {
		log.Error().Err(err).Msg("failed to sign JWT token")
		return "", err
	}

	return signedToken, nil
}

package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type (
	// used for login and register requests
	Request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}

	UserCreds struct {
		UserID   uuid.UUID `json:"user_id"`
		Password []byte    `json:"password"`
		Status   string    `json:"status"`
		Role     string    `json:"role"`
	}

	UserInfo struct {
		UserID  uuid.UUID `json:"user_id"`
		Created time.Time `json:"created"`
		Status  string    `json:"status"`
		Role    string    `json:"role"`
	}

	User struct {
		UserID   uuid.UUID `json:"user_id"`
		Email    string    `json:"email"`
		Password []byte    `json:"password"`
		Created  time.Time `json:"created"`
		Status   string    `json:"status"`
		Role     string    `json:"role"`
	}
)

func NewUser(req *Request) (*User, error) {
	hashedPsw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return nil, err
	}

	return &User{
		UserID:   uuid.New(),
		Email:    req.Email,
		Password: hashedPsw,
		Created:  time.Now(),
		Status:   StatPending,
		Role:     UsrBase,
	}, nil
}

func (u *User) ToUserInfo() *UserInfo {
	return &UserInfo{
		UserID:  u.UserID,
		Created: u.Created,
		Status:  u.Status,
		Role:    u.Role,
	}
}

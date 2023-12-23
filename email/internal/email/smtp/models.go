package smtp

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type EmailType string

const (
	AuthEmail EmailType = "auth"
)

type MailPayload struct {
	UserID string
	Type   EmailType
	Data   interface{}
}

func NewMailPayload(t EmailType, userID string, data interface{}) *MailPayload {
	return &MailPayload{
		UserID: userID,
		Type:   t,
		Data:   data,
	}
}

type TwoFAData struct {
	Code    int
	Retries int
	Status  string
}

func NewTwoFAData(code int, retries int, status string) *TwoFAData {
	return &TwoFAData{
		Code:    code,
		Retries: retries,
		Status:  status,
	}
}

func (t *TwoFAData) Serialize() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal two factor data")
		return "", err
	}

	return string(data), nil
}

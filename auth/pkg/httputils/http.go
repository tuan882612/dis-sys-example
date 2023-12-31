package httputils

import (
	"io"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
)

func Unmarshal(data io.ReadCloser, v interface{}) error {
	if err := jsoniter.NewDecoder(data).Decode(v); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal data")
		return err
	}

	if err := validator.New().Struct(v); err != nil {
		log.Error().Err(err).Msg("failed to validate data")
		return err
	}

	return nil
}

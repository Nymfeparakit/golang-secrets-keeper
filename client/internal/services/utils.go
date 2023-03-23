package services

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkUnavailableError(err error) bool {
	st, ok := status.FromError(err)
	if ok && st.Code() == codes.Unavailable {
		log.Error().Err(err).Msg("remote storage is not available:")
		return true
	}
	return false
}

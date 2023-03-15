package services

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCheckUnavailableError(t *testing.T) {
	tests := []struct {
		name      string
		errArg    error
		expResult bool
	}{
		{
			name:      "unavailable error",
			errArg:    status.Error(codes.Unavailable, "Unavailable"),
			expResult: true,
		},
		{
			name:      "not unavailable error",
			errArg:    status.Error(codes.Internal, "Internal server error"),
			expResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkUnavailableError(tt.errArg)
			assert.Equal(t, tt.expResult, result)
		})
	}
}

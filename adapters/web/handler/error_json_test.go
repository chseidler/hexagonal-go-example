package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonErr(msg)
	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}

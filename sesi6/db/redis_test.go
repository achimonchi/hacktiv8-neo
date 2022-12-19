package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedis_Connect(t *testing.T) {
	client, err := NewRedisStore()
	require.Nil(t, err)
	require.NotNil(t, client)

}

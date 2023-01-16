package pkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "goneo")
	os.Setenv("DB_PASSWORD", "goneo")
	os.Setenv("DB_DBNAME", "goneo")
}

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB()

	require.Nil(t, err)
	require.NotNil(t, db)
}

package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLDAPGet_Success(t *testing.T) {
	username, password := "einstein@ldap.forumsys.com", "password"
	ok, data, err := AuthUsingLDAP(username, password)

	require.Nil(t, err)
	require.True(t, ok)
	require.NotNil(t, data)

	require.Equal(t, "einstein", data.ID)
}

func TestLDAPGet_Failure(t *testing.T) {
	username, password := "einstein@ldap.forumsys.com", "password1asd"
	ok, data, err := AuthUsingLDAP(username, password)

	require.Nil(t, err)
	require.False(t, ok)
	require.Nil(t, data)

}

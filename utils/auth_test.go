package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAuthorizer(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll("testDir")
	})
	require.NoError(t, os.MkdirAll("testDir", os.ModePerm))
	acct, err := NewKeyFile("testDir", "password123")
	require.NoError(t, err)
	_, err = NewAuthorizer(acct.URL.Path, "password123")
	require.NoError(t, err)
}

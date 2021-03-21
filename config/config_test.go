package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigFile(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("test-config.yml")
		os.Remove("go-defi.log")
	})
	err := NewConfig("test-config.yml")
	require.NoError(t, err)
	cfg, err := LoadConfig("test-config.yml")
	require.NoError(t, err)
	logger, err := cfg.ZapLogger()
	require.NoError(t, err)
	logger.Info("hello")
	logger.Sync()
}

func TestConfigEnv(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("go-defi.log")
	})
	cfg, err := LoadConfig("not-a-real-config.yml")
	require.NoError(t, err)
	logger, err := cfg.ZapLogger()
	require.NoError(t, err)
	logger.Info("hello")
	logger.Sync()
}

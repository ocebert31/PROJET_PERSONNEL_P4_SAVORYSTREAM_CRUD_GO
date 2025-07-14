package config

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

const dotenvPath = ".env"

func TestLoadEnv(t *testing.T) {
	t.Run("with .env file", func(t *testing.T) {
		const dotenvContent = "SAUCE_API_PORT=:7777\nSAUCE_API_DSN=test-dsn"
		err := os.WriteFile(dotenvPath, []byte(dotenvContent), 0644)
		assert.NoError(t, err, "Erreur lors de l’écriture du .env de test")
		t.Cleanup(func() {
			_ = os.Remove(dotenvPath)
			_ = os.Unsetenv("SAUCE_API_PORT")
			_ = os.Unsetenv("SAUCE_API_DSN")
		})
		LoadEnv()
		assert.Equal(t, ":7777", os.Getenv("SAUCE_API_PORT"))
		assert.Equal(t, "test-dsn", os.Getenv("SAUCE_API_DSN"))
	})

	t.Run("without .env file", func(t *testing.T) {
		LoadEnv()
	})

	t.Run("with .env file but no SAUCE_API_PORT", func(t *testing.T) {
		const dotenvContent = "SAUCE_API_DSN=test-dsn"
		err := os.WriteFile(dotenvPath, []byte(dotenvContent), 0644)
		assert.NoError(t, err)

		t.Cleanup(func() {
			_ = os.Remove(dotenvPath)
			_ = os.Unsetenv("SAUCE_API_DSN")
		})

		_ = LoadEnv()
		got := GetAPIPort()
		assert.Equal(t, defaultPort, got)
	})
}

func TestGetAPIPort(t *testing.T) {
	t.Run("with port env variable defined", func(t *testing.T) {
		expected := ":9999"
		_ = os.Setenv("SAUCE_API_PORT", expected)
		got := GetAPIPort()
		assert.Equal(t, expected, got)
	})

	t.Run("without port env variable defined", func(t *testing.T) {
		_ = os.Unsetenv("SAUCE_API_PORT")

		t.Cleanup(func() {
			_ = os.Remove(dotenvPath)
			_ = os.Unsetenv("SAUCE_API_PORT")
		})

		got := GetAPIPort()
		assert.Equal(t, defaultPort, got)
	})
}

func TestGetDatabaseDSN(t *testing.T) {
	t.Run("with database env variable defined", func(t *testing.T) {
		expected := "user=test password=1234 dbname=mydb"
		_ = os.Setenv("SAUCE_API_DSN", expected)

		t.Cleanup(func() {
			_ = os.Remove(dotenvPath)
			_ = os.Unsetenv("SAUCE_API_DSN")
		})

		got, err := GetDatabaseDSN()
		assert.NoError(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("without database env variable defined", func(t *testing.T) {
		_ = os.Unsetenv("SAUCE_API_DSN")

		t.Cleanup(func() {
			_ = os.Remove(dotenvPath)
			_ = os.Unsetenv("SAUCE_API_DSN")
		})

		got, err := GetDatabaseDSN()
		assert.Error(t, err)
		assert.Empty(t, got)
	})
}



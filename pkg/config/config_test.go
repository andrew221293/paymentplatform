package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name       string
		setupEnv   func()
		wantConfig *Config
		wantErr    bool
	}{
		{
			name: "Success: All environment variables set",
			setupEnv: func() {
				os.Setenv("POSTGRES_USER", "testuser")
				os.Setenv("POSTGRES_PASSWORD", "testpassword")
				os.Setenv("POSTGRES_DB", "testdb")
				os.Setenv("DB_HOST", "localhost")
				os.Setenv("DB_PORT", "5432")
			},
			wantConfig: &Config{
				DBUser:     "testuser",
				DBPassword: "testpassword",
				DBName:     "testdb",
				DBHost:     "localhost",
				DBPort:     5432,
			},
			wantErr: false,
		},
		{
			name: "Success: Missing or invalid DB_PORT uses default",
			setupEnv: func() {
				os.Setenv("POSTGRES_USER", "testuser")
				os.Setenv("POSTGRES_PASSWORD", "testpassword")
				os.Setenv("POSTGRES_DB", "testdb")
				os.Setenv("DB_HOST", "localhost")
				os.Setenv("DB_PORT", "invalid")
			},
			wantConfig: &Config{
				DBUser:     "testuser",
				DBPassword: "testpassword",
				DBName:     "testdb",
				DBHost:     "localhost",
				DBPort:     5432,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Unsetenv("POSTGRES_USER")
				os.Unsetenv("POSTGRES_PASSWORD")
				os.Unsetenv("POSTGRES_DB")
				os.Unsetenv("DB_HOST")
				os.Unsetenv("DB_PORT")
			}()

			tt.setupEnv()

			gotConfig, err := LoadConfig()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantConfig, gotConfig)
			}
		})
	}
}

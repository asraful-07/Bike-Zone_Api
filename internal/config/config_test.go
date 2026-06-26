package config

import (
	"os"
	"testing"
)

func TestLoadConfigUsesEnvironmentVariablesWithoutDotEnv(t *testing.T) {
	tmpDir := t.TempDir()
	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	defer func() {
		_ = os.Chdir(oldWD)
	}()

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	t.Setenv("DSN", "postgres://user:pass@localhost:5432/db")
	t.Setenv("PORT", "9090")
	t.Setenv("JWT_SECRET_KEY", "test-secret")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	if cfg.DSN != "postgres://user:pass@localhost:5432/db" {
		t.Fatalf("unexpected DSN: %q", cfg.DSN)
	}
	if cfg.PORT != "9090" {
		t.Fatalf("unexpected PORT: %q", cfg.PORT)
	}
	if cfg.JWTSecretKey != "test-secret" {
		t.Fatalf("unexpected JWT secret: %q", cfg.JWTSecretKey)
	}
}

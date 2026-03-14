package config

import "os"

// Config はアプリケーション設定
type Config struct {
	DatabaseURL string
	JWTSecret   string
}

// Load は環境変数から設定を読み込む
func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://cashpilot:cashpilot@localhost:5432/cashpilot?sslmode=disable"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

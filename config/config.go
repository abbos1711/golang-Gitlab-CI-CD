package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Environment                string
	LogLevel                   string
	HttpPort                   string
	CtxTimeout                 int64
	AuthConfigPath             string
	TokenSecretKey             string
	SignKey                    string
	RedisAddr                  string
	Postgres                   Postgres
	AccessTokenExpireDuration  time.Duration
	ResetPasswordTokenDuration time.Duration
}

type Postgres struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	DatabaseURL      string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")

	conf := viper.New()
	conf.AutomaticEnv()

	return Config{
		Environment:                conf.GetString("ENVIRONMENT"),
		CtxTimeout:                 conf.GetInt64("CTX_TIMEOUT"),
		AuthConfigPath:             conf.GetString("AUTH_PATH"),
		LogLevel:                   conf.GetString("LOG_LEVEL"),
		HttpPort:                   conf.GetString("HTTP_PORT"),
		TokenSecretKey:             conf.GetString("TOKEN_SECRET_KEY"),
		SignKey:                    conf.GetString("SIGN_KEY"),
		RedisAddr:                  conf.GetString("REDIS_ADDR"),
		AccessTokenExpireDuration:  conf.GetDuration("ACCESS_TOKEN_DURATION"),
		ResetPasswordTokenDuration: conf.GetDuration("RESET_PASSWORD_TOKEN_DURATION"),

		Postgres: Postgres{
			PostgresDatabase: conf.GetString("POSTGRES_DATABASE"),
			PostgresUser:     conf.GetString("POSTGRES_USER"),
			PostgresPassword: conf.GetString("POSTGRES_PASSWORD"),
			PostgresPort:     conf.GetString("POSTGRES_PORT"),
			PostgresHost:     conf.GetString("POSTGRES_HOST"),
			DatabaseURL:      conf.GetString("DATABASE_URL"),
		},
	}
}

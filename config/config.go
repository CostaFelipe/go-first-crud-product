package config

import "github.com/go-chi/jwtauth"

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	WebServePort string `mapstructure:"WEB_SERVE_PORT"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIN int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

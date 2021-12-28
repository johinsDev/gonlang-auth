package config

import (
	"os"
)

type MaiLConfig struct {
	PORT string
	HOST string
	FROM struct {
		NAME    string
		ADDRESS string
	}
	PASSWORD string
	USERNAME string
}

func Getenv(tag string, defaultValue ...string) string {
	if TAG := os.Getenv(tag); TAG != "" {
		return TAG
	}

	return defaultValue[0]
}

func GetMailConfig() *MaiLConfig {

	return &MaiLConfig{
		PORT:     Getenv("MAIl_PORT", "2525"),
		HOST:     Getenv("MAIL_HOST", "smtp.mailtrap.io"),
		PASSWORD: Getenv("MAIL_PASSWORD", "3c495b7ec5a367"),
		USERNAME: Getenv("MAIL_USERNAME", "d1f01ef63ae5ca"),
		FROM: struct {
			NAME    string
			ADDRESS string
		}{
			NAME: Getenv(
				"MAIL_FROM_NAME",
				"johinsdev",
			), ADDRESS: Getenv("MAIL_FROM_ADDRESS", "johinsdev@gmail.com"),
		},
	}
}

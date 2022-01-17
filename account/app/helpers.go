package app

import (
	"github.com/gookit/config"
	"github.com/johinsDev/authentication/app/providers"
	"github.com/johinsDev/authentication/lib/authentication"
	"github.com/johinsDev/authentication/lib/hash"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Config() *config.Config {
	var c *config.Config

	err := providers.Container().Invoke(func(config *config.Config) {
		c = config

	})

	if err != nil {
		logrus.Error(err)
	}

	return c
}

func Database() *gorm.DB {
	var d *gorm.DB

	err := providers.Container().Invoke(func(database *gorm.DB) {
		d = database
	})

	if err != nil {
		logrus.Error(err)
	}

	return d
}

func Hash() *hash.Hash {
	var h *hash.Hash

	err := providers.Container().Invoke(func(hash *hash.Hash) {
		h = hash
	})

	if err != nil {
		logrus.Error(err)
	}

	return h
}

func Auth() *authentication.Auth {
	var a *authentication.Auth

	err := providers.Container().Invoke(func(auth *authentication.Auth) {
		a = auth
	})

	if err != nil {
		logrus.Error(err)
	}

	return a
}

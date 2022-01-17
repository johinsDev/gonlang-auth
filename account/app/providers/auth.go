package providers

import (
	"github.com/johinsDev/authentication/lib/authentication"
	"github.com/johinsDev/authentication/lib/hash"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewAuth(db *gorm.DB, hasher *hash.Hash) *authentication.Auth {
	logrus.Info("Starting database...")

	return &authentication.Auth{
		Db:     db,
		Hasher: hasher,
	}
}

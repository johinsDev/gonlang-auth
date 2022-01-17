package authentication

import (
	"github.com/johinsDev/authentication/lib/hash"
	"github.com/johinsDev/authentication/models"
	"gorm.io/gorm"
)

type Auth struct {
	Db     *gorm.DB
	Hasher *hash.Hash
}

func (a *Auth) Attempt(username string, password string) (bool, uint) {
	foundUser, user := a.getByUsername(username)

	if !foundUser || !a.hasValidCredentials(user, password) {
		return false, 0
	}

	return true, user.ID
}

func (a *Auth) hasValidCredentials(user *models.User, password string) bool {
	verify, _ := a.Hasher.Verify(user.Password, password)

	return verify
}

func (a *Auth) getByUsername(username string) (bool, *models.User) {
	user := &models.User{}

	res := a.Db.Where("email = ?", username).First(user)

	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return false, nil
	}

	return true, user
}

package hash

import (
	"github.com/sirupsen/logrus"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type BcryptConfig struct {
	Rounds int `mapstructure:"rounds"`
	BaseConfig
}

type Bcrypt struct {
	Config *BcryptConfig
}

func (hasher *Bcrypt) Make(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), hasher.Config.Rounds)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (hasher *Bcrypt) Verify(hashedValue string, plainValue string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(plainValue)); err != nil {

		return false, err
	}

	return true, nil
}

func (hasher *Bcrypt) NeedReHash(hashedValue string) (bool, error) {
	currentCost, err := bcrypt.Cost([]byte(hashedValue))

	if err != nil {
		logrus.Error("Error validating re hash: ", err, hashedValue)
		return true, err
	}

	if currentCost != hasher.Config.Rounds {
		return true, nil
	}

	return false, nil
}

func NewBcrypt(config *BcryptConfig) *Bcrypt {
	return &Bcrypt{
		Config: &BcryptConfig{
			Rounds: config.Rounds,
			BaseConfig: BaseConfig{
				Driver: "brcypt",
			},
		},
	}
}

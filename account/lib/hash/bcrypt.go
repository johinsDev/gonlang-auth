package hash

import (
	crypto "golang.org/x/crypto/bcrypt"
)

type BcryptConfig struct {
	Rounds int `mapstructure:"rounds"`
	BaseConfig
}

type Bcrypt struct {
	Config *BcryptConfig
}

func (hasher *Bcrypt) Make(value string) (string, error) {
	hash, err := crypto.GenerateFromPassword([]byte(value), hasher.Config.Rounds)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (hasher *Bcrypt) Verify(hashedValue string, plainValue string) (bool, error) {
	if err := crypto.CompareHashAndPassword([]byte(hashedValue), []byte(plainValue)); err == nil {
		return false, err
	}

	return true, nil
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

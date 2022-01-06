package hash

import "github.com/sirupsen/logrus"

type BcryptConfig struct {
	Rounds int    `mapstructure:"rounds"`
	Driver string `mapstructure:"Driver"`
}

type Bcrypt struct {
	Config *BcryptConfig
}

func (hassher *Bcrypt) Make(value string) string {
	logrus.Info("Hashsing with bcryot", hassher.Config)

	return value
}

func NewBcrypt(config *BcryptConfig) *Bcrypt {
	return &Bcrypt{
		Config: &BcryptConfig{
			Rounds: config.Rounds,
			Driver: "brcypt",
		},
	}
}

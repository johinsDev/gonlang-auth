package hash

import (
	"github.com/sirupsen/logrus"
	crypto "golang.org/x/crypto/argon2"
)

type ArgonConfig struct {
	Variant  string
	Memory   uint32
	SaltSize string
	Driver   string
	Time     int
	KeyLen   uint32
	Threads  uint8
}

type Argon struct {
	Config *ArgonConfig
}

func (hasher *Argon) Make(value string) (string, error) {
	hashed := crypto.Key(
		[]byte(value),
		[]byte(hasher.Config.SaltSize),
		uint32(hasher.Config.Time),
		hasher.Config.Memory,
		hasher.Config.Threads,
		hasher.Config.KeyLen,
	)

	return string(hashed), nil
}

func (hasher *Argon) Verify(hashedValue string, plainValue string) (bool, error) {
	return true, nil
}

func NewArgon(config *ArgonConfig) *Argon {
	validVariants := map[string]bool{
		"d": true, "i": true, "id": true,
	}

	if _, ok := validVariants[config.Variant]; !ok {
		logrus.Error("Varaint not valid")
	}

	return &Argon{
		Config: &ArgonConfig{
			Memory:   config.Memory,
			KeyLen:   config.KeyLen,
			SaltSize: config.SaltSize,
			Threads:  config.Threads,
			Variant:  config.Variant,
			Driver:   "argon2",
			Time:     config.Time,
		},
	}
}

package hash

import "github.com/sirupsen/logrus"

type ArgonConfig struct {
	Variant     string
	Iterations  int
	Memory      int
	Parallelism int
	SaltSize    int
	Driver      string
}

type Argon struct {
	Config *ArgonConfig
}

func (hassher *Argon) Make(value string) string {
	logrus.Info("Hashsing with argon")
	return value
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
			Memory:      config.Memory,
			Iterations:  config.Iterations,
			SaltSize:    config.SaltSize,
			Parallelism: config.Parallelism,
			Variant:     config.Variant,
			Driver:      "argon2",
		},
	}
}

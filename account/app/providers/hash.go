package providers

import (
	"github.com/gookit/config"
	"github.com/johinsDev/authentication/lib/hash"
	"github.com/sirupsen/logrus"
)

// provider config
// provider logger
// dig
// provider server (port, prefix) move handler to app
// authentication service
// authentication multiple drivers and providers
// mail (config) support multiple drivers
// logger abstract
// config abstract
// redis
// redis provider auth
// throttler
// cors
// readme documentation
// baseModel
// dynamic filter model

func NewHash() *hash.Hash {
	logrus.Info("Starting hash...")

	return hash.NewHasher(&hash.Config{
		DefaultHash: config.DefString("HASH.DRIVER", "bcrypt"),
		List: map[string]interface{}{
			"bcrypt": &hash.BcryptConfig{
				Rounds: config.DefInt("HASH.BCRYPT.ROUNDS", 10),
				BaseConfig: hash.BaseConfig{
					Driver: "bcrypt",
				},
			},
			"fake": &hash.FakeConfig{},
		},
	})
}

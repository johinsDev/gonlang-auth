package hash

import (
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

// @TODO Create fake driver
// implement verify and needsRefash
// real implmentations
// isFaked
// set fake driver
// create extend driver tiwhout touch this
// cache de drivers

type HashDriverContract interface {
	Make(value string) string
}

type Config struct {
	DefaultHash string
	List        map[string]interface{}
}

type Hash struct {
	config *Config
}

func (hasher *Hash) getMappingConfig(name string, mapping interface{}) interface{} {
	if config, ok := hasher.config.List[name]; ok {
		hasher.bindConfig(config, mapping)

		return config
	}

	logrus.Error("Config not found")

	panic("Config not found")
}

func (hasher *Hash) bindConfig(data interface{}, mapping interface{}) {
	var bindConf *mapstructure.DecoderConfig

	bindConf = &mapstructure.DecoderConfig{
		TagName:          "mapstructure",
		WeaklyTypedInput: true,
	}

	bindConf.Result = mapping

	decoder, err := mapstructure.NewDecoder(bindConf)

	if err != nil {
		logrus.Error("Error loading config")
	}

	decoder.Decode(data)
}

func (hasher *Hash) createBcrypt() HashDriverContract {
	config := &BcryptConfig{}

	hasher.getMappingConfig("bcrypt", config)

	return NewBcrypt(config)
}

func (hasher *Hash) createArgon() HashDriverContract {
	config := &ArgonConfig{}

	hasher.getMappingConfig("argon", config)

	return NewArgon(config)
}

func (hasher *Hash) Use(algorithm ...string) HashDriverContract {
	var name string

	if algorithm == nil {
		name = hasher.config.DefaultHash
	} else {
		name = algorithm[0]
	}

	switch name {
	case "argon":
		return hasher.createArgon()
	case "bcrypt":
		return hasher.createBcrypt()
	default:
		return hasher.createBcrypt()
	}
}

func (hasher *Hash) Make(value string) string {
	return hasher.Use().Make(value)
}

func NewHasher(config *Config) *Hash {
	return &Hash{
		config: config,
	}
}

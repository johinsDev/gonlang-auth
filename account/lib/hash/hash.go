package hash

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

type HashDriverContract interface {
	Make(value string) (string, error)
	Verify(hashedValue string, plainValue string) (bool, error)
}

type BaseConfig struct {
	Driver         string                                                    `mapstructure:"driver"`
	Implementation func(config interface{}, hasher *Hash) HashDriverContract `mapstructure:"implementation"`
}

type Config struct {
	DefaultHash string
	List        map[string]interface{}
}

type Hash struct {
	config         *Config
	fakeDriver     HashDriverContract
	MappingsCache  map[string]HashDriverContract
	ExtendedGuards map[string]func(config interface{}, hasher *Hash) HashDriverContract
}

func (hasher *Hash) getMappingConfig(name string) interface{} {
	if config, ok := hasher.config.List[name]; ok {
		return config
	}

	logrus.Error("Config not found")

	panic("Config not found")
}

func (hasher *Hash) BindConfig(data interface{}, mapping interface{}) {
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

func (hasher *Hash) createBcrypt(data interface{}) HashDriverContract {
	config := &BcryptConfig{}

	hasher.BindConfig(data, config)

	return NewBcrypt(config)
}

func (hasher *Hash) createFake(data interface{}) HashDriverContract {
	config := &FakeConfig{}

	hasher.BindConfig(data, config)

	return NewFake(config)
}

func (hasher *Hash) createExtendedDriver(name string, data interface{}) HashDriverContract {
	if cb, ok := hasher.ExtendedGuards[name]; ok {
		return cb(data, hasher)
	}

	return nil
}

func (hasher *Hash) makeMapping(name string) HashDriverContract {
	data := hasher.getMappingConfig(name)

	r := reflect.ValueOf(data)

	f := reflect.Indirect(r).FieldByName("BaseConfig")

	config := &BaseConfig{}

	hasher.BindConfig(f.Interface(), config)

	if config.Implementation != nil {
		return config.Implementation(data, hasher)
	}

	switch config.Driver {
	case "bcrypt":
		return hasher.createBcrypt(data)
	case "fake":
		return hasher.createFake(data)
	default:
		return hasher.createExtendedDriver(name, data)
	}
}

func (hasher *Hash) Restore() *Hash {
	hasher.fakeDriver = nil

	return hasher
}

func (hasher *Hash) isFake() bool {
	return hasher.fakeDriver != nil
}

func (hasher *Hash) Fake(driver ...HashDriverContract) *Hash {
	data := hasher.getMappingConfig("fake")

	if driver == nil {
		hasher.fakeDriver = hasher.createFake(data)
	} else {
		hasher.fakeDriver = driver[0]
	}

	return hasher
}

func (hasher *Hash) Extend(
	name string,
	callback func(data interface{}, hasher *Hash) HashDriverContract,
) *Hash {
	hasher.ExtendedGuards[name] = callback

	return hasher
}

// Proxy functions
func (hasher *Hash) Use(algorithm ...string) HashDriverContract {
	var name string

	if algorithm == nil {
		name = hasher.config.DefaultHash
	} else {
		name = algorithm[0]
	}

	if hasher.fakeDriver != nil {
		return hasher.fakeDriver
	}

	if _, ok := hasher.MappingsCache[name]; !ok {
		driver := hasher.makeMapping(name)

		hasher.MappingsCache[name] = driver

		return driver
	}

	driver := hasher.MappingsCache[name]

	return driver
}

func (hasher *Hash) Make(value string) (string, error) {
	return hasher.Use().Make(value)
}

func (hasher *Hash) Verify(hashedValue string, plainValue string) (bool, error) {
	return hasher.Use().Verify(hashedValue, plainValue)
}

func NewHasher(config *Config) *Hash {
	return &Hash{
		config:         config,
		MappingsCache:  make(map[string]HashDriverContract),
		fakeDriver:     nil,
		ExtendedGuards: make(map[string]func(config interface{}, hasher *Hash) HashDriverContract),
	}
}

package providers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gookit/config"
	"github.com/gookit/config/yaml"
	"github.com/sirupsen/logrus"
)

func NewConfig() *config.Config {

	logrus.Info("Starting config")

	c := config.New("default")

	c.WithOptions(config.ParseEnv)

	c.AddDriver(yaml.Driver)

	files, err := ioutil.ReadDir("config/")

	if err != nil {
		logrus.Error(err)
	}

	filesTransformed := make([]string, 0)

	for _, file := range files {
		ext := filepath.Ext(file.Name())

		if ext == ".yml" {
			filesTransformed = append(filesTransformed, fmt.Sprintf("config/%s", file.Name()))
		}

	}

	err = c.LoadFiles(filesTransformed...)

	if err != nil {
		logrus.Error("err loading files >>>>> ", err)
	}

	return c
}

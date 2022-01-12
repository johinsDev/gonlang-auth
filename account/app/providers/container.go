package providers

import (
	"sync"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var (
	once      sync.Once
	container *dig.Container
)

func Container() *dig.Container {

	if container == nil {
		logrus.Info("Starting container")

		container = dig.New()
	}

	return container
}

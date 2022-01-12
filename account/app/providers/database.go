package providers

import (
	"fmt"

	"github.com/gookit/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(c *config.Config) *gorm.DB {
	logrus.Info("Starting database...")

	host, _ := c.String("DB_HOST")
	port, _ := c.String("DB_PORT")
	username, _ := c.String("DB_USERNAME")
	password, _ := c.String("DB_PASSWORD")
	name, _ := c.String("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host,
		username,
		password,
		name,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Error init database", err)
	}

	return db
}

package backend

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func OpenConnection(filepath string) error {
	conn, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		log.WithError(err).WithField("filepath", filepath).Error("Failed to open DB connection")
		return err
	}
	Conn = conn
	return nil
}

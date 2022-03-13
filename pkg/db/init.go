package db

import (
	"fmt"

	"github.com/marcus-crane/october/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func OpenConnection(filepath string) error {
	conn, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		logger.Log.Errorw(fmt.Sprintf("Failed to open DB connection to %s", filepath), "error", err)
		return err
	}
	Conn = conn
	return nil
}

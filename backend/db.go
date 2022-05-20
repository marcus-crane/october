package backend

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func OpenConnection(filepath string) error {
	conn, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		return err
	}
	Conn = conn
	return nil
}

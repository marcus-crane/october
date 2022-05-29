package backend

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func OpenConnection(filepath string) error {
	conn, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Str("filepath", filepath).Msg("Failed to open DB connection")
		return err
	}
	Conn = conn
	return nil
}

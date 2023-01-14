package kobo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Kobo represents a device, either physically connected or it may be operating on
// a backup locally stored on disc. It doesn't make much difference operationally.
type Kobo struct {
	dbClient   *sqlx.DB `json:"-"`
	Name       string   `json:"device_name"`
	Storage    int      `json:"device_storage"`
	DisplayPPI int      `json:"display_ppi"`
	MountPath  string   `json:"mount_path"`
	DbPath     string   `json:"db_path"`
	Serial     string   `json:"serial"`
	Version    string   `json:"version"`
	DeviceId   string   `json:"deviceId"`
}

// NewKobo can be used to emulate a connection to a device on disc
// with some reasonable defaults. Rather than using default Kobo layout, it is
// possible to store your Kobo device layout in one folder and your DB in another
// but the underlying assumption is that it will try to operate upon a folder
// as if it has epubs and what not in addition to being able to query a database
func NewKobo(mountPath string, dbPath string) Kobo {
	return Kobo{
		Name:      "Direct Connection",
		MountPath: dbPath,
		DbPath:    dbPath,
	}
}

// Connect will do some sanity checking around the configured database path
// and upon passing, will instantiate the connection to the underlying Kobo
// sqlite database
func (k *Kobo) Connect() error {
	if k.DbPath == "" {
		return fmt.Errorf("db path must be specified to create a connection")
	}
	// TODO: Find cgo-less driver
	db, err := sqlx.Connect("sqlite3", k.DbPath)
	if err != nil {
		return err
	}
	k.dbClient = db
	return nil
}

// Ping can be used to check if the connection to the database is still alive
func (k *Kobo) Ping() error {
	return k.dbClient.Ping()
}

// Disconnect can be called to shut down the underlying database connection. This is
// generally done when switching from one database to another.
func (k *Kobo) Disconnect() error {
	return k.dbClient.Close()
}

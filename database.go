package main

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "github.com/wailsapp/wails"
)

var (
  DBConn *gorm.DB
  DBPath string
)

type Database struct {
  runtime *wails.Runtime
}

func OpenDatabase(dbPath string) bool {
  db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
  if err == nil {
    DBConn = db
    return true
  }
  return false
}

func (d *Database) WailsInit(runtime *wails.Runtime) error {
  d.runtime = runtime
  return nil
}

func (d *Database) PromptUserForFilePath() error {
  selectedFile := d.runtime.Dialog.SelectFile()
  DBPath = selectedFile
  OpenDatabase(selectedFile)
  return nil
}

func NewDatabase() *Database {
  return &Database{}
}
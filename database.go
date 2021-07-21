package main

import (
  "github.com/wailsapp/wails"
)

type Database struct {
  runtime *wails.Runtime
}

func (d *Database) WailsInit(runtime *wails.Runtime) error {
  d.runtime = runtime
  return nil
}

func NewDatabase() *Database {
  return &Database{}
}
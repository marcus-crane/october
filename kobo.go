package main

import (
  "fmt"

  "github.com/pgaskin/koboutils/v2/kobo"
  "github.com/wailsapp/wails"
)

var selectedKobo Kobo

type Kobo struct {
  Name       string
  Storage    int
  DisplayPPI int
  MntPath    string
  DbPath     string
  runtime    *wails.Runtime
}

type DetectedKobos struct {
  Kobos      []Kobo
}

func NewKobo() *Kobo {
  return &Kobo{}
}

func (k *Kobo) WailsInit(runtime *wails.Runtime) error {
  k.runtime = runtime
  return nil
}

func (Kobo) DetectKobo() DetectedKobos {
  kobos := DetectedKobos{}
  koboList, err := kobo.Find()
  if err != nil {
    panic(err)
  }
  for _, koboPath := range koboList {
    _, _, deviceId, err := kobo.ParseKoboVersion(koboPath)
    if err != nil {
      panic(err)
    }
    device, found := kobo.DeviceByID(deviceId)
    if !found {
      continue
    }
    kobo := Kobo{
      Name:       device.Name(),
      Storage:    device.StorageGB(),
      DisplayPPI: device.DisplayPPI(),
      MntPath:    koboPath,
      DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", koboPath),
    }
    kobos.Kobos = append(kobos.Kobos, kobo)
  }
  return kobos
}

func (Kobo) SelectKobo(devicePath string) bool {
  _, _, deviceId, err := kobo.ParseKoboVersion(devicePath)
  if err != nil {
    panic(err)
  }
  device, found := kobo.DeviceByID(deviceId)
  if !found {
    panic("device detached?")
  }
  selectedKobo = Kobo{
    Name:       device.Name(),
    Storage:    device.StorageGB(),
    DisplayPPI: device.DisplayPPI(),
    MntPath:    devicePath,
    DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", devicePath),
  }
  return true
}

func (k *Kobo) ConnectToKoboDB() bool {
  return OpenDatabase(selectedKobo.DbPath)
}

func (k *Kobo) SelectLocalDatabase() bool {
  dbPath := k.runtime.Dialog.SelectFile()
  selectedKobo = Kobo{
    Name: "Local Database",
    Storage: 0,
    DisplayPPI: 0,
    MntPath: dbPath,
    DbPath: dbPath,
  }
  return true
}

func (Kobo) GetBasicKoboDetails() Kobo {
  return selectedKobo
}
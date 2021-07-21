package main

import "github.com/pgaskin/koboutils/v2/kobo"

var selectedKobo Kobo

type Kobo struct {
  Name       string
  Storage    int
  DisplayPPI int
  MntPath    string
}

type DetectedKobos struct {
  Kobos      []Kobo
}

func NewKobo() *Kobo {
  return &Kobo{}
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
  }
  return true
}

func (Kobo) GetBasicKoboDetails() Kobo {
  return selectedKobo
}
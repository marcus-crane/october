package main

import (
  _ "embed"

  "github.com/pgaskin/koboutils/v2/kobo"
  "github.com/wailsapp/wails"
)

type Kobo struct {
  Name       string
  Storage    int
  DisplayPPI int
  MntPath    string
}

type DetectedKobos struct {
  Kobos      []Kobo
}

func basic() string {
  return "Hello World!"
}

func detectKobo() DetectedKobos {
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

func selectKobo(devicePath string) bool {
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

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

var selectedKobo Kobo

func main() {

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "Octowise",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Bind(detectKobo)
  app.Bind(selectKobo)
  app.Run()
}

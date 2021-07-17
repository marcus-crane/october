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

func basic() string {
  return "Hello World!"
}

func detectKobo() Kobo {
  kobos, err := kobo.Find()
  if err != nil {
    panic(err)
  }
  for _, koboPath := range kobos {
    _, _, deviceId, err := kobo.ParseKoboVersion(koboPath)
    if err != nil {
      panic(err)
    }
    device, found := kobo.DeviceByID(deviceId)
    if !found {
      return Kobo{}
    }
    return Kobo{
      Name:       device.Name(),
      Storage:    device.StorageGB(),
      DisplayPPI: device.DisplayPPI(),
      MntPath:    koboPath,
    }
  }
  return Kobo{}
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

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
  app.Run()
}

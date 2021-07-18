package main

import (
  _ "embed"
  "fmt"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "gorm.io/gorm/clause"
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

type Bookmark struct {
  BookmarkID string `gorm:"column:BookmarkID"`
  VolumeID   string
  ContentID  string
  StartContainerPath string
  StartContainerChild string
  StartContainerChildIndex string
  StartOffset string
  EndContainerPath string
  EndContainerChildIndex string
  EndOffset string
  Text string
  Annotation string
  ExtraAnnotationData string
  DateCreated string
  ChapterProgress string
  Hidden string
  Version string
  DateModified string
  Creator string
  UUID string
  UserID string
  SyncTime string
  Published string
  ContextString string
}

func (Bookmark) TableName() string {
  return "Bookmark"
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

func getBasicKoboDetails() Kobo {
  return selectedKobo
}

func getHighlightCount() int64 {
  var bookmarks []Bookmark
  dbPath := fmt.Sprintf("%s/.kobo/KoboReader.sqlite", selectedKobo.MntPath)
  db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  result := db.Find(&bookmarks)
  if result.Error != nil {
    panic(err)
  }
  return result.RowsAffected
}

func getMostRecentHighlight() Bookmark {
  var bookmark Bookmark
  dbPath := fmt.Sprintf("%s/.kobo/KoboReader.sqlite", selectedKobo.MntPath)
  db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  result := db.Clauses(clause.OrderBy{
    Expression: clause.Expr{SQL: "RANDOM()",},
  }).Take(&bookmark)
  if result.Error != nil {
    panic(err)
  }
  return bookmark
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
  app.Bind(getBasicKoboDetails)
  app.Bind(getHighlightCount)
  app.Bind(getMostRecentHighlight)
  app.Run()
}

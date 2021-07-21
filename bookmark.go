package main

import (
  _ "embed"
  "fmt"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "gorm.io/gorm/clause"
)

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

func NewBookmark() *Bookmark {
  return &Bookmark{}
}

func (Bookmark) TableName() string {
  return "Bookmark"
}

func (Bookmark) GetHighlightCount() int64 {
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

func (Bookmark) GetMostRecentHighlight() Bookmark {
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

package main

import (
  _ "embed"
)

type Content struct {
  ContentID string `gorm:"column:ContentID"`
  ContentType string `gorm:"column:ContentType"`
  MimeType string `gorm:"column:MimeType"`
  BookID string
  BookTitle string `gorm:"column:BookTitle"`
  ImageId string
  Title string
  Attribution string
  Description string
  DateCreated string
  ShortCoverKey string
  AdobeLocation string `gorm:"column:adobe_location"`
  Publisher string
  IsEncrypted bool
  DateLastRead string
  FirstTimeReading bool
  ChapterIDBookmarked string
  ParagraphBookmarked int
  BookmarkWordOffset int
  NumShortcovers int
  VolumeIndex int `gorm:"column:VolumeIndex"`
  NumPages int `gorm:"column:___NumPages"`
  ReadStatus int
  SyncTime string `gorm:"column:___SyncTime"`
  UserID string `gorm:"column:___UserID"`
  PublicationId string
  FileOffset int `gorm:"column:___FileOffset"`
  FileSize int `gorm:"column:___FileSize"`
  PercentRead string `gorm:"column:___PercentRead"`
  ExpirationStatus int `gorm:"column:___ExpirationStatus"`
  CurrentChapterProgress float32
}

func NewContent() *Content {
  return &Content{}
}

func (Content) TableName() string {
  return "Content"
}

func (Content) GetAllItems() []Content {
  var content []Content
  result := DBConn.Where(
    &Content{ContentType: "6", MimeType: "application/x-kobo-epub+zip", VolumeIndex: -1},
  ).Order("___PercentRead desc, title asc").Find(&content)
  if result.Error != nil {
    panic(result.Error)
  }
  return content
}
package main

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "math/rand"
  "time"
)

var db *gorm.DB

func init() {
  var err error
  db, err = gorm.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/airmeet?parseTime=True&loc=Japan")

  if err != nil {
    panic("failed to connect database")
    return
  }

  db.DB()
}

func CreateEvent(event *Event) {
  //db.NewRecord(event)
  db.Create(&event)
  //db.Save(&event)
}

func GenerateMajor() uint16 {
  rand.Seed(time.Now().UnixNano())

  var major int
  for ;; {
    major = rand.Intn(65535)
    fmt.Println("major = ",major)
    count := -1
    db.Model(&Event{}).Where("major = ?", major).Count(&count)
    fmt.Println("count = ",count)
    if count == 0 {
      break
    }
  }
  fmt.Println("major16 = ",uint16(major))
  return uint16(major)
}

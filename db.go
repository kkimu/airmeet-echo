package main

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "math/rand"
  "time"
  "github.com/k0kubun/pp"
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

func GenerateMajor() int {
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

  return major
}

func GetEvent(major int) (*Event, error) {
  var event Event

  if err := db.Where("major = ?", major).First(&event).Error; err != nil {
    return nil, err
  }
  pp.Println(&event)
  return &event, nil
}

func DeleteEvent(major int) (*Event, error) {
  var event Event
  if err := db.Where("major = ?", major).First(&event).Error; err != nil {
    return nil, err
  }
  if err := db.Where("major = ?", major).Delete(&event).Error; err != nil {
    return nil, err
  }
  pp.Println(event)
  return &event, nil
}

func CreateUser(user *User) {
  //db.NewRecord(event)
  db.Create(&user)
  //db.Save(&event)
}

func EventExist(major int) error {
  var event Event
  if err := db.Where("major = ?", major).First(&event).Error; err != nil {
    return err
  }
  return nil
}

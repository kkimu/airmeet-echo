package main

import (
  "fmt"
  "net/http"
  "time"

  "gopkg.in/go-playground/validator.v8"
  "github.com/labstack/echo"
  "github.com/k0kubun/pp"
  "strconv"
)

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}


// イベントを新規登録
func RegisterEvent(c echo.Context) error {
  fmt.Println("RegisterEvent")
  event := getPostEvent(c)
  pp.Println(event)

  if errs := validate.Struct(event); errs != nil {
    fmt.Println(errs)
    return c.JSON(http.StatusOK, NewError(400, fmt.Sprintf("%s",errs)))
  }

  event.Major = GenerateMajor()
  CreateEvent(event)
	return c.JSON(http.StatusOK, NewSuccess(event))
}

// イベント情報を取得
func GetEventInfo(c echo.Context) error {
  fmt.Println("GetEventInfo")
  major, err := strconv.Atoi(c.Param("major"))
  if err != nil || major < 0 || 65535 < major {
    return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
  }

  event, err := GetEvent(major)
  if err != nil {
    return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s",err)))
  } else {
    return c.JSON(http.StatusOK, NewSuccess(event))
  }
}

// イベントを削除
func RemoveEvent(c echo.Context) error {
  fmt.Println("RemoveEvent")
  major, err := strconv.Atoi(c.Param("major"))
  if err != nil || major < 0 || 65535 < major {
    return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
  }

  event, err := DeleteEvent(major)
  if err != nil {
    return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s",err)))
  } else {
    return c.JSON(http.StatusOK, NewSuccess(event))
  }
}

// Parse the request body, check input data
func getPostEvent(c echo.Context) *Event {
  // リクエストボディをパースして代入
  en,rn,desc,items := c.FormValue("eventName"),c.FormValue("roomName"),c.FormValue("description"),c.FormValue("items")

  return &Event{
      EventName: en,
      RoomName:	rn,
      Description: desc,
      Items: items,
      Active: true,
      CreatedAt: time.Now(),
  }
}

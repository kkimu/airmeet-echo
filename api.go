package main

import (
  "fmt"
  "net/http"
  "time"
  //"gopkg.in/validator.v2"
  "gopkg.in/go-playground/validator.v8"
  "github.com/labstack/echo"
  "github.com/k0kubun/pp"
)

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}


// RegisterEvent return the major id
func RegisterEvent() echo.HandlerFunc{
  return func(c echo.Context) error {
    fmt.Print("aaaaa")
    event := getPostEvent(c)
    pp.Print(event)
    //fmt.Print("ddddd")


    if errs := validate.Struct(event); errs != nil {
      fmt.Print("errrrrrrrrrrrrrr\n")
      fmt.Println(errs)
      /*
      err := errs.(validator.ValidationErrors)["Event.EventName"]
      fmt.Println(err.Field) // output: City
      fmt.Println(err.Tag)   // output: required
      fmt.Println(err.Kind)  // output: string
      fmt.Println(err.Type)  // output: string
      fmt.Println(err.Param) // output:
      fmt.Println(err.Value)
      */
      return c.JSON(http.StatusOK, NewError(400, fmt.Sprintf("%s",errs)))
    }

    /*
    if err := c.Bind(event); err != nil {
      fmt.Print("errrrrrrrrrrrrrr\n")
      //return echo.NewHTTPError(http.StatusBadRequest)
      return err
    }
    */
    //fmt.Print("eeeee")
    CreateEvent(event)
    //return c.JSON(http.StatusOK, event)
  	return c.JSON(http.StatusOK, NewSuccess(event))
  }
}

// Parse the request body, check input data
func getPostEvent(c echo.Context) *Event {
  //fmt.Print("bbbbb")
  // リクエストボディをパースして代入
  en,rn,desc,items := c.Form("eventName"),c.Form("roomName"),c.Form("description"),c.Form("items")
  //fmt.Print("ccccc")
  return &Event{
      EventName: en,
      RoomName:	rn,
      Description: desc,
      Items: items,
      Active: true,
      CreatedAt: time.Now(),
  }
}

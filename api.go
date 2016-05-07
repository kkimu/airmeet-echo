package main

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/go-playground/validator.v8"

	"strconv"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
}

// RegisterEvent イベントを新規登録
func RegisterEvent(c echo.Context) error {
	fmt.Println("RegisterEvent")
	event := getPostEvent(c)
	pp.Println(event)

	if errs := validate.Struct(event); errs != nil {
		fmt.Println(errs)
		return c.JSON(http.StatusOK, NewError(400, fmt.Sprintf("%s", errs)))
	}

	event.Major = GenerateMajor()
	CreateEvent(event)
	return c.JSON(http.StatusOK, NewSuccess(event))
}

// GetEventInfo イベント情報を取得
func GetEventInfo(c echo.Context) error {
	fmt.Println("GetEventInfo")
	major, err := strconv.Atoi(c.Param("major"))
	fmt.Println(major)
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}

	event, err := GetEvent(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(event))
}

// RemoveEvent イベントを削除
func RemoveEvent(c echo.Context) error {
	fmt.Println("RemoveEvent")
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}

	event, err := DeleteEvent(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(event))
}

// Parse the request body, check input data
func getPostEvent(c echo.Context) *Event {
	// リクエストボディをパースして代入
	en, rn, desc, items := c.FormValue("eventName"), c.FormValue("roomName"), c.FormValue("description"), c.FormValue("items")

	return &Event{
		EventName:   en,
		RoomName:    rn,
		Description: desc,
		Items:       items,
		CreatedAt:   time.Now(),
	}
}

// RegisterUser ユーザを新規登録
func RegisterUser(c echo.Context) error {
	fmt.Println("RegisterUser")
	user := getPostUser(c)

	major, err := majorConfirm(c)
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, fmt.Sprintf("%s", err)))
	}
	user.Major = major
	pp.Println(user)

	if errs := validate.Struct(user); errs != nil {
		return c.JSON(http.StatusOK, NewError(400, fmt.Sprintf("%s", errs)))
	}

	CreateUser(user)
	return c.JSON(http.StatusOK, NewSuccess(user))
}

// GetUser ユーザを取得
func GetUser(c echo.Context) error {
	fmt.Println("GetUser")
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}

	event, err := GetEvent(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(event))
}

// RemoveUser ユーザを削除
func RemoveUser(c echo.Context) error {
	fmt.Println("RemoveEvent")
	major, err := strconv.Atoi(c.Param("major"))
	if err != nil || major < 0 || 65535 < major {
		return c.JSON(http.StatusBadRequest, NewError(400, "major is invalid"))
	}

	event, err := DeleteEvent(major)
	if err != nil {
		return c.JSON(http.StatusNotFound, NewError(400, fmt.Sprintf("%s", err)))
	}

	return c.JSON(http.StatusOK, NewSuccess(event))
}

// Parse the request body, check input data
func getPostUser(c echo.Context) *User {
	// リクエストボディをパースして代入
	un, prof, items, img, imgh := c.FormValue("name"), c.FormValue("profile"), c.FormValue("items"), c.FormValue("image"), c.FormValue("image_header")

	return &User{
		UserName:    un,
		Profile:     prof,
		Items:       items,
		Image:       img,
		ImageHeader: imgh,
		CreatedAt:   time.Now(),
	}
}

func majorConfirm(c echo.Context) (int, error) {
	major, err1 := strconv.Atoi(c.Param("major"))
	if err1 != nil || major < 0 || 65535 < major {
		return -1, err1
	}
	if err2 := EventExist(major); err2 != nil {
		return -1, err2
	}
	return major, nil
}

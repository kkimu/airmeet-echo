package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Get("/", hello())
	e.Post("/events", RegisterEvent)
	e.Get("/events/:major", GetEventInfo)
	e.Delete("/events/:major", RemoveEvent)

	e.Post("/users", RegisterUser)
	e.Get("/users/:major/:id", GetUser)
	e.Delete("/users/:id", RemoveUser)
	// Start server
	e.Run(standard.New(":3000"))
}

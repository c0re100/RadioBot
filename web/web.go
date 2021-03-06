package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/c0re100/RadioBot/config"
	"github.com/c0re100/RadioBot/telegram"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	server *echo.Echo
)

// StartServer start http server for vote system
func StartServer() {
	if !config.IsWebEnabled() {
		fmt.Println("Switching to Userbot mode")
		return
	}

	server = echo.New()
	server.HideBanner = true

	//server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/", hello)
	server.POST("/ptcp", recvPtcp)
	server.POST("/reset", resetPtcps)

	server.Logger.Fatal(server.Start(":" + strconv.Itoa(config.GetWebPort())))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func recvPtcp(c echo.Context) error {
	status := c.FormValue("is_join")
	userID := c.FormValue("user_id")

	if status == "" {
		return c.HTML(400, "Field `is_join` is empty.")
	} else if userID == "" {
		return c.HTML(400, "Field `user_id` is empty.")
	}

	if status == "true" {
		telegram.AddPtcp(userID)
		return c.HTML(200, "User added.")
	} else if status == "false" {
		telegram.RemovePtcp(userID)
		return c.HTML(200, "User removed.")
	} else {
		return c.HTML(400, "Field `user_id` is empty or wrong type.")
	}
}

func resetPtcps(c echo.Context) error {
	telegram.ResetPtcps()
	return c.HTML(200, "Resetted.")
}

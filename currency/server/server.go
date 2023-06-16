package server

import (
	"currency/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func StartServer() {
	e.HideBanner = true
	bs, _ := os.ReadFile("../server/banner.txt")
	fmt.Println(string(bs))

	handlers.CurrencyRoutes(e)

	log.Fatal(e.Start(":8080"))
}

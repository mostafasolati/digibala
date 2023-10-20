package main

import (
	echo "github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()
	UserRoutes(server)
	VoucherRoutes(server)
	server.Start("localhost:6060")
}

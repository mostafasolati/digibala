package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func socialRoutes(e *echo.Echo) {
	e.GET("/social", listSocialHandler)
	e.POST("/social", createSocialHandler)
	e.GET("/social/:id", getSocialHandler)
	e.DELETE("/social/:id", deleteSocialHandler)
	e.PUT("/social/:id", updateSocialHandler)
	e.PATCH("/social/:id/logo", updateSocialLogoHandler)
}

func listSocialHandler(c echo.Context) error {
	socials := []*models.Social{}
	return c.JSON(http.StatusOK, socials)
}
func createSocialHandler(c echo.Context) error {
	social := &models.Social{}
	if err := c.Bind(social); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, social)
}
func getSocialHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	social := &models.Social{ID: id}
	return c.JSON(http.StatusOK, social)
}
func deleteSocialHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting social id: ", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}
func updateSocialHandler(c echo.Context) error {
	social := &models.Social{}
	if err := c.Bind(social); err != nil {
		return err
	}
	fmt.Println("Updatig social id: ", social.ID)
	return c.JSON(http.StatusOK, social)
}
func updateSocialLogoHandler(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, models.StatusOK{OK: "Logo updated"})
}

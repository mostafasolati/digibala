package server

import (
	"fmt"
	"net/http"
	"strconv"
	"digibala/models"
	"github.com/labstack/echo/v4"
)

func countryRoutes(e *echo.Echo) {
	e.GET("/country", listCountryHandler)
	e.GET("/country", createCountryHandler)
	e.GET("/country/:id", findCountryHandler)
	e.DELETE("/country/:id", deleteCountryHandler)
	e.PUT("/country", updateCountryHandler)
}

func createCountryHandler(c echo.Context) error {
	country := &models.Country{}
	if err := c.Bind(country); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, country)
}

func listCountryHandler(c echo.Context) error {
	countries := []models.Country{}
	return c.JSON(http.StatusOK, countries)
}

func findCountryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.Param("name")
	country := models.Country{CountryID:id, CountryName:name}
	return c.JSON(http.StatusOK, country)
}

func deleteCountryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Deleting country with id:%v\n", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func updateCountryHandler(c echo.Context) error {
	country := &models.Country{}
	if err:= c.Bind(country); err != nil {
		return err
	}
	fmt.Printf("Updating country id:%v\n", country.CountryID)
	return c.JSON(http.StatusOK, country)
}

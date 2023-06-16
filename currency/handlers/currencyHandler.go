package handlers

import (
	"currency/models"
	"currency/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func listCurrencyHandler(c echo.Context) error {
	currencies, err := service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, currencies)
}

func getCurrencyByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	currency, err := service.FindById(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, "Not Found Currency with Id")
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, currency)
}

func createCurrencyHandler(c echo.Context) error {
	currency := new(models.CurrencyRequest)

	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency data")
	}

	response, err := service.Create(currency)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create currency")
	}

	return c.JSON(http.StatusCreated, response)
}

func updateCurrencyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	currency := new(models.CurrencyRequest)
	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency data")
	}

	response, err := service.Update(id, currency)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update currency")
	}

	return c.JSON(http.StatusOK, response)
}

func deleteCurrencyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid currency ID")
	}

	err = service.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete currency")
	}

	return c.NoContent(http.StatusOK)
}

func CurrencyRoutes(e *echo.Echo) {
	e.GET("/api/currency", listCurrencyHandler)
	e.GET("/api/currency/:id", getCurrencyByIDHandler)
	e.POST("/api/currency", createCurrencyHandler)
	e.PUT("/api/currency/:id", updateCurrencyHandler)
	e.DELETE("/api/currency/:id", deleteCurrencyHandler)
}

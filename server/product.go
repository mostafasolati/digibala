package server

import (
	"digibala/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func productRoutes(e *echo.Echo) {
	e.GET("/product", indexProductHandler)
	e.GET("/product/:id", findProductHandler)
	e.POST("/product", createProductHandler)
	e.PUT("/product/:id", updateProductHandler)
	e.DELETE("/product/:id", deleteProductHandler)
}

// product routes comes here
func indexProductHandler(c echo.Context) error {
	var products []models.Product
	return c.JSON(http.StatusOK, products)
}

func findProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := models.Product{
		ID: id,
	}
	return c.JSON(http.StatusOK, product)
}

func createProductHandler(c echo.Context) error {
	product := &models.Product{}

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, product)
}

func updateProductHandler(c echo.Context) error {
	product := &models.Product{}
	id, _ := strconv.Atoi(c.Param("id"))
	_ = id
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func deleteProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = id

	return c.JSON(http.StatusOK, "product deleted successfully")
}

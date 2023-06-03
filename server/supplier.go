package server

import (
	"digibala/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var suppliers []models.Supplier = []models.Supplier{
	models.Supplier{
		ID: 1,
		CompanyName: "Company 1",
		Address: models.Address{
			ID: 1,
			Address: "Esfahan, ...",
		},
	},
	models.Supplier{
		ID: 2,
		CompanyName: "Company 2",
		Address: models.Address{
			ID: 2,
			Address: "Tehran, ...",
		},
	},
	models.Supplier{
		ID: 3,
		CompanyName: "Company 3",
		Address: models.Address{
			ID: 3,
			Address: "Tehran, ...",
		},
	},

}

// Find supplier and return its index and value
func findSupplier(id int) (int, *models.Supplier) {
	for i, s := range suppliers {
		if s.ID == id {
			return i, &s
		}
	}

	return -1, nil
}

func supplierRoutes(e *echo.Echo) {
	e.GET("/suppliers", listSupplierHandler)
	e.POST("/suppliers", createSupplierHandler)
	e.GET("/suppliers/:id", findSupplierHandler)
	e.DELETE("/suppliers/:id", deleteSupplierHandler)
	e.PUT("/suppliers/:id", updateSupplierHandler)
}

func listSupplierHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, suppliers)
}

func findSupplierHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, supplier := findSupplier(id)
	if supplier == nil {
		return c.JSON(http.StatusNotFound, "Product does not exist!")
	}

	return c.JSON(http.StatusOK, supplier)
}

func createSupplierHandler(c echo.Context) error {
	supplier := &models.Supplier{}
	if err := c.Bind(supplier); err != nil {
		return err
	}
	suppliers = append(suppliers, *supplier)
	return c.JSON(http.StatusCreated, supplier)
}

func deleteSupplierHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	i, supplier := findSupplier(id)
	if supplier == nil {
		return c.JSON(http.StatusNotFound, "Product does not exist!")
	}

	suppliers = append(suppliers[:i], suppliers[i+1:]...)
	return c.JSON(http.StatusNoContent, models.StatusOK{OK: "OK"})
}

func updateSupplierHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	i, s := findSupplier(id)
	if s == nil {
		return c.JSON(http.StatusNotFound, "Product does not exist!")
	}

	supplier := &models.Supplier{}
	if err := c.Bind(supplier); err != nil {
		return err
	}

	supplier.ID = id
	suppliers[i] = *supplier
	return c.JSON(http.StatusOK, supplier)
}

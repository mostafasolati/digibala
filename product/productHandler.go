package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func ProductRoutes(server *echo.Echo) {
	psi := NewProductService()
	//Product Create
	server.POST("/product/create/:Name/:Desc/:CategoryId/:Price/:StorageQuantity", generateCreateProductHandler(psi))
	//Product Update
	server.POST("/product/update/:Name/:Desc/:CategoryId/:Price/:StorageQuantity", generateUpdateProductHandler(psi))
	//Product Delete
	server.GET("/product/delete/:Id", generateDeleteProductHandler(psi))
	//Product Find
	server.GET("/product/find/:Id", generateFindProductHandler(psi))
	//Product List
	server.GET("/product/list", generateListProductHandler(psi))
}

func generateCreateProductHandler(psi ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := &Product{
			Id: lastUsedId,
		}
		lastUsedId++
		if err := c.Bind(product); err != nil {
			return err
		}
		err := psi.Create(product)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	}
}

func generateUpdateProductHandler(psi ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := new(Product)
		if err := c.Bind(product); err != nil {
			return err
		}
		err := psi.Update(product.Id, *product)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	}
}
func generateDeleteProductHandler(psi ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := new(Product)
		if err := c.Bind(product); err != nil {
			return err
		}
		err := psi.Delete(product.Id)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "OK")
	}
}
func generateFindProductHandler(psi ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := new(Product)
		if err := c.Bind(product); err != nil {
			return err
		}
		p, err := psi.Find(product.Id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, p)
	}
}
func generateListProductHandler(psi ServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		product := new(Product)
		if err := c.Bind(product); err != nil {
			return err
		}
		lp, err := psi.List()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, lp)
	}
}

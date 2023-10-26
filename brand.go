package main

import (
	"fmt"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var brands = make([]*Brands, 6)

func init() {
	for i := 1; i <= 5; i++ {
		brands[i] = &Brand{
			Name:        fmt.Sprint("Name", i),
			Description: fmt.Sprint("Description", i),
			Image:       fmt.Sprint("Image", i),
		}
	}
}

type Brands struct {
	Id          int
	Name        string
	Description string
	Image       string
}

type BrandsServiceInterface interface {
	Add(*Brands) error
	Update(*Brands) error
	Delete(int) error
	Find(int) (*Brands, error)
	List() ([]*Brands, error)
}

type brandsService struct{}

func NewBrandService() BrandsServiceInterface {
	return &brandsService{}
}

func (b *brandsService) Add(brand *Brands) error {
	fmt.Println("Add Brands:", brand)
	return nil
}

func (b *brandsService) Update(brand *Brands) error {
	fmt.Println("Update Brands")
	return nil
}

func (b *brandsService) Delete(id int) error {
	fmt.Println("Deleting Brands id:", id)
	return nil
}

func (b *brandsService) Find(id int) (*Brands, error) {
	return brands[id], nil
}

func (b *brandsService) List() ([]*Brands, error) {
	return brands, nil
}

func findBrandsHandler(service BrandsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var err error
		var brand *Brands
		brand, err = service.Find(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, brand)
	}
}

func listBrandsHandler(service BrandsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		brands, err := service.List()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, brands)
	}
}

func AddBrands(service BrandsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		brand := new(Brands)
		if err := c.Bind(brand); err != nil {
			return err
		}
		err := service.Add(brand)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, brand)
	}
}

type Message struct {
	Message string
}

func updateBrands(service BrandsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		brand := new(Brands)
		if err := c.Bind(brand); err != nil {
			return err
		}
		if err := service.Update(brand); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "Error on Update brands"})
		}
		return c.JSON(http.StatusOK, brand)
	}
}

func deleteHandler(service BrandsServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		service.Delete(id)
		return c.JSON(http.StatusOK, Message{Message: "Brands deleted successfully"})
	}
}

func BrandsRoutes(server *echo.Echo) {
	brandService := NewBrandService()
	server.POST("/brand", AddBrands(brandService))
	server.PUT("/brand", updateBrands(brandService))
	server.DELETE("/brand/:id", deleteHandler(brandService))
	server.GET("/brand/:id", findBrandsHandler(brandService))
	server.GET("/brand", listBrandsHandler(brandService))
}

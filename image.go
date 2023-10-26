package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Image struct {
	Id   int    `json:"id"`
	Size string `json:"size"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type imgService struct{}

var Images = make([]*Image, 2)

type ImgServiceInterface interface {
	AddImg(*Image) error
	UpdateImg(*Image) error
	DeleteImg(int) error
	FindImg(int) (*Image, error)
	List() ([]*Image, error)
}

func NewImgServices() ImgServiceInterface {
	return &imgService{}
}

func (img *imgService) AddImg(image *Image) error {
	fmt.Println("Add Image", image)
	return nil
}

func (img *imgService) UpdateImg(image *Image) error {
	if strconv.Itoa(image.Id) == "" {
		return errors.New("Image Id is empty")
	}
	fmt.Println("Update Image", image)
	return nil
}

func (img *imgService) DeleteImg(id int) error {
	fmt.Println("Delete Image", id)
	return nil
}

func (img *imgService) FindImg(id int) (*Image, error) {
	fmt.Println("Find Image", id)
	return Images[id], nil
}

func (img *imgService) List() ([]*Image, error) {
	return Images, nil
}

func addImgHandler(service ImgServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		img := new(Image)
		if err := c.Bind(img); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		}
		if err := service.AddImg(img); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to add image"})
		}
		return c.JSON(http.StatusOK, img)
	}
}

func updateImgHandler(service ImgServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		img := new(Image)
		if err := c.Bind(img); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		}
		if err := service.UpdateImg(img); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Failed to update image"})
		}
		return c.JSON(http.StatusOK, img)
	}
}

func deleteImgHandler(service ImgServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid image ID"})
		}
		if err := service.DeleteImg(id); err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete image"})
		}
		return c.JSON(http.StatusOK, ErrorResponse{Error: "Image deleted successfully"})
	}
}

func listImageHandler(service ImgServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		images, err := service.List()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, images)
	}
}

func findImgHandler(service ImgServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var err error
		var img *Image
		img, err = service.FindImg(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, img)
	}
}

func ImageRoutes(server *echo.Echo) {
	imageService := NewImgServices()
	server.POST("/img", addImgHandler(imageService))
	server.PUT("/img", updateImgHandler(imageService))
	server.DELETE("/img/:id", deleteImgHandler(imageService))
	server.GET("/img/:id", findImgHandler(imageService))
	server.GET("/img", listImageHandler(imageService))
}

package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type DeliveryType struct {
	Id            int
	Name          string
	AvailibleDays []string
	Price         float64
}

var deliveryTypes = make([]*DeliveryType, 3)

func init() {
	deliveryTypes[0] = &DeliveryType{
		Id:            1,
		Name:          "post",
		AvailibleDays: []string{"shanbe", "yekshanbe"},
		Price:         200,
	}
	deliveryTypes[1] = &DeliveryType{
		Id:            2,
		Name:          "tipax",
		AvailibleDays: []string{"yekshanbe", "doshanbe"},
		Price:         100,
	}
	deliveryTypes[2] = &DeliveryType{
		Id:            3,
		Name:          "snappbox",
		AvailibleDays: []string{"doshanbe", "seshanbe"},
		Price:         300,
	}
}

type DeliveryTypeServiceInterface interface {
	Create(*DeliveryType) error
	Update(id int, field string, value string) error
	Delete(int) error
	Find(int) (*DeliveryType, error)
	List() ([]*DeliveryType, error)
}

type DeliveryTypeService struct{}

func NewDeliveryTypeService() DeliveryTypeServiceInterface {
	return &DeliveryTypeService{}
}

func (s *DeliveryTypeService) Create(NewDType *DeliveryType) error {
	deliveryTypes = append(deliveryTypes, NewDType)
	return nil
}

func (s *DeliveryTypeService) Update(id int, field string, value string) error {
	for _, dType := range deliveryTypes {
		if dType.Id == id {
			switch strings.ToLower(field) {
			case "name":
				dType.Name = value
			case "price":
				price, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				dType.Price = price
			case "availibledays":
				days := strings.Split(value, ",")
				dType.AvailibleDays = days
			default:
				return errors.New("invalid field")
			}
			return nil
		}
	}
	return errors.New("delivery type not found")
}

func (s *DeliveryTypeService) Delete(id int) error {
	if id >= 0 && id < len(deliveryTypes) {
		deliveryTypes = append(deliveryTypes[:id], deliveryTypes[id+1:]...)
		return nil
	}
	return errors.New("delivery type not found")
}

func (s *DeliveryTypeService) Find(id int) (*DeliveryType, error) {
	var dType *DeliveryType
	if id >= 0 && id < len(deliveryTypes) {
		dType = deliveryTypes[id]
		return dType, nil
	}
	return nil, errors.New("delivery type not found")
}

func (s *DeliveryTypeService) List() ([]*DeliveryType, error) {
	return deliveryTypes, nil
}

func createDeliveryTypeHandler(service DeliveryTypeServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		dType := new(DeliveryType)
		if err := c.Bind(dType); err != nil {
			return err
		}
		service.Create(dType)
		return c.JSON(http.StatusCreated, dType)
	}
}

func updateDeliveryTypeHandler(service DeliveryTypeServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		field := c.QueryParam("field")
		value := c.QueryParam("value")
		err := service.Update(id, field, value)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, Message{Message: "delivery type updated successfully"})
	}
}

func deleteDeliveryTypeHandler(service DeliveryTypeServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		err := service.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, Message{Message: "delivery type deleted successfully"})
	}
}

func findDeliveryTypeHandler(service DeliveryTypeServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		dType, err := service.Find(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, dType)
	}
}

func listDeliveryTypesHandler(service DeliveryTypeServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		dTypes, err := service.List()
		if err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, dTypes)
	}
}

func DeliveryTypesRoutes(server *echo.Echo) {
	deliveryTypesService := NewDeliveryTypeService()
	server.POST("/dtype", createDeliveryTypeHandler(deliveryTypesService))
	server.PATCH("/dtype/:id/:field/:value", updateDeliveryTypeHandler(deliveryTypesService))
	server.DELETE("/dtype/:id", deleteDeliveryTypeHandler(deliveryTypesService))
	server.GET("/dtype/:id", findDeliveryTypeHandler(deliveryTypesService))
	server.GET("/dtype", listDeliveryTypesHandler(deliveryTypesService))
}

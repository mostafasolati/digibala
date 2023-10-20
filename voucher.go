package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Voucher struct {
	Id              int
	Title           string
	Description     string
	ExpireDate      time.Time
	CategoryId      int
	DiscountPrecent int
	DiscountAmount  int
}

type vocherService struct{}

type VoucherServiceInterface interface {
	Create(*Voucher) error
	Update(*Voucher) error
	Delete(int) error
	Find(int) (*Voucher, error)
	List() ([]*Voucher, error)
}

func NewVoucherService() VoucherServiceInterface {
	return &vocherService{}
}

func (v *vocherService) Create(voucher *Voucher) error {
	fmt.Println("Voucher created: ", voucher)
	return nil
}

func (v *vocherService) Update(voucher *Voucher) error {
	fmt.Println("Voucher updated: ", voucher)
	return nil
}

func (v *vocherService) Delete(id int) error {
	fmt.Println("Voucher deleted: ", id)
	return nil
}

func (v *vocherService) Find(id int) (*Voucher, error) {
	fmt.Println("Voucher found: ", id)
	return nil, nil
}

func (v *vocherService) List() ([]*Voucher, error) {
	fmt.Println("Vouchers found")
	return nil, nil
}

func createVoucherHandler(service VoucherServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		voucher := new(Voucher)
		if err := c.Bind(voucher); err != nil {
			return err
		}
		service.Create(voucher)
		return c.JSON(http.StatusOK, voucher)
	}
}

func updateVoucherHandler(service VoucherServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		voucher := new(Voucher)
		if err := c.Bind(voucher); err != nil {
			return err
		}
		if err := service.Update(voucher); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "Error on updating voucher"})
		}
		return c.JSON(http.StatusOK, voucher)
	}
}

func deleteVoucherHandler(service VoucherServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "Error on deleting voucher"})
		}
		return c.JSON(http.StatusOK, id)
	}
}

func findVoucherHandler(service VoucherServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var voucher *Voucher
		var err error
		voucher, err = service.Find(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, Message{Message: "Voucher not found"})
		}
		return c.JSON(http.StatusOK, voucher)
	}
}

func listVoucherHandler(service VoucherServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		vouchers, err := service.List()
		if err != nil {
			return c.JSON(http.StatusNotFound, Message{Message: "Vouchers not found"})
		}
		return c.JSON(http.StatusOK, vouchers)
	}
}

func VoucherRoutes(server *echo.Echo) {
	vocherService := NewVoucherService()
	server.POST("/voucher", createVoucherHandler(vocherService))
	server.PUT("/voucher", updateVoucherHandler(vocherService))
	server.DELETE("/voucher", deleteVoucherHandler(vocherService))
	server.GET("/voucher/:id", findVoucherHandler(vocherService))
	server.GET("/voucher", listVoucherHandler(vocherService))
}

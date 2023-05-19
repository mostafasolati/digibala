package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func faqRoutes(e *echo.Echo) {
	e.GET("/faqs", listAllFAQHandler)
	e.GET("/faq/:id", findSpecificFAQHandler)
	e.POST("/faq", createFAQHandler)
	e.DELETE("/faqs/:id", deleteFAQHandler)
	e.PUT("/faq/:id", updateSpecificFAQHandler)
}
func listAllFAQHandler(c echo.Context) error {
	faqs := []models.FAQ{}
	return c.JSON(http.StatusOK, faqs)
}
func findSpecificFAQHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	requestedFAQ := &models.FAQ{ID: id}
	return c.JSON(http.StatusOK, requestedFAQ)
}
func createFAQHandler(c echo.Context) error {
	faq := &models.FAQ{}
	if err := c.Bind(faq); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, faq)
}
func deleteFAQHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("FAQ with id %d deleted", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func updateSpecificFAQHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("FAQ with id %d updated", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

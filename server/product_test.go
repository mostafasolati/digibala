package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductHandler(t *testing.T) {
	e := echo.New()
	productRoutes(e)

	product := &models.Product{
		ID:          1,
		Title:       "Design Patter With Go",
		Price:       120000,
		Quantity:    10,
		Description: "This book writes by Mostafa Solati",
		Image:       "./book1.jpg",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(product)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Product)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, product, ret)
}

func TestIndexProductHandler(t *testing.T) {
	e := echo.New()
	productRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	products := []*models.Product{}

	err := json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Len(t, products, 0)

}

func TestUpdateProductHandler(t *testing.T) {
	e := echo.New()
	productRoutes(e)

	product := &models.Product{
		ID:    1,
		Title: "Design Patter With Go",
		Price: 120000,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(product)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/product/1", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Product)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, product, ret)
}

func TestFindProductHandler(t *testing.T) {
	e := echo.New()
	productRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/product/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Product)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, &models.Product{ID: 1}, ret)
}

func TestDeleteProductHandler(t *testing.T) {
	e := echo.New()
	productRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/product/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()

	var message string
	err := json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "product deleted successfully", message)
}

package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListSupplier(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(suppliers)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/suppliers", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	resp := w.Result()
	ret := new([]models.Supplier)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &suppliers, ret)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindSupplier(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	supplier := suppliers[0]
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(suppliers[0])

	w := httptest.NewRecorder()
	target_path := fmt.Sprintf("/suppliers/%d", supplier.ID)
	r := httptest.NewRequest(http.MethodGet, target_path, buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	resp := w.Result()
	ret := new(models.Supplier)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, &supplier, ret)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFindSupplier404Error(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	buf := new(bytes.Buffer)

	w := httptest.NewRecorder()
	target_path := fmt.Sprintf("/suppliers/%d", 4162)
	r := httptest.NewRequest(http.MethodGet, target_path, buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateSupplier(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	supplier := &models.Supplier{
		ID: 6,
		CompanyName: "Company 6",
		Address: models.Address{ID: 6, Address: "Address 6"},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(supplier)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/suppliers", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	resp := w.Result()
	ret := new(models.Supplier)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	_, s := findSupplier(supplier.ID)

	assert.Equal(t, supplier, ret)
	assert.Equal(t, supplier.ID, s.ID)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateSupplier(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	supplier := &models.Supplier{
		CompanyName: "Company 1 ",
		Address: models.Address{ID: 1, Address: "Address 1"},
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(supplier)

	w := httptest.NewRecorder()
	target_path := fmt.Sprintf("/suppliers/%d", 1)
	r := httptest.NewRequest(http.MethodPut, target_path, buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	resp := w.Result()
	ret := new(models.Supplier)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, supplier.CompanyName, ret.CompanyName)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateSupplier404Error(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	buf := new(bytes.Buffer)

	w := httptest.NewRecorder()
	target_path := fmt.Sprintf("/suppliers/%d", 45564)
	r := httptest.NewRequest(http.MethodPut, target_path, buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteSupplier(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)

	w := httptest.NewRecorder()
	target_path := fmt.Sprintf("/suppliers/%d", 1)
	r := httptest.NewRequest(http.MethodDelete, target_path, nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)
	i, _ := findSupplier(1)

	assert.Equal(t, -1, i)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteSupplier404Error(t *testing.T) {
	e := echo.New()
	supplierRoutes(e)
	buf := new(bytes.Buffer)

	w := httptest.NewRecorder()
	target_path := fmt.Sprintf("/suppliers/%d", 45564)
	r := httptest.NewRequest(http.MethodDelete, target_path, buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
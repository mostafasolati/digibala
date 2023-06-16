package handlers

import (
	"bytes"
	"currency/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCurrencyHandler(t *testing.T) {
	e := echo.New()
	CurrencyRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/currency", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new([]*models.CurrencyResponse)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, new([]*models.CurrencyResponse), ret)
}

func TestGetCurrencyByIDHandler(t *testing.T) {
	e := echo.New()
	CurrencyRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/currency/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.CurrencyResponse)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, &models.CurrencyResponse{ID: 1}, ret)
}

func TestCreateCurrencyHandler(t *testing.T) {
	e := echo.New()
	CurrencyRoutes(e)

	currency := &models.CurrencyRequest{
		Code:   "IR",
		Name:   "Rial",
		Symbol: "IR",
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(currency)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/currency", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.CurrencyResponse)
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
	assert.Equal(t, currency, ret)
}
func TestUpdateCurrencyHandler(t *testing.T) {
	e := echo.New()
	CurrencyRoutes(e)

	currency := &models.CurrencyRequest{
		Code:   "IR",
		Name:   "Rial",
		Symbol: "IR",
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(currency)
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/api/currency/1", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.CurrencyResponse)
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, currency, ret)
}
func TestDeleteCurrencyHandler(t *testing.T) {
	e := echo.New()
	CurrencyRoutes(e)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/api/currency/1", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

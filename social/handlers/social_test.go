package server

import (
	"bytes"
	"digibala/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListSocial(t *testing.T) {
	e := echo.New()
	socialRoutes(e)

	socials := []*models.Social{}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/social", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := []*models.Social{}
	err := json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, socials, ret)
}
func TestCreateSocial(t *testing.T) {
	e := echo.New()
	socialRoutes(e)

	social := &models.Social{
		ID:   5,
		Name: "Twitter",
		URL:  "https://twitter.com/hasssanitman",
		Logo: "twitter.png",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(social)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/social", buf)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Social)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, social, ret)
}
func TestGetSocial(t *testing.T) {
	e := echo.New()
	socialRoutes(e)

	social := &models.Social{
		ID: 5,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/social/"+strconv.Itoa(social.ID), nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Social)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, social, ret)
}
func TestDeleteSocial(t *testing.T) {
	e := echo.New()
	socialRoutes(e)

	status := &models.StatusOK{
		OK: "OK",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/social/5", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.StatusOK)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, status, ret)
}
func TestUpdateSocial(t *testing.T) {
	e := echo.New()
	socialRoutes(e)

	social := &models.Social{
		ID: 0,
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/social/5", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.Social)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, social, ret)
}
func TestUpdateLogoSocial(t *testing.T) {
	e := echo.New()
	socialRoutes(e)

	status := &models.StatusOK{
		OK: "Logo updated",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/social/5/logo", nil)
	r.Header.Set("content-type", "application/json")
	e.ServeHTTP(w, r)

	resp := w.Result()
	ret := new(models.StatusOK)
	err := json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, status, ret)
}

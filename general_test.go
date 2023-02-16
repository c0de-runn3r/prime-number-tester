package main

import (
	"net/http"
	"net/http/httptest"
	"prime-number-tester/handlers"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestNumbersRequestSimple(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("[1,2,3,4,5]"))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.HandleNumbersRequest(c)

	assert(t, "[false,true,true,false,true]\n", rec.Body.String())

	assert(t, http.StatusOK, rec.Code)
}

func TestNumbersRequestError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("[a1,2,3,4,5]"))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.HandleNumbersRequest(c)

	assert(t, "{\"error\":\"the given input is invalid. Element on index 0 is not a number\"}\n", rec.Body.String())

	assert(t, http.StatusBadRequest, rec.Code)
}

func TestNumbersRequestError2(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("1,2,3,4,5abc"))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.HandleNumbersRequest(c)

	assert(t, "{\"error\":\"the given input is invalid. Element on index 4 is not a number\"}\n", rec.Body.String())

	assert(t, http.StatusBadRequest, rec.Code)
}

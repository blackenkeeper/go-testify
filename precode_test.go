package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandleWhenOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=moscow", nil)

	resp := httptest.NewRecorder()
	reqHandler := http.HandlerFunc(mainHandle)
	reqHandler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.NotEmpty(t, resp.Body.String())
}

func TestMainHandleWhenCityIncorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=svetlogorsk", nil)

	resp := httptest.NewRecorder()
	reqHandler := http.HandlerFunc(mainHandle)
	reqHandler.ServeHTTP(resp, req)

	respBodyExpect := "wrong city value"

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, respBodyExpect, resp.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(resp, req)

	cafesCount := len(strings.Split(resp.Body.String(), ","))
	assert.Equal(t, totalCount, cafesCount)
}

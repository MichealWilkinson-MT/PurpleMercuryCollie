package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

func TestAddHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/add?x=1&y=1", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addController)
	handler.ServeHTTP(rr, req)

	bodyBytes, _ := ioutil.ReadAll(rr.Body)

	assert.Equal(t, "2\n", string(bodyBytes))
}
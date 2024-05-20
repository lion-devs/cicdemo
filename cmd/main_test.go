package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-cicd/pkg/model"
	cicdRouter "go-cicd/pkg/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestResponse(t *testing.T) {

	router := cicdRouter.SetupRouter()
	res := performRequest(router, "GET", "/__health")

	assert.Equal(t, http.StatusOK, res.Code)

	var response model.Response
	err := json.Unmarshal([]byte(res.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, "ok", response.Message)
}

func TestAlertTest(t *testing.T) {

	router := cicdRouter.SetupRouter()
	res := performRequest(router, "GET", "/alert_test")

	assert.Equal(t, http.StatusServiceUnavailable, res.Code)

	var response model.Response
	err := json.Unmarshal([]byte(res.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, "Service Unavailable Test", response.Message)
}

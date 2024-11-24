package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"vespera-server/database"
	"vespera-server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setup() *gin.Engine {

	database.Connect()
	router := gin.Default()
	// Define routes for registering, login, and protected
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.GET("/protected", handlers.Protected)
	return router
}

func performRequest(router *gin.Engine, method, path string, body interface{}, token string) *http.Response {
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w.Result()
}

func TestUserRegistrationLoginAndProtectedRoute(t *testing.T) {
	// Setup router
	router := setup()

	// Step 1: Register a new user
	registerInput := map[string]string{
		"username": "testuser",
		"email":    "testuser@example.com",
		"password": "password123",
	}

	registerResp := performRequest(router, http.MethodPost, "/register", registerInput, "")
	assert.Equal(t, http.StatusOK, registerResp.StatusCode)

	// Step 2: Login to obtain the JWT token
	loginInput := map[string]string{
		"email":    "testuser@example.com",
		"password": "password123",
	}

	loginResp := performRequest(router, http.MethodPost, "/login", loginInput, "")
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	var loginResponse map[string]interface{}
	err := json.NewDecoder(loginResp.Body).Decode(&loginResponse)
	if err != nil {
		t.Fatalf("Failed to parse login response: %v", err)
	}

	token, ok := loginResponse["token"].(string)
	assert.True(t, ok, "Token should be present in the response")

	// Step 3: Access the protected route with the JWT token
	protectedResp := performRequest(router, http.MethodGet, "/protected", nil, token)
	assert.Equal(t, http.StatusOK, protectedResp.StatusCode)
}

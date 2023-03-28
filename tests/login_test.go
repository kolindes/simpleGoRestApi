package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/kolindes/simpleRestApi/internal/config"
	"github.com/kolindes/simpleRestApi/internal/models"
)

func TestSuccessfullLogin(t *testing.T) {
	config, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	username := "testUser"
	password := "testPassword321!"

	// ----------------------------------------------------------------
	// Register
	regBody := map[string]string{
		"username": username,
		"email":    username + "@example.com",
		"password": password,
	}

	regRequestBody, err := json.Marshal(regBody)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}

	regRequest, err := http.NewRequest("POST", "http://"+config.Main.Host+":"+config.Main.Port+"/register", bytes.NewBuffer(regRequestBody))
	if err != nil {
		t.Fatal(err)
	}
	regRequest.Header.Add("Content-Type", "application/json")

	// Send request -> response
	reqResponse, err := client.Do(regRequest)
	if err != nil {
		t.Fatal(err)
	}
	defer reqResponse.Body.Close()

	// Read body
	regResponseBody, err := io.ReadAll(reqResponse.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code
	if reqResponse.StatusCode != http.StatusCreated && reqResponse.StatusCode != http.StatusConflict {
		t.Errorf("Expected status code %d or %d, but got %d", http.StatusCreated, http.StatusConflict, reqResponse.StatusCode)
	}

	// Unmarshal into a Response struct
	var regResponseModel models.Response

	if err = json.Unmarshal(regResponseBody, &regResponseModel); err != nil {
		t.Fatal(err)
	}

	// ---- auth
	authBody := map[string]string{
		"username": username,
		"email":    username + "@example.com",
		"password": password,
	}

	authRequestBody, err := json.Marshal(authBody)
	if err != nil {
		t.Fatal(err)
	}

	authRequest, err := http.NewRequest("POST", "http://"+config.Main.Host+":"+config.Main.Port+"/login", bytes.NewBuffer(authRequestBody))
	if err != nil {
		t.Fatal(err)
	}
	authRequest.Header.Add("Content-Type", "application/json")

	// Send the request and receive the response
	authResponse, err := client.Do(authRequest)
	if err != nil {
		t.Fatal(err)
	}
	defer reqResponse.Body.Close()

	// Read the response body
	authResponseBody, err := io.ReadAll(authResponse.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if authResponse.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, authResponse.StatusCode)
	}

	// Unmarshal the regResponse body into a Response struct
	var authResponseModel models.Response

	if err = json.Unmarshal(authResponseBody, &authResponseModel); err != nil {
		t.Fatal(err)
	}

	// Check the response data and message
	if authResponseModel.Data == nil || authResponseModel.Message != "authorized" || authResponseModel.Error != "" {
		t.Errorf("Unexpected response: %v", authResponseModel)
	}
}

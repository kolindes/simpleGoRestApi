package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/kolindes/simpleRestApi/internal/config"
	"github.com/kolindes/simpleRestApi/internal/models"
	"github.com/kolindes/simpleRestApi/internal/utils"
)

func TestRegister(t *testing.T) {
	config, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	randomUserName := "user_" + utils.RandomString(5)
	body := map[string]string{
		"username": randomUserName,
		"email":    randomUserName + "@example.com",
		"password": "testPassword321!",
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://"+config.Main.Host+":"+config.Main.Port+"/register", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Send request -> response
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Read body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Check status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}

	// Unmarshal the response body into a Response struct
	var response models.Response

	if err = json.Unmarshal(responseBody, &response); err != nil {
		t.Fatal(err)
	}

	// Check response
	if response.Data == nil || response.Message != "registered" || response.Error != "" {
		t.Errorf("Unexpected response: %v", response)
	}
}

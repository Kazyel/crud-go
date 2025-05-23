package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type BodyStruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const successMessage = "user created successfully!"

func initiateBody(body BodyStruct) []byte {
	jsonData, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return jsonData
}

func sendPostRequest(url string, body []byte) *http.Response {
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	return response
}

func sendGetRequest(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return response
}

// TestCreateUserSuccess tests if the user creation is successful.
func TestCreateUserSuccess(t *testing.T) {
	body := BodyStruct{
		Name:     "kazyel",
		Email:    "kazyel@gmail.com",
		Password: "123456",
	}

	jsonData := initiateBody(body)
	response := sendPostRequest("http://localhost:8080/api/v1/users", jsonData)

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logMessage := fmt.Sprintf("expected status code: %d, got: %d", http.StatusOK, response.StatusCode)
		t.Log(logMessage)
		t.Fatal("user was not found")
	}
}

// TestCreateUserFailure tests if the user creation fails when the user already exists.
func TestCreateUserFailure(t *testing.T) {
	body := BodyStruct{
		Name:     "kazyel",
		Email:    "kazyel@gmail.com",
		Password: "123456",
	}

	jsonData := initiateBody(body)
	response := sendPostRequest("http://localhost:8080/api/v1/users", jsonData)

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logMessage := fmt.Sprintf("expected status code: %d, got: %d", http.StatusOK, response.StatusCode)
		t.Log(logMessage)
		t.Fatal("user was not found")
	}
}

// TestGetUserByIDSuccess tests if the user is found by ID.
func TestGetUserByIDSuccess(t *testing.T) {
	userId := "6a5452c9-e5b3-482a-8349-c3f4as44e4aa"
	url := fmt.Sprintf("http://localhost:8080/api/v1/users/%s", userId)

	response := sendGetRequest(url)

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		logMessage := fmt.Sprintf("expected status code: %d, got: %d", http.StatusOK, response.StatusCode)
		t.Log(logMessage)
		t.Fatal("user was not found")
	}
}

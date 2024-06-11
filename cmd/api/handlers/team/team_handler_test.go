package team

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"testing"
)

type CreateTeamRequest struct {
	Name string `json:"name"`
}

func NewCreateTeamRequest(name string) *CreateTeamRequest {
	return &CreateTeamRequest{Name: name}
}

type CreateTeamResponseSuccess struct {
	Name string `json:"name"`
}

func GetSuccessResponse(name string) CreateTeamResponseSuccess {
	return CreateTeamResponseSuccess{Name: name}
}

type CreateTeamResponseFail struct {
	Error string `json:"error"`
}

func GetErrorResponse(error string) CreateTeamResponseFail {
	return CreateTeamResponseFail{Error: error}
}

func getRandName() string {
	length := 10
	buff := make([]byte, length)
	_, err := rand.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:length]
}

func TestCreateTeamSuccess(t *testing.T) {
	url := "http://localhost:8080/createteam"
	method := "POST"
	name := getRandName()
	expectedResponse := GetSuccessResponse(name)
	payload := NewCreateTeamRequest(name)
	payloadBytes, _ := json.Marshal(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to create request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute request %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("Failed to close response body %v", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", res.StatusCode)
	}

	var apiResponse CreateTeamResponseSuccess
	err = json.Unmarshal(body, &apiResponse)

	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Errorf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

func TestCreateTeamFailure(t *testing.T) {
	url := "http://localhost:8080/createteam"
	method := "POST"

	expectedResponse := GetErrorResponse("Could not create team")
	payload := NewCreateTeamRequest("team1")
	payloadStr, _ := json.Marshal(payload)

	//payload := strings.NewReader(`{` + " " + `"name": "team8"` + " " + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadStr))

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("Failed to close response body %v", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status OK, got %v", res.StatusCode)
	}

	var apiResponse CreateTeamResponseFail
	err = json.Unmarshal(body, &apiResponse)

	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Errorf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

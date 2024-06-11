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

func GetCreateTeamSuccessResponse(name string) CreateTeamResponseSuccess {
	return CreateTeamResponseSuccess{Name: name}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func GetErrorResponse(error string) ErrorResponse {
	return ErrorResponse{Error: error}
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func GetSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{message}
}

type AddMemberSuccessResponse struct {
	TeamID uint `json:"team_id"`
	UserID uint `json:"user_id"`
}

func GetAddMemberSuccessResponse(teamId uint, userId uint) AddMemberSuccessResponse {
	return AddMemberSuccessResponse{TeamID: teamId, UserID: userId}
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
	expectedResponse := GetCreateTeamSuccessResponse(name)
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

	var apiResponse ErrorResponse
	err = json.Unmarshal(body, &apiResponse)

	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Errorf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

func TestAddMemberSuccess(t *testing.T) {
	url := "http://localhost:8080/teams/8/members/6"
	method := "POST"

	expectedResponse := GetAddMemberSuccessResponse(8, 6)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
		t.Fatalf("Failed to read response body%v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected Status OK, got %v", res.StatusCode)
		return
	}

	var apiResponse AddMemberSuccessResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal body: %v", err)
	}
	//fmt.Printf("response body: %+v\n, expected: %+v\n", apiResponse, expectedResponse)

	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Errorf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

func TestAddMemberFailure(t *testing.T) {
	url := "http://localhost:8080/teams/0/members/6"
	method := "POST"

	expectedResponse := GetErrorResponse("Could not add member")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
			t.Fatalf("Failed to close response body: %v", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status OK, got %v", res.StatusCode)
		return
	}
	var apiResponse ErrorResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal body %v\n", err)
	}

	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Fatalf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

func TestRemoveMemberSuccess(t *testing.T) {
	url := "http://localhost:8080/teams/8/members/6"
	method := "DELETE"

	expectedResponse := GetSuccessResponse("Member removed")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
			t.Fatalf("Failed to close response body: %v", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body %v\n", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected Status OK, got %v\n", res.StatusCode)
		return
	}

	var apiResponse SuccessResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal body %v\n", err)
	}
	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Errorf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

func TestRemoveMemberFailure(t *testing.T) {
	url := "http://localhost:8080/teams/0/members/6"
	method := "DELETE"

	expectedResponse := GetErrorResponse("Could not remove member")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
		t.Fatalf("Failed to read response body %v", err)
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status OK, got %v\n", res.StatusCode)
		return
	}

	var apiResponse ErrorResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal body %v\n", err)
	}
	if !reflect.DeepEqual(apiResponse, expectedResponse) {
		t.Fatalf("Expected %v, got %v", expectedResponse, apiResponse)
	}
}

func TestMakeAdminSuccess(t *testing.T) {
	url := "http://localhost:8080/make_admin/6"
	method := "POST"
	expectedResponse := GetSuccessResponse("User is now an admin")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
			t.Fatalf("Failed to close response body")
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 status OK, got %v", res.StatusCode)
		return
	}

	var apiResponse SuccessResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal body %v\n", err)
	}

	if !reflect.DeepEqual(expectedResponse, apiResponse) {
		t.Errorf("Expected %v, got %v\n", expectedResponse, apiResponse)
	}
}

func TestMakeAdminFailure(t *testing.T) {
	url := "http://localhost:8080/make_admin/0"
	method := "POST"
	expectedResponse := GetErrorResponse("Could not make admin")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
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
			t.Fatalf("Failed to close response body")
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected 500 StatusInternalServerError, got %v", res.StatusCode)
		return
	}

	var apiResponse ErrorResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal body %v\n", err)
	}

	if !reflect.DeepEqual(expectedResponse, apiResponse) {
		t.Errorf("Expected %v, got %v\n", expectedResponse, apiResponse)
	}
}

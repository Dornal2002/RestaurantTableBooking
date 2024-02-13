package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/internal/app/pkg/dto"
	"testing"
)

type MockAdminService struct{}

func (m *MockAdminService) AdminSignup(ctx context.Context, req dto.AdminSignUpRequest) error {
	return nil
}

func (m *MockAdminService) AdminLogin(ctx context.Context, req dto.AdminLoginRequest) error {
	return nil
}

func (m *MockAdminService) GetAdmin(ctx context.Context) ([]dto.AdminResponse, error) {
	return []dto.AdminResponse{{AdminID: 1, Name: "test"}}, nil
}

func TestSignUpHandler(t *testing.T) {
	adminSvc := &MockAdminService{}
	handler := SignUpHandler(adminSvc)

	reqBody := dto.AdminSignUpRequest{Name: "test", Password: "password"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	handler(res, req)

	if res.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d, got %d", http.StatusAccepted, res.Code)
	}
}

func TestLoginHandler(t *testing.T) {
	adminSvc := &MockAdminService{}
	handler := LoginHandler(adminSvc)

	reqBody := dto.AdminLoginRequest{Email: "test", Password: "password"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	handler(res, req)

	if res.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d, got %d", http.StatusAccepted, res.Code)
	}
}

func TestGetUsersHandler(t *testing.T) {
	adminSvc := &MockAdminService{}
	handler := GetUsersHandler(adminSvc)

	req := httptest.NewRequest("GET", "/users", nil)
	res := httptest.NewRecorder()

	handler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	var response []dto.AdminResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	if len(response) != 1 {
		t.Errorf("Expected 1 admin user, got %d", len(response))
	}
}

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

type MockBookingService struct{}

func (m *MockBookingService) CreateUserBooking(details dto.BookingDetails) (dto.BookingDetails, error) {
	if details.CustomerName == "test" {
		return dto.BookingDetails{}, nil
	}
	return dto.BookingDetails{BookingID: 1, CustomerName: "test"}, nil
}

func (m *MockBookingService) GetSlots(ctx context.Context) ([]dto.SlotResponse, error) {
	return []dto.SlotResponse{{SlotId: 1, StartTime: "09:00", EndTime: "10:00"}}, nil
}

func TestCreateBookingHandler(t *testing.T) {
	bookSvc := &MockBookingService{}
	handler := CreateBooking(bookSvc)

	reqBody := dto.BookingDetails{CustomerName: "test", SlotId: 1, TableId: 1}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/create-booking", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	handler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	if res.Body.String() != "Booking Done Successfully" {
		t.Errorf("Expected response 'Booking Done Successfully', got %s", res.Body.String())
	}
}

func TestGetSlotsHandler(t *testing.T) {
	bookSvc := &MockBookingService{}
	handler := GetSlots(bookSvc)

	req := httptest.NewRequest("GET", "/slots", nil)
	res := httptest.NewRecorder()

	handler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	var response []dto.SlotResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	if len(response) != 1 {
		t.Errorf("Expected 1 slot, got %d", len(response))
	}
}

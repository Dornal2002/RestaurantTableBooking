package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project/internal/app/booking/mocks"
	"project/internal/app/pkg/dto"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCreateBookingHandler(t *testing.T) {
	bookSvc := mocks.NewService(t)
	handler := CreateBooking(bookSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Success for booking details",
			input: `{
						"customer_name":"booking",
						"contact_no":"9328783491",
						"date":"1-2-2003",
						"slot_id":1,
						"table_id":2

					}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateUserBooking", mock.Anything).Return(dto.BookingDetails{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Fail for booking details",
			input: `{
				"customer_name":"",
				"contact_no":"",
				"date":"",
				"slot_id":0,
				"table_id":0
					}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for booking details",
			input: `{
				"customer_name":"",
				"date":"",
				"slot_id":0,
				"table_id":0
						
					}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for booking details",
			input: `{
				"contact_no":"",
				"date":"",
				"slot_id":0,
				"table_id":0
						
					}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for booking details",
			input: `{
				"name":"",
				"contact_no":"",
				"slot_id":0,
				"table_id":0
						
					}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for booking details",
			input: `{
				"name":"",
				"contact_no":"",
				"date":""
						
					}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(bookSvc)

			req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

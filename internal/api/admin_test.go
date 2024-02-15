package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project/internal/app/admin/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestLoginHandler(t *testing.T) {
	adminSvc := mocks.NewAdminService(t)
	userLoginHandler := LoginHandler(adminSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mock *mocks.AdminService)
		expectedStatusCode int
	}{
		{
			name: "Success for user login",
			input: `{
						"email": "admin@gmail.com",
						"password": "Abc@123456"   
					}`,
			setup: func(mockSvc *mocks.AdminService) {
				mockSvc.On("AdminLogin", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Success for user login",
			input: `{
						"email": "admin@gmail.com",
						"password": "Abc@123456"   
					}`,
			setup: func(mockSvc *mocks.AdminService) {
				mockSvc.On("AdminLogin", mock.Anything, mock.Anything).Return(errors.New("Error occured")).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "Fail for user login",
			input: `{
						"email": "admin"						
					}`,
			setup: func(mockSvc *mocks.AdminService) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for user login",
			input: `{
						"password": "12345"   
					}`,
			setup: func(mockSvc *mocks.AdminService) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for user login",
			input: `{
						"email": "",
						"password": ""   
					}`,
			setup: func(mockSvc *mocks.AdminService) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Fail for user login",
			input: `{
						"email": "admin,
						"password": "Abc"   
					}`,
			setup: func(mockSvc *mocks.AdminService) {
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(adminSvc)

			req, err := http.NewRequest("POST", "/admin/login", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(userLoginHandler)
			handler.ServeHTTP(rr, req)

			fmt.Println("Error")

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

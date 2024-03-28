package admin

// import (
// 	"context"
// 	"project/internal/app/pkg/dto"
// 	"testing"
// )

// type MockAdminRepo struct{}

// // Mock implementation of AdminSignup function
// func (m *MockAdminRepo) AdminSignup(ctx context.Context, user dto.AdminSignUpRequest) error {
// 	return nil
// }

// // Mock implementation of AdminLogin function
// func (m *MockAdminRepo) AdminLogin(ctx context.Context, user dto.AdminLoginRequest) error {
// 	return nil
// }

// // Mock implementation of GetAdmin function
// func (m *MockAdminRepo) GetAdmin(ctx context.Context) ([]dto.AdminResponse, error) {
// 	return []dto.AdminResponse{{AdminID: 1, Name: "test"}}, nil
// }

// func TestAdminSignup_Success(t *testing.T) {
// 	// Arrange
// 	adminRepo := &MockAdminRepo{}
// 	adminSvc := NewService(adminRepo)

// 	// Act
// 	err := adminSvc.AdminSignup(context.Background(), dto.AdminSignUpRequest{Name: "test", Password: "password"})

// 	// Assert
// 	if err != nil {
// 		t.Errorf("AdminSignup returned an error: %v", err)
// 	}
// }

// func TestAdminLogin_Success(t *testing.T) {
// 	// Arrange
// 	adminRepo := &MockAdminRepo{}
// 	adminSvc := NewService(adminRepo)

// 	// Act
// 	err := adminSvc.AdminLogin(context.Background(), dto.AdminLoginRequest{Email: "test", Password: "password"})

// 	// Assert
// 	if err != nil {
// 		t.Errorf("AdminLogin returned an error: %v", err)
// 	}
// }

// // func TestGetAdmin_Success(t *testing.T) {
// // 	// Arrange
// // 	adminRepo := &MockAdminRepo{}
// // 	adminSvc := NewService(adminRepo)

// // 	// Act
// // 	users, err := adminSvc.GetAdmin(context.Background())

// // 	// Assert
// // 	if err != nil {
// // 		t.Errorf("GetAdmin returned an error: %v", err)
// // 	}

// // 	if len(users) != 1 || users[0].AdminID != 1 || users[0].Name != "test" {
// // 		t.Errorf("GetAdmin returned unexpected result: %v", users)
// // 	}
// // }

// // Add more test cases as needed...

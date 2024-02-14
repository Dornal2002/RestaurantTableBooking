package adminBookings

// import (
// 	"context"
// 	"project/internal/app/pkg/dto"
// 	"testing"
// )

// type MockAdminBookingRepo struct{}

// func (m *MockAdminBookingRepo) AdminAssignTable(admin dto.AdminAssignTable) (dto.AdminAssignTable, error) {
// 	return dto.AdminAssignTable{}, nil
// }

// func (m *MockAdminBookingRepo) AdminCancelTable(admin dto.CancelTable, tableId int) (dto.CancelTable, error) {
// 	return dto.CancelTable{}, nil
// }

// func (m *MockAdminBookingRepo) AdminUpdateTable(admin dto.UpdateTable, tableId int) (dto.UpdateTable, error) {
// 	return dto.UpdateTable{}, nil
// }

// func (m *MockAdminBookingRepo) AdminGetDetails(ctx context.Context) ([]dto.GetTable, error) {
// 	return []dto.GetTable{{TableID: 1}}, nil
// }

// func TestAdminAssignTable(t *testing.T) {
// 	adminRepo := &MockAdminBookingRepo{}
// 	adminSvc := NewService(adminRepo)

// 	adminData := dto.AdminAssignTable{SlotId: 1, TableID: 1}
// 	_, err := adminSvc.AdminAssignTable(adminData)
// 	if err != nil {
// 		t.Errorf("AdminAssignTable returned an error: %v", err)
// 	}
// }

// func TestAdminCancelTable(t *testing.T) {
// 	adminRepo := &MockAdminBookingRepo{}
// 	adminSvc := NewService(adminRepo)

// 	adminData := dto.CancelTable{TableID: 1}
// 	_, err := adminSvc.AdminCancelTable(adminData, 1)
// 	if err != nil {
// 		t.Errorf("AdminCancelTable returned an error: %v", err)
// 	}
// }

// func TestAdminUpdateTable(t *testing.T) {
// 	adminRepo := &MockAdminBookingRepo{}
// 	adminSvc := NewService(adminRepo)

// 	adminData := dto.UpdateTable{BookingID: 1}
// 	_, err := adminSvc.AdminUpdateTable(adminData, 1)
// 	if err != nil {
// 		t.Errorf("AdminUpdateTable returned an error: %v", err)
// 	}
// }

// func TestAdminGetDetails(t *testing.T) {
// 	adminRepo := &MockAdminBookingRepo{}
// 	adminSvc := NewService(adminRepo)

// 	details, err := adminSvc.AdminGetDetails(context.Background())
// 	if err != nil {
// 		t.Errorf("AdminGetDetails returned an error: %v", err)
// 	}

// 	if len(details) != 1 || details[0].TableID != 1 {
// 		t.Errorf("AdminGetDetails returned unexpected result: %v", details)
// 	}
// }

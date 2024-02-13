package booking

// import (
// 	"context"
// 	"project/internal/app/pkg/dto"
// 	"testing"
// )

// type MockBookingRepo struct{}

// func (m *MockBookingRepo) InsertBookingDetails(user dto.BookingDetails) (dto.BookingDetails, error) {
// 	return dto.BookingDetails{}, nil
// }

// func (m *MockBookingRepo) GetSlotDetails(ctx context.Context, book dto.BookingDetails) ([]dto.SlotResponse, error) {
// 	return []dto.SlotResponse{{SlotId: 1, StartTime: "09:00", EndTime: "10:00"}}, nil
// }

// func TestCreateUserBooking_ValidData(t *testing.T) {
// 	bookingRepo := &MockBookingRepo{}
// 	bookingSvc := NewService(bookingRepo)

// 	userDetails := dto.BookingDetails{CustomerName: "John Doe", ContactNo: "1234567890"}
// 	_, err := bookingSvc.CreateUserBooking(userDetails)
// 	if err != nil {
// 		t.Errorf("CreateUserBooking returned an error for valid data: %v", err)
// 	}
// }

// func TestCreateUserBooking_InvalidData(t *testing.T) {
// 	bookingRepo := &MockBookingRepo{}
// 	bookingSvc := NewService(bookingRepo)

// 	userDetails := dto.BookingDetails{CustomerName: "", ContactNo: "123"}
// 	_, err := bookingSvc.CreateUserBooking(userDetails)
// 	if err == nil || err.Error() != "invalid user data" {
// 		t.Errorf("CreateUserBooking did not return expected error for invalid data: %v", err)
// 	}
// }

// func TestGetSlots(t *testing.T) {
// 	bookingRepo := &MockBookingRepo{}
// 	bookingSvc := NewService(bookingRepo)

// 	slots, err := bookingSvc.GetSlots(context.Background())
// 	if err != nil {
// 		t.Errorf("GetSlots returned an error: %v", err)
// 	}

// 	if len(slots) != 1 || slots[0].ID != 1 || slots[0].StartTime != "09:00" || slots[0].EndTime != "10:00" {
// 		t.Errorf("GetSlots returned unexpected result: %v", slots)
// 	}
// }

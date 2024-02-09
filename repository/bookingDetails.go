package repository

import (
	"project/internal/app/pkg/dto"
)

type BookingStorer interface {
	// CreateUserBooking([]dto.BookingDetails) (int, error)
	InsertBookingDetails(bookDeteails dto.BookingDetails) (BookingTableDetails, error)
}

type BookingTableDetails struct {
	BookingID    int64
	CustomerName string
	ContactNo    string
	NoOfPeople   int
	SelectDate   string
	StartTime    string
	EndTime      string
}

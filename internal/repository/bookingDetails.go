package repository

import (
	"context"
	"project/internal/app/pkg/dto"
)

type BookingStorer interface {
	InsertBookingDetails(bookDetails dto.BookingDetails) (BookingTableDetails, error)
	GetSlotDetails(ctx context.Context, bookDetails dto.BookingDetails) ([]dto.SlotResponse, error)
	// removeBookedSlot(slots []dto.SlotResponse, slotID, tableID int) []dto.SlotResponse
}

type BookingTableDetails struct {
	BookingID    int64
	CustomerName string
	ContactNo    string
	Date         string
	SlotId       int
	TableId      int
}

type SlotDetails struct {
	SlotId    int
	StartTime string
	EndTime   string
	// Availability bool
	TableId []int
}

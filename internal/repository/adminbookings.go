package repository

import (
	"context"
	"project/internal/app/pkg/dto"
)

type AdminBookingsStorer interface {
	AdminAssignTable(admin dto.AdminAssignTable) (dto.AdminAssignTable, error)
	AdminCancelTable(admin dto.CancelTable, bookingId int64) (dto.CancelTable, error)
	AdminUpdateTable(admin dto.UpdateTable, bookingId int64) (dto.UpdateTable, error)
	AdminGetDetails(ctx context.Context) ([]dto.BookingDetails, error)
}

type AssignTableDetails struct {
	BookingID    int64
	CustomerName string
	ContactNo    string
	Date         string
	SlotID       int
	TableID      int
}

type CancelTableDetails struct {
	BookingID int
}

type UpdateTableDetails struct {
	BookingID    int64
	CustomerName string
	ContactNo    string
	Date         string
	SlotId       int
	TableID      int
}

type GetTableDetails struct {
	BookingID    int64
	CustomerName string
	ContactNo    string
	Date1        string
	SlotId       int
	TableID      int
}

package repository

import (
	"context"
	"project/internal/app/pkg/dto"
)

type AdminBookingsStorer interface {
	AdminAssignTable(admin dto.AdminAssignTable) (dto.AdminAssignTable, error)
	AdminCancelTable(admin dto.CancelTable, tableId int) (dto.CancelTable, error)
	AdminUpdateTable(admin dto.UpdateTable, tableID int) (dto.UpdateTable, error)
	AdminGetDetails(ctx context.Context) ([]dto.GetTable, error)
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
	TableID int
}

type UpdateTableDetails struct {
	BookingID int
	TableID   int
}

type GetTableDetails struct {
	BookingID    int64
	CustomerName string
	ContactNo    string
	Date         string
	SlotId       int
	TableID      int
}

package repository

import "project/internal/app/pkg/dto"

type TableStorer interface {
	InsertTableDetails(adminDetails dto.TableDetails) (TableDetails, error)
}

type TableDetails struct {
	TableID      int
	TableNo      int
	Availability bool
	BookingID    int
}

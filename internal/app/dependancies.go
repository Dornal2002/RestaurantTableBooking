package app

import (
	"database/sql"
	"project/internal/app/admin"
	"project/internal/app/booking"
	repository "project/repository/boltdb"
)

type Dependencies struct {
	BookingService booking.Service
	AdminService   admin.AdminService
}

func NewServices(db *sql.DB) Dependencies {

	//intialize repo dependencies
	bookingRepo := repository.NewBookingRepo(db)
	adminRepo := repository.NewAdminRepo(db)

	//intialize service dependencies
	bookingService := booking.NewService(bookingRepo)
	adminService := admin.NewService(adminRepo)
	return Dependencies{
		BookingService: bookingService,
		AdminService:   adminService,
	}
}

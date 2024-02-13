package app

import (
	"database/sql"
	"project/internal/app/admin"
	"project/internal/app/adminBookings"
	"project/internal/app/booking"
	repository "project/internal/repository/boltdb"
)

type Dependencies struct {
	BookingService      booking.Service
	AdminService        admin.AdminService
	AdminBookingService adminBookings.AdminBookingService
}

func NewServices(db *sql.DB) Dependencies {

	//intialize repo dependencies
	bookingRepo := repository.NewBookingRepo(db)
	adminRepo := repository.NewAdminRepo(db)
	adminBookingRepo := repository.NewAdminBookingRepo(db)

	//intialize service dependencies
	bookingService := booking.NewService(bookingRepo)
	adminService := admin.NewService(adminRepo)
	adminBookingService := adminBookings.AdminBookingService(adminBookingRepo)
	return Dependencies{
		BookingService:      bookingService,
		AdminService:        adminService,
		AdminBookingService: adminBookingService,
	}
}

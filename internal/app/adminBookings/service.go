package adminBookings

import (
	"context"
	"errors"
	"fmt"
	"log"
	"project/internal/app/pkg/dto"
	"project/internal/repository"
)

type AdminBookingService interface {
	AdminAssignTable(admin dto.AdminAssignTable) (dto.AdminAssignTable, error)
	AdminCancelTable(admin dto.CancelTable, bookingId int64) (dto.CancelTable, error)
	AdminUpdateTable(admin dto.UpdateTable, bookingId int64) (dto.UpdateTable, error)
	AdminGetDetails(ctx context.Context) ([]dto.BookingDetails, error)
}

type service struct {
	AdminBookingRepo repository.AdminBookingsStorer
}

func NewService(AdminBookingRepo repository.AdminBookingsStorer) AdminBookingService {
	return &service{
		AdminBookingRepo: AdminBookingRepo,
	}
}

func (at *service) AdminAssignTable(admin dto.AdminAssignTable) (dto.AdminAssignTable, error) {
	bkd := dto.AdminAssignTable{}
	fmt.Println(admin)
	_, err := at.AdminBookingRepo.AdminAssignTable(admin)
	if err != nil {
		fmt.Println("Error occured while assigning table", err.Error())
		return bkd, err
	}
	fmt.Println("Table Assigned Successfully")
	return bkd, nil
}

func (at *service) AdminCancelTable(admin dto.CancelTable, bookingId int64) (dto.CancelTable, error) {
	act := dto.CancelTable{}

	if admin.BookingID == 0 {
		return act, errors.New("booking id is invalid")
	}
	_, err := at.AdminBookingRepo.AdminCancelTable(admin, bookingId)
	// act = MapRepoObjectToDto1(actDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return act, err
	}
	fmt.Println("Table Cancelled Successfully")

	return act, nil

}

func (aut *service) AdminUpdateTable(admin dto.UpdateTable, bookingId int64) (dto.UpdateTable, error) {
	log.Print("in service")
	act := dto.UpdateTable{}

	if admin.BookingID <= 0 {
		return act, errors.New("booking id is invalid")
	}
	log.Print("sevice update info: ")
	log.Println(admin)
	_, err := aut.AdminBookingRepo.AdminUpdateTable(admin, bookingId)
	// act = MapRepoObjectToDto2(actDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return act, err
	}

	

	return act, nil

}

func (gt *service) AdminGetDetails(ctx context.Context) ([]dto.BookingDetails, error) {
	// details := dto.GetTable{}
	userList, err := gt.AdminBookingRepo.AdminGetDetails(ctx)
	if err != nil {
		log.Println("error in admin get details service")
		return userList, err
	}
	return userList, nil
}

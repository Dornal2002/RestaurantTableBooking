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
	AdminCancelTable(admin dto.CancelTable, tableId int) (dto.CancelTable, error)
	AdminUpdateTable(admin dto.UpdateTable, tableId int) (dto.UpdateTable, error)
	AdminGetDetails(ctx context.Context) ([]dto.GetTable, error)
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
	if admin.CustomerName == "" || admin.ContactNo == "" {
		return bkd, errors.New("invalid user data")
	}
	fmt.Println(admin)
	_, err := at.AdminBookingRepo.AdminAssignTable(admin)
	if err != nil {
		fmt.Println("Error occured while assigning table", err.Error())
		return bkd, err
	}
	fmt.Println("Table Assigned Successfully")
	return bkd, nil
}

func (at *service) AdminCancelTable(admin dto.CancelTable, tableId int) (dto.CancelTable, error) {
	act := dto.CancelTable{}

	if admin.TableID <= 0 {
		return act, errors.New("table id is invalid")
	}
	_, err := at.AdminBookingRepo.AdminCancelTable(admin, tableId)
	// act = MapRepoObjectToDto1(actDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return act, err
	}
	fmt.Println("Table Cancelled Successfully")

	return act, nil

}

func (aut *service) AdminUpdateTable(admin dto.UpdateTable, tableId int) (dto.UpdateTable, error) {
	act := dto.UpdateTable{}

	if admin.BookingID <= 0 {
		return act, errors.New("booking id is invalid")
	}
	_, err := aut.AdminBookingRepo.AdminUpdateTable(admin, tableId)
	// act = MapRepoObjectToDto2(actDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return act, err
	}

	return act, nil

}

func (gt *service) AdminGetDetails(ctx context.Context) ([]dto.GetTable, error) {
	// details := dto.GetTable{}
	userList, err := gt.AdminBookingRepo.AdminGetDetails(ctx)
	if err != nil {
		log.Println("error in admin get details service")
		return userList, err
	}
	return userList, nil
}

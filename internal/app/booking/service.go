package booking

import (
	"context"
	"errors"
	"fmt"
	"log"
	"project/internal/app/pkg/dto"
	"project/internal/repository"
)

type service struct {
	BookingRepo repository.BookingStorer
}
type Service interface {
	CreateUserBooking(user dto.BookingDetails) (dto.BookingDetails, error)
	GetSlots(ctx context.Context) ([]dto.SlotResponse, error)
}

func NewService(bookingRepo repository.BookingStorer) Service {
	return &service{
		BookingRepo: bookingRepo,
	}
}
func (bs *service) CreateUserBooking(user dto.BookingDetails) (dto.BookingDetails, error) {
	bkd := dto.BookingDetails{}
	log.Println(user)
	if user.CustomerName == "" || len(user.ContactNo) != 10 {
		return bkd, errors.New("invalid user data")
	}
	bkdDB, err := bs.BookingRepo.InsertBookingDetails(user)
	bkd = MapRepoObjectToDto(bkdDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return bkd, err
	}

	return bkd, nil
}

// return repository.InitializeDatabse()

func (bs *service) GetSlots(ctx context.Context) ([]dto.SlotResponse, error) {

	//here we are getting a list of slots with associcated tables
	book := dto.BookingDetails{}
	fmt.Println(book)

	slot, err := bs.BookingRepo.GetSlotDetails(ctx, book)
	fmt.Println(slot)
	if err != nil {
		log.Println("error in getslots service")
		return slot, err
	}
	return slot, nil

	//TODO: make a call to booking table SELECT * FROM bookings WHERE restaurantID=?, Date=?
	//Here you will get all the booking for a particular date and a particular restauratn
	//Now remove these values, from the above list of slots

}

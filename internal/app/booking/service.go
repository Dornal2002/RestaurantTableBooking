package booking

import (
	"errors"
	"fmt"
	"project/internal/app/pkg/dto"
	"project/repository"
)

type service struct {
	BookingRepo repository.BookingStorer
}
type Service interface {
	CreateUserBooking(user dto.BookingDetails) (dto.BookingDetails, error)
}

func NewService(bookingRepo repository.BookingStorer) Service {
	return &service{
		BookingRepo: bookingRepo,
	}
}
func (bs *service) CreateUserBooking(user dto.BookingDetails) (dto.BookingDetails, error) {
	bkd := dto.BookingDetails{}
	if user.CustomerName == "" || user.ContactNo == "" || user.NoOfPeople <= 0 {
		return bkd, errors.New("invalid user data")
	}
	fmt.Println(user)
	bkdDB, err := bs.BookingRepo.InsertBookingDetails(user)
	bkd = MapRepoObjectToDto(bkdDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return bkd, err
	}

	return bkd, nil
	// return repository.InitializeDatabse()
}

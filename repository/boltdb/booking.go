package repository

import (
	"database/sql"
	"fmt"
	"project/internal/app/pkg/dto"
	"project/repository"
)

type bookingStore struct {
	BaseRepository
}

type Service interface {
	InsertBookingDetails(bookDetails dto.BookingDetails) (repository.BookingTableDetails, error)
}

func NewBookingRepo(db1 *sql.DB) repository.BookingStorer {
	return &bookingStore{
		BaseRepository: BaseRepository{db1},
	}
}

func (bs *bookingStore) InsertBookingDetails(bookDetails dto.BookingDetails) (repository.BookingTableDetails, error) {
	b1 := repository.BookingTableDetails{}
	fmt.Println(bookDetails)
	query := `INSERT INTO table_bookings (customer_name, contact_no, no_of_people, select_date, start_time, end_time) VALUES (?, ?, ?, ?, ?, ?)`

	statement, err := bs.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating table", err.Error())
		return b1, err
	}
	res, err := statement.Exec(bookDetails.CustomerName, bookDetails.ContactNo, bookDetails.NoOfPeople, bookDetails.SelectDate, bookDetails.StartTime, bookDetails.EndTime)
	if err != nil {
		fmt.Println("Error occured in inserting data")
		return b1, err
	}
	b1.BookingID, err = res.LastInsertId()
	b1.CustomerName = bookDetails.CustomerName
	b1.NoOfPeople = bookDetails.NoOfPeople
	b1.SelectDate = bookDetails.SelectDate
	b1.StartTime = bookDetails.StartTime
	b1.EndTime = bookDetails.EndTime
	b1.ContactNo = bookDetails.ContactNo
	if err != nil {
		fmt.Println("error occured in fetching data")
		return b1, err
	}

	return b1, nil
}

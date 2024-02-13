package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"project/internal/app/pkg/dto"
	"project/internal/repository"
)

var db *sql.DB

type AdminBookingStore struct {
	BaseRepository
}

func NewAdminBookingRepo(db1 *sql.DB) repository.AdminBookingsStorer {
	return &AdminBookingStore{
		BaseRepository: BaseRepository{db1},
	}
}

func (bs *AdminBookingStore) AdminAssignTable(abd dto.AdminAssignTable) (dto.AdminAssignTable, error) {
	b1 := dto.AdminAssignTable{}
	fmt.Println(abd)
	query := `INSERT INTO table_bookings (customer_name, contact_no, date, slot_id,table_id) VALUES (?, ?, ?, ?, ?)`

	statement, err := bs.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating table", err.Error())
		return b1, err
	}
	_, err = statement.Exec(abd.CustomerName, abd.ContactNo, abd.Date, abd.SlotId, abd.TableID)
	if err != nil {
		fmt.Println("Error occured in inserting data")
		return b1, err
	}
	// booking_id, err := res.LastInsertId()
	// b1.BookingID = int(booking_id)
	// b1.CustomerName = abd.CustomerName
	// b1.NoOfPeople = abd.NoOfPeople
	// b1.SelectDate = abd.SelectDate
	// b1.StartTime = abd.StartTime
	// b1.EndTime = abd.EndTime
	// b1.ContactNo = abd.ContactNo
	// if err != nil {
	// 	fmt.Println("error occured in fetching data")
	// 	return b1, err
	// }

	return b1, nil
}

func (bs *AdminBookingStore) AdminCancelTable(admin dto.CancelTable, tableID int) (dto.CancelTable, error) {
	act := dto.CancelTable{}
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tables WHERE table_id=?", tableID).Scan(&count)
	if err != nil {
		fmt.Println("Error occured while fetching the data in repository")
		return act, err
	}
	if count == 0 {
		fmt.Println("Table not found", http.StatusNotFound)
		return act, nil
	}

	// Cancel table booking
	_, err = db.Exec("UPDATE tableDetails SET availability=?, booking_id=? WHERE table_id=?", true, nil, tableID)
	if err != nil {
		fmt.Println("Error occured while updating the data in repository")
		return act, err
	}

	_, err = db.Exec("DELETE from table_bookings WHERE table_id ?", 0)

	if err != nil {
		fmt.Println("Error occured while deleting the data in repository for table bookings")
		return act, err

	}
	fmt.Printf("Table %v booking cancelled", tableID)
	return act, nil
}

func (bs *AdminBookingStore) AdminUpdateTable(admin dto.UpdateTable, tableID int) (dto.UpdateTable, error) {
	aut := dto.UpdateTable{}
	booking := repository.BookingTableDetails{}
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tables WHERE table_id=?", tableID).Scan(&count)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error occured in retriewing data", err)
		return aut, err
	}
	if count == 0 {
		fmt.Println("error occured in when to accse data", err)
		return aut, err
	}

	// Update table booking details
	_, err = db.Exec("UPDATE table_bookings SET booking_id=?, customer_name=? ,contact_no=?,date=?,start_time=?,end_time=?,WHERE table_id=?", booking.BookingID, booking.CustomerName, booking.ContactNo, booking.Date, booking.SlotId, booking.TableId)
	// http.Error(err.Error(), http.StatusInternalServerError)
	return aut, err

	// w.WriteHeader(http.StatusOK)
	// fmt.Printf("Table %v booking updated", tableID)
	// return aut, nil
}

func (bs *AdminBookingStore) AdminGetDetails(ctx context.Context) ([]dto.GetTable, error) {
	usersList := make([]dto.GetTable, 0)
	query := "SELECT * FROM table_bookings"
	rows, err := bs.BaseRepository.DB.Query(query)
	// fmt.Println(rows)
	if err != nil {
		return usersList, err
	}
	defer rows.Close()
	for rows.Next() {
		user := dto.GetTable{}
		err = rows.Scan(&user.BookingID, &user.CustomerName, &user.ContactNo, &user.Date, &user.SlotId, &user.TableID)
		if err != nil {
			log.Println(err)
		}
		// log.Println(user)
		usersList = append(usersList, user)
	}

	return usersList, nil

}

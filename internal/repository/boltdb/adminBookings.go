package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"project/internal/app/pkg/dto"
	"project/internal/repository"
)

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

func (bs *AdminBookingStore) AdminCancelTable(admin dto.CancelTable, bookingId int64) (dto.CancelTable, error) {
	act := dto.CancelTable{}
	_, err := bs.BaseRepository.DB.Exec("DELETE FROM table_bookings where booking_id=?", bookingId)
	if err != nil {
		fmt.Println("Error occured while deleting the data in repository", err)
		return act, err
	}

	// Cancel table booking
	// _, err = db.Exec("UPDATE tableDetails SET availability=?, booking_id=? WHERE table_id=?", true, nil, tableID)
	// if err != nil {
	// 	fmt.Println("Error occured while updating the data in repository")
	// 	return act, err
	// }

	// _, err = db.Exec("DELETE from table_bookings WHERE table_id ?", 0)

	// if err != nil {
	// 	fmt.Println("Error occured while deleting the data in repository for table bookings")
	// 	return act, err

	// }
	fmt.Println("table is cancelled")
	return act, nil
}

func (bs *AdminBookingStore) AdminUpdateTable(admin dto.UpdateTable, bookingId int64) (dto.UpdateTable, error) {
	aut := dto.UpdateTable{}
	// booking := repository.BookingTableDetails{}
	var count int
	log.Println("update info: ")
	log.Println(admin)
	err := bs.BaseRepository.DB.QueryRow("SELECT COUNT(*) FROM table_bookings WHERE booking_id=?", bookingId).Scan(&count)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error occured in retriewing data", err)
		return aut, err
	}
	if count == 0 {
		fmt.Println("Booking_id is not available")
		return aut, err
	}

	// Update table booking details
	_, err = bs.BaseRepository.DB.Exec("UPDATE table_bookings SET  customer_name=? ,contact_no=?,date=?,slot_id=?,table_id=? WHERE booking_id=?", admin.CustomerName, admin.ContactNo, admin.Date, admin.SlotId, admin.TableID, bookingId)
	// http.Error(err.Error(), http.StatusInternalServerError)
	return aut, err

	// w.WriteHeader(http.StatusOK)
	// fmt.Printf("Table %v booking updated", tableID)
	// return aut, nil
}

func (bs *AdminBookingStore) AdminGetDetails(ctx context.Context) ([]dto.BookingDetails, error) {
	usersList := make([]dto.BookingDetails, 0)
	query := "SELECT * FROM table_bookings"
	rows, err := bs.BaseRepository.DB.Query(query)
	// fmt.Println(rows)
	if err != nil {
		return usersList, err
	}

	defer rows.Close()
	for rows.Next() {
		user := dto.BookingDetails{}
		err = rows.Scan(&user.BookingID, &user.Id, &user.Date, &user.SlotId, &user.TableId)
		if err != nil {
			log.Println(err)
		}
		usersList = append(usersList, user)
	}

	return usersList, nil

}

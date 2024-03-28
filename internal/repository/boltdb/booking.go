package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"project/internal/app/pkg/dto"
	"project/internal/repository"
	"strconv"
	"strings"
)

type bookingStore struct {
	BaseRepository
}

func NewBookingRepo(db1 *sql.DB) repository.BookingStorer {
	return &bookingStore{
		BaseRepository: BaseRepository{db1},
	}
}

func (bs *bookingStore) InsertBookingDetails(bookDetails dto.BookingDetails) (repository.BookingTableDetails, error) {
	b1 := repository.BookingTableDetails{}
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM table_bookings WHERE slot_id = %d AND table_id = %d`, bookDetails.SlotId, bookDetails.TableId)
	rows, err := bs.BaseRepository.DB.Query(query)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return b1, err
	}
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Printf("Error querying database: %v", err)
			return b1, err
		}
	}
	if count == 0 {
		query := `INSERT INTO table_bookings (customer_id,date,slot_id,table_id) VALUES (?, ?, ?, ?)`
		statement, err := bs.DB.Prepare(query)

		if err != nil {
			fmt.Println("Error while creating table", err.Error())
			return b1, err
		}
		res, err := statement.Exec(bookDetails.Id, bookDetails.Date, bookDetails.SlotId, bookDetails.TableId)
		if err != nil {
			fmt.Println("Error occured in inserting data")
			return b1, err
		}
		// fmt.Println(res)
		b1.BookingID, err = res.LastInsertId()
		// b1.CustomerName = bookDetails.CustomerName
		// b1.ContactNo = bookDetails.ContactNo
		b1.Id = int64(bookDetails.Id)
		b1.Date = bookDetails.Date
		b1.SlotId = bookDetails.SlotId
		b1.TableId = bookDetails.TableId

		if err != nil {
			fmt.Println("error occured in fetching data")
			return b1, err
		}

		fmt.Println("Table Booked Successfully")
	} else {
		fmt.Println("Slot is Not Available")
	}

	return b1, nil
}

func (bs *bookingStore) GetSlotDetails(ctx context.Context, booking dto.BookingDetails) ([]dto.SlotResponse, error) {
	slotList := make([]dto.SlotResponse, 0)
	query := "SELECT * FROM time_slots"
	rows, err := bs.BaseRepository.DB.Query(query)
	if err != nil {
		return slotList, err
	}
	defer rows.Close()
	for rows.Next() {
		slots := dto.SlotResponse{}
		var tableIds string
		err := rows.Scan(&slots.SlotId, &slots.StartTime, &slots.EndTime, &tableIds)
		if err != nil {
			return slotList, err
		}
		if tableIds != "" {
			tables := strings.Split(tableIds[1:len(tableIds)-1], ",")

			bookTableIds := []int{}
			query, err := bs.BaseRepository.DB.Query("select distinct table_id from table_bookings where slot_id=?", slots.SlotId)
			if err != nil {
				fmt.Println(err)
				return slotList, err
			}

			for query.Next() {
				id := 0
				err := query.Scan(&id)
				if err != nil {
					fmt.Println(err)
					return slotList, err
				}

				bookTableIds = append(bookTableIds, id)
			}

			allTableIds := []int{}

			for _, t := range tables {
				id, err := strconv.Atoi(t)
				if err != nil {
					return slotList, err
				}

				// slots.TableId = append(slots.TableId, id)

				allTableIds = append(allTableIds, id)
			}
			// fmt.Println("All tables", allTableIds)
			// fmt.Println("=> ", bookTableIds, allTableIds)

			result := removeCommon(allTableIds, bookTableIds)
			// fmt.Println("result ", result)
			slots.TableId = result
		}

		if len(slots.TableId) != 0 {
			slotList = append(slotList, slots)
		}
	}

	return slotList, nil
}

func removeCommon(arr1, arr2 []int) []int {
	result := []int{}

	// Create a map to store unique elements of array2
	uniqueElements := make(map[int]bool)
	for _, num := range arr2 {
		uniqueElements[num] = true
	}

	// Iterate through array1 and add elements that are not in array2 to the result
	for _, num := range arr1 {
		if !uniqueElements[num] {
			result = append(result, num)
		}
	}

	return result
}

package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabse() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "TableBooking.db")

	if err != nil {
		fmt.Println("Error occured while creating databse", err.Error())
		return nil, err
	}
	query := "CREATE TABLE IF NOT EXISTS table_bookings (booking_id INTEGER PRIMARY KEY AUTOINCREMENT,customer_name TEXT,contact_no TEXT ,no_of_people INTEGER,select_date DATE,start_time TIMESTAMP,end_time TIMESTAMP)"

	statement, err := db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating table_bookings table ", err.Error())
		return nil, err
	}

	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS admin_data (admin_id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT,contact_no TEXT ,email TEXT,password TEXT,access_token TEXT)"

	statement, err = db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating admin_data table", err.Error())
		return nil, err
	}
	statement.Exec()
	return db, nil

}

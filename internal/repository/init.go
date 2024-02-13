package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDatabse() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "TableBooking.db")

	if err != nil {
		fmt.Println("Error occured while creating database", err.Error())
		return nil, err
	}
	query := "CREATE TABLE IF NOT EXISTS table_bookings (booking_id INTEGER PRIMARY KEY AUTOINCREMENT,customer_name TEXT,contact_no TEXT ,date TEXT,slot_id INTEGER,table_id INTEGER,FOREIGN KEY (booking_id) REFERENCES time_slots(id),FOREIGN KEY (booking_id) REFERENCES tableDetails(table_id))"

	statement, err := db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating table_bookings table ", err.Error())
		return nil, err
	}

	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS tableDetails (table_id INTEGER PRIMARY KEY AUTOINCREMENT,table_no INTEGER NOT NULL UNIQUE, availability BOOLEAN,booking_id INTEGER)"

	statement, err = db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating Table table", err.Error())
		return nil, err
	}
	statement.Exec()
	seedTableData()

	query = "CREATE TABLE IF NOT EXISTS admin_data (admin_id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT,contact_no TEXT ,email TEXT,password TEXT)"

	statement, err = db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating admin_data table", err.Error())
		return nil, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS restaurant (id INTEGER PRIMARY KEY, name TEXT,address TEXT)"

	statement, err = db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating RESTAURANT table", err.Error())
		return nil, err
	}
	statement.Exec()
	seedRestoTable()

	query = "CREATE TABLE IF NOT EXISTS time_slots(id INTEGER,start_time TEXT ,end_time TEXT ,table_id INTEGER NOT NULL,FOREIGN KEY (id) REFERENCES tableDetails(table_id),PRIMARY KEY (id,table_id))"

	statement, err = db.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating the Time slots table", err.Error())
		return nil, err
	}
	statement.Exec()
	seedTimeSlots()

	return db, nil
}

func seedTableData() {
	query := "INSERT INTO tableDetails (table_no,availability,booking_id) VALUES(?,?,?)"
	statement, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting table data: " + err.Error())
		return
	}
	statement.Exec(1, true, 101)
	statement.Exec(2, true, 102)
	statement.Exec(3, true, 103)
	statement.Exec(4, true, 104)
	statement.Exec(5, true, 105)
	statement.Exec(6, true, 106)
	statement.Exec(7, false, 107)
	statement.Exec(8, true, 108)
	statement.Exec(9, false, 109)
	statement.Exec(10, true, 110)

}

func seedRestoTable() {
	query := "INSERT INTO restaurant(id,name,address) VALUES(?,?,?)"
	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println("Error occured in inserting restaurant data", err.Error())
		return
	}

	stmt.Exec(1, "resto", "solapur")

}

func seedTimeSlots() {
	query := "INSERT INTO time_slots(id,start_time,end_time,table_id) VALUES(?,?,?,?)"

	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Println("Error occured intinserting time_slots table", err.Error())
		return
	}

	stmt.Exec(1, "8:00:00", "10:00:00", "[1,2,3,4,5,6,7,8,9,10]")
	// stmt.Exec(1, "8:00:00", "10:00:00", 2)
	// stmt.Exec(1, "8:00:00", "10:00:00", 3)
	// stmt.Exec(1, "8:00:00", "10:00:00", 4)
	// stmt.Exec(1, "8:00:00", "10:00:00", 5)
	// stmt.Exec(1, "8:00:00", "10:00:00", 6)
	// stmt.Exec(1, "8:00:00", "10:00:00", 7)
	// stmt.Exec(1, "8:00:00", "10:00:00", 8)
	// stmt.Exec(1, "8:00:00", "10:00:00", 9)
	// stmt.Exec(1, "8:00:00", "10:00:00", 10)

	stmt.Exec(2, "10:15:00", "12:15:00", "[1,2,3,4,5,6,7,8,9,10]")
	// stmt.Exec(2, "10:15:00", "12:15:00", 2)
	// stmt.Exec(2, "10:15:00", "12:15:00", 3)
	// stmt.Exec(2, "10:15:00", "12:15:00", 4)
	// stmt.Exec(2, "10:15:00", "12:15:00", 5)
	// stmt.Exec(2, "10:15:00", "12:15:00", 6)
	// stmt.Exec(2, "10:15:00", "12:15:00", 7)
	// stmt.Exec(2, "10:15:00", "12:15:00", 8)
	// stmt.Exec(2, "10:15:00", "12:15:00", 9)
	// stmt.Exec(2, "10:15:00", "12:15:00", 10)

	stmt.Exec(3, "12:30:00", "02:30:00", "[1,2,3,4,5,6,7,8,9,10]")
	// stmt.Exec(3, "12:30:00", "02:30:00", 2)
	// stmt.Exec(3, "12:30:00", "02:30:00", 3)
	// stmt.Exec(3, "12:30:00", "02:30:00", 4)
	// stmt.Exec(3, "12:30:00", "02:30:00", 5)
	// stmt.Exec(3, "12:30:00", "02:30:00", 6)
	// stmt.Exec(3, "12:30:00", "02:30:00", 7)
	// stmt.Exec(3, "12:30:00", "02:30:00", 8)
	// stmt.Exec(3, "12:30:00", "02:30:00", 9)
	// stmt.Exec(3, "12:30:00", "02:30:00", 10)

	stmt.Exec(4, "3:00:00", "5:00:00", "[1,2,3,4,5,6,7,8,9,10]")
	// stmt.Exec(4, "3:00:00", "5:00:00", 2)
	// stmt.Exec(4, "3:00:00", "5:00:00", 3)
	// stmt.Exec(4, "3:00:00", "5:00:00", 4)
	// stmt.Exec(4, "3:00:00", "5:00:00", 5)
	// stmt.Exec(4, "3:00:00", "5:00:00", 6)
	// stmt.Exec(4, "3:00:00", "5:00:00", 7)
	// stmt.Exec(4, "3:00:00", "5:00:00", 8)
	// stmt.Exec(4, "3:00:00", "5:00:00", 9)
	// stmt.Exec(4, "3:00:00", "5:00:00", 10)

	stmt.Exec(5, "6:00:00", "8:00:00", "[1,2,3,4,5,6,7,8,9,10]")
	// stmt.Exec(5, "6:00:00", "8:00:00", 2)
	// stmt.Exec(5, "6:00:00", "8:00:00", 3)
	// stmt.Exec(5, "6:00:00", "8:00:00", 4)
	// stmt.Exec(5, "6:00:00", "8:00:00", 5)
	// stmt.Exec(5, "6:00:00", "8:00:00", 6)
	// stmt.Exec(5, "6:00:00", "8:00:00", 7)
	// stmt.Exec(5, "6:00:00", "8:00:00", 8)
	// stmt.Exec(5, "6:00:00", "8:00:00", 9)
	// stmt.Exec(5, "6:00:00", "8:00:00", 10)

}

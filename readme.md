# restaurant table booking system

## Problem statement

The project includes two services

1. The admin service - Can create account, signIn account, Assign table for customer, remove/cancel table, update table, list the table booked table details.

2. The user service - Where user can book table by entering the details.

## Problem description 

The project includes the user where user gets the details of the avilable slots and tables and the user can enter details like user name, contact number, date, slot number, table number . Once he/she enters the details theirs booking is confirmed.

The admin module includes creating the admin account. Once admin creates account he/she can login and the access of assigning table for customer, cancelling the table and updating the table. The admin also gets the list of details of how many tables are booked.

## Setup

This project uses the sqlite database to handle database queries
There are restaurant and table detaiils already seeded into the database.

Firstly, run the following command to download all dependencies
```bash
go mod download
```


1. Run following command to start table_booking Application

```bash
go mod download  // go run main.go
```


2. Run following command to run unit test cases
```bash
make test  // go test ./...
```

3. Run following command to check test coverage
```bash
make test-cover

#you can also check code test coverage on top. Click on codeccov badge to check more about test coverage
```

## APIs

1.Get slots :`GET http://localhost:8080/getslots`
2.Creating bookings : `POST http://localhost:8080/bookings`
3.Admin SignUp :`POST http://localhost:8080/admin/create`
4.Admin Login:`POST http://localhost:8080/admin/login`
5.Get Admin list :`GET http://localhost:8080/getuser`
6.Admin Assign Table: `PUT http://localhost:8080/assign_table`
7.Admin Cancel Table :`DELETE http://localhost:8080/cancel_table`
8.Admin Update Table:`PUT http://localhost:8080/update_table/{booking_id}`
9.Get booking Details :`GET http://localhost:8080/get_details`


## Project Structure
.
├── cmd
│   ├── main.go
│   └── TableBooking.db
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── adminBookings.go
│   │   ├── admin.go
│   │   ├── admin_test.go
│   │   ├── booking.go
│   │   ├── booking_test.go
│   │   └── route.go
│   ├── app
│   │   ├── admin
│   │   │   ├── domain.go
│   │   │   ├── mocks
│   │   │   │   └── AdminService.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── adminBookings
│   │   │   ├── domain.go
│   │   │   ├── mocks
│   │   │   │   └── AdminBookingService.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── booking
│   │   │   ├── domain.go
│   │   │   ├── mocks
│   │   │   │   └── Service.go
│   │   │   ├── service.go
│   │   │   └── service_test.go
│   │   ├── dependancies.go
│   │   └── pkg
│   │       └── dto
│   │           ├── adminBookingDetails.go
│   │           ├── adminDetails.go
│   │           ├── bookingDetails.go
│   │           └── table.go
│   └── repository
│       ├── adminbookings.go
│       ├── adminDetails.go
│       ├── boltdb
│       │   ├── adminBookings.go
│       │   ├── admin.go
│       │   ├── base.go
│       │   └── booking.go
│       ├── bookingDetails.go
│       ├── init.go
│       ├── mocks
│       │   ├── AdminBookingsStorer.go
│       │   ├── AdminStorer.go
│       │   ├── BookingStorer.go
│       │   └── TableStorer.go
│       └── tableDetails.go
├── readme.md
└── TableBooking.db




package repository

import (
	"database/sql"
	"fmt"
	"project/internal/app/pkg/dto"
	"project/repository"
)

type adminStore struct {
	BaseRepository
}

type AdminService interface {
	InsertAdminDetails(adminDetails dto.AdminDetails) (repository.AdminPersonalDetails, error)
}

func NewAdminRepo(db1 *sql.DB) repository.AdminStorer {
	return &adminStore{
		BaseRepository: BaseRepository{db1},
	}
}

func (bs *adminStore) InsertAdminDetails(adminDetails dto.AdminDetails) (repository.AdminPersonalDetails, error) {
	b1 := repository.AdminPersonalDetails{}
	fmt.Println(adminDetails)
	query := `INSERT INTO admin_data (name, contact_no,email,password,access_token) VALUES (?, ?, ?, ?, ?)`

	statement, err := bs.DB.Prepare(query)

	if err != nil {
		fmt.Println("Error while creating table", err.Error())
		return b1, err
	}
	res, err := statement.Exec(adminDetails.Name, adminDetails.ContactNo, adminDetails.Email, adminDetails.Password, adminDetails.AccessToken)
	if err != nil {
		fmt.Println("Error occured in inserting data")
		return b1, err
	}
	b1.AdminID, err = res.LastInsertId()
	b1.Name = adminDetails.Name
	b1.ContactNo = adminDetails.ContactNo
	b1.Email = adminDetails.Email
	b1.Password = adminDetails.Password
	b1.AccessToken = adminDetails.AccessToken
	if err != nil {
		fmt.Println("error occured in fetching data")
		return b1, err
	}

	return b1, nil
}

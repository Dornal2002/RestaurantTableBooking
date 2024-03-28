package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"project/internal/app/pkg/dto"
	"project/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type adminStore struct {
	BaseRepository
}

func NewAdminRepo(db1 *sql.DB) repository.AdminStorer {
	return &adminStore{
		BaseRepository: BaseRepository{db1},
	}
}

func (as *adminStore) AdminSignup(ctx context.Context, user dto.AdminSignUpRequest) (int32, error) {

    query := "INSERT INTO admin_data (name, contact_no, email, password, role) VALUES (?, ?, ?, ?, ?)"
    statement, err := as.BaseRepository.DB.Prepare(query)
    if err != nil {
        errMsg := fmt.Errorf("failed to prepare insert statement: %v", err)
        return 0, errMsg // Return 0 as ID and the error
    }
    defer statement.Close()

    result, err := statement.Exec(user.Name, user.ContactNo, user.Email, user.Password, user.Role)
    if err != nil {
        errMsg := fmt.Errorf("failed to execute insert query: %v", err)
        return 0, errMsg // Return 0 as ID and the error
    }

    insertedID, err := result.LastInsertId()
    if err != nil {
        errMsg := fmt.Errorf("failed to get last inserted ID: %v", err)
        return 0, errMsg // Return 0 as ID and the error
    }

    return int32(insertedID), nil // Return the inserted ID and no error
}

func (as *adminStore) AdminLogin(ctx context.Context, user dto.AdminLoginRequest) (dto.LoginResponse, error){

	query := fmt.Sprintf("SELECT password,admin_id,role FROM admin_data WHERE email = \"%s\"", user.Email)
	rows, err := as.BaseRepository.DB.Query(query)
	var adminId int32
	if err != nil {
		fmt.Println("Email is incorrect: " + err.Error())
		return dto.LoginResponse{},err
	}
	var password string
	var role string
	for rows.Next() {
		rows.Scan(&password,&adminId,&role)
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return dto.LoginResponse{}, err
	}
	var loginResp dto.LoginResponse
	loginResp.Id = int64(adminId)
	loginResp.Role = role
	return loginResp,nil

}

func (as *adminStore) GetAdmin(ctx context.Context) ([]dto.AdminResponse, error) {

	usersList := make([]dto.AdminResponse, 0)
	query := "SELECT admin_id,name,contact_no,email FROM admin_data"
	rows, err := as.BaseRepository.DB.Query(query)
	log.Println(rows)
	if err != nil {
		log.Printf("error in getting rows : %v", err)
		return usersList, err
	}

	for rows.Next() {
		user := dto.AdminResponse{}
		rows.Scan(&user.AdminID, &user.Name, &user.ContactNo, &user.Email)
		log.Println(user)
		usersList = append(usersList, user)
	}
	log.Println(usersList)
	return usersList, nil
}

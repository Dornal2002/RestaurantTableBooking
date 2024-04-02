package repository

import (
	"context"
	"database/sql"
	"errors"
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
	query := "SELECT admin_id,name,contact_no,email,role FROM admin_data"
	rows, err := as.BaseRepository.DB.Query(query)
	log.Println(rows)
	if err != nil {
		log.Printf("error in getting rows : %v", err)
		return usersList, err
	}

	for rows.Next() {
		user := dto.AdminResponse{}
		rows.Scan(&user.AdminID, &user.Name, &user.ContactNo, &user.Email,&user.Role)
		log.Println(user)
		usersList = append(usersList, user)
	}
	log.Println(usersList)
	return usersList, nil
}

func (as *adminStore) GetUserById(ctx context.Context, id int64) (dto.AdminResponse, error) {
	usersList := dto.AdminResponse{}

	// Prepare the SQL query with parameter binding
	query := "SELECT admin_id, name, contact_no, email, role FROM admin_data WHERE admin_id = ?"

	// Execute the query with the provided id
	rows, err := as.BaseRepository.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Printf("error in executing query: %v", err)
		return usersList, err
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into AdminResponse struct
	for rows.Next() {
		user := dto.AdminResponse{}
		err := rows.Scan(&user.AdminID, &user.Name, &user.ContactNo, &user.Email, &user.Role)
		if err != nil {
			log.Printf("error in scanning row: %v", err)
			return usersList, err
		}
		return usersList, nil
	}

	if err := rows.Err(); err != nil {
		log.Printf("error in iterating rows: %v", err)
		return usersList, err
	}

	return dto.AdminResponse{},errors.New("user Not found")
}

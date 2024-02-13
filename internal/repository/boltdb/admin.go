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

func (as *adminStore) AdminSignup(ctx context.Context, user dto.AdminSignUpRequest) error {

	query := "INSERT INTO admin_data (name,contact_no,email,password) VALUES(?,?,?,?)"
	statement, err := as.BaseRepository.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return err
	}
	_, err = statement.Exec(user.Name, user.ContactNo, user.Email, user.Password)
	if err != nil {
		fmt.Println("error occured in executing insert query: " + err.Error())
		return err
	}

	return nil
}

func (as *adminStore) AdminLogin(ctx context.Context, user dto.AdminLoginRequest) error {

	query := fmt.Sprintf("SELECT password FROM admin_data WHERE email = \"%s\"", user.Email)
	rows, err := as.BaseRepository.DB.Query(query)
	if err != nil {
		fmt.Println("Email is incorrect: " + err.Error())
		return err
	}
	var password string
	for rows.Next() {
		rows.Scan(&password)
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil

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

package dto

type AdminDetails struct {
	AdminID     int64  `json:"admin_id"`
	Name        string `json:"name"`
	ContactNo   string `json:"contact_no"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

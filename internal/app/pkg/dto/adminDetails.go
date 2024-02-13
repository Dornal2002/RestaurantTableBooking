package dto

type AdminSignUpRequest struct {
	AdminID     int64  `json:"admin_id"`
	Name        string `json:"name"`
	ContactNo   string `json:"contact_no"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminResponse struct {
	AdminID     int64  `json:"admin_id"`
	Name        string `json:"name"`
	ContactNo   string `json:"contact_no"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

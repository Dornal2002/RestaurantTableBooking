package dto

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type AdminSignUpRequest struct {
	AdminID   int64  `json:"admin_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminResponse struct {
	AdminID   int64  `json:"admin_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AdminLoginResp struct{
	Token  string `json:"token"`
}
func (ar *AdminLoginRequest) Validate() error {
	if len(ar.Email) <= 0 || len(ar.Password) <= 0 {
		return fmt.Errorf("please provide anything as input")
	}
	if !isValidEmail(ar.Email) {
		return fmt.Errorf("please provide a valid Email address")
	}
	if !isValidPassword(ar.Password) {
		return fmt.Errorf("please provide a valid password ")
	}
	return nil
}

func (req *AdminSignUpRequest) ValidateAdmin() error {
	if len(req.Name) <= 0 {
		return fmt.Errorf("name field cannot be empty")
	}
	if len(req.ContactNo) <= 0 {
		return fmt.Errorf("mobile field cannot be empty")
	}
	if len(req.ContactNo) != 10 {
		return fmt.Errorf("mobile field must be of 10 digits only")
	}
	// Check each character is a digit
	for _, char := range req.ContactNo {
		if !unicode.IsDigit(char) {
			return fmt.Errorf("mobile field must be contains digits only")
		}
	}
	if len(req.Email) <= 0 {
		return fmt.Errorf("email field cannot be empty")
	}
	if !isValidEmail(req.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(req.Password) <= 0 {
		return fmt.Errorf("password field cannot be empty")
	}
	if len(req.Password) < 3 || len(req.Password) > 20 {
		return fmt.Errorf("length of the password field must be between 3 and 20 characters")
	}
	if !isValidPassword(req.Password) {
		return fmt.Errorf("please provide a proper password credentials")
	}
	// if len(req.Role) <= 0 || (strings.EqualFold(req.Role, "Customer") && strings.EqualFold(req.Role, "Admin")) {
	// 	return fmt.Errorf("role field can't be empty")
	// }
	// if !isValidateRole(Role(req.Role)) {
	// 	return fmt.Errorf(" invalid role, accepted roles are customer and admin only")
	// }

	return nil
}

// func isValidateRole(role Role) bool {
// 	switch role {
// 	case Customer, Admin:
// 		return true
// 	default:
// 		return false
// 	}
// }

func isValidEmail(email string) bool {
	// Basic Checks
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(regex).MatchString(email) {
		return false
	}
	if strings.Contains(email, " ") {
		return false
	}
	if strings.Count(email, "@") != 1 {
		return false
	}
	return true
}

func isValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}

	upperCase := false
	lowerCase := false
	digit := false
	specialChar := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			upperCase = true
		} else if unicode.IsLower(char) {
			lowerCase = true
		} else if unicode.IsDigit(char) {
			digit = true
		} else if !unicode.IsSpace(char) {
			specialChar = true
		}
	}
	return upperCase && lowerCase && digit && specialChar
}

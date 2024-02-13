package admin

// import (
// 	"fmt"
// 	"project/internal/app/pkg/dto"
// 	"regexp"
// )

// func validateUser(user dto.AdminSignUpRequest) bool {
// 	if len(user.Name) < 2 {
// 		return false
// 	} else if len(user.ContactNo) < 10 {
// 		return false
// 	} else if !(isValidEmail(user.Email) && isValidPassword(user.Password)) {
// 		return false
// 	}
// 	return true
// }

// func isValidEmail(email string) bool {
// 	// Regular expression for validating email addresses
// 	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
// 	re := regexp.MustCompile(emailRegex)
// 	if !re.MatchString(email) {
// 		fmt.Println("invalid email")
// 	}
// 	return re.MatchString(email)
// }
// func isValidPassword(password string) bool {
// 	if len(password) < 10 {
// 		return false
// 	}
// 	passwordRegex := `^(.*[a-zA-Z].*[a-zA-Z].*[a-zA-Z].*[a-zA-Z].*[a-zA-Z])(.*[!@#$%^&*()-_+=?])(.*[0-9].*[0-9].*[0-9])`
// 	re := regexp.MustCompile(passwordRegex)
// 	if !re.MatchString(password) {
// 		fmt.Println("invalid password")
// 	}
// 	return re.MatchString(password)
// }

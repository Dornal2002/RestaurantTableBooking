package middleware

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
	"time"

	"project/internal"
	"project/internal/app/pkg/dto"

	"github.com/dgrijalva/jwt-go"
)

// Role constants
const (
	RoleCustomer    = "customer"
	RoleDeliveryBoy = "deliveryboy"
	RoleAdmin       = "admin"
)

// User struct to represent user information
type User struct {
	Username string
	Roles    []string
}

// JWT middleware to verify and extract user roles
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from Authorization header
		path := r.URL.Path
		if path == "/user/signup" || path == "/user/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized: Token missing", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(internal.JWTKEY), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user roles from JWT claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Create User object with extracted roles
		// user := User{
		// 	Username: claims["username"].(string),
		// 	Roles:    claims["role"],
		// }
		role := claims["role"].(string)

		// Check access based on user roles
		if role == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		if !hasAccess(role, path) {
			http.Error(w, "Forbidden: Insufficient privileges", http.StatusForbidden)
			return
		}

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

func hasAccess(role string, path string) bool {
	switch path {
	case "/users":
		return role == "admin"

	case "/user/{id}":
		return role == "deliveryboy" || role == "admin"

	default:
		return false
	}
}

// /*****
func RequireAuth(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			http.Error(w, "Unauthorized: Token missing", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(internal.JWTKEY), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user roles from JWT claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Create User object with extracted roles
		// user := User{
		// 	Username: claims["username"].(string),
		// 	Roles:    claims["role"],
		// }
		// userId:=claims["userid"].(float64)

		role := claims["role"].(string)

		if !slices.Contains(roles, role) {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		// r.Header.Set("user_id", Id.String())
		// r.Header.Set("role", Role)

		next.ServeHTTP(w, r)

	})

}
func GenerateJWT(adminId int32) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(time.Hour * 1)

	// Create claims
	claims := jwt.MapClaims{
		"userid": adminId,
		// "role":   user.Role,
		"exp":    expirationTime.Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(internal.JWTKEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

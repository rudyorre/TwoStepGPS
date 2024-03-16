package auth

import (
	"backend/db"
	"backend/models"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Secret []byte
	DB     *db.DB
}

func NewAuthService(secret []byte, db *db.DB) (*AuthService, error) {
	if len(secret) == 0 {
		return nil, errors.New("secret cannot be empty")
	}
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &AuthService{Secret: secret, DB: db}, nil
}

func (a *AuthService) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// Parse and validate the request body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the username and password
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password must not be empty", http.StatusBadRequest)
		return
	}
	if len(user.Password) < 5 {
		http.Error(w, "Password must be at least 5 characters long", http.StatusBadRequest)
		return
	}

	// Check if the username is already taken
	exists, err := a.DB.UserExists(user.Username)
	if err != nil {
		http.Error(w, "Error checking if user exists: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Username already taken", http.StatusBadRequest)
		return
	}

	// Hash the user's password
	err = a.hashPassword(&user)
	if err != nil {
		http.Error(w, "Error hashing password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the user in the database
	err = a.DB.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a token for the user and respond
	a.createTokenAndRespond(user.Username, w)
}

func (a *AuthService) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Parse and validate the request body
	w.Header().Set("Content-Type", "application/json")
	var responseUser models.User
	err := json.NewDecoder(r.Body).Decode(&responseUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch the user from the database
	dbUser, err := a.DB.GetUserByUsername(responseUser.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error fetching user"})
		return
	}

	// Verify the password
	err = a.verifyPassword(dbUser.PasswordHash, responseUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
		return
	}

	// Create a token for the user and respond
	a.createTokenAndRespond(responseUser.Username, w)
}

func (a *AuthService) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := a.verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"username": username})
}

func (a *AuthService) createTokenAndRespond(username string, w http.ResponseWriter) {
	// Create a token for the user
	tokenString, err := a.createToken(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error creating token"})
		return
	}

	// Return a response body with the username and token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Operation successful",
		"username": username,
		"token":    tokenString,
	})
}

// createToken generates a new JWT token for the given username.
// It uses the HS256 signing method and sets the "username" claim to the provided username.
// The token expires after 14 days.
// It returns the generated token string and any error encountered during the process.
func (a *AuthService) createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24 * 14).Unix(),
		})

	tokenString, err := token.SignedString(a.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// verifyToken verifies the authenticity of a JWT token and returns the username associated with it.
// It takes a token string as input and returns the username and an error (if any).
func (a *AuthService) verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.Secret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("username claim not found")
	}
	return username, nil
}

// hashPassword hashes the user's password using bcrypt and updates the user's PasswordHash field.
// It takes a pointer to a User struct as input and returns an error if any.
func (a *AuthService) hashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	return nil
}

// verifyPassword compares a hashed password with a plain-text password and returns an error if
// they don't match.
func (a *AuthService) verifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

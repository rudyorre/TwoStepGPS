package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/auth"
	"backend/db"
	"backend/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var secretKey = []byte("secret-key")

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "Welcome to the the protected area")
}

func main() {
	_ = godotenv.Load(".env.local")
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	authService, err := auth.NewAuthService([]byte("secret-key"), db)
	if err != nil {
		log.Fatal(err)
	}

	deviceService := &DeviceService{APIKey: os.Getenv("ONESTEPGPS_API_KEY")}

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := models.Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})
	router.HandleFunc("/display-names", deviceService.getDisplayNames)
	router.HandleFunc("/device-locations", deviceService.getDeviceLocations)

	router.HandleFunc("/login", authService.HandleLogin)
	router.HandleFunc("/protected", ProtectedHandler)
	router.HandleFunc("/profile", authService.HandleGetProfile)
	router.HandleFunc("/signup", authService.HandleSignUp)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}

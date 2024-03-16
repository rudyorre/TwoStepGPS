package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"backend/auth"
	"backend/db"
	"backend/handlers"
	"backend/models"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

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

	deviceService := &handlers.DeviceService{APIKey: os.Getenv("ONESTEPGPS_API_KEY")}

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := models.Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})
	router.HandleFunc("/display-names", deviceService.HandleGetDisplayNames)
	router.HandleFunc("/device-locations", deviceService.HandleGetDeviceLocations)

	router.HandleFunc("/login", authService.HandleLogin)
	router.HandleFunc("/signup", authService.HandleSignUp)
	router.HandleFunc("/profile", authService.AuthMiddleware(authService.HandleGetProfile))
	router.HandleFunc(
		"/update-username",
		authService.AuthMiddleware(authService.HandleUpdateUsername),
	)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}

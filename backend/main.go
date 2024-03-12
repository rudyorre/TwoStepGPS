package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load(".env.local")

	deviceService := &DeviceService{APIKey: os.Getenv("ONESTEPGPS_API_KEY")}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})
	mux.HandleFunc("/display-names", deviceService.getDisplayNames)
	mux.HandleFunc("/device-locations", deviceService.getDeviceLocations)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}

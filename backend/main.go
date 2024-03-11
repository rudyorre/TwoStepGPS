package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env.local")

	deviceService := &DeviceService{APIKey: os.Getenv("ONESTEPGPS_API_KEY")}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/display-names", deviceService.getDisplayNames)
	http.HandleFunc("/device-locations", deviceService.getDeviceLocations)
	http.ListenAndServe(":8080", nil)
}

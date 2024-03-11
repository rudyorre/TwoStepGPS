package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Device struct {
	DeviceID    string `json:"device_id"`
	DisplayName string `json:"display_name"`
}

type APIResponse struct {
	ResultList []Device `json:"result_list"`
}

type Response struct {
	Message string `json:"message"`
}

var ONESTEPGPS_API_KEY string

// getDisplayNames retrieves the display names of devices from a remote API and returns them as a
// JSON response.
func getDisplayNames(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(
		"https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" +
			os.Getenv("ONESTEPGPS_API_KEY"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse APIResponse
	json.Unmarshal(body, &apiResponse)

	var displayNames []string
	for _, device := range apiResponse.ResultList {
		displayNames = append(displayNames, device.DisplayName)
	}

	displayNamesJson, _ := json.Marshal(displayNames)
	w.Header().Set("Content-Type", "application/json")
	w.Write(displayNamesJson)
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/display-names", getDisplayNames)
	http.ListenAndServe(":8080", nil)
}

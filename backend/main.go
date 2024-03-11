package main

import (
	"encoding/json"
	"io"
	"net/http"
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

// getDisplayNames retrieves the display names of devices from a remote API and returns them as a
// JSON response.
func getDisplayNames(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=oUFJLAi7YeGOgFyAr8d4_TkX_2GpFJnJi64xUem_d2k")
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Hello, World!"}
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/display-names", getDisplayNames)
	http.ListenAndServe(":8080", nil)
}

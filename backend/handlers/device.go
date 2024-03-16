package handlers

import (
	"backend/models"
	"encoding/json"
	"io"
	"net/http"
)

type DeviceService struct {
	APIKey string
}

// getDisplayNames retrieves the display names of devices from a remote API and returns them as a
// JSON response.
func (d *DeviceService) HandleGetDisplayNames(w http.ResponseWriter, r *http.Request) {
	var apiResponse models.APIResponse
	err := d.fetchAndUnmarshal(
		"https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="+
			d.APIKey, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var displayNames []string
	for _, device := range apiResponse.ResultList {
		displayNames = append(displayNames, device.DisplayName)
	}

	displayNamesJson, _ := json.Marshal(displayNames)
	w.Header().Set("Content-Type", "application/json")
	w.Write(displayNamesJson)
}

// getDeviceLocations retrieves the latest device locations from the OneStepGPS API
// and writes the locations as JSON to the HTTP response.
func (d *DeviceService) HandleGetDeviceLocations(w http.ResponseWriter, r *http.Request) {
	var apiResponse models.APIResponse
	err := d.fetchAndUnmarshal(
		"https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="+
			d.APIKey, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var locations []models.Device
	for _, device := range apiResponse.ResultList {
		locations = append(locations, models.Device{
			DeviceID:    device.DeviceID,
			DisplayName: device.DisplayName,
			Latitude:    device.LatestDevicePoint.Latitude,
			Longitude:   device.LatestDevicePoint.Longitude,
			Altitude:    device.LatestDevicePoint.Altitude,
			Angle:       device.LatestDevicePoint.Angle,
		})
	}

	locationsJson, _ := json.Marshal(locations)
	w.Header().Set("Content-Type", "application/json")
	w.Write(locationsJson)
}

// fetchAndUnmarshal fetches data from the specified URL and unmarshals it into the provided value.
// It returns an error if there was a problem fetching the data or unmarshaling it.
func (d *DeviceService) fetchAndUnmarshal(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

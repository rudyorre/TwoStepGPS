package main

import (
	"encoding/json"
	"net/http"
)

// getDisplayNames retrieves the display names of devices from a remote API and returns them as a
// JSON response.
func (d *DeviceService) getDisplayNames(w http.ResponseWriter, r *http.Request) {
	var apiResponse APIResponse
	err := fetchAndUnmarshal(
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
func (d *DeviceService) getDeviceLocations(w http.ResponseWriter, r *http.Request) {
	var apiResponse APIResponse
	err := fetchAndUnmarshal(
		"https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="+
			d.APIKey, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var locations []DeviceLocation
	for _, device := range apiResponse.ResultList {
		locations = append(locations, DeviceLocation{
			DeviceID:    device.DeviceID,
			DisplayName: device.DisplayName,
			Latitude:    device.LatestDevicePoint.Latitude,
			Longitude:   device.LatestDevicePoint.Longitude,
		})
	}

	locationsJson, _ := json.Marshal(locations)
	w.Header().Set("Content-Type", "application/json")
	w.Write(locationsJson)
}

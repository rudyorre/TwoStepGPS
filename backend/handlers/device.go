package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type DeviceService struct {
	APIKey string
	DB     *db.DB
}

func NewDeviceService(APIKey string, db *db.DB) (*DeviceService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &DeviceService{APIKey: APIKey, DB: db}, nil
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
			IsHidden:    false,
		})
	}

	locationsJson, _ := json.Marshal(locations)
	w.Header().Set("Content-Type", "application/json")
	w.Write(locationsJson)
}

func (d *DeviceService) HandleHideDevice(w http.ResponseWriter, r *http.Request, username string) {
	// Parse and validate the request body
	var reqBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceID, ok := reqBody["device_id"].(string)
	if !ok {
		http.Error(w, "Invalid device_id", http.StatusBadRequest)
		return
	}

	hide, ok := reqBody["hide"].(bool)
	if !ok {
		http.Error(w, "Invalid hide", http.StatusBadRequest)
		return
	}

	err = d.DB.HideDevice(username, deviceID, hide)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *DeviceService) HandleGetHiddenDevices(w http.ResponseWriter, r *http.Request, username string) {
	hiddenDevices, err := d.DB.GetHiddenDevices(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hiddenDevicesJson, _ := json.Marshal(hiddenDevices)
	w.Header().Set("Content-Type", "application/json")
	w.Write(hiddenDevicesJson)
}

func (d *DeviceService) HandleChangeColor(w http.ResponseWriter, r *http.Request, username string) {
	// Parse and validate the request body
	var reqBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceID, ok := reqBody["device_id"].(string)
	if !ok {
		http.Error(w, "Invalid device_id", http.StatusBadRequest)
		return
	}

	color, ok := reqBody["color"].(string)
	if !ok {
		http.Error(w, "Invalid color", http.StatusBadRequest)
		return
	}

	err = d.DB.ChangeColor(username, deviceID, color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (d *DeviceService) HandleGetDeviceSettings(w http.ResponseWriter, r *http.Request, username string) {
	var apiResponse models.APIResponse
	err := d.fetchAndUnmarshal(
		"https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="+
			d.APIKey, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the device settings from the database
	deviceSettingsMap, err := d.DB.GetDeviceSettings(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var locations []models.Device
	for _, device := range apiResponse.ResultList {
		deviceSettings, ok := deviceSettingsMap[device.DeviceID]
		if !ok {
			continue
		}
		locations = append(locations, models.Device{
			DeviceID:    device.DeviceID,
			DisplayName: device.DisplayName,
			Latitude:    device.LatestDevicePoint.Latitude,
			Longitude:   device.LatestDevicePoint.Longitude,
			Altitude:    device.LatestDevicePoint.Altitude,
			Angle:       device.LatestDevicePoint.Angle,
			IsHidden:    deviceSettings.IsHidden,
			Color:       deviceSettings.Color,
			Nickname:    deviceSettings.Nickname,
		})
	}

	locationsJson, _ := json.Marshal(locations)
	w.Header().Set("Content-Type", "application/json")
	w.Write(locationsJson)
}

func (d *DeviceService) HandleChangeNickname(w http.ResponseWriter, r *http.Request, username string) {
	// Parse and validate the request body
	var reqBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deviceID, ok := reqBody["device_id"].(string)
	if !ok {
		http.Error(w, "Invalid device_id", http.StatusBadRequest)
		return
	}

	nickname, ok := reqBody["nickname"].(string)
	if !ok {
		http.Error(w, "Invalid nickname", http.StatusBadRequest)
		return
	}

	err = d.DB.ChangeNickname(username, deviceID, nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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

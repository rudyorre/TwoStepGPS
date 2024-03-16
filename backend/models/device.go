package models

type Device struct {
	DeviceID    string  `json:"device_id"`
	DisplayName string  `json:"display_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Altitude    float64 `json:"altitude"`
	Angle       float64 `json:"angle"`
}

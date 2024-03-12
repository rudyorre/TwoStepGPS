package main

type Device struct {
	DeviceID          string `json:"device_id"`
	DisplayName       string `json:"display_name"`
	ActiveState       string `json:"active_state"`
	LatestDevicePoint struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
		Altitude  float64 `json:"altitude"`
		Angle     float64 `json:"angle"`
	} `json:"latest_device_point"`
}

type DeviceLocation struct {
	DeviceID    string  `json:"device_id"`
	DisplayName string  `json:"display_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Altitude    float64 `json:"altitude"`
	Angle       float64 `json:"angle"`
}

type APIResponse struct {
	ResultList []Device `json:"result_list"`
}

type Response struct {
	Message string `json:"message"`
}

type DeviceService struct {
	APIKey string
}

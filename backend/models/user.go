package models

type User struct {
	ID             int              `json:"id"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	PasswordHash   string           `json:"-"`
	DeviceSettings []DeviceSettings `json:"device_settings"`
}

type DeviceSettings struct {
	ID       int    `json:"id"`
	DeviceID string `json:"device_id"`
	Username string `json:"username"`
	IsHidden bool   `json:"is_hidden"`
	Nickname string `json:"nickname"`
	Color    string `json:"color"`
}

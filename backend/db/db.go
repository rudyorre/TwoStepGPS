package db

// The DeviceSettings table seed script:
// CREATE TABLE DeviceSettings (
//     ID SERIAL PRIMARY KEY,
//     DeviceID VARCHAR(255),
//     Username VARCHAR(255),
//     IsHidden BOOLEAN DEFAULT false,
//     Nickname VARCHAR(255),
//     Color VARCHAR(255),
//     CONSTRAINT unique_username_deviceid UNIQUE (Username, DeviceID)
// );

import (
	"backend/models"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func NewDB() (*DB, error) {
	var connStr string
	if os.Getenv("ENV") == "development" {
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASS"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_NAME"),
		)
	} else {
		connStr = os.Getenv("DATABASE_URL")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

func (d *DB) CreateUser(user models.User) error {
	_, err := d.Conn.Exec(
		"INSERT INTO users (username, password_hash) VALUES ($1, $2);",
		user.Username,
		user.PasswordHash,
	)
	return err
}

func (d *DB) UserExists(username string) (bool, error) {
	var exists bool
	err := d.Conn.QueryRow(
		"SELECT exists (SELECT 1 FROM users WHERE username=$1)",
		username,
	).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}

func (d *DB) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := d.Conn.QueryRow(
		"SELECT * FROM users WHERE username=$1;",
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *DB) UpdateUsername(oldUsername string, newUsername string) error {
	_, err := d.Conn.Exec(
		"UPDATE users SET username=$1 WHERE username=$2;",
		newUsername,
		oldUsername,
	)
	return err
}

func (d *DB) AddDeviceSetting(device models.DeviceSettings) error {
	_, err := d.Conn.Exec(
		"INSERT INTO DeviceSettings (DeviceID, Username, IsHidden, Nickname, Color) VALUES ($1, $2, $3, $4, $5);",
		device.DeviceID,
		device.Username,
		false,
		device.Nickname,
		device.Color,
	)
	return err
}

// HideDevice updates the visibility status of a device for a specific user.
// It sets the IsHidden field in the DeviceSettings table to the provided value.
// Parameters:
//   - username: The username of the user.
//   - deviceID: The ID of the device.
//   - hide: The visibility status to set for the device (true for hidden, false for visible).
//
// Returns:
//   - error: An error if the update operation fails, nil otherwise.
func (d *DB) HideDevice(username string, deviceID string, hide bool) error {
	_, err := d.Conn.Exec(
		`INSERT INTO DeviceSettings (Username, DeviceID, IsHidden) 
		VALUES ($1, $2, $3)
		ON CONFLICT (Username, DeviceID) 
		DO UPDATE SET IsHidden = $3;`,
		username,
		deviceID,
		hide,
	)
	return err
}

// GetHiddenDevices retrieves the hidden devices for a given username from the database.
// It returns a slice of device IDs and an error, if any.
func (d *DB) GetHiddenDevices(username string) ([]string, error) {
	rows, err := d.Conn.Query(
		"SELECT DeviceID FROM DeviceSettings WHERE Username=$1 AND IsHidden=true;",
		username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hiddenDeviceIDs []string
	for rows.Next() {
		var deviceID string
		err := rows.Scan(&deviceID)
		if err != nil {
			return nil, err
		}
		hiddenDeviceIDs = append(hiddenDeviceIDs, deviceID)
	}

	return hiddenDeviceIDs, nil
}

func (d *DB) ChangeColor(username string, deviceID string, color string) error {
	_, err := d.Conn.Exec(
		`INSERT INTO DeviceSettings (Username, DeviceID, Color) 
		VALUES ($1, $2, $3)
		ON CONFLICT (Username, DeviceID) 
		DO UPDATE SET Color = $3;`,
		username,
		deviceID,
		color,
	)
	return err
}

func (d *DB) GetDeviceSettings(username string) (map[string]models.DeviceSettings, error) {
	// Fetch device settings for the user from the database
	rows, err := d.Conn.Query("SELECT DeviceID, IsHidden, Color, Nickname FROM DeviceSettings WHERE Username = $1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a map to store the device settings
	deviceSettingsMap := make(map[string]models.DeviceSettings)
	for rows.Next() {
		var deviceSettings models.DeviceSettings
		var isHidden sql.NullBool
		var color sql.NullString
		var nickname sql.NullString

		err := rows.Scan(&deviceSettings.DeviceID, &isHidden, &color, &nickname)
		if err != nil {
			return nil, err
		}

		// Check if the values are valid and, if not, set default values
		if isHidden.Valid {
			deviceSettings.IsHidden = isHidden.Bool
		} else {
			deviceSettings.IsHidden = false
		}
		if color.Valid {
			deviceSettings.Color = color.String
		} else {
			deviceSettings.Color = "#AA4A44"
		}

		if nickname.Valid {
			deviceSettings.Nickname = nickname.String
		} else {
			deviceSettings.Nickname = ""
		}
		deviceSettingsMap[deviceSettings.DeviceID] = deviceSettings
	}

	return deviceSettingsMap, nil
}

func (d *DB) ChangeNickname(username string, deviceID string, nickname string) error {
	_, err := d.Conn.Exec(
		`INSERT INTO DeviceSettings (Username, DeviceID, Nickname) 
		VALUES ($1, $2, $3)
		ON CONFLICT (Username, DeviceID) 
		DO UPDATE SET Nickname = $3;`,
		username,
		deviceID,
		nickname,
	)
	return err
}

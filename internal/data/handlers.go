package data

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PinModel struct {
	DB *sql.DB
}

type Pin struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          float64   `json:"price"`
	Category       string    `json:"category"`
	Image_url      string    `json:"image_url"`
	Average_rating float64   `json:"average_rating"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(storedHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	return err == nil
}

func (m *PinModel) RegisterUser(username, password string) (int, error) {
	passwordHash, err := HashPassword(password)
	if err != nil {
		return 0, err
	}

	var userID int
	err = m.DB.QueryRow("INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id", username, passwordHash).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (m *PinModel) AuthenticateUser(username, password string) (bool, error) {
	var storedHash string
	err := m.DB.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("invalid username or password")
		}
		return false, err
	}

	if !CheckPassword(storedHash, password) {
		return false, errors.New("invalid username or password")
	}

	return true, nil
}

package model

type UserData struct {
	UserID         int64  `json:"user_id"`
	PassportNumber string `json:"passport_number"`
	PassportSerie  string `json:"passport_serie"`
}

package model

type GetAllUsersResp struct {
	UserID         int64  `json:"user_id" db:"user_id"`
	PassportNumber string `json:"passport_number" db:"passport_number"`
	PassportSerie  string `json:"passport_serie" db:"passport_serie"`
}

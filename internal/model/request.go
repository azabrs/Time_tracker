package model

type GetAllUsersReq struct {
	PageSize int64  `json:"page_size"`
	Page     int64  `json:"page"`
	Sort     Order  `json:"sort"`
	Filter   Filter `json:"filter"`
}

type Order struct {
	Asc bool   `json:"asc"`
	By  string `json:"by"`
}

type Filter struct {
	UserID         []int64  `json:"user_id"`
	PassportNumber []string `json:"passport_number"`
	PassportSerie  []string `json:"passport_serie"`
}

type StartTaskDescrReq struct {
	Description string   `json:"description"`
	UserTask    UserTask `json:"user_task"`
}

package time_tracker

import (
	"Time-tracker/internal/model"
)

type Service interface {
	GetUsers(model.GetAllUsersReq) ([]model.GetAllUsersResp, error)
}

package time_tracker_db

import "Time-tracker/internal/model"

type Repository interface {
	GetUsers(model.GetAllUsersReq) ([]model.GetAllUsersResp, error)
}

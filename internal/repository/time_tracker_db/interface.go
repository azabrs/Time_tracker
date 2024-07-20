package time_tracker_db

import "Time-tracker/internal/model"

type Repository interface {
	GetUsers(req model.GetAllUsersReq) ([]model.GetAllUsersResp, error)
	DeleteUser(userId int64) error
	ChangeUserData(userData model.UserData) error
	IsUserExist(userId int64) error
	AddTaskDescription(descr string)
	FinishTask(req model.UserTask) error
	IsActive(req model.UserTask) (bool, error)
	StartTask(req model.UserTask) error
	IsStarted(userTask model.UserTask) (bool, error)
}

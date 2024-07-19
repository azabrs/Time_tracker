package model

import "time"

type UserTask struct {
	UserId     int64     `json:"user_id"`
	TaskId     int64     `json:"task_id"`
	StartTime  time.Time `json:"start_time"`
	FinishTime time.Time `json:"finish_time"`
}

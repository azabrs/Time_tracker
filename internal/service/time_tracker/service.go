package time_tracker

import "Time-tracker/internal/repository/time_tracker_db"

type service struct {
	rep time_tracker_db.Repository
}

func New(rep time_tracker_db.Repository) *service {
	return &service{rep: rep}
}

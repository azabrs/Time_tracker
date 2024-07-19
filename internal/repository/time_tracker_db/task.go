package time_tracker_db

import (
	"Time-tracker/internal/model"
	"github.com/doug-martin/goqu/v9"
	"time"
)

func (s repository) IsStarted(userTask model.UserTask) (bool, error) {
	query := goqu.From(userTaskTable)
	query = query.Where(goqu.Ex{"task_id": userTask.TaskId})
	query = query.Where(goqu.Ex{"user_id": userTask.UserId})
	sqlQuery, _, err := query.Select("*").ToSQL()
	var t []time.Time
	if err != nil {
		return false, err
	}
	err = s.db.Select(&t, sqlQuery)
	if err != nil {
		return false, err
	}
	if len(t) == 0 {
		return false, nil
	}

	return true, nil

}

func (s repository) StartTask(req model.UserTask) error {
	query := goqu.From(userTaskTable)
	sqlQuery, _, err := query.Insert().Rows(req).OnConflict(goqu.DoNothing()).ToSQL()
	if err != nil {
		return err
	}
	_, err = s.db.Exec(sqlQuery, query)
	if err != nil {
		return err
	}
	return nil
}

func (s repository) FinishTask(req model.UserTask) error {
	query := goqu.From(userTaskTable)
	sqlQuery, _, err := query.Insert().Rows(req).OnConflict(goqu.DoNothing()).ToSQL()
	if err != nil {
		return err
	}
	_, err = s.db.Exec(sqlQuery, query)
	if err != nil {
		return err
	}
	return nil
}

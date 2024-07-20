package time_tracker_db

import (
	"Time-tracker/internal/model"
	"github.com/doug-martin/goqu/v9"
	"time"
)

func (s *repository) IsStarted(userTask model.UserTask) (bool, error) {
	query := goqu.From(userTaskTable)
	query = query.Where(goqu.Ex{"task_id": userTask.TaskId})
	query = query.Where(goqu.Ex{"user_id": userTask.UserId})
	sqlQuery, _, err := query.Select("start_time").ToSQL()
	if err != nil {
		return false, err
	}
	var t []time.Time
	err = s.db.Select(&t, sqlQuery)
	if err != nil {
		return false, err
	}
	if len(t) == 0 {
		return false, nil
	}

	return true, nil

}

func (s *repository) StartTask(req model.UserTask) error {
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

func (s *repository) IsActive(req model.UserTask) (bool, error) {
	query := goqu.From(userTaskTable)
	query = query.Where(goqu.Ex{"task_id": req.TaskId})
	query = query.Where(goqu.Ex{"user_id": req.UserId})
	query = query.Where(goqu.I("finish_time").IsNull())
	sqlQuery, _, err := query.Select("start_time").ToSQL()
	if err != nil {
		return false, err
	}
	var t []time.Time
	err = s.db.Select(&t, sqlQuery)
	if err != nil {
		return false, err
	}
	if len(t) == 0 {
		return false, nil
	}

	return true, nil
}

func (s *repository) FinishTask(req model.UserTask) error {
	query := goqu.From(userTaskTable)
	req.FinishTime = time.Now().UTC()
	query = query.Where(goqu.Ex{"task_id": req.TaskId})
	query = query.Where(goqu.Ex{"user_id": req.UserId})
	sqlQuery, _, err := query.Update().Set(
		goqu.Ex{"finish_time": req.FinishTime}).ToSQL()
	if err != nil {
		return err
	}
	_, err = s.db.Exec(sqlQuery, query)
	if err != nil {
		return err
	}
	return nil
}

func (s *repository) AddTaskDescription(descr string) (int64, error) {
	sqlQuery, _, err := goqu.Insert(taskTable).Rows(struct {
		description string
	}{description: descr}).Returning("task_id").ToSQL()
	if err != nil {
		return -1, err
	}
	var taskId int64
	err = s.db.Select(&taskId, sqlQuery)
	if err != nil {
		return -1, err
	}
	return taskId, nil
}

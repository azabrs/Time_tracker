package time_tracker_db

import (
	"Time-tracker/internal/model"
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

const (
	taskTable     = "tasks"
	userTable     = "users"
	userTaskTable = "task_user"
)

func (s *repository) GetUsers(req model.GetAllUsersReq) ([]model.GetAllUsersResp, error) {
	query := goqu.From(userTable)
	if len(req.Filter.UserID) > 0 {
		query = query.Where(goqu.Ex{"user_id": req.Filter.UserID})
	}
	if len(req.Filter.Description) > 0 {
		query = query.Where(goqu.Ex{"description": req.Filter.Description})
	}
	if len(req.Filter.PassportNumber) > 0 {
		query = query.Where(goqu.Ex{"passport_number": req.Filter.PassportNumber})
	}
	if len(req.Filter.PassportSerie) > 0 {
		query = query.Where(goqu.Ex{"passport_serie": req.Filter.PassportSerie})
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 25
	}

	if req.Sort.Asc {
		query = query.Order(goqu.I(req.Sort.By).Asc(), goqu.I("user_id").Desc())
	} else {
		query = query.Order(goqu.I(req.Sort.By).Desc(), goqu.I("user_id").Desc())
	}

	if req.Page != -1 {
		query = query.Limit(uint(req.PageSize)).Offset(uint((req.Page - 1) * req.PageSize))
	} else {
		query = query.Limit(100000)
	}
	var resp []model.GetAllUsersResp
	sqlQuery, _, err := query.Select("user_id", "passport_serie", "passport_number").ToSQL()
	if err != nil {
		return []model.GetAllUsersResp{}, nil
	}
	if err := s.db.Select(&resp, sqlQuery); err != nil {
		return []model.GetAllUsersResp{}, fmt.Errorf("get data: %w", err)
	}
	return resp, nil
}

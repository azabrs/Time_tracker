package time_tracker

import "Time-tracker/internal/model"

func (s *service) GetUsers(req model.GetAllUsersReq) ([]model.GetAllUsersResp, error) {
	resp, err := s.rep.GetUsers(req)
	if err != nil {
		return []model.GetAllUsersResp{}, nil
	}
	return resp, nil
}

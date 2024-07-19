package handler

import (
	"Time-tracker/internal/model"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var buf model.GetAllUsersReq
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = jsoniter.Unmarshal(body, &buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.Service.GetUsers(buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respB, err := jsoniter.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(respB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

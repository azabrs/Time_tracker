package handler

import "net/http"

type Handler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

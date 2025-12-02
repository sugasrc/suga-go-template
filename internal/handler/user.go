package handler

import (
	"net/http"
)

type UsersResponse struct {
	Message string `json:"message"`
}

func Users(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, UsersResponse{Message: "respond with a resource"})
}

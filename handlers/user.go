package handlers

import (
	"net/http"

	"github.com/neil-berg/go-rest/data"
)

// GetUsers returns a list of sample users
func (handler *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := data.GetUsers()
	err := users.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// AddUser adds a user to the list of sample users
func (handler *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey{}).(data.User)
	addedUser := data.AddUser(&user)
	handler.logger.Println(addedUser)
}

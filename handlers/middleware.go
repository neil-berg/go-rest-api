package handlers

import (
	"context"
	"net/http"

	"github.com/neil-berg/go-rest/data"
)

type RecipeKey struct{}

// ParseJSONRecipe parses JSON data from the request and deserializes it to a
// recipe struct
func (handler *Handler) ParseJSONRecipe(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recipe := data.Recipe{}

		err := recipe.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = recipe.Validate()
		if err != nil {
			handler.logger.Println("Data failed validation", err)
			http.Error(w, "Error reading recipe", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), RecipeKey{}, recipe)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

type UserKey struct{}

// ParseJSONUser parses incoming user data on the request body and deserializes
// it to a user struct
func (handler *Handler) ParseJSONUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := data.User{}

		err := user.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to marshal user JSON", http.StatusBadRequest)
		}

		err = user.Validate()
		if err != nil {
			handler.logger.Println("User data failed validation", err)
			http.Error(w, "Error reading user", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), UserKey{}, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

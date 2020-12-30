package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neil-berg/go-rest/data"
)

// Handler is the basic shape of handlers that take a logger
type Handler struct {
	logger *log.Logger
}

// CreateHandler returns a new Handler object
func CreateHandler(logger *log.Logger) *Handler {
	return &Handler{logger}
}

// GetRecipes fetches all recipes from the database
func (handler *Handler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipeList := data.GetRecipes()
	err := recipeList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// AddRecipe adds a new recipe to the database
func (handler *Handler) AddRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := r.Context().Value(RecipeKey{}).(data.Recipe)
	addedRecipe := data.AddRecipe(&recipe)
	handler.logger.Printf("Added recipe [id: %s]", addedRecipe.ID)
}

// UpdateRecipe updates an existing recipe in the database
func (handler *Handler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	recipe := r.Context().Value(RecipeKey{}).(data.Recipe)
	updatedRecipe, err := data.UpdateRecipe(id, &recipe)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusInternalServerError)
	}
	handler.logger.Printf("Updated recipe [id: %s]", updatedRecipe.ID)
}

type RecipeKey struct{}

// BodyParserJSON parses JSON data from the request and deserializes it to a recipe struct
func (handler *Handler) BodyParserJSON(next http.Handler) http.Handler {
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
			http.Error(w, "Error reading product", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), RecipeKey{}, recipe)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

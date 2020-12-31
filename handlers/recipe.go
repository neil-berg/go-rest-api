// Package handlers Recipe API
//
// Documentation for Recipe API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neil-berg/go-rest/data"
)

// swagger:response
type recipesResponse struct {
	// in:body
	Body []data.Recipe
}

// GetRecipes fetches all recipes from the database
// swagger:route GET /recipes recipes recipeList
// Returns a list of recipes
// responses:
//	200: recipesResponse
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

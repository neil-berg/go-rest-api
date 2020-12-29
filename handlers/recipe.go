package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/neil-berg/go-rest/data"
)

type Recipes struct {
	l *log.Logger
}

func NewRecipes(l *log.Logger) *Recipes {
	return &Recipes{l}
}

func (recipes *Recipes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		recipes.getRecipes(w, r)
		return
	}

	if r.Method == http.MethodPost {
		recipes.addRecipe(w, r)
		return
	}

	if r.Method == http.MethodPut {
		// Expect the ID in the URI
		re := regexp.MustCompile(`/([0-9]+)`)
		match := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(match) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(match[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		id := match[0][1]
		recipes.updateRecipe(id, w, r)
		return
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (recipes *Recipes) getRecipes(w http.ResponseWriter, r *http.Request) {
	recipeList := data.GetRecipes()
	err := recipeList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (recipes *Recipes) addRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := &data.Recipe{}
	err := recipe.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	addedRecipe := data.AddRecipe(recipe)
	fmt.Printf("Added recipe: %#v", addedRecipe)
}

func (recipes *Recipes) updateRecipe(id string, w http.ResponseWriter, r *http.Request) {
	recipe := &data.Recipe{}
	err := recipe.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	updatedRecipe, err := data.UpdateRecipe(id, recipe)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusInternalServerError)
	}
	fmt.Println("Updated recipe: %#v", updatedRecipe)
}

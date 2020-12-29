package handlers

import (
	"log"
	"net/http"

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

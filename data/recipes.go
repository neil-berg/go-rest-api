package data

import (
	"encoding/json"
	"io"
	"time"
)

type Recipe struct {
	ID           int
	Name         string
	Ingredients  []string
	Instructions string
	CreatedAt    string
	UpdatedAt    string
}

type Recipes []*Recipe

var recipeList = Recipes{
	&Recipe{
		ID:   1,
		Name: "Chocolate Chip Cookies",
		Ingredients: []string{
			"1 cup floud",
			"1/4 cup chocolate chips",
		},
		Instructions: "Mix butter into bowl. Stir in chips.",
		CreatedAt:    time.Now().UTC().String(),
		UpdatedAt:    time.Now().UTC().String(),
	},
	&Recipe{
		ID:   2,
		Name: "Chicken Pot Pie",
		Ingredients: []string{
			"2 chicken breasts",
			"3 cups peas",
		},
		Instructions: "Make dough and fold in chicken pieces.",
		CreatedAt:    time.Now().UTC().String(),
		UpdatedAt:    time.Now().UTC().String(),
	},
}

// GetRecipes returns the list of static recipes
func GetRecipes() Recipes {
	return recipeList
}

// ToJSON method converts recipes to JSON
func (r *Recipes) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(r)
}

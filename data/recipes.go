package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// Recipe is the shape of a recipe
type Recipe struct {
	ID           string   `json:"id"`
	Name         string   `json:"name" validate:"required"`
	Ingredients  []string `json:"ingredients" validate:"required"`
	Instructions string   `json:"instructions" validate:"required"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
}

// Recipes are a slice of recipe pointers
type Recipes []*Recipe

var recipeList = Recipes{
	&Recipe{
		ID:   "1",
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
		ID:   "2",
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

// Validate validates the recipe struct fields
func (recipe *Recipe) Validate() error {
	validate := validator.New()
	return validate.Struct(recipe)
}

// GetRecipes returns the list of static recipes
func GetRecipes() Recipes {
	return recipeList
}

// AddRecipe adds a recipe to the DB
func AddRecipe(recipe *Recipe) *Recipe {
	uuid, _ := uuid.NewRandom()
	recipe.ID = uuid.String()
	recipe.CreatedAt = time.Now().UTC().String()
	recipe.UpdatedAt = time.Now().UTC().String()

	recipeList = append(recipeList, recipe)
	return recipe
}

// UpdateRecipe updates an existing recipe in the DB
func UpdateRecipe(id string, recipe *Recipe) (*Recipe, error) {
	index, err := findRecipe(id)
	if err != nil {
		return nil, err
	}

	recipeList[index] = recipe
	return recipe, nil
}

var ErrorRecipeNotFound = fmt.Errorf("Recipe Not Found")

func findRecipe(id string) (int, error) {
	for i, r := range recipeList {
		if r.ID == id {
			return i, nil
		}
	}
	return -1, ErrorRecipeNotFound
}

// ToJSON method converts recipes to JSON
func (r *Recipes) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(r)
}

// FromJSON decodes JSON data
func (recipe *Recipe) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(recipe)
}

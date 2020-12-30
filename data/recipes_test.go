package data

import "testing"

// TestValidateRecipe checks the validator struct tags
func TestValidateRecipe(t *testing.T) {
	recipe := &Recipe{
		Name: "Test Recipe",
		Ingredients: []string{
			"Test ingredient 1",
			"Test ingredient 2",
		},
		Instructions: "Test instructions",
	}

	err := recipe.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

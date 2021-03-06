package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User defines the shape of a user
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Users is a collection of user pointers
type Users []*User

// SampleUsers are a sample of users
var sampleUsers = Users{
	&User{
		ID:       "1",
		Email:    "one@example.com",
		Password: "",
	},
	&User{
		ID:       "2",
		Email:    "two@example.com",
		Password: "",
	},
}

// Validate checks fields of the user struct
func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)

}

// FromJSON deserializes JSON users in a request
func (user *User) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(user)
}

// ToJSON serializes a user struct into JSON
func (users Users) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(users)
}

// GetUsers returns a list of sample users
func GetUsers() Users {
	return sampleUsers
}

// AddUser adds a new user to the sample users
func AddUser(user *User) *User {
	uuid, _ := uuid.NewRandom()
	user.ID = uuid.String()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hash)
	sampleUsers = append(sampleUsers, user)
	return user
}

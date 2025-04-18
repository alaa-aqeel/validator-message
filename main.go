package main

import (
	"encoding/json"
	"fmt"

	"github.com/alaa-aqeel/validator-message/validator"
)

type CreateUserDto struct {
	Name                 string `json:"name" validate:"required"`
	Username             string `json:"username" validate:"required,min=3,max=10"`
	Password             string `json:"password" validate:"required,min=6,max=20"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}

func main() {

	v := validator.InitlizeValidator()
	jsonData := []byte(`{
		"name": "alaa aqeel",
		"username": "al",
		"password": "123",
		"password_confirmation": "123456"
	}`)

	var user CreateUserDto
	if err := json.Unmarshal(jsonData, &user); err != nil {
		fmt.Println("Invalid JSON:", err)
		return
	}

	errors := v.Struct(user)
	if errors.HasErrors() {
		fmt.Println(errors)
	}
}

package handler

import "encoding/json"

type CreateUserRequest struct {
	FirstName       *string      `json:"FirstName"`
	LastName        *string      `json:"LastName"`
	Age             *json.Number `json:"Age"`
	PhoneNumber     *string      `json:"PhoneNumber"`
	IsPhoneVerified *bool        `json:"IsPhoneVerified"`
}

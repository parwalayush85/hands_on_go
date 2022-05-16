package handler

type getUserByIdResponse struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Age             int    `json:"age"`
	PhoneNumber     string `json:"Phone number"`
	IsPhoneVerified bool   `json:"Verified Phone number"`
}

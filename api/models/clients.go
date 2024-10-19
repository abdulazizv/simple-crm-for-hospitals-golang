package models

type ClientsRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Age         int    `json:"age,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type ClientsReq struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Age          int    `json:"age,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
type ClientUpdateReq struct {
	Id          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Age         int    `json:"age,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
type ClientsResponse struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Age         int    `json:"age,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type ClientsList struct {
	Clients []ClientsResponse `json:"clients"`
}

type ClientLoginReq struct {
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
}

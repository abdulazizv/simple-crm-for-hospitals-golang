package models

type ClinicReq struct {
	Name        string `json:"name,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type ClinicRes struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type UpdateClinicReq struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type ClinicList struct {
	Clinics []ClinicRes `json:"clinics"`
}

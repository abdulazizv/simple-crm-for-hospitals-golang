package models

type ServicesRequest struct {
	ClinicId int     `json:"clinic_id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

type ServicesResponse struct {
	Id       int     `json:"id,omitempty"`
	ClinicId int     `json:"clinic_id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

type UpdateServicesReq struct {
	Id    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type ServicesRes struct {
	Id       int              `json:"id,omitempty"`
	ClinicId int              `json:"clinic_id,omitempty"`
	Name     string           `json:"name,omitempty"`
	Price    float64          `json:"price,omitempty"`
	Doctors  []DoctorResponse `json:"doctors"`
}

type ServicesList struct {
	Services []ServicesResponse `json:"services"`
}

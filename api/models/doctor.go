package models

type DoctorReqForUI struct {
	ClinicId    int    `json:"clinic_id,omitempty"`
	ServiceId   int    `json:"service_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	WorkDay     string `json:"work_day,omitempty"`
	Floor       int    `json:"floor,omitempty"`
	RoomNumber  int    `json:"room_number,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Experience  int    `json:"experience,omitempty"`
}

type DoctorRequest struct {
	ClinicId     int    `json:"clinic_id,omitempty"`
	ServiceId    int    `json:"service_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
	EndTime      string `json:"end_time,omitempty"`
	WorkDay      string `json:"work_day,omitempty"`
	Floor        int    `json:"floor,omitempty"`
	RoomNumber   int    `json:"room_number,omitempty"`
	ImageLink    string `json:"image_link,omitempty"`
	Experience   int    `json:"experience,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type UpdateDoctor struct {
	Id          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	WorkDay     string `json:"work_day,omitempty"`
	Floor       int    `json:"floor,omitempty"`
	RoomNumber  int    `json:"room_number,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Experience  int    `json:"experience,omitempty"`
}

type DoctorResponse struct {
	Id          int    `json:"id,omitempty"`
	ClinicId    int    `json:"clinic_id,omitempty"`
	ServiceId   int    `json:"service_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	WorkDay     string `json:"work_day,omitempty"`
	Floor       int    `json:"floor,omitempty"`
	RoomNumber  int    `json:"room_number,omitempty"`
	ImageLink   string `json:"image_link,omitempty"`
	Experience  int    `json:"experience,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

type DoctorsList struct {
	Doctors []DoctorResponse `json:"doctors"`
}
type DoctorResInfo struct {
	Id        int    `json:"id,omitempty"`
	ClinicId  int    `json:"clinic_id,omitempty"`
	ServiceId int    `json:"service_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ImageLink string `json:"image_link,omitempty"`
	WorkDay   string `json:"work_day,omitempty"`
}
type Doctors struct {
	Doctor []DoctorResInfo `json:"doctors"`
}

type DoctorLoginReq struct {
	PhoneNumber string `json:"phone_number"`
}
type GetCustomersOfDoctor struct {
	DoctorID  int    `json:"doctor_id,omitempty"`
	ClientID  int    `json:"client_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Status    string `json:"status,omitempty"`
}
type DoctorLoginRes struct {
	Id          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	ClinicId    int    `json:"clinic_id"`
	PhoneNumber string `json:"phone_number,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type CheckfieldReq struct {
	Field string `json:"field,omitempty"`
	Value string `json:"value,omitempty"`
}
type CheckFieldRes struct {
	Exists bool `json:"exists"`
}

package models

type KorikRequest struct {
	ClientId    int    `json:"client_id,omitempty"`
	DoctorId    int    `json:"doctor_id,omitempty"`
	DiagnosisId int    `json:"diagnosis_id,omitempty"`
	FileUrl     string `json:"file_url,omitempty"`
}

type UpdateKorikRequest struct {
	Id          int    `json:"id,omitempty"`
	ClientId    int    `json:"client_id,omitempty"`
	DoctorId    int    `json:"doctor_id,omitempty"`
	DiagnosisId int    `json:"diagnosis_id,omitempty"`
	FileUrl     string `json:"file_url,omitempty"`
	Count       int    `json:"count,omitempty"`
}

type KorikResponse struct {
	Id          int    `json:"id,omitempty"`
	ClientId    int    `json:"client_id,omitempty"`
	DoctorId    int    `json:"doctor_id,omitempty"`
	DiagnosisId int    `json:"diagnosis_id,omitempty"`
	FileUrl     string `json:"file_url,omitempty"`
	Count       int    `json:"count,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

type ClientsKorikResponse struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type DoctorKorikResponse struct {
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

type DiagnosisKorikResponse struct {
	Id        int    `json:"id"`
	Reason    string `json:"reason"`
	ServiceId int    `json:"service_id"`
	CreatedAt string `json:"created_at"`
}

type KorikGetResponse struct {
	Korik     KorikResponse          `json:"korik"`
	Doctor    DoctorKorikResponse    `json:"doctor"`
	Client    ClientsKorikResponse   `json:"client"`
	Diagnosis DiagnosisKorikResponse `json:"diagnosis"`
}

type KoriksGetListResponse struct {
	Koriks []KorikGetResponse `json:"koriks"`
}

type KoriksList struct {
	Koriks []KorikResponse `json:"koriks"`
}

/*
 client_id INT NOT NULL REFERENCES clients(id),
    doctor_id INT NOT NULL REFERENCES doctors(id),
    diagnosis_id INT NOT NULL REFERENCES diagnosis(id),
    file_url TEXT NOT NULL,
    count INT,
*/

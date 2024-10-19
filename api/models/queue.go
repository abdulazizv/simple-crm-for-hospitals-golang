package models

type QueueRes struct {
	ID          int `json:"id,omitempty"`
	ClientID    int `json:"client_id,omitempty"`
	DoctorID    int `json:"doctor_id,omitempty"`
	QueueNumber int `json:"queue_number,omitempty"`
}

type QueueReq struct {
	ClientID int `json:"client_id,omitempty"`
	DoctorID int `json:"doctor_id,omitempty"`
}

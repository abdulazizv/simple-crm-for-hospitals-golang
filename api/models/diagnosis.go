package models

type DiagnosisRequest struct {
	Reason    string `json:"reason"`
	ServiceId int    `json:"service_id"`
}

type DiagnosisResponse struct {
	Id        int    `json:"id"`
	Reason    string `json:"reason"`
	ServiceId int    `json:"service_id"`
	CreatedAt string `json:"created_at"`
}

package postgres

import (
	"fmt"

	"gitlab.com/backend/api/models"
)

func (r storageRepo) CreateKorik(req *models.KorikRequest) (*models.KorikResponse, error) {
	res := models.KorikResponse{}
	query := `
		INSERT INTO korik(client_id,doctor_id,diagnosis_id,file_url)
		VALUES($1,$2,$3,$4) RETURNING id,client_id,doctor_id,diagnosis_id,file_url,created_at
    `

	err := r.db.QueryRow(query, req.ClientId, req.DoctorId, req.DiagnosisId, req.FileUrl).
		Scan(&res.Id, &res.ClientId, &res.DoctorId, &res.DiagnosisId, &res.FileUrl, &res.CreatedAt)

	if err != nil {
		return &models.KorikResponse{}, err
	}
	return &res, nil
}

func (r storageRepo) GetKorik(id int) (*models.KorikGetResponse, error) {
	res := models.KorikGetResponse{}
	query := `
		SELECT korik.id,korik.client_id,korik.doctor_id,korik.diagnosis_id,korik.file_url,COALESCE(korik.count,0),clients.id,clients.first_name,clients.last_name,clients.phone_number,doctors.id,doctors.first_name,doctors.last_name,doctors.phone_number,doctors.start_time,doctors.end_time,doctors.work_day,doctors.floor,doctors.room_number,doctors.image_link,doctors.experience,diagnosis.id,diagnosis.reason,diagnosis.service_id,diagnosis.created_at FROM korik JOIN clients ON clients.id = korik.client_id JOIN doctors ON doctors.id = korik.doctor_id JOIN diagnosis ON diagnosis.id = korik.diagnosis_id WHERE korik.id=$1
    `
	err := r.db.QueryRow(query, id).Scan(&res.Korik.Id, &res.Korik.ClientId, &res.Korik.DoctorId, &res.Korik.DiagnosisId, &res.Korik.FileUrl, &res.Korik.Count, &res.Client.Id, &res.Client.FirstName, &res.Client.LastName, &res.Client.PhoneNumber, &res.Doctor.Id, &res.Doctor.FirstName, &res.Doctor.LastName, &res.Doctor.PhoneNumber, &res.Doctor.StartTime, &res.Doctor.EndTime, &res.Doctor.WorkDay, &res.Doctor.Floor, &res.Doctor.RoomNumber, &res.Doctor.ImageLink, &res.Doctor.Experience, &res.Diagnosis.Id, &res.Diagnosis.Reason, &res.Diagnosis.ServiceId, &res.Diagnosis.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r storageRepo) GetKoriks() (*models.KoriksGetListResponse, error) {
	var res = models.KoriksGetListResponse{}
	rows, err := r.db.Query(`SELECT korik.id,korik.client_id,korik.doctor_id,korik.diagnosis_id,korik.file_url,COALESCE(korik.count,0),clients.id,clients.first_name,clients.last_name,clients.phone_number,doctors.id,doctors.first_name,doctors.last_name,doctors.phone_number,doctors.start_time,doctors.end_time,doctors.work_day,doctors.floor,doctors.room_number,doctors.image_link,doctors.experience,diagnosis.id,diagnosis.reason,diagnosis.service_id,diagnosis.created_at FROM korik JOIN clients ON clients.id = korik.client_id JOIN doctors ON doctors.id = korik.doctor_id JOIN diagnosis ON diagnosis.id = korik.diagnosis_id`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := models.KorikGetResponse{}
		err = rows.Scan(&temp.Korik.Id, &temp.Korik.ClientId, &temp.Korik.DoctorId, &temp.Korik.DiagnosisId, &temp.Korik.FileUrl, &temp.Korik.Count, &temp.Client.Id, &temp.Client.FirstName, &temp.Client.LastName, &temp.Client.PhoneNumber, &temp.Doctor.Id, &temp.Doctor.FirstName, &temp.Doctor.LastName, &temp.Doctor.PhoneNumber, &temp.Doctor.StartTime, &temp.Doctor.EndTime, &temp.Doctor.WorkDay, &temp.Doctor.Floor, &temp.Doctor.RoomNumber, &temp.Doctor.ImageLink, &temp.Doctor.Experience, &temp.Diagnosis.Id, &temp.Diagnosis.Reason, &temp.Diagnosis.ServiceId, &temp.Diagnosis.CreatedAt)

		if err != nil {
			return &models.KoriksGetListResponse{}, err
		}
		fmt.Println(temp.Korik.Count)
		res.Koriks = append(res.Koriks, temp)
	}
	return &res, nil
}

func (r storageRepo) UpdateKorik(req *models.UpdateKorikRequest) (*models.KorikResponse, error) {
	var res models.KorikResponse

	query := `
		UPDATE 
			korik SET client_id=$1,doctor_id=$2,diagnosis_id=$3,file_url=$4,count=$5
	WHERE id=$6 RETURNING id,client_id,doctor_id,diagnosis_id,file_url,count	
    `
	err := r.db.QueryRow(query, req.ClientId, req.DoctorId, req.DiagnosisId, req.FileUrl, req.Count, req.Id).
		Scan(&res.Id, &res.ClientId, &res.DoctorId, &res.DiagnosisId, &res.FileUrl, &res.Count)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r storageRepo) DeleteKorik(id int) error {
	_, err := r.db.Exec(`DELETE FROM korik WHERE id=$1`, id)

	if err != nil {
		return err
	}
	return nil
}

func (r storageRepo) GetKorikByUserId(id int) (*models.KorikGetResponse, error) {
	res := models.KorikGetResponse{}
	query := `
		SELECT korik.id,korik.client_id,korik.doctor_id,korik.diagnosis_id,korik.file_url,COALESCE(korik.count,0),clients.id,clients.first_name,clients.last_name,clients.phone_number,doctors.id,doctors.first_name,doctors.last_name,doctors.phone_number,doctors.start_time,doctors.end_time,doctors.work_day,doctors.floor,doctors.room_number,doctors.image_link,doctors.experience,diagnosis.id,diagnosis.reason,diagnosis.service_id,diagnosis.created_at FROM korik JOIN clients ON clients.id = korik.client_id JOIN doctors ON doctors.id = korik.doctor_id JOIN diagnosis ON diagnosis.id = korik.diagnosis_id WHERE korik.client_id=$1
    `
	err := r.db.QueryRow(query, id).Scan(&res.Korik.Id, &res.Korik.ClientId, &res.Korik.DoctorId, &res.Korik.DiagnosisId, &res.Korik.FileUrl, &res.Korik.Count, &res.Client.Id, &res.Client.FirstName, &res.Client.LastName, &res.Client.PhoneNumber, &res.Doctor.Id, &res.Doctor.FirstName, &res.Doctor.LastName, &res.Doctor.PhoneNumber, &res.Doctor.StartTime, &res.Doctor.EndTime, &res.Doctor.WorkDay, &res.Doctor.Floor, &res.Doctor.RoomNumber, &res.Doctor.ImageLink, &res.Doctor.Experience, &res.Diagnosis.Id, &res.Diagnosis.Reason, &res.Diagnosis.ServiceId, &res.Diagnosis.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

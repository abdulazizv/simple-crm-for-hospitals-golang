package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/backend/api/models"
)

// Create Doctor
func (r storageRepo) CreateDoctor(req *models.DoctorRequest) (*models.DoctorResponse, error) {
	var res = models.DoctorResponse{}

	query := `
INSERT INTO 
   doctors(
           clinic_id,service_id,first_name,last_name,phone_number,start_time,end_time,work_day,floor,room_number,image_link,experience
   )
   VALUES(
          $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
   )
   RETURNING id,clinic_id,service_id,first_name,last_name,phone_number,start_time,end_time,work_day,floor,room_number,image_link,experience`

	err := r.db.QueryRow(query, req.ClinicId, req.ServiceId, req.FirstName, req.LastName, req.PhoneNumber, req.StartTime, req.EndTime, req.WorkDay, req.Floor, req.RoomNumber, req.ImageLink, req.Experience).
		Scan(&res.Id, &res.ClinicId, &res.ServiceId, &res.FirstName, &res.LastName, &res.PhoneNumber, &res.StartTime, &res.EndTime, &res.WorkDay, &res.Floor, &res.RoomNumber, &res.ImageLink, &res.Experience)

	if err != nil {
		return &models.DoctorResponse{}, err
	}

	return &res, nil
}

func (r storageRepo) GetDoctor(id int) (*models.DoctorResponse, error) {
	res := models.DoctorResponse{}

	query := `
SELECT 
    id, clinic_id, service_id, first_name, last_name, 
    phone_number, start_time, end_time, work_day, 
    floor, room_number, image_link, experience, created_at
FROM 
    doctors WHERE id=$1`

	err := r.db.QueryRow(query, id).Scan(&res.Id, &res.ClinicId, &res.ServiceId, &res.FirstName, &res.LastName, &res.PhoneNumber, &res.StartTime, &res.EndTime, &res.WorkDay, &res.Floor, &res.RoomNumber, &res.ImageLink, &res.Experience, &res.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &res, nil
}

func (r storageRepo) GetDoctorsList() (*models.DoctorsList, error) {
	res := models.DoctorsList{}

	query := `SELECT doctors.id,doctors.clinic_id,doctors.service_id,doctors.first_name,doctors.last_name,doctors.phone_number,doctors.start_time,doctors.end_time,doctors.work_day,doctors.floor,doctors.room_number,doctors.image_link,doctors.experience,doctors.created_at FROM doctors JOIN clinics ON clinics.id = doctors.clinic_id`

	rows, err := r.db.Query(query)

	if err != nil {
		return &models.DoctorsList{}, err
	}

	for rows.Next() {
		temp := models.DoctorResponse{}
		err := rows.Scan(&temp.Id, &temp.ClinicId, &temp.ServiceId, &temp.FirstName, &temp.LastName, &temp.PhoneNumber, &temp.StartTime, &temp.EndTime, &temp.WorkDay, &temp.Floor, &temp.RoomNumber, &temp.ImageLink, &temp.Experience, &temp.CreatedAt)
		if err != nil {
			return &models.DoctorsList{}, err
		}
		res.Doctors = append(res.Doctors, temp)
	}
	return &res, nil
}

func (r storageRepo) GetDoctorForLogin() (*models.DoctorLoginRes, error) {
	var res models.DoctorLoginRes

	err := r.db.QueryRow(`SELECT id,first_name,last_name,phone_number FROM doctors`).Scan(&res.Id, &res.FirstName, &res.LastName, &res.PhoneNumber)

	if err != nil {
		return &models.DoctorLoginRes{}, err
	}
	return &res, nil
}

func (r storageRepo) GetDoctorByPhoneNumber(phoneNumber string) (*models.DoctorLoginRes, error) {
	res := models.DoctorLoginRes{}
	err := r.db.QueryRow(`
SELECT 
    id, clinic_id, first_name, last_name, phone_number 
FROM 
    doctors WHERE phone_number=$1`, phoneNumber).
		Scan(&res.Id, &res.ClinicId, &res.FirstName, &res.LastName, &res.PhoneNumber)

	if err != nil {
		return &models.DoctorLoginRes{}, err
	}

	return &res, nil
}

func (r storageRepo) UpdateDoctor(req *models.UpdateDoctor) error {
	query := `
UPDATE doctors SET
	first_name=$1,last_name=$2,phone_number=$3,
	start_time=$4,end_time=$5,work_day=$6,
	floor=$7,room_number=$8,image_link=$9,experience=$10 WHERE id=$11
    `
	_, err := r.db.Exec(query,
		req.FirstName, req.LastName, req.PhoneNumber,
		req.StartTime, req.EndTime, req.WorkDay, req.Floor,
		req.RoomNumber, req.ImageLink, req.Experience, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r storageRepo) DeleteDoctor(id int) error {
	_, err := r.db.Exec(`DELETE FROM doctors WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r storageRepo) CheckField(req *models.CheckfieldReq) (*models.CheckFieldRes, error) {
	res := &models.CheckFieldRes{Exists: false}
	query := fmt.Sprintf("SELECT 1 FROM doctors WHERE %s=$1", req.Field)
	var temp = 0
	err := r.db.QueryRow(query, req.Value).Scan(&temp)
	if errors.Is(err, sql.ErrNoRows) {
		return res, nil
	} else if err != nil {
		return res, err
	}
	if temp == 1 {
		res.Exists = true
		return res, nil
	}
	return res, nil
}

func (r storageRepo) GetDoctorsSearch(id int, keyword string) (*models.DoctorsList, error) {
	var res models.DoctorsList
	query := `
SELECT 
    d.id, d.clinic_id, d.first_name, d.last_name, d.image_link, s.name
FROM 
    doctors d JOIN clinics c ON d.clinic_id=c.id JOIN services s ON s.id=d.service_id 
WHERE 
    d.clinic_id=$1 AND (d.first_name ILIKE $2 OR s.name ILIKE $2)`
	rows, err := r.db.Query(query, id, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.DoctorResponse{}
		err = rows.Scan(&temp.Id, &temp.ClinicId, &temp.FirstName, &temp.LastName, &temp.ImageLink, &temp.ServiceName)
		if err != nil {
			return nil, err
		}
		res.Doctors = append(res.Doctors, temp)
	}
	return &res, nil
}

func (r storageRepo) GetDoctorsByClinicId(id int) (*models.DoctorsList, error) {
	res := models.DoctorsList{}
	query := `
SELECT 
    id, clinic_id, first_name, last_name, phone_number, image_link 
FROM 
    doctors WHERE clinic_id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.DoctorResponse{}
		err = rows.Scan(&temp.Id, &temp.ClinicId, &temp.FirstName, &temp.LastName, &temp.PhoneNumber, &temp.ImageLink)
		if err != nil {
			return nil, err
		}
		res.Doctors = append(res.Doctors, temp)
	}
	return &res, nil
}

func (r storageRepo) GetDoctorsByService(clinicID int, key string) (*models.Doctors, error) {
	var res = models.Doctors{}
	query := `
SELECT 
	d.id, d.clinic_id, d.service_id, d.first_name, d.last_name, d.image_link, d.work_day  
FROM 
    doctors d JOIN services s ON s.id=d.service_id WHERE d.clinic_id=$1 AND s.name=$2`
	rows, err := r.db.Query(query, clinicID, key)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.DoctorResInfo{}
		err = rows.Scan(
			&temp.Id, &temp.ClinicId, &temp.ServiceId, &temp.FirstName,
			&temp.LastName, &temp.ImageLink, &temp.WorkDay)
		if err != nil {
			return nil, err
		}
		res.Doctor = append(res.Doctor, temp)
	}
	return &res, nil
}

func (r storageRepo) GetDoctorsByServiceID(id int) (*models.DoctorsList, error) {
	res := models.DoctorsList{}
	query := `
SELECT id, clinic_id, service_id, first_name, last_name, 
    phone_number, start_time, end_time, work_day, 
    floor, room_number, image_link, experience 
FROM doctors WHERE service_id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.DoctorResponse{}
		err = rows.Scan(&temp.Id, &temp.ClinicId, &temp.ServiceId, &temp.FirstName,
			&temp.LastName, &temp.PhoneNumber, &temp.StartTime, &temp.EndTime, &temp.WorkDay,
			&temp.Floor, &temp.RoomNumber, &temp.ImageLink, &temp.Experience)
		if err != nil {
			return nil, err
		}
		res.Doctors = append(res.Doctors, temp)
	}
	return &res, nil
}

func (r storageRepo) GetCustomersByDoctorID(id int) ([]*models.GetCustomersOfDoctor, error) {
	res := []*models.GetCustomersOfDoctor{}
	query := `
SELECT 
    q.doctor_id, c.id, c.first_name, c.last_name, q.Status 
FROM 
    clients c INNER JOIN queue q ON q.client_id=c.id WHERE q.doctor_id=$1 AND DATE(q.created_at) = CURRENT_DATE`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.GetCustomersOfDoctor{}
		err = rows.Scan(&temp.DoctorID, &temp.ClientID, &temp.LastName, &temp.LastName, &temp.Status)
		if err != nil {
			return nil, err
		}
		res = append(res, &temp)
	}
	return res, nil
}

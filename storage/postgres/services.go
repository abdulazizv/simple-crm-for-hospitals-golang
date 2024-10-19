package postgres

import (
	"gitlab.com/backend/api/models"
)

func (r storageRepo) CreateServices(req *models.ServicesRequest) (*models.ServicesResponse, error) {
	var res = models.ServicesResponse{}
	query := `
		INSERT INTO services(clinic_id,name,price) VALUES($1,$2,$3)
		RETURNING id,clinic_id,name,price
	`

	err := r.db.QueryRow(query, req.ClinicId, req.Name, req.Price).Scan(&res.Id, &res.ClinicId, &res.Name, &res.Price)

	if err != nil {
		return &models.ServicesResponse{}, err
	}
	return &res, nil
}

func (r storageRepo) GetService(id int) (*models.ServicesRes, error) {
	res := models.ServicesRes{}
	query := `
		SELECT id,clinic_id,name,price FROM services WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(&res.Id, &res.ClinicId, &res.Name, &res.Price)
	if err != nil {
		return &models.ServicesRes{}, err
	}
	query1 := `
SELECT 
	id, clinic_id, service_id, first_name, last_name, 
	phone_number, start_time, end_time, work_day, floor, 
	room_number, image_link, experience
FROM doctors WHERE service_id=$1
`

	rows, err := r.db.Query(query1, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.DoctorResponse{}
		err = rows.Scan(
			&temp.Id, &temp.ClinicId, &temp.ServiceId, &temp.FirstName,
			&temp.LastName, &temp.PhoneNumber, &temp.StartTime, &temp.EndTime,
			&temp.WorkDay, &temp.Floor, &temp.RoomNumber, &temp.ImageLink, &temp.Experience)
		if err != nil {
			return nil, err
		}
		res.Doctors = append(res.Doctors, temp)
	}
	return &res, nil
}

func (r storageRepo) GetServicesList() (*models.ServicesList, error) {
	res := models.ServicesList{}
	query := `SELECT * FROM services`
	rows, err := r.db.Query(query)

	if err != nil {
		return &models.ServicesList{}, err
	}

	for rows.Next() {
		temp := models.ServicesResponse{}
		err = rows.Scan(&temp.Id, &temp.ClinicId, &temp.Name, &temp.Price)
		if err != nil {
			return &models.ServicesList{}, err
		}
		res.Services = append(res.Services, temp)
	}

	return &res, nil
}

func (r storageRepo) DeleteService(id int) error {
	_, err := r.db.Exec(`DELETE FROM services WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r storageRepo) UpdateServices(req *models.UpdateServicesReq) (*models.ServicesResponse, error) {
	var res models.ServicesResponse
	query := `
	  	UPDATE
			services SET name=$1, price=$2 WHERE id=$3
	  	RETURNING id, clinic_id, name, price
    `
	err := r.db.QueryRow(query, req.Name, req.Price, req.Id).
		Scan(&res.Id, &res.ClinicId, &res.Name, &res.Price)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

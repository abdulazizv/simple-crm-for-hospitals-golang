package postgres

import (
	"database/sql"
	"gitlab.com/backend/api/models"
)

func (r storageRepo) CreateClinic(req *models.ClinicReq) (*models.ClinicRes, error) {
	res := models.ClinicRes{}
	query := `
INSERT INTO 
    clinics(name, image_link, phone_number, address) 
VALUES 
    ($1, $2, $3, $4) RETURNING id, name, image_link, phone_number, address`
	err := r.db.QueryRow(query, req.Name, req.ImageLink, req.PhoneNumber, req.Address).
		Scan(&res.Id, &res.Name, &res.ImageLink, &res.PhoneNumber, &res.Address)
	if err != nil {
		return &models.ClinicRes{}, err
	}
	return &res, nil
}

func (r storageRepo) GetClinic(id int) (*models.ClinicRes, error) {
	res := models.ClinicRes{}
	err := r.db.QueryRow(`SELECT id, name,phone_number, image_link, address FROM clinics WHERE id=$1`, id).
		Scan(&res.Id, &res.Name, &res.ImageLink, &res.PhoneNumber, &res.Address)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &res, nil
}

func (r storageRepo) GetList() (*models.ClinicList, error) {
	var res = models.ClinicList{}
	rows, err := r.db.Query(`SELECT id, name, image_link, phone_number, address FROM clinics`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.ClinicRes{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.ImageLink, &temp.PhoneNumber, &temp.Address)
		if err != nil {
			return nil, err
		}
		res.Clinics = append(res.Clinics, temp)
	}
	return &res, nil
}

func (r storageRepo) UpdateClinics(req *models.UpdateClinicReq) (*models.ClinicRes, error) {
	var res models.ClinicRes
	query := `UPDATE clinics SET name=$1,image_link=$2,address=$3,phone_number=$4
              WHERE id=$5 RETURNING id,name,image_link,address,phone_number`

	err := r.db.QueryRow(query, req.Name, req.ImageLink, req.Address, req.PhoneNumber, req.Id).
		Scan(&res.Id, &res.Name, &res.ImageLink, &res.Address, &res.PhoneNumber)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r storageRepo) DeleteClinics(id int) error {
	_, err := r.db.Exec(`DELETE FROM clinics WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

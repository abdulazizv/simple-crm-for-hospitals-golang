package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/backend/api/models"
)

func (r storageRepo) CreateClient(req *models.ClientsReq) (*models.ClientsResponse, error) {
	var res models.ClientsResponse
	query := `
INSERT INTO 
    clients(first_name, last_name, age, phone_number, refresh_token) 
VALUES 
    ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, age, phone_number`
	err := r.db.QueryRow(query, req.FirstName, req.LastName, req.Age, req.PhoneNumber, req.RefreshToken).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.Age, &res.PhoneNumber)
	if err != nil {
		return &models.ClientsResponse{}, err
	}
	return &res, nil
}

func (r storageRepo) GetClient(id int) (*models.ClientsResponse, error) {
	var res = models.ClientsResponse{}
	query := `SELECT id, first_name, last_name, age, phone_number FROM clients WHERE id=$1`
	err := r.db.QueryRow(query, id).Scan(&res.Id, &res.FirstName, &res.LastName, &res.Age, &res.PhoneNumber)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &res, nil
}

func (r storageRepo) GetClientForLogin(name, phone string) (*models.ClientsResponse, error) {
	var res = models.ClientsResponse{}
	err := r.db.QueryRow(`
SELECT 
    id, first_name, last_name, age, phone_number 
FROM 
    clients WHERE first_name=$1 AND phone_number=$2`, name, phone).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.Age, &res.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (r storageRepo) GetClients() (*models.ClientsList, error) {
	var res = models.ClientsList{}
	rows, err := r.db.Query(`SELECT id, first_name, last_name, age, phone_number FROM clients`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.ClientsResponse{}
		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Age, &temp.PhoneNumber)
		if err != nil {
			return nil, err
		}
		res.Clients = append(res.Clients, temp)
	}
	return &res, nil
}

func (r storageRepo) UpdateClient(req *models.ClientUpdateReq) (*models.ClientsResponse, error) {
	var res models.ClientsResponse
	query := `
UPDATE 
    clients SET first_name=$1, last_name=$2, age=$3, phone_number=$4
WHERE 
    id=$5 RETURNING id, first_name, last_name, age, phone_number`
	err := r.db.QueryRow(query, req.FirstName, req.LastName, req.Age, req.PhoneNumber, req.Id).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.Age, &res.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r storageRepo) DeleteClient(id int) error {
	_, err := r.db.Exec(`DELETE FROM clients WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r storageRepo) CheckFieldClient(req *models.CheckfieldReq) (*models.CheckFieldRes, error) {
	res := &models.CheckFieldRes{Exists: false}
	query := fmt.Sprintf("SELECT 1 FROM clients WHERE %s=$1", req.Field)
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

package postgres

import "gitlab.com/backend/api/models"

func (r storageRepo) CreateQueue(doctorId, clientID int) (*models.QueueRes, error) {
	var res = models.QueueRes{}
	query := `
INSERT INTO 
    queue(doctor_id, client_id, status, queue_number) 
VALUES
    ($1, $2, $3, $4)  
RETURNING 
	id, doctor_id, client_id, queue_number`
	err := r.db.QueryRow(query, doctorId, clientID, "process", 1).Scan(&res.ID, &res.DoctorID, &res.ClientID, &res.QueueNumber)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r storageRepo) CancelQueue(doctorId, clientId int) error {
	_, err := r.db.Exec(`DELETE FROM queue WHERE doctor_id=$1 AND client_id=$2`, doctorId, clientId)
	if err != nil {
		return err
	}
	return nil
}

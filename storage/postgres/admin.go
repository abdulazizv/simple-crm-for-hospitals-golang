package postgres

import "gitlab.com/backend/api/models"

// Create Admin
func (r storageRepo) CreateAdmin(req *models.AdminReq) (*models.AdminRes, error) {
	var res = models.AdminRes{}
	query := `
INSERT INTO 
    admins( 
        	name, username, password, refresh_token
        	) 
VALUES
    (
     $1, $2, $3, $4) 
RETURNING 
    id, name, username`
	err := r.db.QueryRow(query, req.Name, req.UserName, req.Password, req.RefreshToken).
		Scan(&res.Id, &res.Name, &res.UserName)
	if err != nil {
		return &models.AdminRes{}, err
	}
	return &res, nil
}

// Get Admin
func (r storageRepo) GetAdmin(id int) (*models.AdminRes, error) {
	res := models.AdminRes{}
	query := `
SELECT id, name, username FROM admins WHERE id=$1 `
	err := r.db.QueryRow(query, id).Scan(&res.Id, &res.Name, &res.UserName)
	if err != nil {
		return &models.AdminRes{}, err
	}
	return &res, nil
}

func (r storageRepo) GetAdminList() (*models.ListAdmin, error) {
	res := models.ListAdmin{}
	query := `
SELECT id, name, username FROM admins`
	rows, err := r.db.Query(query)
	if err != nil {
		return &models.ListAdmin{}, err
	}
	for rows.Next() {
		temp := models.AdminRes{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.UserName)
		if err != nil {
			return &models.ListAdmin{}, err
		}
		res.Admins = append(res.Admins, temp)
	}
	return &res, nil
}

func (r storageRepo) GetAdminForLogin() (*models.AdminRes, error) {
	var res models.AdminRes
	err := r.db.QueryRow(`SELECT id, name, username FROM admins`).Scan(&res.Id, &res.Name, &res.UserName)
	if err != nil {
		return &models.AdminRes{}, err
	}
	return &res, nil
}

func (r storageRepo) GetAdminByUsername(username string) (*models.AdminRes, error) {
	res := models.AdminRes{}
	err := r.db.QueryRow(`SELECT id, name, username, password FROM admins WHERE username=$1`, username).Scan(&res.Id, &res.Name, &res.UserName, &res.Password)
	if err != nil {
		return &models.AdminRes{}, err
	}
	return &res, nil
}

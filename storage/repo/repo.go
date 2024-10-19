package repo

import "gitlab.com/backend/api/models"

type ClinicI interface {
	//admin
	CreateAdmin(req *models.AdminReq) (*models.AdminRes, error)
	GetAdmin(id int) (*models.AdminRes, error)
	GetAdminList() (*models.ListAdmin, error)
	GetAdminForLogin() (*models.AdminRes, error)
	GetAdminByUsername(username string) (*models.AdminRes, error)

	//doctor
	CreateDoctor(req *models.DoctorRequest) (*models.DoctorResponse, error)
	GetDoctor(id int) (*models.DoctorResponse, error)
	GetDoctorsList() (*models.DoctorsList, error)
	GetDoctorForLogin() (*models.DoctorLoginRes, error)
	GetDoctorByPhoneNumber(phoneNumber string) (*models.DoctorLoginRes, error)
	UpdateDoctor(req *models.UpdateDoctor) error
	DeleteDoctor(id int) error
	GetDoctorsByClinicId(id int) (*models.DoctorsList, error)
	GetDoctorsSearch(id int, keyword string) (*models.DoctorsList, error)
	GetDoctorsByService(clinicId int, keyword string) (*models.Doctors, error)
	GetDoctorsByServiceID(id int) (*models.DoctorsList, error)
	GetCustomersByDoctorID(id int) ([]*models.GetCustomersOfDoctor, error)

	//services
	CreateServices(req *models.ServicesRequest) (*models.ServicesResponse, error)
	GetService(id int) (*models.ServicesRes, error)
	GetServicesList() (*models.ServicesList, error)
	DeleteService(id int) error
	UpdateServices(req *models.UpdateServicesReq) (*models.ServicesResponse, error)

	// clinic
	CreateClinic(req *models.ClinicReq) (*models.ClinicRes, error)
	GetClinic(id int) (*models.ClinicRes, error)
	GetList() (*models.ClinicList, error)
	UpdateClinics(req *models.UpdateClinicReq) (*models.ClinicRes, error)
	DeleteClinics(id int) error

	//korik
	CreateKorik(req *models.KorikRequest) (*models.KorikResponse, error)
	GetKorik(id int) (*models.KorikGetResponse, error)
	GetKoriks() (*models.KoriksGetListResponse, error)
	UpdateKorik(req *models.UpdateKorikRequest) (*models.KorikResponse, error)
	DeleteKorik(id int) error
	GetKorikByUserId(id int) (*models.KorikGetResponse, error)

	// client
	CreateClient(req *models.ClientsReq) (*models.ClientsResponse, error)
	GetClient(id int) (*models.ClientsResponse, error)
	GetClients() (*models.ClientsList, error)
	GetClientForLogin(name, phone string) (*models.ClientsResponse, error)
	UpdateClient(req *models.ClientUpdateReq) (*models.ClientsResponse, error)
	DeleteClient(id int) error

	// checkfield
	CheckField(req *models.CheckfieldReq) (*models.CheckFieldRes, error)
	CheckFieldClient(req *models.CheckfieldReq) (*models.CheckFieldRes, error)

	// queue
	CreateQueue(docterID, clientID int) (*models.QueueRes, error)
	CancelQueue(doctorID, clientID int) error
}

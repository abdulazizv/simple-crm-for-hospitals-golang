package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
	"strconv"
)

// Doctor Login
// @ID logindoctor
// @Router /doctor/login [POST]
// @Summary Login Doctor
// @Description Login Doctor
// @Tags LOGIN
// @Accept json
// @Produce json
// @Param login body models.DoctorLoginReq true "DoctorLogin"
// @Success 200 {object} http.Response{data=models.DoctorLoginRes} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) LoginDoctor(c *gin.Context) {
	var body models.DoctorLoginReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
	}
	isPhone, err := h.Storage.Clinic().CheckField(&models.CheckfieldReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})
	if !isPhone.Exists {
		h.handleResponse(c, http.NotFound, models.ResponseMessage{Message: "The phone number was not found."})
		return
	}
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetDoctorByPhoneNumber(body.PhoneNumber)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.JwtHandler.Sub = res.Id
	h.JwtHandler.Aud = []string{"backend"}
	h.JwtHandler.SigninKey = h.Cfg.SigningKey
	h.JwtHandler.Log = h.Log
	h.JwtHandler.Role = "doctor"
	tokens, err := h.JwtHandler.GenerateAuthJWT()

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	res.AccessToken = tokens[0]

	h.handleResponse(c, http.OK, res)
}

// Create Doctor
// @Security BearerAuth
// @ID createdoctor
// @Router /doctor [POST]
// @Summary Create Doctor
// @Description Create Doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param body body models.DoctorReqForUI true "CreateAdminInfo"
// @Success 201 {object} http.Response{data=models.DoctorResponse} "Admin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) RegisterDoctor(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	var body models.DoctorReqForUI
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	isPhone, err := h.Storage.Clinic().CheckField(&models.CheckfieldReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	if isPhone.Exists {
		h.handleResponse(c, http.Conflict, models.ResponseMessage{
			Message: "Phone number already exists",
		})
		return
	}

	h.JwtHandler.Aud = []string{"backend"}
	h.JwtHandler.Role = "doctor"
	h.JwtHandler.SigninKey = h.Cfg.SigningKey
	h.JwtHandler.Log = h.Log
	tokens, err := h.JwtHandler.GenerateAuthJWT()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	var info = models.DoctorRequest{
		ClinicId:     body.ClinicId,
		ServiceId:    body.ServiceId,
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		PhoneNumber:  body.PhoneNumber,
		StartTime:    body.StartTime,
		EndTime:      body.EndTime,
		WorkDay:      body.WorkDay,
		Floor:        body.Floor,
		RoomNumber:   body.RoomNumber,
		ImageLink:    body.ImageLink,
		Experience:   body.Experience,
		RefreshToken: tokens[1],
	}
	res, err := h.Storage.Clinic().CreateDoctor(&info)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Get Doctor
// @ID getdoctorbyid
// @Router /doctor/{id} [GET]
// @Summary Get Doctor
// @Description Get Doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param id path int true "doctor_id"
// @Success 200 {object} http.Response{data=models.DoctorResponse} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetDoctor(c *gin.Context) {
	id := c.Param("id")
	doctorId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Storage.Clinic().GetDoctor(doctorId)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	if data != nil {
		h.handleResponse(c, http.OK, data)
	} else {
		h.handleResponse(c, http.NotFound, models.ResponseMessage{
			"info not found",
		})
	}
}

// Get Doctor
// @Security BearerAuth
// @ID getcustomersbydoctorid
// @Router /doctor/customers/{doctor_id} [GET]
// @Summary Get Doctor's customers
// @DescriptionGet Doctor's customers
// @Tags Doctor
// @Accept json
// @Produce json
// @Param doctor_id path int true "doctor_id"
// @Success 200 {object} http.Response{data=models.GetCustomersOfDoctor} "customer data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetCustomersByDoctor(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	id := c.Param("doctor_id")
	doctorID, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetCustomersByDoctorID(doctorID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)

}

// Get All Doctor
// @ID getalldoctor
// @Router /doctors/search/{clinic_id} [GET]
// @Summary Get all Doctor
// @Description Get all  Doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param clinic_id path int true "ClinicID"
// @Param keyword query string true "keyword"
// @Success 200 {object} http.Response{data=[]models.DoctorsList} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetAllDoctor(c *gin.Context) {
	id := c.Param("clinic_id")
	clinicID, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	keyword := c.Query("keyword")
	res, err := h.Storage.Clinic().GetDoctorsSearch(clinicID, keyword)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// Get List Doctor
// @ID getlistdoctor
// @Router /doctors/{clinic_id} [GET]
// @Summary Get all Doctor
// @Description Get all  Doctors List
// @Tags Doctor
// @Accept json
// @Produce json
// @Param clinic_id path int true "ClinicID"
// @Success 200 {object} http.Response{data=[]models.DoctorsList} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetDoctorsList(c *gin.Context) {
	id := c.Param("clinic_id")
	clinicID, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetDoctorsByClinicId(clinicID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// PUT Doctors
// @Security BearerAuth
// @ID putdoctors
// @Router /doctors [PUT]
// @Summary PUT Doctors
// @Description PUT Doctors
// @Tags Doctor
// @Accept json
// @Produce json
// @Param body body models.UpdateDoctor true "UpdateClient"
// @Success 200 {object} http.Response{data=models.ResponseMessage} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateDoctor(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	body := models.UpdateDoctor{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	err = h.Storage.Clinic().UpdateDoctor(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{Message: "Successfully updated"})
}

// DELETE Doctor
// @Security BearerAuth
// @ID deletedoctor
// @Router /doctors/delete/{id} [DELETE]
// @Summary Delete Doctors
// @Description DELETE Doctors
// @Tags Doctor
// @Accept json
// @Produce json
// @Param id path int true "doctorId"
// @Success 200 {object} http.Response{data=models.ResponseMessage} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteDoctor(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	id := c.Param("id")
	doctorId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	err = h.Storage.Clinic().DeleteDoctor(doctorId)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{Message: "successfully deleted"})
}

// Get Doctors ByService
// @ID getdoctorsbyservicename
// @Router /doctor/service/{clinic_id} [GET]
// @Summary Get Doctor by ServiceName
// @Description Get Doctor by ServiceName
// @Tags Doctor
// @Accept json
// @Produce json
// @Param clinic_id path int true "clinicID"
// @Param servicename query string true "ServiceName"
// @Success 200 {object} http.Response{data=models.Doctors} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetDoctorsByService(c *gin.Context) {
	id := c.Param("clinic_id")
	clinicId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	serviceName := c.Query("servicename")
	res, err := h.Storage.Clinic().GetDoctorsByService(clinicId, serviceName)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// Get Doctors ByServiceID
// @ID getdoctorsbyserviceid
// @Router /doctors/service/{service_id} [GET]
// @Summary Get Doctor by ServiceID
// @Description Get Doctor by ServiceID
// @Tags Doctor
// @Accept json
// @Produce json
// @Param service_id path int true "serviceID"
// @Success 200 {object} http.Response{data=models.DoctorsList} "Doctor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetDoctorsByServiceId(c *gin.Context) {
	id := c.Param("service_id")
	serviceID, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetDoctorsByServiceID(serviceID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

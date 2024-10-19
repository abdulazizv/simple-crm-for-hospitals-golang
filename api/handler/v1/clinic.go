package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
	"strconv"
)

// Create Clinic
// @Security BearerAuth
// @ID createclinic
// @Router /clinic [POST]
// @Summary Create Clinic
// @Description Create Clinic
// @Tags Clinic
// @Accept json
// @Produce json
// @Param body body models.ClinicReq true "CreateClinicInfo"
// @Success 201 {object} http.Response{data=models.ClinicRes} "Clinic data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateClinic(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	var body models.ClinicReq

	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().CreateClinic(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// GET Clinic
// @ID getclinic
// @Router /clinic/{id} [GET]
// @Summary GET Clinic
// @Description GET Clinic
// @Tags Clinic
// @Accept json
// @Produce json
// @Param id path int true "GetClinicInfo"
// @Success 201 {object} http.Response{data=models.ClinicRes} "Clinic data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetClinic(c *gin.Context) {
	id := c.Param("id")
	clinicID, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	info, err := h.Storage.Clinic().GetClinic(clinicID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	if info != nil {
		h.handleResponse(c, http.OK, info)
	} else {
		h.handleResponse(c, http.NotFound, models.ResponseMessage{
			Message: "info not found",
		})
	}
}

// GET Clinics
// @ID getclinics
// @Router /clinics [GET]
// @Summary GET Clinics
// @Description GET Clinics
// @Tags Clinic
// @Accept json
// @Produce json
// @Success 201 {object} http.Response{data=[]models.ClinicRes} "Clinics data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetClinicsList(c *gin.Context) {
	res, err := h.Storage.Clinic().GetList()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// PUT Clinics
// @Security BearerAuth
// @ID putclinic
// @Router /clinics [PUT]
// @Summary PUT Clinics
// @Description PUT Clinics
// @Tags Clinic
// @Accept json
// @Produce json
// @Param body body models.UpdateClinicReq true "UpdateClinics"
// @Success 200 {object} http.Response{data=models.ClinicRes} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateClinics(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	body := models.UpdateClinicReq{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Storage.Clinic().UpdateClinics(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

// DELETE Clinics
// @ID deleteclinics
// @Security BearerAuth
// @Router /clinics/delete/{id} [DELETE]
// @Summary PUT Clinics
// @Description PUT Clinics
// @Tags Clinic
// @Accept json
// @Produce json
// @Param id path int true "clinicID"
// @Success 200 {object} http.Response{data=models.ResponseMessage} "Clinic data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteClinic(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	id := c.Param("id")
	clientID, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	err = h.Storage.Clinic().DeleteClinics(clientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{Message: "successfully deleted"})
}

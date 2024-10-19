package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
	"strconv"
)

// Create Services
// @Security BearerAuth
// @ID createservice
// @Router /service [POST]
// @Summary Create Service
// @Description Create Service
// @Tags Service
// @Accept json
// @Produce json
// @Param body body models.ServicesRequest true "CreateServiceInfo"
// @Success 201 {object} http.Response{data=models.ServicesResponse} "Service data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateServices(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	var body models.ServicesRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().CreateServices(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Get Services
// @ID getservicebyid
// @Router /service/{id} [GET]
// @Summary Get Service
// @Description Get Service
// @Tags Service
// @Accept json
// @Produce json
// @Param id path int true "service_id"
// @Success 200 {object} http.Response{data=models.ServicesResponse} "Service data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetServices(c *gin.Context) {
	id := c.Param("id")
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Storage.Clinic().GetService(serviceId)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

// Get All Services
// @ID getservices
// @Router /services [GET]
// @Summary Get all Service
// @Description Get all Service
// @Tags Service
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=[]models.ServicesResponse} "Service data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetAllServices(c *gin.Context) {
	res, err := h.Storage.Clinic().GetServicesList()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// PUT Services
// @Security BearerAuth
// @ID putservices
// @Router /service [PUT]
// @Summary PUT Services
// @Description PUT Services
// @Tags Service
// @Accept json
// @Produce json
// @Param body body models.UpdateServicesReq true "UpdateServices"
// @Success 200 {object} http.Response{data=models.ServicesResponse} "Service data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateServices(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	body := models.UpdateServicesReq{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Storage.Clinic().UpdateServices(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

// DELETE Services
// @Security BearerAuth
// @ID service
// @Router /service/delete/{id} [DELETE]
// @Summary PUT Service
// @Description PUT Service
// @Tags Service
// @Accept json
// @Produce json
// @Param id path int true "serviceID"
// @Success 200 {object} http.Response{data=models.ResponseMessage} "Service data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteServices(c *gin.Context) {
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
	err = h.Storage.Clinic().DeleteService(clientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{Message: "successfully deleted"})
}

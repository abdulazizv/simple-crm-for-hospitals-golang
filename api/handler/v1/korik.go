package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
	"strconv"
)

// Create Korik
// @Security BearerAuth
// @ID createkorik
// @Router /korik [POST]
// @Summary Create Korik
// @Description Create Korik
// @Tags Korik
// @Accept json
// @Produce json
// @Param body body models.KorikRequest true "CreateKorikInfo"
// @Success 201 {object} http.Response{data=models.KorikResponse} "Korik data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateKorik(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	var body models.KorikRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().CreateKorik(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Get Koriks
// @ID getkorikbyid
// @Security BearerAuth
// @Router /korik/{id} [GET]
// @Summary Get Korik
// @Description Get Korik
// @Tags Korik
// @Accept json
// @Produce json
// @Param id path int true "korik_id"
// @Success 200 {object} http.Response{data=models.KorikResponse} "Korik data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetKorik(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}

	id := c.Param("id")
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Storage.Clinic().GetKorik(serviceId)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

// Get All Korik
// @ID getkorik
// @Security BearerAuth
// @Router /korik [GET]
// @Summary Get all Korik
// @Description Get all Korik
// @Tags Korik
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=[]models.KoriksList} "Service data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetAllKoriks(c *gin.Context) {
	_, err := GetClaims(*h, c)

	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetKoriks()

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

// PUT Korik
// @Security BearerAuth
// @ID putkorik
// @Router /korik [PUT]
// @Summary PUT Korik
// @Description PUT Korik
// @Tags Korik
// @Accept json
// @Produce json
// @Param body body models.UpdateKorikRequest true "UpdateKorik"
// @Success 200 {object} http.Response{data=models.KorikResponse} "Korik data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateKorik(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	body := models.UpdateKorikRequest{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Storage.Clinic().UpdateKorik(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

// DELETE Korik
// @ID korik
// @Security BearerAuth
// @Router /korik/delete/{id} [DELETE]
// @Summary PUT Korik
// @Description PUT Korik
// @Tags Korik
// @Accept json
// @Produce json
// @Param id path int true "korikID"
// @Success 200 {object} http.Response{data=models.ResponseMessage} "Korik data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteKorik(c *gin.Context) {
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
	err = h.Storage.Clinic().DeleteKorik(clientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{Message: "successfully deleted"})
}

// Get Koriks
// @ID getkorikbyuserid
// @Security BearerAuth
// @Router /korik/user/{id} [GET]
// @Summary Get Korik
// @Description Get Korik
// @Tags Korik
// @Accept json
// @Produce json
// @Param id path int true "user_id"
// @Success 200 {object} http.Response{data=models.KorikResponse} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetKorikByUserId(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}

	id := c.Param("id")
	serviceId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	data, err := h.Storage.Clinic().GetKorikByUserId(serviceId)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, data)
}

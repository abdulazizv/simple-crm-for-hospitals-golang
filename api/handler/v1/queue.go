package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
)

// Client Queue
// @ID createqueue
// @Security BearerAuth
// @Router /queue [POST]
// @Summary Create Queue Request
// @Description Create Queue Request
// @Tags Queue
// @Accept json
// @Produce json
// @Param body body models.QueueReq true "QueueCreate"
// @Success 201 {object} http.Response{data=models.QueueRes} "Queue data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateQueue(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	fmt.Println("Hello World")
	var body models.QueueReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().CreateQueue(body.DoctorID, body.ClientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Cancel Queue
// @ID cancelqueue
// @Security BearerAuth
// @Router /queue/cancel [DELETE]
// @Summary Cancel Queue Request
// @Description Cancel Queue Request
// @Tags Queue
// @Accept json
// @Produce json
// @Param body body models.QueueReq true "QueueCreate"
// @Success 201 {object} http.Response{data=models.ResponseMessage} "Cancel data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CancelQueue(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	var body models.QueueReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	err = h.Storage.Clinic().CancelQueue(body.DoctorID, body.ClientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{
		Message: "queue successfully canceled",
	})
}

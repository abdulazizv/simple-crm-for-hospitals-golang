package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
	"strconv"
)

// Client Login
// @ID loginclient
// @Router /client/login [POST]
// @Summary Login Client
// @Description Login Client
// @Tags LOGIN
// @Accept json
// @Produce json
// @Param login body models.ClientLoginReq true "ClientLogin"
// @Success 200 {object} http.Response{data=models.ClientsResponse} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) LoginClient(c *gin.Context) {
	var body models.ClientLoginReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	isPhone, err := h.Storage.Clinic().CheckFieldClient(&models.CheckfieldReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	if !isPhone.Exists {
		h.handleResponse(c, http.NotFound, models.ResponseMessage{
			Message: "phone number was not found",
		})
		return
	}
	isName, err := h.Storage.Clinic().CheckFieldClient(&models.CheckfieldReq{
		Field: "first_name",
		Value: body.FirstName,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	if !isName.Exists {
		h.handleResponse(c, http.NotFound, models.ResponseMessage{
			Message: "first_name was not found",
		})
		return
	}
	res, err := h.Storage.Clinic().GetClientForLogin(body.FirstName, body.PhoneNumber)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.JwtHandler.Sub = res.Id
	h.JwtHandler.Aud = []string{"clinic"}
	h.JwtHandler.SigninKey = h.Cfg.SigningKey
	h.JwtHandler.Log = h.Log
	h.JwtHandler.Role = "client"
	tokens, err := h.JwtHandler.GenerateAuthJWT()
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res.AccessToken = tokens[0]
	h.handleResponse(c, http.OK, res)
}

// Client Create
// @ID createclient
// @Router /client [POST]
// @Summary Create Client
// @Description Create Client
// @Tags Client
// @Accept json
// @Produce json
// @Param login body models.ClientsRequest true "ClientCreate"
// @Success 201 {object} http.Response{data=models.ClientsResponse} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) CreateClient(c *gin.Context) {
	var body models.ClientsRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	isPhone, err := h.Storage.Clinic().CheckFieldClient(&models.CheckfieldReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	if isPhone.Exists {
		h.handleResponse(c, http.Conflict, models.ResponseMessage{
			Message: "phone number already exists",
		})
		return
	}
	h.JwtHandler.Aud = []string{"client"}
	h.JwtHandler.SigninKey = h.Cfg.SigningKey
	h.JwtHandler.Log = h.Log
	h.JwtHandler.Role = "client"
	tokens, err := h.JwtHandler.GenerateAuthJWT()
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	info := models.ClientsReq{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Age:          body.Age,
		PhoneNumber:  body.PhoneNumber,
		RefreshToken: tokens[1],
	}
	res, err := h.Storage.Clinic().CreateClient(&info)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	res.AccessToken = tokens[0]
	h.handleResponse(c, http.Created, res)
}

// Client Get
// @Security BearerAuth
// @ID getclient
// @Router /client/{id} [GET]
// @Summary GET Client
// @Description GET Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path int true "ClientID"
// @Success 201 {object} http.Response{data=models.ClientsResponse} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetClient(c *gin.Context) {
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
	res, err := h.Storage.Clinic().GetClient(clientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	if res != nil {
		h.handleResponse(c, http.OK, res)
	} else {
		h.handleResponse(c, http.NotFound, models.ResponseMessage{
			Message: "info not found",
		})
	}
}

// Get Clients
// @ID getclients
// @Router /clients [GET]
// @Summary GET Clients
// @Description GET Clients
// @Tags Client
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=models.ClientsList} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetClients(c *gin.Context) {
	data, err := h.Storage.Clinic().GetClients()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}
	h.handleResponse(c, http.OK, data)
}

// PUT Clients
// @Security BearerAuth
// @ID putclients
// @Router /client [PUT]
// @Summary PUT Clients
// @Description PUT Clients
// @Tags Client
// @Accept json
// @Produce json
// @Param body body models.ClientUpdateReq true "UpdateClient"
// @Success 200 {object} http.Response{data=models.ClientsList} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) UpdateClient(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	body := models.ClientUpdateReq{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Storage.Clinic().UpdateClient(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

// DELETE Clients
// @Security BearerAuth
// @ID deleteclients
// @Router /client/delete/{id} [DELETE]
// @Summary PUT Clients
// @Description PUT Clients
// @Tags Client
// @Accept json
// @Produce json
// @Param id path int true "clientID"
// @Success 200 {object} http.Response{data=models.ResponseMessage} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) DeleteClient(c *gin.Context) {
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
	err = h.Storage.Clinic().DeleteClient(clientID)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, models.ResponseMessage{Message: "successfully deleted"})
}

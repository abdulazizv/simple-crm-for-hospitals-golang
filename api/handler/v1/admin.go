package v1

import (
	"gitlab.com/backend/pkg/etc"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
)

// Admin Login
// @ID loginadmin
// @Router /admin/login [POST]
// @Summary Login Admin
// @Description Login Admin
// @Tags LOGIN
// @Accept json
// @Produce json
// @Param login body models.AdminLoginReq true "AdminLogin"
// @Success 200 {object} http.Response{data=models.AdminRes} "Admin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) LoginAdmin(c *gin.Context) {
	var body models.AdminReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
	}
	info, err := h.Storage.Clinic().GetAdminByUsername(body.UserName)
	if !etc.CheckPasswordHash(body.Password, info.Password) {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetAdminForLogin()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.JwtHandler.Sub = res.Id
	h.JwtHandler.Aud = []string{"backend"}
	h.JwtHandler.SigninKey = h.Cfg.SigningKey
	h.JwtHandler.Log = h.Log
	h.JwtHandler.Role = "admin"
	tokens, err := h.JwtHandler.GenerateAuthJWT()

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	res.AccessToken = tokens[0]

	h.handleResponse(c, http.OK, res)

}

// Create Admin
// @Security BearerAuth
// @ID createadmin
// @Router /admin [POST]
// @Summary Create Admin
// @Description Create Admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param body body models.AdminReq true "CreateAdminInfo"
// @Success 201 {object} http.Response{data=models.AdminRes} "Admin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) RegisterAdmin(c *gin.Context) {
	//_, err := GetClaims(*h, c)
	//if err != nil {
	//	h.handleResponse(c, http.Unauthorized, err.Error())
	//	return
	//}
	var body models.AdminReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	body.Password, err = etc.HashPassword(body.Password)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	h.JwtHandler.Aud = []string{"backend"}
	h.JwtHandler.Role = "admin"
	h.JwtHandler.SigninKey = h.Cfg.SigningKey
	h.JwtHandler.Log = h.Log
	tokens, err := h.JwtHandler.GenerateAuthJWT()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	body.RefreshToken = tokens[1]
	res, err := h.Storage.Clinic().CreateAdmin(&body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, res)
}

// Get Admin
// @Security BearerAuth
// @ID getadminbyid
// @Router /admin/{id} [GET]
// @Summary Get Admin
// @Description Get Admin
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path int true "admin_id"
// @Success 200 {object} http.Response{data=models.AdminRes} "Admin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetAdmin(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	id := c.Param("id")
	adminId, err := strconv.Atoi(id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	data, err := h.Storage.Clinic().GetAdmin(adminId)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)
}

// Get All Admin
// @Security BearerAuth
// @ID getadmin
// @Router /admins [GET]
// @Summary Get all Admin
// @Description Get all  Admin
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=[]models.AdminRes} "Admin data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) GetAllAdmin(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		h.handleResponse(c, http.Unauthorized, err.Error())
		return
	}
	res, err := h.Storage.Clinic().GetAdminList()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, res)
}

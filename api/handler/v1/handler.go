package v1

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gitlab.com/backend/api/handler/v1/http"
	t "gitlab.com/backend/api/tokens"
	"gitlab.com/backend/config"
	"gitlab.com/backend/pkg/logger"
	"gitlab.com/backend/storage"
)

type handlerV1 struct {
	Cfg        *config.Config
	Storage    storage.StorageI
	Log        logger.Logger
	JwtHandler t.JWTHandler
}

type HandlerV1Option struct {
	Cfg        *config.Config
	Storage    storage.StorageI
	Log        logger.Logger
	JwtHandler t.JWTHandler
}

func New(optoins *HandlerV1Option) *handlerV1 {
	return &handlerV1{
		Cfg:        optoins.Cfg,
		Storage:    optoins.Storage,
		Log:        optoins.Log,
		JwtHandler: optoins.JwtHandler,
	}
}

func (h *handlerV1) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.Log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *handlerV1) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.Cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *handlerV1) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.Cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func GetClaims(h handlerV1, c *gin.Context) (*t.CustomClaims, error) {
	var (
		claims = t.CustomClaims{}
	)
	strToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) { return []byte(h.Cfg.SigningKey), nil })

	if err != nil {
		h.Log.Error("invalid access token")
		return nil, err
	}
	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Exp = rawClaims["exp"].(float64)
	aud := cast.ToStringSlice(rawClaims["aud"])
	claims.Aud = aud
	claims.Role = rawClaims["role"].(string)
	claims.Token = token
	return &claims, nil
}

package v1

import (
	"context"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/backend/api/handler/v1/http"
	"gitlab.com/backend/api/models"
)

// Engineer File Upload
// @ID file-upload
// @Router /media/file-upload [POST]
// @Summary File Upload
// @Description File Upload
// @Tags Media
// @Accept json
// @Produce json
// @Param file 	formData file true "image-upload"
// @Success 201 {object} http.Response{data=models.FileResponse} "File-upload"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *handlerV1) FileUpload(c *gin.Context) {
	file := models.File{}
	err := c.ShouldBind(&file)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	ext := filepath.Ext(file.File.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".pdf" && ext != ".psd" {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	mediaType := "image"

	fileName := uuid.New().String() + filepath.Ext(file.File.Filename)

	f, err := file.File.Open()

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	// Aws started
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	res, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(h.Cfg.ClinicBucket),
		Key:    aws.String(fileName),
		ACL:    "public-read",
		Body:   f,
	})

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	data := models.FileResponse{
		Url:       res.Location,
		MediaType: mediaType,
	}
	h.handleResponse(c, http.Created, data)

}

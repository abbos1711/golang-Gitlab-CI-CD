package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	l "gitlab.com/tizim-back/pkg/logger"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// File upload
// @Security ApiKeyAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Router /v1/file-upload [post]
// @Success 200 {object} string
func (h *handlerV1) UploadFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		h.log.Error("Error while uploading file", l.Any("post", err))
		return
	}
	if filepath.Ext(file.File.Filename) != ".png" && filepath.Ext(file.File.Filename) != ".jpg" && filepath.Ext(file.File.Filename) != ".jpeg" {
		fmt.Println(filepath.Ext(file.File.Filename) != ".png" || filepath.Ext(file.File.Filename) != ".jpg")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couln't find matching file format",
		})
		h.log.Error("Error while getting uploading img", l.Any("file-upload", err))
		return
	}
	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	if _, err := os.Stat(dst + "/media"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media", os.ModePerm)
	}

	filePath := "/media/" + fileName
	err = c.SaveUploadedFile(file.File, dst+filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Couln't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting customer by email", l.Any("post", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": c.Request.Host + filePath,
	})
}

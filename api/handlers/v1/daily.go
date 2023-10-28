package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "gitlab.com/tizim-back/api/models"
	"gitlab.com/tizim-back/pkg/logger"

	//	logger "gitlab.com/tizim-back/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create workers attendace
// @Summary      Create attendance
// @Description  Create attendence
// @Tags         Daily
// Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        req   	body 	  models.DailyReq true  "request"
// @Success      200  	{object}  models.DailyRes
// @Router 		/v1/daily [post]
func (h *handlerV1) CreateAttendance(c *gin.Context) {
	var (
		body        models.DailyReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json: ", logger.Error(err))
		return
	}
	request := models.DailyReq{	
		Id: body.Id,
	}

	response, err := h.Storage.Daily().CreateAttendance(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create attandance", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get daily attendance portion 
// @Summary      Get attendance portion
// @Description  Get attendance portion
// @Tags         Daily
// Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200  	{object}  models.AttendancePortion
// @Router 		/v1/daily/portion [get]
func (h *handlerV1) GetAttendancePortion(c *gin.Context) {

	response, err := h.Storage.Daily().GetAttendancePortion()
	fmt.Println("Responce ----->", response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get daily attendance portion", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
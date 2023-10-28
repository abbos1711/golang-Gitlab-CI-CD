package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "gitlab.com/tizim-back/api/models"
	"gitlab.com/tizim-back/pkg/logger"

	//	logger "gitlab.com/tizim-back/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Worker
// @Summary      Create worker
// @Description  Create worker
// @Tags         Worker
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        worker   body 	  models.WorkerCreate true  "worker"
// @Success      200  	{object}  models.WorkerResp
// @Router 		/v1/worker [post]
func (h *handlerV1) CreateWorker(c *gin.Context) {
	var (
		body        models.WorkerCreate
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
	request := models.WorkerCreate{
		Img_url:    body.Img_url,
		Name:       body.Name,
		Surname:    body.Surname,
		Position:   body.Position,
		Department: body.Department,
		Gender:     body.Gender,
		Contact:    body.Contact,
		Birthday:   body.Birthday,
		ComeTime:   body.ComeTime,
	}

	response, err := h.Storage.Worker().CreateWorker(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create worker", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Worker
// @Summary 	Delete worker
// @Description This finction deletes worker by id
// @Tags 		Worker
// @Security 	ApiKeyAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "id"
// @Success 	200 {object} models.ResponseOk
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/worker/{id} [delete]
func (h *handlerV1) DeleteWorker(c *gin.Context) {
	id := c.Param("id")

	err := h.Storage.Worker().DeleteWorker(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err,
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOk{
		Message: "Deleted Successfully",
	})
}

// Update Worker
// @Summary 	Update worker
// @Description This finction updates worker by id
// @Tags		Worker
// @Security 	ApiKeyAuth
// @Accept 		json
// @Produce 	json
// @Param 		worker body models.WorkerUpdate true "worker"
// @Success 	200 {object} models.ResponseOk
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/worker/update [put]
func (h *handlerV1) UpdateWorker(c *gin.Context) {

	var (
		body        models.WorkerUpdate
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: err})
		return
	}

	response, err := h.Storage.Worker().UpdateWorker(&models.WorkerUpdate{
		Id:         body.Id,
		Img_url:    body.Img_url,
		Name:       body.Name,
		Surname:    body.Surname,
		Position:   body.Position,
		Department: body.Department,
		Gender:     body.Gender,
		Contact:    body.Contact,
		Birthday:   body.Birthday,
		ComeTime:   body.ComeTime,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GEt All Workers
// @Summary      Get all workers
// @Description  This function gets all workerers
// @Tags         Worker
// Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200  	{object}  models.AllWorkers
// @Router 		/v1/workers [get]
func (h *handlerV1) GetAllWorkers(c *gin.Context) {

	response, err := h.Storage.Worker().GetAllWorkers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all workers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Worker
// @Summary 	Get worker
// @Description This finction gets worker by id
// @Tags		Worker
// Security 	ApiKeyAuth
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "id"
// @Success 	200 {object} models.WorkerResp
// @Router 		/v1/worker/{id} [get]
func (h *handlerV1) GetWorker(c *gin.Context) {

	id := c.Param("id")

	response, err := h.Storage.Worker().GetWorker(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Workers by Gender
// @Summary      Get workers by gender
// @Description  This function gets all workerers
// @Tags         Worker
// Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param 		 gender path string true "gender"
// @Success      200  	{object}  models.AllWorkersFilter
// @Router 		/v1/workers/{gender} [get]
func (h *handlerV1) GetWorkersByGender(c *gin.Context) {

	gender := c.Param("gender")

	response, err := h.Storage.Worker().GetWorkersByGender(gender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to get workers by gender", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Workers At Work
// @Summary      Get workers at work
// @Description  This function gets all workerers
// @Tags         Worker
// Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200  	{object}  models.AllWorkersFilter
// @Router 		/v1/workers/at-work [get]
func (h *handlerV1) GetWorkersAtWork(c *gin.Context) {

	response, err := h.Storage.Worker().GetWorkersAtWork()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all workers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Top Workers 
// @Summary      Get top workers 
// @Description  This function gets top workerers
// @Tags         Worker
// Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200  	{object}  []models.TopWorkers
// @Router 		/v1/workers-top [get]
func (h *handlerV1) GetTopWorkers(c *gin.Context) {

	response, err := h.Storage.Worker().GetTopWorkers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get best workers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}


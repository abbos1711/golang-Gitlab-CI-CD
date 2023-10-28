package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/tizim-back/pkg/logger"
)


// @Summary      Get all workers by date
// @Description  This function gets all workers by date
// Security 	 ApiKeyAuth
// @Tags 		 Worker-History
// @Accept       json
// @Produce      json
// @Param 		 date path string true "date"
// @Success      200  	{object}  models.WorkersByMonthResp
// @Router 		/v1/get-workers-by-month/{date} [get]
func (h *handlerV1) GetWorkersByMonth(c *gin.Context) {
	date := c.Param("date")

	response, err := h.Storage.WorkerHistory().GetWorkersByMonth(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all workers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary      Get all workers by two date
// @Description  This function gets all workers by two date
// Security 	 ApiKeyAuth
// @Tags 		 Worker-History
// @Accept       json
// @Produce      json
// @Param 		 date1 path string true "date1"
// @Param 		 date2 path string true "date2"
// @Success      200  	{object}  models.WorkersByTwoDateResp
// @Router 		/v1/get-workers-by-two-date/{date} [get]
func (h *handlerV1) GetWorkersByTwoDate(c *gin.Context) {
	date1 := c.Param("date1")
	date2 := c.Param("date2")

	response, err := h.Storage.WorkerHistory().GetWorkersByTwoDate(date1, date2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all workers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}


// @Summary      Get all workers by day
// @Description  This function gets all workers by day
// Security 	 ApiKeyAuth
// @Tags 		 Worker-History
// @Accept       json
// @Produce      json
// @Param 		 date   path string true "date"
// @Success      200  	{object}  models.WorkersByDayResp
// @Router 		/v1/get-workers-by-day/{date} [get]
func (h *handlerV1) GetWorkersByDay(c *gin.Context) {
	date := c.Param("date")

	response, err := h.Storage.WorkerHistory().GetWorkersByDay(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get all workers", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

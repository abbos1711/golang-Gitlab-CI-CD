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
func (h *handlerV1) GetAllWorkersByMonth(c *gin.Context) {
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


// // @Summary      Get all workers by month
// // @Description  This function gets all workers by month
// // Security 	 ApiKeyAuth
// // @Tags 		 Worker-History
// // @Accept       json
// // @Produce      json
// // @Param 		 date   path string true "date"
// // @Success      200  	{object}  models.WorkersByMonth
// // @Router 		/v1/get-workers-by-month/{date} [get]
// func (h *handlerV1) GetAllWorkersByMonth(c *gin.Context) {
// 	date := c.Param("date")

// 	response, err := h.Storage.WorkerHistory().GetWorkersByMonth(date)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to get all workers", logger.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

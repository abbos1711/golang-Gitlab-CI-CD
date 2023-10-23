package repo

import (
	"gitlab.com/tizim-back/api/models"
)

type WorkerHistoryStorageI interface {
	GetWorkersByMonth(date string) (*models.WorkersByMonthResp, error)
	//GetWorkersByMonth(date string) (*models.WorkersByMonth, error)
}//

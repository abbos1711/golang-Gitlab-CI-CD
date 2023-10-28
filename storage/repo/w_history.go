package repo

import (
	"gitlab.com/tizim-back/api/models"
)

type WorkerHistoryStorageI interface {
	GetWorkersByMonth(date string) (*models.WorkersByMonthResp, error)
	GetWorkersByTwoDate(date1, date2 string) (*models.WorkersByTwoDateResp, error)
	GetWorkersByDay(date string) (*models.WorkersByDayResp, error)

}

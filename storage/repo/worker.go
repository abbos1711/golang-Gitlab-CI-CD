package repo

import (
	"gitlab.com/tizim-back/api/models"
)

type WorkerStorageI interface {
	CreateWorker(*models.WorkerCreate) (*models.WorkerResp, error)
	DeleteWorker(id string) error
	UpdateWorker(*models.WorkerUpdate) (*models.WorkerResp, error)
	GetAllWorkers() (*models.AllWorkers, error)
	GetWorker(id string) (*models.WorkerResp, error)
	GetWorkersByGender (gender string)(*models.AllWorkersFilter, error)
	GetWorkersAtWork ()(*models.AllWorkersFilter, error)
	GetTopWorkers()(*models.TopWorkers, error)
}

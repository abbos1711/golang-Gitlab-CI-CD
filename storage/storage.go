package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"gitlab.com/tizim-back/storage/postgres"
	"gitlab.com/tizim-back/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
	Worker() repo.WorkerStorageI
	WorkerHistory() repo.WorkerHistoryStorageI
	Daily() repo.DailyStorageI
}

type storagePg struct {
	db                *pgxpool.Pool
	userRepo          repo.UserStorageI
	workerRepo        repo.WorkerStorageI
	workerHistoryRepo repo.WorkerHistoryStorageI
	dailyRepo         repo.DailyStorageI
}

func NewStoragePg(db *pgxpool.Pool) *storagePg {
	return &storagePg{
		db:                db,
		userRepo:          postgres.NewUserRepo(db),
		workerRepo:        postgres.NewWorkerRepo(db),
		workerHistoryRepo: postgres.NewWorkerHistoryRepo(db),
		dailyRepo:         postgres.NewDailyRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s storagePg) Worker() repo.WorkerStorageI {
	return s.workerRepo
}

func (s storagePg) WorkerHistory() repo.WorkerHistoryStorageI {
	return s.workerHistoryRepo
}

func (s storagePg) Daily() repo.DailyStorageI {
	return s.dailyRepo
}

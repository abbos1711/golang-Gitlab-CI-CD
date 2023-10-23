package repo

import (
	"gitlab.com/tizim-back/api/models"
)

type DailyStorageI interface {
	CreateAttendance(*models.DailyReq) (*models.DailyRes, error)
	GetAttendancePortion()(*models.AttendancePortion, error)
}

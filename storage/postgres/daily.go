package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/tizim-back/api/models"
)

type dailyRepo struct {
	db *pgxpool.Pool
}

func NewDailyRepo(db *pgxpool.Pool) *dailyRepo {
	return &dailyRepo{
		db: db,
	}
}

func (r *dailyRepo) CreateAttendance(req *models.DailyReq) (*models.DailyRes, error) {

	var res models.DailyRes
	var comeTime, date time.Time
	var leaveTime, workDuration sql.NullTime

	query := `
		INSERT INTO daily 
			(worker_id,
			come_time,
			status
		)SELECT
			$1,
			$2,
			true
		WHERE EXISTS (
			SELECT 
				1
			FROM 
				workers
			WHERE 
				workers.id = $1
			AND
				workers.deleted_at IS NULL
		)RETURNING 
			w_date,
			worker_id,
			come_time,
			leave_time,
			w_hours,
			status
	`
	err := r.db.QueryRow(context.Background(), query,
		req.Id,
		req.Time,
	).Scan(
		&date,
		&res.WorkerId,
		&comeTime,
		&leaveTime,
		&workDuration,
		&res.Status,
	)

	if err != nil {
		log.Println("Error in creting workers: ", err)
		return &models.DailyRes{}, err
	}

	if leaveTime.Valid {
		res.LeaveTime = leaveTime.Time.Format(time.TimeOnly)
	} else {
		res.LeaveTime = ""
	}

	if leaveTime.Valid {
		res.WorkDuration = workDuration.Time.Format(time.DateOnly)
	} else {
		res.LeaveTime = ""
	}

	res.Date = date.Format(time.DateOnly)
	res.ComeTime = comeTime.Format(time.TimeOnly)

	return &res, nil
}

func (r dailyRepo) GetAttendancePortion() (*models.AttendancePortion, error) {
	var res models.AttendancePortion

	query := `
		SELECT 
			(COUNT(*) * 100.0) / (SELECT COUNT(*) FROM workers WHERE deleted_at IS NULL) AS portion
		FROM 
			daily
		WHERE 
			w_date = CURRENT_DATE;

	`
	err := r.db.QueryRow(context.Background(), query).Scan(&res.Portion)

	if err != nil {
		log.Println("Error while getting attandance portion: ", err)
		return &models.AttendancePortion{}, err
	}

	return &res, nil
}

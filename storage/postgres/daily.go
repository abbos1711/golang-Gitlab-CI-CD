package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
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

func durationCounter(from, untill time.Time) int {

	hour := from.Hour()
	minute := from.Minute()
	late_min := (hour*60 + minute) - (untill.Hour()*60 + untill.Minute())
	return late_min
}

func (r *dailyRepo) CreateAttendance(req *models.DailyReq) (*models.DailyRes, error) {

	var res models.DailyRes
	var check sql.NullBool
	var comeTime, date, wComeTime time.Time

	err := r.db.QueryRow(context.Background(), "SELECT status FROM daily WHERE worker_id = $1 AND w_date = CURRENT_DATE", req.Id).Scan(&check)
	if err == pgx.ErrNoRows { 
		// Worker is entering for the first time today

		// getting worker's come_time
		err := r.db.QueryRow(context.Background(), `SELECT come_time FROM workers WHERE id = $1`, req.Id).Scan(&wComeTime)
		if err != nil {
			log.Println("Error in getting come_time of daily: ", err)
			return &models.DailyRes{}, err
		}

		late_min := durationCounter(time.Now(), wComeTime)

		query := `
			INSERT INTO daily(
				worker_id,
				come_time,
				status,
				late_min 
			)SELECT
				$1,
				CURRENT_TIME,
				TRUE,
				$2
			FROM workers
			WHERE workers.id = $1
			RETURNING 
				w_date,
				worker_id,
				come_time,
				status,
				late_min
		`

		err = r.db.QueryRow(context.Background(), query, req.Id, late_min).Scan(
			&date,
			&res.WorkerId,
			&comeTime,
			&res.Status,
			&res.LateMinute,
		)

		if err != nil {
			log.Println("Error in creating attendance: ", err)
			return &models.DailyRes{}, err
		}

		res.Date = date.Format(time.DateOnly)
		res.ComeTime = comeTime.Format(time.TimeOnly)

		return &res, nil

	} else if check.Bool {

		var exit sql.NullTime
		err := r.db.QueryRow(context.Background(), "SELECT exit FROM break_up WHERE worker_id = $1 AND date = CURRENT_DATE", req.Id).Scan(&exit)
		if err == pgx.ErrNoRows {

			err := r.db.QueryRow(context.Background(), `SELECT come_time FROM daily WHERE worker_id = $1`, req.Id).Scan(&wComeTime)
			if err != nil {
				log.Println("Error in getting come_time of daily: ", err)
				return &models.DailyRes{}, err
			}

			work_hour := durationCounter(time.Now(), wComeTime)

			query1 := `
				INSERT INTO break_up( 
					date,
					worker_id,
					exit
				)VALUES(
					CURRENT_DATE,
					$1,
					CURRENT_TIME
				)
			`
			_, err = r.db.Exec(context.Background(), query1, req.Id)
			if err != nil {
				log.Println("Error in inserting data in break_up: ", err)
				return &models.DailyRes{}, err
			}

			query2 := `
				UPDATE daily SET 
					leave_time = CURRENT_TIME, 
					w_hours = $1, 
					status = FALSE 
				WHERE 
					worker_id = $2 AND w_date = CURRENT_DATE
			`
			_, err = r.db.Exec(context.Background(), query2, work_hour, req.Id)
			if err != nil {
				log.Println("Error in updating data in daily table : ", err)
				return &models.DailyRes{}, err
			}
			return &models.DailyRes{}, nil

		} else if exit.Valid {

			err := r.db.QueryRow(context.Background(), `SELECT back FROM break_up WHERE worker_id = $1`, req.Id).Scan(&wComeTime)
			if err != nil {
				log.Println("Error in getting come_time of daily: ", err)
				return &models.DailyRes{}, err
			}

			work_hour := durationCounter(time.Now(), wComeTime)

			query2 := `
				UPDATE daily SET 
					leave_time = CURRENT_TIME, 
					w_hours = w_hours + $1,
					status = FALSE 
				WHERE 
					worker_id = $2 AND w_date = CURRENT_DATE
			`
			_, err = r.db.Exec(context.Background(), query2, work_hour, req.Id)
			if err != nil {
				log.Println("Error in updating data in daily table : ", err)
				return &models.DailyRes{}, err
			}

			return &models.DailyRes{}, nil
		}

	} else if !check.Bool {

		query := `
			UPDATE daily
			SET
				status = TRUE
			WHERE
				worker_id = $1 AND w_date = CURRENT_DATE 
				
		`
		_, err = r.db.Exec(context.Background(), query, req.Id)
		if err != nil {
			log.Println("Error in updating status in daily table: ", err)
			return &models.DailyRes{}, err
		}

		query1 := `
			UPDATE break_up SET 
				back = CURRENT_TIME
			WHERE 
				worker_id = $1 AND date = CURRENT_DATE
		`
		_, err = r.db.Exec(context.Background(), query1, req.Id)
		if err != nil {
			log.Println("Error in updating data in daily table : ", err)
			return &models.DailyRes{}, err
		}
		return &models.DailyRes{}, nil
	}

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
	if err := r.db.QueryRow(context.Background(), query).Scan(&res.Portion); err != nil {
		log.Println("Error while getting attandance portion: ", err)
		return &models.AttendancePortion{}, err
	}

	return &res, nil
}

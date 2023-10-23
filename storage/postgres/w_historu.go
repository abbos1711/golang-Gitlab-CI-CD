package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/tizim-back/api/models"
)

type workerHistoryRepo struct {
	db *pgxpool.Pool
}

func NewWorkerHistoryRepo(db *pgxpool.Pool) *workerHistoryRepo {
	return &workerHistoryRepo{
		db: db,
	}
}

func (r *workerHistoryRepo) GetWorkersByMonth(date string) (*models.WorkersByMonthResp, error) {
	response := models.WorkersByMonthResp{}
	//var in, out time.Time//
	query := `
		SELECT
    		workers.id,
    		workers.img_url,
    		workers.name,
    		workers.surname,
			COUNT(daily.come_time) AS work_day_month,
			COALESCE(AVG(EXTRACT(EPOCH FROM daily.come_time) - '00:00:00'::interval), 0) AS average_time
			COALESCE(AVG(EXTRACT(EPOCH FROM COALESCE(daily.leave_time, '00:00:00') - '00:00:00'::interval)), 0) AS average_absence_time
		FROM
    		workers
		LEFT JOIN (
    		SELECT
        		worker_id,
        		come_time
    		FROM
        		daily
    		WHERE
        		EXTRACT(MONTH FROM come_time) = $1
		) AS daily ON workers.id = daily.worker_id
		GROUP BY
    		workers.id, workers.img_url, workers.name, workers.surname
		ORDER BY
    		workers.id
	`

	rows, err := r.db.Query(context.Background(), query, date)
	if err != nil {
		return &models.WorkersByMonthResp{}, err
	}
	defer rows.Close()

	for rows.Next() {
		worker := models.WorkersByMonth{}

		err := rows.Scan(
			&worker.Id,
			&worker.Img_url,
			&worker.Name,
			&worker.Surname,
			&worker.WorkDayMonth,
			&worker.MiddleComeTime,
			&worker.MiddleLeaveTime,
		)
		if err != nil {
			log.Println("Error while scanning rows: ", err)
			return nil, err
		}

		response.WorkersResp = append(response.WorkersResp, worker)
	}

	return &response, nil
}

func (r *workerHistoryRepo) GetWorkersBy(date string) (*models.WorkersByMonth, error) {
	response := models.WorkersByMonth{}
	// var in, out, hours time.Time
	// var workDay int
	// query := ``

	// rows, err := r.db.Query(context.Background(), query, date, date)
	// if err != nil {
	// 	return &models.WorkersByDateResp{}, err
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	worker := models.WorkersResp{}

	// 	err := rows.Scan(
	// 	)
	// 	if err != nil {
	// 		log.Println("Error while scanning rows: ", err)
	// 		return nil, err
	// 	}

	// 	response.WorkersResp = append(response.WorkersResp, worker)
	// }

	return &response, nil
}

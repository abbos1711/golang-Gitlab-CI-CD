package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/tizim-back/api/models"
)

type workerRepo struct {
	db *pgxpool.Pool
}

func NewWorkerRepo(db *pgxpool.Pool) *workerRepo {
	return &workerRepo{
		db: db,
	}
}

func (r *workerRepo) CreateWorker(worker *models.WorkerCreate) (*models.WorkerResp, error) {

	workerResp := models.WorkerResp{}

	var created_at, come_time, birthday time.Time

	id := uuid.New().String()

	query := `
	INSERT INTO workers(
		id,
        img_url,
        name,
		surname,
        position,
        department,
		gender,
        contact,
        birthday,
		come_time
	) VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	RETURNING 
		id, 
        img_url, 
        name, 
		surname,
        position, 
        department,
		gender,
        contact, 
		birthday,
		come_time,
		created_at
	`
	err := r.db.QueryRow(context.Background(), query,
		id,
		worker.Img_url,
		worker.Name,
		worker.Surname,
		worker.Position,
		worker.Department,
		worker.Gender,
		worker.Contact,
		worker.Birthday,
		worker.ComeTime,
	).Scan(
		&workerResp.Id,
		&workerResp.Img_url,
		&workerResp.Name,
		&workerResp.Surname,
		&workerResp.Position,
		&workerResp.Department,
		&workerResp.Gender,
		&workerResp.Contact,
		&birthday,
		&come_time,
		&created_at,
	)

	if err != nil {
		log.Println("Error in creting workers: ", err)
		return &models.WorkerResp{}, err
	}

	workerResp.ComeTime = come_time.Format(time.TimeOnly)
	workerResp.Birthday = birthday.Format(time.DateOnly)
	workerResp.CreatedAt = created_at.Format(time.RFC1123)

	return &workerResp, nil
}

func (r *workerRepo) DeleteWorker(id string) error {

	query := `
		UPDATE workers SET 
			deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	_, err := r.db.Exec(context.Background(), query, id)

	if err != nil {
		log.Println("Error while deleting worker: ", err)
		return err
	}

	return nil
}

func (r *workerRepo) UpdateWorker(newData *models.WorkerUpdate) (*models.WorkerResp, error) {

	updatedWorker := models.WorkerResp{}
	var created_at, updated_at, come_time, birthday time.Time

	query := `
		UPDATE workers SET 
			img_url = $1,
			name = $2,
			surname = $3,
			position = $4,
			department = $5,
			gender = $6,
			contact = $7,
			birthday = $8,
			come_time = $9,
			updated_at = CURRENT_TIMESTAMP
		WHERE 
			id  = $10 AND deleted_at IS NULL
		RETURNING 
			id, 
			img_url, 
			name, 
			surname,
			position, 
			department,
			gender,
			contact, 
			birthday,		
			come_time,
			created_at,
			updated_at
	`
	err := r.db.QueryRow(context.Background(), query,
		newData.Img_url,
		newData.Name,
		newData.Surname,
		newData.Position,
		newData.Department,
		newData.Gender,
		newData.Contact,
		newData.Birthday,
		newData.ComeTime,
		newData.Id,
	).Scan(
		&updatedWorker.Id,
		&updatedWorker.Img_url,
		&updatedWorker.Name,
		&updatedWorker.Surname,
		&updatedWorker.Position,
		&updatedWorker.Department,
		&updatedWorker.Gender,
		&updatedWorker.Contact,
		&birthday,
		&come_time,
		&created_at,
		&updated_at,
	)

	if err != nil {
		log.Println("Error while updating worker info: ", err)
		return nil, err
	}

	updatedWorker.ComeTime = come_time.Format(time.TimeOnly)
	updatedWorker.Birthday = birthday.Format(time.DateOnly)
	updatedWorker.CreatedAt = created_at.Format(time.RFC1123)
	updatedWorker.UpdatedAt = updated_at.Format(time.RFC1123)

	return &updatedWorker, nil
}

func (r *workerRepo) GetAllWorkers() (*models.AllWorkers, error) {
	var workers models.AllWorkers

	query := `
		SELECT
			(SELECT COUNT(*) FROM workers WHERE deleted_at IS NULL) AS total_amount,
			(SELECT COUNT(*) FROM workers WHERE gender = 'male' AND deleted_at IS NULL) AS male,
			(SELECT COUNT(*) FROM workers WHERE gender = 'female' AND deleted_at IS NULL) AS female,
			id,
			name,
			surname,
			img_url,
			position,
			department,
			gender,
			contact,
			birthday,
			come_time,
			created_at,
			updated_at
		FROM
			workers
		WHERE
			deleted_at IS NULL
		GROUP BY 
			id

	`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		log.Println("Error while getting all workers: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var w models.WorkerResp

		var created_at, come_time, birthday time.Time

		var updatedAtNullable sql.NullTime

		err := rows.Scan(
			&workers.Total_amount,
			&workers.Male,
			&workers.Female,
			&w.Id,
			&w.Name,
			&w.Surname,
			&w.Img_url,
			&w.Position,
			&w.Department,
			&w.Gender,
			&w.Contact,
			&birthday,
			&come_time,
			&created_at,
			&updatedAtNullable,
		)

		if err != nil {
			log.Println("Error while scanning rows: ", err)
			return nil, err
		}

		if updatedAtNullable.Valid {
			w.UpdatedAt = updatedAtNullable.Time.Format(time.RFC1123)
		} else {
			w.UpdatedAt = ""
		}

		w.ComeTime = come_time.Format(time.TimeOnly)
		w.Birthday = birthday.Format(time.DateOnly)
		w.CreatedAt = created_at.Format(time.RFC1123)

		workers.Workers = append(workers.Workers, w)
	}

	return &workers, nil
}

func (r *workerRepo) GetWorker(id string) (*models.WorkerResp, error) {
	var res models.WorkerResp

	var created_at, come_time, birthday time.Time
	var updatedAtNullable sql.NullTime

	query := `
		SELECT 
			id,
			name, 
			surname,
			img_url,
			position, 
			department,
			gender,
			contact,
			birthday,
			come_time,
			created_at,
			updated_at
		FROM 
			workers  
		WHERE 
			deleted_at IS NULL AND id = $1
	`
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&res.Id,
		&res.Name,
		&res.Surname,
		&res.Img_url,
		&res.Position,
		&res.Department,
		&res.Gender,
		&res.Contact,
		&birthday,
		&come_time,
		&created_at,
		&updatedAtNullable,
	)

	if err != nil {
		log.Println("Error while getting worker: ", err)
		return nil, err
	}

	if updatedAtNullable.Valid {
		res.UpdatedAt = updatedAtNullable.Time.Format(time.RFC1123)
	} else {
		res.UpdatedAt = ""
	}

	res.ComeTime = come_time.Format(time.TimeOnly)
	res.Birthday = birthday.Format(time.DateOnly)
	res.CreatedAt = created_at.Format(time.RFC1123)

	return &res, err
}

func (r workerRepo) GetWorkersByGender(gender string) (*models.AllWorkersFilter, error) {
	var res models.AllWorkersFilter

	query := `
		SELECT
			id,
			name,
			surname,
			img_url,
			position,
			department,
			gender,
			contact,
			birthday,
			come_time,
			created_at,
			updated_at
		FROM
			workers
		WHERE
			deleted_at IS NULL AND gender = $1
		GROUP BY 
			id

	`
	rows, err := r.db.Query(context.Background(), query, gender)
	if err != nil {
		log.Println("Error while getting all workers: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var w models.WorkerResp

		var created_at, come_time, birthday time.Time

		var updatedAtNullable sql.NullTime

		err := rows.Scan(
			&w.Id,
			&w.Name,
			&w.Surname,
			&w.Img_url,
			&w.Position,
			&w.Department,
			&w.Gender,
			&w.Contact,
			&birthday,
			&come_time,
			&created_at,
			&updatedAtNullable,
		)

		if err != nil {
			log.Println("Error while scanning rows: ", err)
			return nil, err
		}

		if updatedAtNullable.Valid {
			w.UpdatedAt = updatedAtNullable.Time.Format(time.RFC1123)
		} else {
			w.UpdatedAt = ""
		}

		w.ComeTime = come_time.Format(time.TimeOnly)
		w.Birthday = birthday.Format(time.DateOnly)
		w.CreatedAt = created_at.Format(time.RFC1123)

		res.Workers = append(res.Workers, w)

	}

	return &res, nil
}

func (r *workerRepo) GetWorkersAtWork() (*models.AllWorkersFilter, error) {
	var workers models.AllWorkersFilter

	query := `
		SELECT
			workers.id,
			workers.name,
			workers.surname,
			workers.img_url,
			workers.position,
			workers.department,
			workers.gender,
			workers.contact,
			workers.birthday,
			workers.come_time,
			workers.created_at,
			workers.updated_at
		FROM
			workers
		JOIN
			daily ON workers.id = daily.worker_id
		WHERE
			workers.deleted_at IS NULL
			AND daily.w_date = CURRENT_DATE
		GROUP BY 
			workers.id
	`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		log.Println("Error while getting workers at work: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var w models.WorkerResp

		var created_at, come_time, birthday time.Time

		var updatedAtNullable sql.NullTime

		err := rows.Scan(
			&w.Id,
			&w.Name,
			&w.Surname,
			&w.Img_url,
			&w.Position,
			&w.Department,
			&w.Gender,
			&w.Contact,
			&birthday,
			&come_time,
			&created_at,
			&updatedAtNullable,
		)

		if err != nil {
			log.Println("Error while scanning rows: ", err)
			return nil, err
		}

		if updatedAtNullable.Valid {
			w.UpdatedAt = updatedAtNullable.Time.Format(time.RFC1123)
		} else {
			w.UpdatedAt = ""
		}

		w.ComeTime = come_time.Format(time.TimeOnly)
		w.Birthday = birthday.Format(time.DateOnly)
		w.CreatedAt = created_at.Format(time.RFC1123)

		workers.Workers = append(workers.Workers, w)
	}

	return &workers, nil
}

package repository

import (
	"movie-services/model/db"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type (
	ApiLogRepository interface {
		SaveLog(*db.ApiLog) (db.ApiLog, error)
	}

	apiLogRepository struct {
		DB *sqlx.DB
	}
)

func NewApiLogRepository(db *sqlx.DB) ApiLogRepository {
	al := apiLogRepository{}
	al.DB = db

	return al
}

func (al apiLogRepository) SaveLog(apiLog *db.ApiLog) (db.ApiLog, error) {
	query := `insert into api_log (
					id,
					request_url,
					request_parameter,
					status_code,
					response,
					created_at
			) values (uuid_generate_v4(), $1, $2, $3, $4, now())`

	err := al.DB.QueryRow(
		query,
		apiLog.RequestURL,
		apiLog.RequestParameter,
		apiLog.StatusCode,
		apiLog.Response,
	).Scan(&apiLog.ID)

	if err != nil {
		log.Error(err)
		// return *apiLog, err
	}

	return *apiLog, nil
}

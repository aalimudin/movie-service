package db

import (
	"time"

	"github.com/google/uuid"
)

type (
	ApiLog struct {
		ID               uuid.UUID `db:"id"`
		RequestURL       string    `db:"request_url"`
		RequestParameter string    `db:"request_parameter"`
		StatusCode       int       `db:"status_code"`
		Response         string    `db:"response"`
		CreatedAt        time.Time `db:"created_at"`
	}
)

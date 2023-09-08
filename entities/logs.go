package entities

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type LogsPayload struct {
	ID                 string         `json:"id"`
	Fullname           string         `json:"fullname"`
	Email              string         `json:"email"`
	Event              string         `json:"event"`
	UserAgent          string         `json:"user_agent"`
	HttpStatusCode     int            `json:"http_status_code"`
	HttpMethod         string         `json:"http_method"`
	ClientRequestData  types.JSONText `json:"client_request_data"`
	ClientResponseData types.JSONText `json:"client_response_data"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at" db:"updated_at"`
}

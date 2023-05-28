package webhook

import (
	"database/sql"
	"github.com/rs/zerolog"
	"net/http"
)

type (
	Server struct {
		client *http.Client
		logger *zerolog.Logger
		db     *sql.DB
	}
)

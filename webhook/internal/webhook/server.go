package webhook

import "net/http"

type (
	Server struct {
		client *http.Client
		logger *zerolog.Logger
	}
)

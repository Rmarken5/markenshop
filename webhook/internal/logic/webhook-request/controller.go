package webhook_request

import (
	"github.com/rmarken5/markenshop/database/repository"
	pull_request "github.com/rmarken5/markenshop/webhook/internal/models/pull-request"
	"github.com/rmarken5/markenshop/webhook/internal/models/push"
	"github.com/rs/zerolog"
)

type (
	RequestController interface {
		WritePushToDatabase(push *push.Push) error
		WritePullRequestToDatabase(pr *pull_request.PullRequest) error
	}

	Controller struct {
		logger *zerolog.Logger
		repo   repository.Repository
		parser Parse
	}
)

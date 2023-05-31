package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rmarken5/markenshop/database/models"
	"github.com/rmarken5/markenshop/database/models/push"
	"github.com/rs/zerolog"
)

//go:generate mockgen -destination=./mock_repository.go -package=repository . Repository

type (
	PushReader interface {
		ReadAll(offset, limit int) ([]*push.Push, error)
		Read(id string) (*push.Push, error)
	}
	PushWriter interface {
		Insert(push *push.Push) (string, error)
	}

	PusherReader interface {
		ReadAll(offset, limit int) ([]*push.Pusher, error)
		Read(id string) (*push.Pusher, error)
	}

	RepoBuilder struct {
		modifiers []modifier
	}

	repository struct {
		db     *sqlx.DB
		logger *zerolog.Logger
	}

	modifier func(repo *repository)
)

func (r *RepoBuilder) WithDB(db *sqlx.DB) *RepoBuilder {
	r.modifiers = append(r.modifiers, func(repo *repository) {
		repo.db = db
	})
	return r
}

func (r *RepoBuilder) WithLogger(logger *zerolog.Logger) *RepoBuilder {
	r.modifiers = append(r.modifiers, func(repo *repository) {
		repo.logger = logger
	})
	return r
}

func (r *RepoBuilder) Build() Repository {
	repo := &repository{}
	for _, m := range r.modifiers {
		m(repo)
	}
	return repo
}

func (r *repository) InsertEvent(e models.Event) error {
	panic("Not yet implemented")
}

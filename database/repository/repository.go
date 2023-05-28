package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rmarken5/markenshop/database/models"
	"github.com/rs/zerolog"
)

type (
	Repository interface {
		InsertEvent(event models.Event) error
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

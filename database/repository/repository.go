package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rmarken5/markenshop/database/models/push"
	"github.com/rs/zerolog"
)

//go:generate mockgen -destination=./mock_repository.go -package=repository . Repository

type (
	Repository interface {
		PushReader
		PusherWriter
		PusherReader
		PusherWriter
		SenderReader
		SenderWriter
	}
	PushReader interface {
		ReadAllPush(offset, limit int) ([]*push.Push, error)
		ReadPush(id string) (*push.Push, error)
	}
	PushWriter interface {
		InsertPush(push *push.Push) (string, error)
	}

	PusherReader interface {
		ReadAllPusher(offset, limit int) ([]*push.Pusher, error)
		ReadPusherByName(name string) (*push.Pusher, error)
	}
	PusherWriter interface {
		InsertPusher(pusher *push.Pusher) (string, error)
	}

	SenderReader interface {
		ReadAllSender(offset, limit int) ([]*push.Sender, error)
		ReadSenderByName(name string) (*push.Sender, error)
	}
	SenderWriter interface {
		InsertSender(pusher *push.Sender) (string, error)
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

func (r *repository) ReadAllPush(offset, limit int) ([]*push.Push, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) ReadPush(id string) (*push.Push, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) InsertPush(push *push.Push) (string, error) {
	var key string
	r.logger.Info().Msgf("Getting ready to insert push: %v", push)

	queryx, err := r.db.Queryx(insertPush, push)
	if err != nil {
		r.logger.Err(err).Msg("Trying to insert push")
		return "", err
	}
	err = queryx.Scan(&key)
	if err != nil {
		r.logger.Err(err).Msg("Trying to scan key into variable")
		return "", err
	}
	return key, nil
}

func (r *repository) ReadAllSender(offset, limit int) ([]*push.Sender, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) ReadSenderByName(name string) (*push.Sender, error) {
	var sender push.Sender
	err := r.db.Select(&sender, getSenderByName, name)
	if err != nil {
		return nil, err
	}
	return &sender, nil
}

func (r *repository) InsertSender(sender *push.Sender) (string, error) {
	var key string
	r.logger.Info().Msgf("Getting ready to insert sender: %v", sender)

	queryx, err := r.db.Queryx(insertSender, sender)
	if err != nil {
		r.logger.Err(err).Msg("Trying to insert sender")
		return "", err
	}
	err = queryx.Scan(&key)
	if err != nil {
		r.logger.Err(err).Msg("Trying to scan key into variable")
		return "", err
	}
	return key, nil
}

func (r *repository) ReadAllPusher(offset, limit int) ([]*push.Pusher, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) ReadPusherByName(name string) (*push.Pusher, error) {
	var pusher push.Pusher
	err := r.db.Select(&pusher, getSenderByName, name)
	if err != nil {
		return nil, err
	}
	return &pusher, nil
}

func (r *repository) InsertPusher(pusher *push.Pusher) (string, error) {
	var key string
	r.logger.Info().Msgf("Getting ready to insert pusher: %v", pusher)
	queryx, err := r.db.Queryx(insertPusher, pusher)
	if err != nil {
		r.logger.Err(err).Msg("Trying to insert sender")
		return "", err
	}
	err = queryx.Scan(&key)
	if err != nil {
		r.logger.Err(err).Msg("Trying to scan key into variable")
		return "", err
	}
	return key, nil
}

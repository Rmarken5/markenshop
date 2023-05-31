package webhook_request

import (
	"errors"
	"github.com/rmarken5/markenshop/database/repository"
	pull_request "github.com/rmarken5/markenshop/webhook/internal/models/pull-request"
	"github.com/rmarken5/markenshop/webhook/internal/models/push"
	"github.com/rs/zerolog"
	"io"
	"strings"
)

const PUSH = "push"
const PULL_REQUEST = "pull_request"

type (
	RequestController interface {
		ProcessWebhookRequest(eventType string, push io.Reader) error
	}

	Controller struct {
		logger *zerolog.Logger
		repo   repository.Repository
		parser Parse
	}
)

func (c *Controller) ProcessWebhookRequest(eventType string, r io.Reader) error {
	switch strings.ToLower(eventType) {
	case PUSH:
		ps, err := c.parser.ToPush(r)
		if err != nil {
			c.logger.Err(err).Msg("parsing to push")
			return err
		}
		err = c.processPush(ps)
		if err != nil {
			c.logger.Err(err).Msg("processing push")
			return err
		}
	default:
		return errors.New("Not Supported")
	}
	return nil
}

func (c *Controller) processPush(push *push.Push) error {

	senderName := push.Sender.Login

	return nil
}

func (c *Controller) processPullRequst(pr *pull_request.PullRequest) error {
	// TODO: Implement me
	panic("implement me")
}

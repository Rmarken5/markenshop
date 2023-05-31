package webhook_request

import (
	"encoding/json"
	pullrequest "github.com/rmarken5/markenshop/webhook/internal/models/pull-request"
	"github.com/rmarken5/markenshop/webhook/internal/models/push"
	"github.com/rs/zerolog"
	"io"
)

type (
	Parse interface {
		ToPush(r io.Reader) (*push.Push, error)
		ToPullRequest(r io.Reader) (*pullrequest.PullRequest, error)
	}

	Parser struct {
		logger *zerolog.Logger
	}
)

var instance *Parser

// CreateParser  creates a Singleton using the provided logger. ALl calls subsequent calls will panic.
func CreateParser(logger *zerolog.Logger) Parse {
	if instance != nil {
		panic("call GetParserInstance to get instance.")
	}
	instance = &Parser{logger: logger}
	return instance
}

func GetParserInstance() Parse {
	if instance == nil {
		panic("Parser instance has not been initialized. Call CreateParser.")
	}
	return instance
}

func (p *Parser) ToPush(r io.Reader) (*push.Push, error) {
	push := &push.Push{}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, push)
	if err != nil {
		return nil, err
	}

	return push, nil
}

func (p *Parser) ToPullRequest(r io.Reader) (*pullrequest.PullRequest, error) {
	pullReq := &pullrequest.PullRequest{}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, pullReq)
	if err != nil {
		return nil, err
	}

	return pullReq, nil
}

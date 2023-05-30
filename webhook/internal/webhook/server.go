package webhook

import (
	"github.com/rmarken5/markenshop/database/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type (
	eventHandler struct {
		client         *http.Client
		logger         *zerolog.Logger
		db             repository.Repository
		allowedMethods map[string]bool
	}

	ServerBuilder struct {
		modifiers []modifier
	}

	modifier func(s *eventHandler)
)

func (sb *ServerBuilder) WithClient(client *http.Client) *ServerBuilder {
	sb.modifiers = append(sb.modifiers, func(s *eventHandler) {
		s.client = client
	})
	return sb
}

func (sb *ServerBuilder) WithLogger(logger *zerolog.Logger) *ServerBuilder {
	l := logger.With().Str("module", "http eventHandler").Logger()
	sb.modifiers = append(sb.modifiers, func(s *eventHandler) {
		s.logger = &l
	})
	return sb
}

func (sb *ServerBuilder) WithDB(db repository.Repository) *ServerBuilder {
	sb.modifiers = append(sb.modifiers, func(s *eventHandler) {
		s.db = db
	})
	return sb
}

func (sb *ServerBuilder) AllowedMethods(allowedMethods ...string) *ServerBuilder {
	sb.modifiers = append(sb.modifiers, func(s *eventHandler) {
		for _, method := range allowedMethods {
			s.allowedMethods[method] = true
		}
	})
	return sb
}

func (sb *ServerBuilder) BuildEventHandler() http.Handler {
	srv := &eventHandler{}
	srv.allowedMethods = make(map[string]bool)
	for _, mod := range sb.modifiers {
		mod(srv)
	}
	return srv
}

func (s *eventHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	s.logger.Info().Msgf("Handling request: %v", request.RequestURI)

	switch request.Method {
	case http.MethodPost:
		err := s.handlePost(responseWriter, request)
		if err != nil {
			return
		}
	default:
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}
	defer request.Body.Close()
	s.logger.Info().Msgf("Request handled: %v", request.RequestURI)
}

func (s *eventHandler) handlePost(responseWriter http.ResponseWriter, request *http.Request) error {
	requestBytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Err(err).Msg("when trying to read request body")
	}
	s.logger.Info().Msgf("request headers: %v", request.Header)
	s.logger.Info().Msgf("request body: %v", string(requestBytes))

	responseWriter.WriteHeader(http.StatusCreated)
	return nil
}

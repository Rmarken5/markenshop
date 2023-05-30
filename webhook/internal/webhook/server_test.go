package webhook

import (
	"github.com/golang/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rmarken5/markenshop/database/repository"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestEventHandler_ServeHTTP(t *testing.T) {
	// create a mocked HTTP request
	req, err := http.NewRequest("POST", "/example-url", strings.NewReader("example body"))
	if err != nil {
		t.Fatal(err)
	}
	logger := zerolog.Nop()

	ctrl := gomock.NewController(t)
	// create a fake response recorder
	res := httptest.NewRecorder()

	// create a new event handler with modifiers
	builder := ServerBuilder{}
	handler := builder.WithLogger(&logger).
		WithDB(repository.NewMockRepository(ctrl)).
		AllowedMethods(http.MethodPost).
		BuildEventHandler()

	// call ServeHTTP() method to handle the request
	handler.ServeHTTP(res, req)

	// assert that the status code is correct
	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestEventHandler_HandlePost(t *testing.T) {
	// prepare test data
	body := strings.NewReader("example request body")
	response := httptest.NewRecorder()
	logger := zerolog.Nop()

	ctrl := gomock.NewController(t)

	// create a new event handler with modifiers
	builder := ServerBuilder{}
	handler := builder.WithLogger(&logger).
		WithDB(repository.NewMockRepository(ctrl)).
		AllowedMethods(http.MethodPost).
		BuildEventHandler()
	// call the handlePost() method
	err := handler.(*eventHandler).handlePost(response, &http.Request{Method: http.MethodPost, Body: io.NopCloser(body)})
	assert.NoError(t, err)

	// assert that the response body is correct
	assert.Equal(t, http.StatusCreated, response.Code)
}

package webhook_request

import (
	"bytes"
	"encoding/json"
	"errors"
	pullrequest "github.com/rmarken5/markenshop/webhook/internal/models/pull-request"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"

	"github.com/rmarken5/markenshop/webhook/internal/models/push"
	"github.com/rs/zerolog"
)

func TestCreateAndGetParserInstance(t *testing.T) {
	t.Parallel()
	logger := zerolog.Nop()
	parser := CreateParser(&logger)

	if parser == nil {
		t.Errorf("CreateParser returned nil")
	}

	p := GetParserInstance()

	assert.Equal(t, parser, p)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when it was expected to")
		}
	}()

	CreateParser(&logger)
}

func TestToPush(t *testing.T) {
	t.Parallel()
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		p       *Parser
		args    args
		want    *push.Push
		wantErr bool
	}{
		{
			name: "valid input",
			p:    &Parser{},
			args: args{
				r: strings.NewReader(`{"head_commit":{"message":"test","author":{}},"ref":"refs/heads/main"}`),
			},
			want: &push.Push{
				HeadCommit: push.HeadCommit{
					Message: "test",
					Author:  push.Author{},
				},
				Ref: "refs/heads/main",
			},
			wantErr: false,
		},
		{
			name: "invalid input - read error",
			p:    &Parser{},
			args: args{
				r: &errorReader{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid input - json unmarshalling error",
			p:    &Parser{},
			args: args{
				r: strings.NewReader(`invalid json`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.p.ToPush(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.ToPush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !isEqual(got, tt.want) {
				t.Errorf("Parser.ToPush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPullRequest(t *testing.T) {
	t.Parallel()
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		p       *Parser
		args    args
		want    *pullrequest.PullRequest
		wantErr bool
	}{
		{
			name: "valid input",
			p:    &Parser{},
			args: args{
				r: strings.NewReader(`{"number":5}`),
			},
			want: &pullrequest.PullRequest{
				Number: 5,
			},
			wantErr: false,
		},
		{
			name: "invalid input - read error",
			p:    &Parser{},
			args: args{
				r: &errorReader{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid input - json unmarshalling error",
			p:    &Parser{},
			args: args{
				r: strings.NewReader(`invalid json`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.p.ToPullRequest(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.ToPullRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !isEqual(got, tt.want) {
				t.Errorf("Parser.ToPullRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(a, b interface{}) bool {
	bytesA, _ := json.Marshal(a)
	bytesB, _ := json.Marshal(b)

	return bytes.Equal(bytesA, bytesB)
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read errorerror")
}

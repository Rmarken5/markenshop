package main

import (
	"github.com/rmarken5/markenshop/webhook/internal/webhook"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout)
	_ = loadConfigurations()

	httpClient := &http.Client{}
	sb := webhook.ServerBuilder{}
	sb.WithClient(httpClient).WithLogger(&logger)

	mux := http.ServeMux{}

	mux.Handle("/wh/handler", sb.BuildEventHandler())

	logger.Fatal().Err(http.ListenAndServe(":8999", &mux))

}

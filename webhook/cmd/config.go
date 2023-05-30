package main

import (
	"errors"
	"os"
)

type (
	config struct {
		postgres database
	}

	database struct {
		host, user, password, database, port string
	}
)

func loadConfigurations() config {
	return config{
		postgres: loadDatabaseConfigs(),
	}
}

func loadDatabaseConfigs() database {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		panic(errors.New("db host is empty"))
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		panic(errors.New("db user is empty"))
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		panic(errors.New("db password is empty"))
	}

	databaseName := os.Getenv("POSTGRES_NAME")
	if databaseName == "" {
		panic(errors.New("db name is empty"))
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		panic(errors.New("db port is empty"))
	}

	return database{
		host:     host,
		user:     user,
		password: password,
		database: databaseName,
		port:     port,
	}
}

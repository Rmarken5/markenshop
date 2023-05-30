module github.com/rmarken5/markenshop/webhook

go 1.20

require (
	github.com/rmarken5/markenshop/database v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.29.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/sys v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect

)

replace github.com/rmarken5/markenshop/database => ../database

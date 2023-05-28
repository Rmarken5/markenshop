module github.com/rmarken5/markenshop/webhook

go 1.20

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	golang.org/x/sys v0.8.0 // indirect
	github.com/rmarken5/markenshop/database
)

replace (
	github.com/rmarken5/markenshop/database latest => ../database latest
)

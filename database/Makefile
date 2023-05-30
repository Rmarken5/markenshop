.PHONEY:
build-database:
	docker build . -t rmarken5/markenshop-database:latest

.PHONY: build-database
run-database:
	docker run -rm -d -p 5432:5432 -v /var/lib/webhook/postgresql/data:/var/lib/webhok/postgresql/data rmarken5/markenshop-database:latest


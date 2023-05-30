.PHONY: build-webhook
build-webhook:
	docker build -f webhook/Dockerfile . -t webhook

.PHONY: run-rebhook
run-webhook:
	docker run --rm -p 8999:8999 webhook:latest

.PHONY:
build-database:
	$(MAKE) -C ./database build-database

.PHONY: build-database
run-database:
	$(MAKE) -C ./database run-database

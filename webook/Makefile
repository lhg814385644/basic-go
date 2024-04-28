.PHONY: docker
docker:
	@rm webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -o webook .
	@docker rmi -f webook:dev || true
	@docker build -t  hgfly/webook:dev .
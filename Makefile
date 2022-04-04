test/run:
	@go test -coverprofile cover.out

test/html:
	@make test/run
	@go tool cover -html=cover.out


run:
	go run cmd/main.go

##@ Testing

.PHONY: test
test: ## Run unit test
	go test ./... -cover -race

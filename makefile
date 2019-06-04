test_dir = $(shell go list ./... | grep -v mock)

test: 
	@echo "=== RUNNING TEST ==="
	@go test -cover -race $(test_dir) | tee test.out
	@scripts/test/test.sh test.out 30
	@echo "done"

lint:
	@echo "=== RUNNING GOLINT ==="
	@golint -set_exit_status cmd/... internal/...
	@echo "done"

vet:
	@echo "=== RUNNING GO VET ==="
	@go vet ./cmd/... ./internal/...
	@echo "done"

run-shortener:
	@echo "=== RUNNING SHORTENER ==="
	@refresh run -c scripts/refresh/config-shortener.yml
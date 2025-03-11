# Makefile

run:
	@echo "starting risk management server"
	go run cmd/main.go

test:
	go test ./... -timeout 5m

enum:
	@echo "Generating enums"
	go get github.com/dmarkham/enumer
	go install github.com/dmarkham/enumer
	cd internal/enum && $$(go env GOPATH)/bin/enumer -json -type=RiskState

swagger:
	@echo "Generating swagger docs of apis"
	@echo "gopath $$(go env GOPATH)"
	go get github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag
	cd pkg/api && $$(go env GOPATH)/bin/swag init --parseDependency --parseInternal -g server.go --parseDepth 2
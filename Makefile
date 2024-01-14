download:
	@echo Download go.mod dependencies
	@go mod download
 
install-tools: download
	@echo Installing tools from tools.go
	@cat backend/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

generate:
	@echo Generating api server and models
	@go generate -x ./...

run:
	@go run backend/main.go
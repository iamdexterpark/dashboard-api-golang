REPO=github.com/ddexterpark/dashboard-api-golang
MODULE=dashboard-api-golang

default: compile

dependencies:
	wget https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json

	@if [ test ! $(which swagger); then \
			echo "Installing go-swagger"; \
			brew install go-swagger; \
			exit 1; \
		fi
lint:
	@echo "Formatting with gofmt"
	gofmt ./client

test: lint
	go test ./... -coverprofile=coverage.out
	go vet

clean:
	go clean
	rm spec2.json

compile: dependencies
	swagger generate client -f ./spec2.json -A meraki-api-golang --principal meraki

.PHONY: dependencies lint test compile clean
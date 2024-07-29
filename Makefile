GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: coverage-check
coverage-check: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml

coverage-report: 
	go tool cover -html=cover.out -o=cover.html

coverage:
	$(MAKE) $(MAKEFLAGS) coverage-check; rc=$$? \
        ; $(MAKE) $(MAKEFLAGS) coverage-report \
        ; exit $$rc

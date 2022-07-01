#----SERVICES----

##----GREET----

tidy-greet:
	cd services/greet && make tidy

build-greet:
	cd services/greet && make build

watch-greet:
	cd services/greet && watch notify

dev-greet:
	cd services/greet && make dev

test-greet:
	cd services/greet && make test-handlers

lint-greet:
	cd services/greet && make lint

lint-fix-greet:
	cd services/greet && make lint-fix

#----INSTALLATION----

install-lint:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	golangci-lint --version

install-gotest:
	curl https://gotest-release.s3.amazonaws.com/gotest_linux > gotest && chmod +x gotest

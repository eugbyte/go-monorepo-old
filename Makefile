#----SERVICES----

##----NOTIFY----

tidy-notify:
	cd services/notify && make tidy

build-notify:
	cd services/notify && make build

watch-notify:
	cd services/notify && watch notify

dev-notify:
	cd services/notify && make dev

func-start-notify:
	cd services/notify && make func-start

test-notify:
	cd services/notify && make test-handlers

lint-notify:
	cd services/notify && make lint

lint-fix-notify:
	cd services/notify && make lint-fix

#----INSTALLATION----

install-lint:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	golangci-lint --version

install-gotest:
	curl https://gotest-release.s3.amazonaws.com/gotest_linux > gotest && chmod +x gotest

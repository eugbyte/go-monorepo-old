#----DEVELOPMENT----

tidy:
	go mod tidy -e

build:
	export GO111MODULE=on
	go build main.go

dev:
	go run main.go

func-start:
	func start

watch:
	when-changed -r -1 -s "./" make build

clean:
	rm -rf ./bin ./vendor go.sum

#----TESTING----
test-handlers:
	gotest -v ./hello_handler/... || go clean -testcache
	go clean -testcache

test:
	make test-handlers

#----LINT----

lint:
	@if [ -z `which golangci-lint 2> /dev/null` ]; then \
			echo "Need to install golangci-lint, execute \"make lint-install\"";\
			exit 1;\
	fi
	golangci-lint run

lint-fix:
	golangci-lint run --fix

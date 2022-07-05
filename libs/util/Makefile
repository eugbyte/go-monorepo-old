tidy:
	go mod tidy -e
download:
	go mod download

#----LINT----
lint:
	@if [ -z `which golangci-lint 2> /dev/null` ]; then \
			echo "Need to install golangci-lint, execute \"make lint-install\"";\
			exit 1;\
	fi
	golangci-lint run
lint-fix:
	golangci-lint run --fix

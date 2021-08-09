build:
	go build -o ./bin/main ./cmd/ova-moment-api/main.go

run:
	if ! [ -f ./bin/main ]; then \
	    make build; \
	fi
	./bin/main

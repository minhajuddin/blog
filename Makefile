all: test

test: errcheck
	go test -v

build: errcheck
	go build

errcheck:
	errcheck blog

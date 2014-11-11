all: test errcheck

test:
	go test -v

build: errcheck
	go build

errcheck:
	errcheck blog

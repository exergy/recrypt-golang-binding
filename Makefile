ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

clean:
	rm -rf ./lib/hello/target
	rm ./lib/hello/Cargo.lock ./lib/libhello.so go-rust

library:
	$(MAKE) -C lib/hello build

build:
	cp lib/hello/target/release/libhello.so ./lib
	go build -ldflags="-r $(ROOT_DIR)lib" -o recrypt

test: library
	cp lib/hello/target/release/libhello.so ./lib
	go test -v -ldflags="-r $(ROOT_DIR)lib"

all: library build

run: build
	./recrypt

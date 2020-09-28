.PHONY: build test bpftest coverhtml bench gen clean

PKG_NAME=$(shell basename `pwd`)
PKG_LIST := ./pkg/...

all: build

build:
	go build -o bin/c_map ./cmd/create_map

gen:
	./build_elf.sh

clean:
	rm bin/*

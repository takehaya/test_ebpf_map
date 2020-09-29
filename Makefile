.PHONY: build test bpftest coverhtml bench gen clean

PKG_NAME=$(shell basename `pwd`)
PKG_LIST := ./pkg/...

all: build

build:
	go build -o bin/c_map ./cmd/create_map
	go build -o bin/e_c_map ./cmd/elffrom_createmap

gen:
	./build_elf.sh

clean:
	rm bin/*

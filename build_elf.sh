#!/bin/bash

prog_dir=./build/bpfprog
rm -rf $prog_dir
mkdir -p $prog_dir
cp ./Makefile-bpf $prog_dir/Makefile
cp -r ./src $prog_dir
cp -r ./include $prog_dir
pushd $prog_dir
make EXTRA_CFLAGS="$*"
popd

cp $prog_dir/src/*.o ./obj/

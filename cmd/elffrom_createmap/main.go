package main

import (
	"fmt"
	"os"

	"github.com/cilium/ebpf"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

func LoadElf(filepath string) (*ebpf.Collection, *ebpf.ProgramSpec, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	// Read ELF
	spec, err := ebpf.LoadCollectionSpecFromReader(f)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(spec))

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return coll, spec.Programs["dummymain"], nil
}

func main() {
	coll, spec, err := LoadElf("./createmap.o")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(coll))
	innerMap := coll.Maps["inner_map"]
	outerArrMap := coll.Maps["outer_arr"]

	if err := innerMap.Put(uint32(0), uint32(4242)); err != nil {
		fmt.Println("Can't put inner map:", err)
	}

	if err := outerArrMap.Put(uint32(0), innerMap); err != nil {
		fmt.Println("Can't put inner map:", err)
	}

	// fmt.Printf("%# v\n", pretty.Formatter(innerMap))
	if err := outerArrMap.Lookup(uint32(0), &innerMap); err != nil {
		fmt.Println("Can't lookup 0:", err)
	}
	var v uint32
	if err := innerMap.Lookup(uint32(0), &v); err != nil {
		fmt.Println("Can't lookup 0:", err)
	}
	fmt.Println(v)

	prog, err := ebpf.NewProgramWithOptions(spec, ebpf.ProgramOptions{
		LogLevel: 2,
		LogSize:  102400 * 1024,
	})

	fmt.Printf("%# v\n", pretty.Formatter(prog))
	if err != nil {
		panic(err)
	}
}

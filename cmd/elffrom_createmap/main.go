package main

import (
	"fmt"
	"os"

	"github.com/cilium/ebpf"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

func LoadElf(filepath string) (*ebpf.Collection, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Read ELF
	fmt.Println("	ebpf.LoadCollectionSpecFromReader(f)	")

	spec, err := ebpf.LoadCollectionSpecFromReader(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(spec))

	fmt.Println("	coll, err := ebpf.NewCollection(spec)	")
	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return coll, nil
}

func main() {
	coll, err := LoadElf("./obj/createmap.o")
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

	fmt.Printf("%# v\n", pretty.Formatter(innerMap))
	if err := outerArrMap.Lookup(uint32(0), &innerMap); err != nil {
		fmt.Println("Can't lookup 0:", err)
	}
	var v uint32
	if err := innerMap.Lookup(uint32(0), &v); err != nil {
		fmt.Println("Can't lookup 0:", err)
	}
	fmt.Println(v)

}

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
}

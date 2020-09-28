package main

import (
	"github.com/cilium/ebpf"
	"fmt"
	"github.com/kr/pretty"
)

func main(){
	innerSpec := &ebpf.MapSpec{
		Type:       ebpf.Array,
		KeySize:    4,
		ValueSize:  4,
		MaxEntries: 3,
	}
	outerArrSpec := &ebpf.MapSpec{
		Type:       ebpf.ArrayOfMaps,
		KeySize:    4,
		ValueSize:  4,
		MaxEntries: 1,
		InnerMap: innerSpec,
	}
	outerHashSpec := &ebpf.MapSpec{
		Type:       ebpf.HashOfMaps,
		KeySize:    4,
		ValueSize:  4,
		MaxEntries: 5,
		InnerMap: innerSpec,
	}
	fmt.Println(outerArrSpec)
	fmt.Println(outerHashSpec)
	outerArrMap, err := ebpf.NewMap(outerArrSpec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(outerArrMap))
}

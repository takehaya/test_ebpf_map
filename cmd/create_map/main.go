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

	innerMap, err := ebpf.NewMap(innerSpec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(innerMap))

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

package main

import (
	"github.com/cilium/ebpf"
	"fmt"
)

func main(){
	inner := &ebpf.MapSpec{
		Type:       ebpf.Array,
		KeySize:    4,
		ValueSize:  4,
		MaxEntries: 3,
	}
	outerArr := &ebpf.MapSpec{
		Type:       ebpf.ArrayOfMaps,
		KeySize:    4,
		ValueSize:  4,
		MaxEntries: 1,
		InnerMap: inner,
	}
	outerHash := &ebpf.MapSpec{
		Type:       ebpf.HashOfMaps,
		KeySize:    4,
		ValueSize:  4,
		MaxEntries: 5,
		InnerMap: inner,
	}
	fmt.Println(outerArr)
	fmt.Println(outerHash)
	
}

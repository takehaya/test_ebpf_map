module github.com/takehaya/test_ebpf_map

go 1.14

require (
	github.com/cilium/ebpf v0.0.0-20200917151652-9b4cff7de01e
	github.com/kr/pretty v0.2.1
	github.com/pkg/errors v0.9.1
)

replace github.com/cilium/ebpf v0.0.0-20200917151652-9b4cff7de01e => /home/vagrant/go/src/github.com/cilium/ebpf

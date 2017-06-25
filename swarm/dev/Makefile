.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm     github.com/nekonium/go-nekonium/cmd/swarm
	go build -o bin/gnekonium github.com/nekonium/go-nekonium/cmd/gnekonium
	go build -o bin/bootnode github.com/nekonium/go-nekonium/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/nekonium/go-nekonium/swarm/...

# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gnekonium android ios gnekonium-cross evm all test clean
.PHONY: gnekonium-linux gnekonium-linux-386 gnekonium-linux-amd64 gnekonium-linux-mips64 gnekonium-linux-mips64le
.PHONY: gnekonium-linux-arm gnekonium-linux-arm-5 gnekonium-linux-arm-6 gnekonium-linux-arm-7 gnekonium-linux-arm64
.PHONY: gnekonium-darwin gnekonium-darwin-386 gnekonium-darwin-amd64
.PHONY: gnekonium-windows gnekonium-windows-386 gnekonium-windows-amd64

GOBIN = build/bin
GO ?= latest

gnekonium:
	build/env.sh go run build/ci.go install ./cmd/gnekonium
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gnekonium\" to launch gnekonium."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gnekonium.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gnekonium.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

gnekonium-cross: gnekonium-linux gnekonium-darwin gnekonium-windows gnekonium-android gnekonium-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-*

gnekonium-linux: gnekonium-linux-386 gnekonium-linux-amd64 gnekonium-linux-arm gnekonium-linux-mips64 gnekonium-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-*

gnekonium-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gnekonium
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep 386

gnekonium-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gnekonium
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep amd64

gnekonium-linux-arm: gnekonium-linux-arm-5 gnekonium-linux-arm-6 gnekonium-linux-arm-7 gnekonium-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep arm

gnekonium-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gnekonium
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep arm-5

gnekonium-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gnekonium
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep arm-6

gnekonium-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gnekonium
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep arm-7

gnekonium-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gnekonium
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep arm64

gnekonium-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gnekonium
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep mips

gnekonium-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gnekonium
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep mipsle

gnekonium-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gnekonium
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep mips64

gnekonium-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gnekonium
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-linux-* | grep mips64le

gnekonium-darwin: gnekonium-darwin-386 gnekonium-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-darwin-*

gnekonium-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gnekonium
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-darwin-* | grep 386

gnekonium-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gnekonium
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-darwin-* | grep amd64

gnekonium-windows: gnekonium-windows-386 gnekonium-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-windows-*

gnekonium-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gnekonium
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-windows-* | grep 386

gnekonium-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gnekonium
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gnekonium-windows-* | grep amd64

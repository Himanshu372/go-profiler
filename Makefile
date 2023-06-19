create-binaries:
	GOOS="darwin" go build -o go-profiler-darwin-bin ./cmd/
	GOOS="linux" go build -o go-profiler-linux-bin ./cmd/
	GOOS="windows" go build -o go-profiler-windows-bin ./cmd/
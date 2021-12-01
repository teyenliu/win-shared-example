.PHONY: dlonwindows dlonlinux slonwindows slonlinux xgo-build


dlonwindows:
	go build -ldflags -v -x -installsuffix cgo -o example.exe .

dlonlinux:
	go build -ldflags -v -x -installsuffix cgo -o example .

slonwindows:
	go build -ldflags="-extldflags '-static -lstdc++'" -o example_static.exe .

slonlinux:
	go build -ldflags="-extldflags '-static -lstdc++'" -o example_static .

xgo-build:
	xgo -targets windows/amd64 github.com/teyenliu/win-shared-example
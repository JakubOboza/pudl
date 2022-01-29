.PHONY: deps build test build-release package-release release

build:
	go build -o bin/pudl main.go


test:
	go test -v ./... -count=1

deps:
	go get

release: build-release package-release
	@echo "Release build and packaged"

build-release:
	GOOS=darwin  GOARCH=amd64 go build -o release/osx-amd64/pudl main.go
	GOOS=darwin  GOARCH=arm64 go build -o release/osx-arm64/pudl main.go
	GOOS=linux   GOARCH=amd64 go build -o release/linux-amd64/pudl main.go
	GOOS=linux   GOARCH=386   go build -o release/linux-i386/pudl main.go
	GOOS=windows GOARCH=amd64 go build -o release/win-amd64/pudl.exe main.go

package-release:
	tar -czvf release/pudl.osx-amd64.tar.gz --directory=release/osx-amd64/ pudl
	tar -czvf release/pudl.osx-arm64.tar.gz --directory=release/osx-arm64/ pudl
	tar -czvf release/pudl.linux-amd64.tar.gz --directory=release/linux-amd64/ pudl
	tar -czvf release/pudl.linux-i386.tar.gz --directory=release/linux-i386/ pudl
	zip -j release/pudl.win-amd64.zip release/win-amd64/pudl.exe
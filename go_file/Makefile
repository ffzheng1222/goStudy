export tag=v1.0

root:
	export ROOT=/data/gitProject/goStudy/go_file

build:
	echo "building go file operate binary"
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/goFile fContentModify.go  fNameReplace.go  main.go 

VERSION=1.0.0
APPNAME=delocalize
NV=$(shell glide novendor)

build-cross:
	GOOS=linux GOARCH=amd64 go build -o bin/${APPNAME}-${VERSION}/linux/amd64/${APPNAME} main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/${APPNAME}-${VERSION}/darwin/amd64/${APPNAME} main.go
	GOOS=windows GORARCH=amd64 go build -o bin/${APPNAME}-${VERSION}/windows/amd64/${APPNAME}.exe main.go

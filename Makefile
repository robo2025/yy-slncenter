# Binary name
BINARY=slncenter
VERSION=0.1
LDFLAGS='-w -s'

build:
	# build for local os
	rm -f ./${BINARY}
	go clean
	GOPATH=${GOPATH}:`pwd` go build -o ${BINARY} -ldflags ${LDFLAGS} src/*.go

docker:
	# Build for docker
	rm -f ./${BINARY}
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} src/*.go
	upx ./${BINARY}

linux:
	# Make release for linux
	rm -rf release/ && mkdir release/
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} src/*.go
	upx ./${BINARY}
	tar czvf release/${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
	rm -f ./${BINARY}

windows:
	# Make release for windows
	rm -rf release/ && mkdir release/
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe -ldflags ${LDFLAGS} src/*.go
	upx ./${BINARY}.exe
	tar czvf release/${BINARY}-win64-${VERSION}.tar.gz ./${BINARY}.exe
	rm -f ./${BINARY}.exe

# Cleans our projects: deletes binaries
clean:
	go clean
	rm -f ${BINARY}
	rm -f ${BINARY}.exe
	rm -rf release/

.PHONY: build linux docker windows clean

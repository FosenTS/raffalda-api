# Build
buildForLinux:
	GOOS=linux GOARCH=amd64 go build -o application cmd/app/main.go

buildForMacARM:
	GOOS=darwin GOARCH=arm64 go build -o application cmd/app/main.go

buildForWindows:
	set GOOS=windows \
	set	GOARCH=amd64 \
	go build -o application cmd/app/main.go

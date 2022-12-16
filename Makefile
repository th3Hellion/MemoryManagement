build:
	GO111MODULE="on" && go build -o memory

run:
	GO111MODULE="on" && go run *.go

compile:
	GOARCH=amd64 GOOS=darwin go build -o memorydarwin64
	GOARCH=arm64 GOOS=darwin go build -o memorydarwinARM64
	GOARCH=amd64 GOOS=linux go build -o memorylinux64
	GOARCH=amd64 GOOS=windows go build -o memorywindows64
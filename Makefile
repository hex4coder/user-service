watch:
	CompileDaemon -build="go build cmd/main.go" -exclude-dir=docker -command=./main

build:
	CGO_ENABLED=0 go build cmd/main.go
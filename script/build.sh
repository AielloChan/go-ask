version=${1:-1.0.0}

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/go-ask_darwin_amd64_$version main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/go-ask_linux_amd64_$version main.go
version=${1:-1.0.0}

rm -rf dist
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/go-ask_darwin_amd64_$version main.go
cp dist/go-ask_darwin_amd64_$version dist/go-ask_darwin_amd64_latest
cp dist/go-ask_darwin_amd64_$version dist/go-ask
tar zcvf dist/go-ask_$version.tar.gz dist/go-ask
rm dist/go-ask
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/go-ask_linux_amd64_$version main.go
cp dist/go-ask_linux_amd64_$version dist/go-ask_linux_amd64_latest
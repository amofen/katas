## Starter Commands 
```go
go init example.com/mymodule
go run .
go mod tidy
go mod edit --replace example.com/mymodule=../mymodule
go test 
go test -v
go build
go list -f '{{.target}}'
go env -w GOBIN=/custom/path/bin
go install
```
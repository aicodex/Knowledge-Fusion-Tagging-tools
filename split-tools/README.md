### 编译命令

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o app.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o app_mac main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app_linux main.go
```
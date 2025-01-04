# URL Store
> A simple API and, eventually, TUI to manage urls I want to share

## Languages in use
- golang
- pkl

## Building
```shell
pkl-gen-go cmd/api/config/AppConfig.pkl --base-path github.com/chaezwav/urlstore
go build -o ./bin/api ./cmd/api
```

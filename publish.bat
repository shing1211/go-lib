go version
del .\go.mod
go mod init github.com/shing1211/go-lib
go mod tidy
set GOPROXY=proxy.golang.org
go list -m github.com/shing1211/go-lib
go list -m github.com/shing1211/go-lib@latest
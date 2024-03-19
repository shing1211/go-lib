go version
del .\go.mod
go mod init github.com/shing1211/go-lib
go mod tidy
set GOPROXY=proxy.golang.org
go list -m -u github.com/shing1211/go-lib@v0.1.4

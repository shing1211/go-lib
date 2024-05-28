go version
del go.mod
del go.sum
go mod init github.com/shing1211/go-lib
go mod tidy
set GOPROXY=proxy.golang.org
git tag v0.1.18
git push origin v0.1.18
go list -m github.com/shing1211/go-lib@v0.1.18

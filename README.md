# About

[![Build status](https://github.com/rgl/example-go-windows-app/workflows/Build/badge.svg)](https://github.com/rgl/example-go-windows-app/actions?query=workflow%3ABuild)

This is a Example Go Windows App application.

## Develop

Install [Chocolatey](https://chocolatey.org/install).

Install [MSYS2](https://community.chocolatey.org/packages/msys2) and [Go](https://community.chocolatey.org/packages/go):

```batch
choco install -y msys2 --params="'/NoPath'"
choco install -y go
```

Execute the following commands in a MSYS2 `bash` session.

Build and execute the application:

```bash
go generate ./...
CGO_ENABLED=0 GOAMD64=v3 go build -trimpath -ldflags "-s -w -H=windowsgui"
go version -m example-go-windows-app.exe
pwsh -Command '(Get-Item example-go-windows-app.exe).VersionInfo | Format-List'
./example-go-windows-app.exe
```

Build a release snapshot:

```bash
make release-snapshot
```

Show the executable version information:

```
pwsh -Command '(Get-Item (Resolve-Path dist/*_amd64_v3/example-go-windows-app.exe)).VersionInfo | Format-List'
```

## References

* https://thesvg.org/icon/gopher

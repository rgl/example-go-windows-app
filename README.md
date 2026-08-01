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

Install the dependencies:

```bash
winget install --exact --id MichalTrojnara.osslsigncode
```

To use the updated `PATH` environment variable, which will now include the
newly installed applications, exit the shell session, and open a new one.

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

Verify the code signature:

```bash
osslsigncode verify \
    -in dist/*_amd64_v3/example-go-windows-app.exe \
    -CAfile example-code-signing-ca-crt.pem
osslsigncode verify \
    -in dist/*_amd64v3.msix \
    -CAfile example-code-signing-ca-crt.pem
```

Show the code signature:

**NB** The signature verification will fail when your host does not trust the
`example-code-signing` CA.

```bash
pwsh -Command 'Import-Certificate example-code-signing-ca-crt.pem -CertStoreLocation Cert:/LocalMachine/Root'
pwsh -Command 'Get-AuthenticodeSignature (Resolve-Path dist/*_amd64_v3/example-go-windows-app.exe) | Format-List'
pwsh -Command 'Get-AuthenticodeSignature (Resolve-Path dist/*_amd64v3.msix) | Format-List'
```

Show the msix app manifest:

```bash
unzip -p dist/*_amd64v3.msix AppxManifest.xml
```

## References

* https://learn.microsoft.com/en-us/windows/apps/desktop/modernize/package-identity-overview
* https://thesvg.org/icon/gopher

# go-osv-scanner
Query deps.dev for vulnerability information on dependent modules and output.

Build
----------
```console
make build
```

Command
----------
```console
Usage:
  scanner [flags]

Flags:
  -h, --help         help for scanner
  -m, --mod string   Specify the path to the go.mod file
  -v, --verbose      Output vulnerability details
```

Examples
----------
```console
$ scanner -m go.mod
Your module name: scanner
No vulnerabilities were found in github.com/edoardottt/depsdev:v0.0.8
No vulnerabilities were found in github.com/spf13/cobra:v1.8.0
No vulnerabilities were found in golang.org/x/mod:v0.14.0
Vulnerability Detection!: golang.org/x/net:v0.16.0
GHSA-4374-p667-p6c8
GHSA-qppj-fm5r-hxr3
GO-2023-2102
No vulnerabilities were found in golang.org/x/mod:v0.14.0
No vulnerabilities were found in github.com/avast/retry-go:v3.0.0+incompatible
No vulnerabilities were found in github.com/inconshreveable/mousetrap:v1.1.0
No vulnerabilities were found in github.com/spf13/pflag:v1.0.5
```

Licens
----------
This repository is under [Apache2.0 License](https://github.com/siroa/go-osv-scanner/blob/main/LICENSE). 
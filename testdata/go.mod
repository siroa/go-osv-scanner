module test_data

go 1.21.4

replace project1 => ../project1

require (
    project1 v0.0.0
	github.com/edoardottt/depsdev v0.0.8
	github.com/spf13/cobra v1.8.0
	golang.org/x/mod v0.14.0
	golang.org/x/net v0.16.0
)

require (
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

# nullz

Custom types, encoders, and decoders for integrating [msgpack](https://github.com/vmihailenco/msgpack) with generated [sqlboiler](https://github.com/volatiletech/sqlboiler) models that have nullable fields.

Hell yea, nullz :metal:


### Description

nullz is a lib of types, bundled with a cli converter tool. The types map to, and are to be used in place of, their respective sqlboiler types in your app's sqlboiler-generated model file(s). Each nullz type has an implemented `msgpack.CustomEncoder` and `msgpack.CustomDecoder` method, which are called implicitly by msgpack. Then the model file(s) are passed through the converter, invoked with `nullz`.


## Before we get started

If `$GOBIN` is not set
```bash
$ export GOBIN=$GOPATH/bin
```

If `$GO111MODULE` is not set
```bash
$ export GO111MODULE=on
```

>This package assumes your code is organized as a conventional golang [workspace](https://golang.org/doc/gopath_code.html).


## Usage

1. bring in lib as dependency where necessary in your application. 
2. run the converter on the sqlboiler model file.


Bring in as dep
```bash
$ go get github.com/parkerduckworth/nullz
```

Import where necessary
```go
import "github.com/parkerduckworth/nullz"
```

Install binary (from package directory in `$GOPATH`)
```bash
$ git clone https://github.com/parkerduckworth/nullz.git $GOPATH/src/github.com/parkerduckworth/nullz
$ go install $GOPATH/src/github.com/parkerduckworth/nullz/cmd/nullz/nullz.go
```

Test installation & Help
```bash
$ nullz --help
NAME:
   nullz - convert nullable sqlboiler models for integration with msgpak

USAGE:
   nullz [global options] command [command options] [arguments...]

VERSION:
   <version>

AUTHOR:
   parkerduckworth <duckw0rth@protonmail.com>

COMMANDS:
   convert  convert file provided as [PATH]
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```

Convert model file
```bash
$ nullz convert [PATH] > `output-filename`
```

## License

This project is licensed under the [MIT License](LICENSE)

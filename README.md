# nullz

Custom types, encoders, and decoders for integrating [msgpack](https://github.com/vmihailenco/msgpack) with generated [sqlboiler](https://github.com/volatiletech/sqlboiler) models that have nullable fields.

Hell yea, nullz :metal:


## Usage

`nullz` is just a lib of types and functions, bundled with a converter tool.

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

Install binary (from package directory)
```bash
$ cd <gopath>/nullz && go install
```

Test installation
```bash
$ nullz
NAME:
   nullz - convert nullable sqlboiler models for integration with msgpak

USAGE:
   nullz [global options] command [command options] [arguments...]

COMMANDS:
   convert  convert file provided as `PATH`
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

```

Convert model file
```bash
$ nullz convert `PATH` > <output-filename>
```

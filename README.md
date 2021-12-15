# RMDSUM
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/rmdsum/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/rmdsum?status.png)](http://godoc.org/github.com/pedroalbanese/rmdsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/rmdsum)](https://goreportcard.com/report/github.com/pedroalbanese/rmdsum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/rmdsum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/rmdsum)](https://github.com/pedroalbanese/rmdsum/releases)
### RIPEMD Recursive Hasher (ISO/IEC 10118-3:2004)
<PRE>
Usage of rmdsum:
rmdsum [-v] [-c <hash.rmd>] [-r] <file.ext>
  -c string
        Check hashsum file.
  -r    Process directories recursively.
  -v    Verbose mode. (for CHECK command)</PRE>
  
### Examples:

#### Generate hashsum list:
```sh
$ ./rmdsum [-r] "*.*" > hash.txt
```
##### Always works in binary mode. 

#### Check hashsum file:
```sh
$ ./rmdsum [-v] -c hash.txt
```
##### Exit code is always 0 in verbose mode. 

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Lab.

# RMDSUM
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/rmdsum/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gosttk/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gosttk/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/rmdsum?status.png)](http://godoc.org/github.com/pedroalbanese/rmdsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/rmdsum)](https://goreportcard.com/report/github.com/pedroalbanese/rmdsum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/rmdsum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/rmdsum)](https://github.com/pedroalbanese/rmdsum/releases)
### RIPEMD Recursive Hasher (ISO/IEC 10118-3:2004)
<PRE>
Usage of rmdsum:
rmdsum [-v] [-b N] [-c &lt;hash.ext&gt;] [-r] &lt;file.ext&gt;
  -b int
        Bits: 128, 160, 256 and 320. (default 160)
  -c string
        Check hashsum file.
  -r    Process directories recursively.
  -v    Verbose mode. (The exit code is always 0 in this mode)</PRE>
  
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
##### Copyright (c) 2020-2021 Pedro F. Albanese - ALBANESE Research Lab.

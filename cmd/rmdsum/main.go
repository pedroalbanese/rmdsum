package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pedroalbanese/go-ripemd"
)

var (
	bits      = flag.Int("b", 160, "Bits: 128, 160, 256 and 320.")
	check     = flag.String("c", "", "Check hashsum file.")
	recursive = flag.Bool("r", false, "Process directories recursively.")
	verbose   = flag.Bool("v", false, "Verbose mode. (The exit code is always 0 in this mode)")
)

func main() {
	flag.Parse()

	if (len(os.Args) < 2) || (*bits != 128 && *bits != 160 && *bits != 256 && *bits != 320) {
		fmt.Fprintln(os.Stderr, "RMDSUM Copyright (c) 2020-2021 ALBANESE Research Lab")
		fmt.Fprintln(os.Stderr, "ISO/IEC 10118-3 RIPEMD Recursive Hasher written in Go\n")
		fmt.Fprintln(os.Stderr, "Usage of", os.Args[0]+":")
		fmt.Fprintf(os.Stderr, "%s [-v] [-b N] [-c <hash.rmd>] [-r] <file.ext>\n", os.Args[0])

		flag.PrintDefaults()
		os.Exit(1)
	}

	Files := strings.Join(flag.Args(), " ")

	if Files == "-" {
		var h hash.Hash
		if *bits == 128 {
			h = ripemd.New128()
		} else if *bits == 160 {
			h = ripemd.New160()
		} else if *bits == 256 {
			h = ripemd.New256()
		} else if *bits == 320 {
			h = ripemd.New320()
		}
		io.Copy(h, os.Stdin)
		fmt.Println(hex.EncodeToString(h.Sum(nil)) + " (stdin)")
		os.Exit(0)
	}

	if *check == "" && *recursive == false {
		for _, wildcard := range flag.Args() {
			files, err := filepath.Glob(wildcard)
			if err != nil {
				log.Fatal(err)
			}
			for _, match := range files {
				var h hash.Hash
				if *bits == 128 {
					h = ripemd.New128()
				} else if *bits == 160 {
					h = ripemd.New160()
				} else if *bits == 256 {
					h = ripemd.New256()
				} else if *bits == 320 {
					h = ripemd.New320()
				}
				f, err := os.Open(match)
				if err != nil {
					log.Fatal(err)
				}
				file, err := os.Stat(match)
				if file.IsDir() {
				} else {
					if _, err := io.Copy(h, f); err != nil {
						log.Fatal(err)
					}
					fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
				}
				f.Close()
			}
		}
		os.Exit(0)
	}

	if *check == "" && *recursive == true {
		err := filepath.Walk(filepath.Dir(Files),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				file, err := os.Stat(path)
				if file.IsDir() {
				} else {
					for _, match := range flag.Args() {
						filename := filepath.Base(path)
						pattern := filepath.Base(match)
						matched, err := filepath.Match(pattern, filename)
						if err != nil {
							fmt.Println(err)
						}
						if matched {
							var h hash.Hash
							if *bits == 128 {
								h = ripemd.New128()
							} else if *bits == 160 {
								h = ripemd.New160()
							} else if *bits == 256 {
								h = ripemd.New256()
							} else if *bits == 320 {
								h = ripemd.New320()
							}
							f, err := os.Open(path)
							if err != nil {
								log.Fatal(err)
							}
							if _, err := io.Copy(h, f); err != nil {
								log.Fatal(err)
							}
							f.Close()
							fmt.Println(hex.EncodeToString(h.Sum(nil)), "*"+f.Name())
						}
					}
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	if *check != "" {
		var file io.Reader
		var err error
		if *check == "-" {
			file = os.Stdin
		} else {
			file, err = os.Open(*check)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string

		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		for _, eachline := range txtlines {
			lines := strings.Split(string(eachline), " *")
			if strings.Contains(string(eachline), " *") {
				var h hash.Hash
				if *bits == 128 {
					h = ripemd.New128()
				} else if *bits == 160 {
					h = ripemd.New160()
				} else if *bits == 256 {
					h = ripemd.New256()
				} else if *bits == 320 {
					h = ripemd.New320()
				}
				_, err := os.Stat(lines[1])
				if err == nil {
					f, err := os.Open(lines[1])
					if err != nil {
						log.Fatal(err)
					}
					io.Copy(h, f)
					f.Close()

					if *verbose {
						if hex.EncodeToString(h.Sum(nil)) == lines[0] {
							fmt.Println(lines[1]+"\t", "OK")
						} else {
							fmt.Println(lines[1]+"\t", "FAILED")
						}
					} else {
						if hex.EncodeToString(h.Sum(nil)) == lines[0] {
						} else {
							os.Exit(1)
						}
					}
				} else {
					if *verbose {
						fmt.Println(lines[1]+"\t", "Not found!")
					} else {
						os.Exit(1)
					}
				}
			}
		}
	}
}

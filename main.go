package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	targetUrl *string = flag.String("u", "", "Type of link to output")
	genType   *string = flag.String("t", "md", "Type of link to output")
	outDir    *string = flag.String("o", "", "Directory to output QR code")
	version   string
	revision  string
)

func usage() {
	format := `
                  _ _       _
  __ _  ___ _ __ | (_)_ __ | | __
 / _' |/ _ \ '_ \| | | '_ \| |/ /
| (_| |  __/ | | | | | | | |   <
 \__, |\___|_| |_|_|_|_| |_|_|\_\
 |___/   Version: %s-%s

Usage:
  genlink [flags] [values]
Flags:
	-u (required)  URL
	-t             Type of link to output
	    md:        Markdown (default)
	    html:      HTML a tag
	    html-bl:   HTML a tag with 'target="_blank"'
	    qr:        QR code image
	-o             Absolute path to directory to output QR code
	               Use this flag in combination with '-t qr'
	               Default is current directory

Author:
  michimani <michimani210@gmail.com>
`
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, version, revision))
}

func main() {
	flag.Usage = usage
	flag.Parse()

	os.Exit(run())
}

func run() int {
	res, err := Generate(*targetUrl, *genType, *outDir)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	fmt.Println(res)

	return 0
}

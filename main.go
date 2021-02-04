package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	format := `Usage:
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
	fmt.Fprintln(os.Stderr, format)
}

var targetUrl *string = flag.String("u", "", "Type of link to output")
var genType *string = flag.String("t", "md", "Type of link to output")
var outDir *string = flag.String("o", "", "Directory to output QR code")

func main() {
	flag.Usage = usage
	flag.Parse()

	os.Exit(run())
}

func run() int {
	fmt.Println("Target URL    : ", *targetUrl)
	fmt.Println("Generate Type : ", *genType)
	fmt.Print("\n\n")

	res, err := Generate(*targetUrl, *genType, *outDir)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	fmt.Println(res)

	return 0
}

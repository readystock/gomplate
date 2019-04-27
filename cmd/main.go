package main

import (
	"flag"
	"fmt"
	"github.com/readystock/gomplate"
	"log"
	"os"
	"strings"
)

var (
	templateName = flag.String("template", "", "name of the template used to generate go code")
	typeNames    = flag.String("type", "", "comma-separated list of type names; must be set")
	output       = flag.String("output", "", "output file name; default srcdir/<type>_string.go")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of gomplate:\n")
	fmt.Fprintf(os.Stderr, "\tgomplate -template [name] -type T -output [result] [directory]\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("stringer: ")
	flag.Usage = Usage
	flag.Parse()

	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	types := strings.Split(*typeNames, ",")

	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}
	gomplate.GenerateTemplate(*templateName, types, *output, args)
}

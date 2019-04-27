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
	typeName     = flag.String("type", "", "comma-separated list of type names; must be set")
	output       = flag.String("output", "", "output file name; default srcdir/<type>_generated.go")
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

	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	out := *output

	if len(out) == 0 {
		out = fmt.Sprintf("%s.%s.generated.go", strings.ToLower(*typeName), strings.ReplaceAll(strings.ToLower(*templateName), ".gt", ""))
	}

	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}
	gomplate.GenerateTemplate(*templateName, *typeName, out, args)
}

package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/go-frenchspace"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		flag.CommandLine.Output().Write([]byte("Usage of frenchsapce:\n"))
		flag.CommandLine.Output().Write([]byte("\n"))
		flag.CommandLine.Output().Write([]byte("frenchsapce input1.html input2.txt ...\n"))
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	for _, n := range flag.Args() {
		in, err := ioutil.ReadFile(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to read %q\n", n)
			os.Exit(1)
		}

		f, err := os.Create(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to create %q\n", n)
			os.Exit(1)
		}
		defer f.Close()

		if strings.HasSuffix(n, ".html") {
			err = frenchspace.Stream(f, bytes.NewReader(in))
		} else {
			_, err = f.WriteString(frenchspace.Text(string(in)))
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		fmt.Println(n)
	}
}

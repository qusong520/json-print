package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	pretty := flag.Bool("p", false, "pretty print")
	escape := flag.Bool("e", false, "quote print")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	filepath := flag.Arg(0)
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read file fail: %v\n", err)
		os.Exit(2)
	}

	jsonM := make(map[string]interface{})
	if err := json.Unmarshal(bs, &jsonM); err != nil {
		fmt.Fprintf(os.Stderr, "parse json fail: %v\n", err)
		os.Exit(3)
	}

	if *pretty {
		r, err := json.MarshalIndent(jsonM, "", "    ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "print json fail: %v\n", err)
			os.Exit(4)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(r))
		os.Exit(0)
	} else {
		r, err := json.Marshal(jsonM)
		if err != nil {
			fmt.Fprintf(os.Stderr, "print json fail: %v\n", err)
			os.Exit(5)
		}

		s := string(r)
		if *escape {
			s = strconv.Quote(s)
		}

		fmt.Fprintf(os.Stdout, "%s\n", s)
		os.Exit(0)
	}
}

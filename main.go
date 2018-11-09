package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var version = "undefined"

func main() {
	pretty := flag.Bool("p", false, "pretty print")
	escape := flag.Bool("e", false, "quote print")
	vv := flag.Bool("v", false, "print version")
	flag.Parse()

	if *vv {
		fmt.Fprintln(os.Stdout, version)
		os.Exit(0)
	}

	var readFromStdinFlag bool
	var bs []byte
	if flag.NArg() < 1 {
		inBS, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read from stdin fail: %v\n", err)
			os.Exit(2)
		}
		bs = inBS
		readFromStdinFlag = true
	} else {
		filepath := flag.Arg(0)
		fileBS, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read file fail: %v\n", err)
			os.Exit(2)
		}
		bs = fileBS
	}

	var jsonObj interface{} = make(map[string]interface{})
	if len(bs) > 0 && bs[0] == '[' {
		jsonObj = make([]map[string]interface{}, 0)
	}

	if err := json.Unmarshal(bs, &jsonObj); err != nil {
		fmt.Fprintf(os.Stderr, "parse json fail: %v\n", err)
		os.Exit(3)
	}

	if *pretty {
		r, err := json.MarshalIndent(jsonObj, "", "    ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "print json fail: %v\n", err)
			os.Exit(4)
		}
		if readFromStdinFlag {
			fmt.Fprintln(os.Stdout)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(r))
		os.Exit(0)
	} else {
		r, err := json.Marshal(jsonObj)
		if err != nil {
			fmt.Fprintf(os.Stderr, "print json fail: %v\n", err)
			os.Exit(5)
		}

		s := string(r)
		if *escape {
			s = strconv.Quote(s)
		}

		if readFromStdinFlag {
			fmt.Fprintln(os.Stdout)
		}
		fmt.Fprintf(os.Stdout, "%s\n", s)
		os.Exit(0)
	}
}

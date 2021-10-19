package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/smiyaguchi/bf/pkg/parser"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	r := os.Stdin
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	p := parser.New()
	rows, err := p.Parse(string(input))
	if err != nil {
		return err
	}
	b, err := json.Marshal(rows)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

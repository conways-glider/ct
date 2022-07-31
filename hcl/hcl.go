package hcl

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/conways-glider/ct/config"
	"github.com/conways-glider/ct/json"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	hcljson "github.com/hashicorp/hcl/json/parser"
)

func Encode(config *config.Config, in interface{}) ([]byte, error) {
	json, err := json.Encode(config, in)
	if err != nil {
		return nil, fmt.Errorf("could not encode json: %w", err)
	}

	ast, err := hcljson.Parse(json)
	if err != nil {
		return nil, fmt.Errorf("could not encode hcl: %w", err)
	}

	out := bytes.Buffer{}
	writer := bufio.NewWriter(&out)
	printer.Fprint(writer, ast)
	writer.Flush()
	return out.Bytes(), nil
}

func Decode(config *config.Config) (interface{}, error) {
	var out interface{}

	in, err := ioutil.ReadAll(config.InputReader)
	if err != nil {
		return nil, fmt.Errorf("could not read input stream: %w", err)
	}
	// f, err := hcl.ParseBytes(in)
	err = hcl.Decode(out, string(in))
	if err != nil {
		return nil, fmt.Errorf("could not decode hcl: %w", err)
	}

	fmt.Println("HCL DECODE")
	fmt.Printf("%#v\n", out)

	return out, nil
}

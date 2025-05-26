package hcl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

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
	err = printer.Fprint(writer, ast)
	if err != nil {
		return nil, fmt.Errorf("could not print hcl: %w", err)
	}

	err = writer.Flush()
	if err != nil {
		return nil, fmt.Errorf("flush err: %w", err)
	}

	return out.Bytes(), nil
}

func Decode(config *config.Config) (interface{}, error) {
	var out interface{}

	in, err := io.ReadAll(config.InputReader)
	if err != nil {
		return nil, fmt.Errorf("could not read input stream: %w", err)
	}

	err = hcl.Decode(&out, string(in))
	if err != nil {
		return nil, fmt.Errorf("could not decode hcl: %w", err)
	}

	return out, nil
}

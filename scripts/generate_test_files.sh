#!/bin/bash

set -euo pipefail

export TOML='[config]
test = "value"
number = 501
array = [1,2,3]
[config.subconfig]
float = 1.5
'

rm -rf test_resources/files/
mkdir -p test_resources/files/

echo "$TOML" | ./bin/ct -i toml -o test_resources/files/toml.toml
echo "$TOML" | ./bin/ct -i toml -o test_resources/files/toml_indent.toml --indent

echo "$TOML" | ./bin/ct -i toml -o test_resources/files/json.json
echo "$TOML" | ./bin/ct -i toml -o test_resources/files/json_indent.json --indent

echo "$TOML" | ./bin/ct -i toml -o test_resources/files/yaml.yaml
echo "$TOML" | ./bin/ct -i toml -o test_resources/files/yml.yml

echo "$TOML" | ./bin/ct -i toml -o test_resources/files/hcl.hcl

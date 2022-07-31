# ct

ct is a lightweight, fast, and simple config file transformer. Currently, it supports the following file types:

- JSON
- YAML
- TOML

## Quick Links

- [Usage](#usage)
- [Installation](#installation)
- [License](#license)
- [Contribution](#contribution)

## Usage

### Flags

```
  -e, --escape-html                 Escapes HTML (JSON only)
  -f, --force                       Force overwrite of output file
  -h, --help                        help for ct
      --indent                      Indent output (JSON & TOML only)
  -i, --input string                Input file or extension (e.g. example.toml or toml) (accepted extensions: toml, yaml, json)
  -o, --output string               Output file or extension  (e.g. example.json or json) (accepted extensions: toml, yaml, json)
  -p, --output-permissions uint32   File permissions for output file (default 644)
  -v, --version                     version for ct
```

### Examples
| Input              | Output               | Command                                         |
|--------------------|----------------------|-------------------------------------------------|
| Pipe (toml format) | Stdout (json format) | `cat example.toml \| ct -i toml -o json`        |
| File (json format) | Stdout (yaml format) | `ct -i example.json -o yaml`                    |
| Pipe (yaml format) | File (toml format)   | `cat example.yml \| ct -i yaml -o example.toml` |
| File (toml format) | File (json format)   | `ct -i example.toml -o example.json`            |

## Installation

TODO

## License

Licensed under either of

 * Apache License, Version 2.0
   ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
 * MIT license
   ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.

# ct

ct is a lightweight, fast, and simple config file transformer. Currently, it supports the following file types:

- JSON
- YAML
- TOML
- HCL (v1)

## Quick Links

- [Usage](#usage)
- [Installation](#installation)
- [License](#license)
- [Contribution](#contribution)

## Usage

### Flags

```text
  -e, --escape-html                 Escapes HTML (JSON only)
  -f, --force                       Force overwrite of output file
  -h, --help                        help for ct
      --indent                      Indent output (JSON & TOML only)
  -i, --input string                Input file or extension (e.g. example.toml or toml) (accepted extensions: toml, yaml, json, hcl)
  -o, --output string               Output file or extension  (e.g. example.json or json) (accepted extensions: toml, yaml, json, hcl)
  -p, --output-permissions uint32   File permissions for output file (default 644)
  -v, --version                     version for ct
```

### Examples

| Input              | Output               | Command                                         |
|--------------------|----------------------|-------------------------------------------------|
| Pipe (toml format) | Stdout (hcl format)  | `cat example.toml \| ct -i toml -o hcl`        |
| File (json format) | Stdout (yaml format) | `ct -i example.json -o yaml`                    |
| Pipe (yaml format) | File (toml format)   | `cat example.yml \| ct -i yaml -o example.toml` |
| File (hcl format)  | File (json format)   | `ct -i example.hcl -o example.json`            |

## Installation

### Homebrew

You can install directly with:

```bash
brew install --cask conways-glider/tap/ct
```

Or, you can tap and install with:

```bash
brew tap conways-glider/tap
brew install --cask ct
```

### Manual

Binaries for macOS, Linux, and Windows are [attached to each release](https://github.com/conways-glider/ct/releases).

### Building

To build from source:

```bash
$ git clone https://github.com/conways-glider/ct
$ cd ct
$ make build
$ ./bin/ct --version
0.0.0
```

## License

Licensed under either of

- Apache License, Version 2.0 ([LICENSE-APACHE](LICENSE-APACHE) or <http://www.apache.org/licenses/LICENSE-2.0>)

- MIT license ([LICENSE-MIT](LICENSE-MIT) or <http://opensource.org/licenses/MIT>)

at your option.

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
dual licensed as above, without any additional terms or conditions.

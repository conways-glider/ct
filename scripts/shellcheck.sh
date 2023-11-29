#!/bin/bash

set -euo pipefail

find . -type f -name '*.sh' -print0 | xargs -0 shellcheck

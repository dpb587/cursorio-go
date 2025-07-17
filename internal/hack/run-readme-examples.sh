#!/bin/bash

set -euo pipefail

cd examples

go mod tidy

go run ./line-dump <<<'A-𝄞-Clef' \
  > line-dump/readme-output.txt

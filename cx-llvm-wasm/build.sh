#!/usr/bin/env bash
# Builds cx from skycoin-cx (layout a). run.sh builds each program
# with cx build --target=llvm --llvm-linker=tinygo for wasip1.
set -eu
"$(dirname "$0")/../src/cx-tools.sh" a

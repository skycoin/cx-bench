#!/usr/bin/env bash
# Builds cx from skycoin-cx (layout a) and the pad object of the layout b
# builds. run.sh builds each program with cx build --target=llvm.
set -eu
source "$(dirname "$0")/../src/form.sh"
"$ROOT/src/cx-tools.sh" a
mkdir -p "$OUT/cx-llvm"
${CX_CLANG:-clang} -O2 -c -o "$OUT/cx-llvm/pad.o" "$ROOT/src/layout/pad.c"

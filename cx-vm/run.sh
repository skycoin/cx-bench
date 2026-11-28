#!/usr/bin/env bash
# cx --executor=vm, the flat vm, running the .cx source. The program is
# parsed and lowered in the Load phase. LAYOUT=b runs cx built with the vm
# loop moved. The size and the zero are ++ arguments, which cx passes on.
source "$(dirname "$0")/../src/form.sh"
FORM_NOT_RUN=$CX_RECURSIVE
FORM_NOT_RUN_REASON="recursive, the CX forms run its loop variant (decision D145)"
FORM_CAPS="layout"
form_cmd() { printf '%s\n' "$OUT/cx-tools/$LAYOUT/cx" --executor=vm "$ROOT/src/cx/$1.cx" "++$2" "++0"; }
form_cmd_empty() { printf '%s\n' "$OUT/cx-tools/$LAYOUT/cx" --executor=vm "$ROOT/src/cx/empty.cx"; }
form_main "$@"

#!/usr/bin/env bash
# cx.wasm, cx itself built for GOOS=js, running the .cx source on its vm
# under node. VARIANT=default runs node as it ships, turbofan with
# --no-liftoff. The size and the zero are ++ arguments, as in cx-vm.
source "$(dirname "$0")/../src/form.sh"
FORM_NOT_RUN=$CX_RECURSIVE
FORM_NOT_RUN_REASON="recursive, the CX forms run its loop variant (decision D145)"
FORM_CAPS=""
FORM_VARIANTS="default turbofan"
start() { printf '%s\n' "$(command -v node)" --no-warnings $(node_flags) "$(go env GOROOT)/lib/wasm/wasm_exec_node.js" "$OUT/cx-tools/wasm/cx.wasm" --executor=vm; }
form_cmd() {
	start
	printf '%s\n' "$ROOT/src/cx/$1.cx" "++$2" "++0"
}
form_cmd_empty() { start; echo "$ROOT/src/cx/empty.cx"; }
form_main "$@"

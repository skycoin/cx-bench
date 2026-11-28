#!/usr/bin/env bash
# cx build --target=llvm linked by TinyGo for wasm32-wasip1, under
# node:wasi. VARIANT=default runs node as it ships, turbofan with
# --no-liftoff. The size and the zero are arguments, as in cx-llvm.
source "$(dirname "$0")/../src/form.sh"
FORM_CAPS=""
FORM_VARIANTS="default turbofan"
module() {
	local source=$1 wasm=$2
	if [[ ! -f $wasm ]]; then
		mkdir -p "$(dirname "$wasm")"
		CX_TINYGO_FLAGS="-opt=2 -no-debug" "$OUT/cx-tools/a/cx" build --target=llvm --llvm-linker=tinygo --llvm-tinygo-target=wasip1 "$source" -o "$wasm" >&2 || return 1
		rm -f "$wasm.ll"
	fi
	printf '%s\n' "$(command -v node)" --no-warnings $(node_flags) "$ROOT/src/wasi/wasi.mjs" "$wasm"
}
form_cmd() {
	module "$ROOT/src/cx/$1.cx" "$OUT/cx-llvm-wasm/$1-$(cx_key "$1").wasm" || return 1
	printf '%s\n' "$2" 0
}
form_cmd_empty() { module "$ROOT/src/cx/empty.cx" "$OUT/cx-llvm-wasm/empty.wasm"; }
form_main "$@"

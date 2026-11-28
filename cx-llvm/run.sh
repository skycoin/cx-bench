#!/usr/bin/env bash
# cx build --target=llvm, a native program through clang. LAYOUT=b links
# src/layout/pad.c first. The size and the zero are arguments, so LLVM sees
# neither at build time.
source "$(dirname "$0")/../src/form.sh"
FORM_CAPS="layout"
native() {
	local source=$1 program=$2
	if [[ ! -x $program ]]; then
		mkdir -p "$(dirname "$program")"
		if [[ $LAYOUT == b ]]; then
			CXBENCH_PAD=$OUT/cx-llvm/pad.o CX_CLANG=$ROOT/cx-llvm/clang-pad.sh "$OUT/cx-tools/a/cx" build --target=llvm "$source" -o "$program" >&2 || return 1
		else
			"$OUT/cx-tools/a/cx" build --target=llvm "$source" -o "$program" >&2 || return 1
		fi
		rm -f "$program.ll"
	fi
	echo "$program"
}
form_cmd() {
	native "$ROOT/src/cx/$1.cx" "$OUT/cx-llvm/$LAYOUT/$1-$(cx_key "$1")" || return 1
	printf '%s\n' "$2" 0
}
form_cmd_empty() { native "$ROOT/src/cx/empty.cx" "$OUT/cx-llvm/$LAYOUT/empty"; }
form_main "$@"

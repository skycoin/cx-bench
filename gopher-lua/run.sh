#!/usr/bin/env bash
# gopher-lua, a Lua 5.1 interpreter in Go, running src/lua51. primes,
# factorialiter and factorialrec need 64 bit integers and are left out.
source "$(dirname "$0")/../src/form.sh"
FORM_NOT_RUN="primes factorialiter factorialrec"
FORM_NOT_RUN_REASON="no Lua 5.1 version, it needs 64 bit integers"
FORM_CAPS="markers selftime server"
form_cmd() {
	printf '%s\n' "$OUT/gopher-lua/gopher-lua" "$ROOT/src/lua51/bench.lua" "$1" "$2"
}
form_cmd_empty() { printf '%s\n' "$OUT/gopher-lua/gopher-lua" "$ROOT/src/lua51/bench.lua" empty; }
form_cmd_server() { printf '%s\n' "$OUT/gopher-lua/gopher-lua" "$ROOT/src/lua51/bench.lua" -server; }
form_main "$@"

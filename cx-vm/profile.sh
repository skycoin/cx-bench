#!/usr/bin/env bash
# Runs one program once under perf record on the cx binary, then with
# cx -debug-profile, which times each compile and run phase and profiles the
# parse only (a pprof of the vm run loop would be a small addition to cx).
# CX has no server mode, so this is a one-shot run of two passes.
# Usage: profile.sh <name|number> <size>. Writes results/profiles.
set -eu
source "$(dirname "$0")/../src/form.sh"
profile_setup "$@"
mapfile -t argv < <(bash "$(dirname "$0")/run.sh" --cmd "$PROFILE_NAME" "$PROFILE_SIZE")
profile_perf "${argv[@]}"
"${argv[0]}" -debug-profile=1000 "${argv[@]:1}" >"$PROFILE_BASE.debug-profile.txt" 2>&1
mv "${argv[0]}_parse_cpu.pprof" "$PROFILE_BASE.parse_cpu.pprof" 2>/dev/null || true
rm -f "${argv[0]}_parse_mem.pprof"
echo "wrote $PROFILE_BASE.debug-profile.txt and .parse_cpu.pprof"

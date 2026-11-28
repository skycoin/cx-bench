#!/usr/bin/env bash
# Runs one native program once under perf record. Build with
# --llvm-debug for CX line names, see skycoin-cx docs/LlvmTarget.md.
# CX has no server mode, so this is a one-shot run of two passes.
# Usage: profile.sh <name|number> <size>. Writes results/profiles.
set -eu
source "$(dirname "$0")/../src/form.sh"
profile_setup "$@"
mapfile -t argv < <(bash "$(dirname "$0")/run.sh" --cmd "$PROFILE_NAME" "$PROFILE_SIZE")
profile_perf "${argv[@]}"

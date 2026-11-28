#!/usr/bin/env bash
# Runs one program once under node --cpu-prof. CX has no server mode.
# Usage: profile.sh <name|number> <size>. Writes a .cpuprofile, which Chrome
# DevTools (Performance, load profile) or speedscope opens.
set -eu
source "$(dirname "$0")/../src/form.sh"
profile_setup "$@"
mapfile -t argv < <(bash "$(dirname "$0")/run.sh" --cmd "$PROFILE_NAME" "$PROFILE_SIZE")
dir=$(mktemp -d)
"${argv[0]}" --cpu-prof --cpu-prof-dir="$dir" "${argv[@]:1}"  >/dev/null
mv "$dir"/*.cpuprofile "$PROFILE_BASE.cpuprofile" && rmdir "$dir"
echo "wrote $PROFILE_BASE.cpuprofile"

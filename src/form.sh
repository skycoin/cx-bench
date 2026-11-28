# shellcheck shell=bash
# form.sh is sourced by every form's build.sh, run.sh and profile.sh. It
# holds the paths, the benchmark names and the run.sh options the harness
# uses, so each form only says how to build and how to start a program.
#
# A form's run.sh sets FORM_CAPS (any of markers selftime server layout),
# optionally FORM_VARIANTS and FORM_NOT_RUN (benchmarks it leaves out, with
# FORM_NOT_RUN_REASON), defines form_cmd <name> <size>, form_cmd_empty
# and, with server in FORM_CAPS, form_cmd_server, then calls form_main.
# Each prints the command to start, one argument per line.

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
OUT=${CXBENCH_OUT:-$ROOT/out}
LAYOUT=${LAYOUT:-a}
BENCH_NAMES=(count fib array struct memtest primes ackermann digitalroot factorialiter factorialrec structs structserialize fibiter ackermanniter)

# CX_RECURSIVE are the recursive benchmarks. The CX forms run only their loop
# variants, fibiter, ackermanniter and factorialiter (decision D145).
CX_RECURSIVE="fib ackermann factorialrec"

# bench_name turns a benchmark number or a name into the name.
bench_name() {
	if [[ $1 =~ ^[0-9]+$ ]]; then
		(($1 >= 1 && $1 <= ${#BENCH_NAMES[@]})) || return 1
		echo "${BENCH_NAMES[$1 - 1]}"
		return
	fi
	local name
	for name in "${BENCH_NAMES[@]}"; do
		[[ $name == "$1" ]] && echo "$name" && return
	done
	return 1
}

# cx_key prints a hash of the CX program $1 and the cx revision, which names
# what a form builds from it, so that is rebuilt when either changes.
cx_key() {
	cat "$ROOT/src/cx/$1.cx" "$OUT/cx-tools/rev.txt" 2>/dev/null | sha256sum | cut -c1-12
}

# node_flags prints the node flags of VARIANT: default, or turbofan for
# TurboFan only (no Liftoff baseline compiler).
node_flags() {
	case ${VARIANT:-default} in
	turbofan) echo "--no-liftoff" ;;
	esac
}

form_has() { [[ " $FORM_CAPS " == *" $1 "* ]]; }

# form_runs fails with "not run: <reason>" on stderr for a benchmark in
# FORM_NOT_RUN. The harness reports that benchmark as skipped, not failed.
form_runs() {
	if [[ " ${FORM_NOT_RUN:-} " == *" $1 "* ]]; then
		echo "not run: $1, ${FORM_NOT_RUN_REASON:-left out by this form}" >&2
		return 1
	fi
}

form_main() {
	local argv
	case ${1:-} in
	--caps) echo "$FORM_CAPS" ;;
	--variants) echo "${FORM_VARIANTS:-default}" ;;
	--cmd)
		if [[ ${2:-} == empty ]]; then
			form_cmd_empty
		else
			local name
			name=$(bench_name "${2:-}") || { echo "unknown benchmark ${2:-}" >&2; exit 2; }
			form_runs "$name" || exit 1
			form_cmd "$name" "${3:?size}"
		fi
		;;
	--cmd-server)
		form_has server || exit 3
		form_cmd_server
		;;
	-server)
		form_has server || { echo "no server mode in this form" >&2; exit 3; }
		mapfile -t argv < <(form_cmd_server)
		exec "${argv[@]}"
		;;
	empty)
		mapfile -t argv < <(form_cmd_empty)
		exec "${argv[@]}"
		;;
	"" | -h | --help)
		echo "usage: run.sh <name|number> <size> [-time=ms|us] | empty | -server | --caps | --variants | --cmd <name|number|empty> [size] | --cmd-server" >&2
		exit 2
		;;
	*)
		local name
		name=$(bench_name "$1") || { echo "unknown benchmark $1" >&2; exit 2; }
		form_runs "$name" || exit 1
		mapfile -t argv < <(form_cmd "$name" "${2:?size}")
		((${#argv[@]})) || exit 1
		if [[ -n ${3:-} ]] && form_has selftime; then
			argv+=("$3")
		fi
		exec "${argv[@]}"
		;;
	esac
}

# profile_setup reads the arguments of a profile.sh, <name|number> <size>,
# and sets PROFILE_NAME, PROFILE_SIZE, PROFILE_INPUT (server mode lines that
# run the benchmark twice, then 0) and PROFILE_BASE, the path without the
# extension under PROFILE_DIR (results/profiles by default).
profile_setup() {
	PROFILE_NAME=$(bench_name "${1:-}") || { echo "usage: profile.sh <name|number> <size>" >&2; exit 2; }
	PROFILE_SIZE=${2:?usage: profile.sh <name|number> <size>}
	local number dir
	for number in "${!BENCH_NAMES[@]}"; do
		[[ ${BENCH_NAMES[$number]} == "$PROFILE_NAME" ]] && break
	done
	PROFILE_INPUT=$(printf '%d %d us\n%d %d us\n0\n' $((number + 1)) "$PROFILE_SIZE" $((number + 1)) "$PROFILE_SIZE")
	dir=${PROFILE_DIR:-$ROOT/results/profiles}
	mkdir -p "$dir"
	PROFILE_BASE=$dir/$(basename "$(cd "$(dirname "$0")" && pwd)")${VARIANT:+-$VARIANT}-$PROFILE_NAME-$PROFILE_SIZE
}

# profile_perf records the command after it with perf, user space only with
# the cpu-clock software event, which needs no hardware counters and works
# at kernel.perf_event_paranoid=2, and writes the flat report beside the
# data. Call graphs need a lower paranoid setting, then PERF_FLAGS=-g adds
# them. The program's stdin is PROFILE_INPUT.
profile_perf() {
	# shellcheck disable=SC2086
	perf record -q -e cpu-clock:u -F 2000 ${PERF_FLAGS:-} -o "$PROFILE_BASE.perf.data" -- "$@" <<<"$PROFILE_INPUT" >/dev/null
	perf report -i "$PROFILE_BASE.perf.data" --stdio --no-children 2>/dev/null | grep -v '^$' | head -80 >"$PROFILE_BASE.perf.txt"
	echo "wrote $PROFILE_BASE.perf.data and .perf.txt, view with perf report -i $PROFILE_BASE.perf.data"
}

#!/bin/bash

scriptDir=$(dirname "$(readlink -f "$0")")
rootDir=${scriptDir}/..

sugolver="${rootDir}/sugolver"
[ ! -f "${sugolver}" ] && exit 1

NR=1000 # number of runs

resPerf=$(mktemp -t "perf.XXXXXX")
resStats=$(mktemp -t "stats.XXXXXX")

getCycles() {
    grep cycles "${resPerf}" |
        cat -e |
        sed 's/M-.//g' |
        sed 's/@//g' |
        sed 's/cycles.*//' |
        sed 's/ //g'
}

getTime() {
    grep time "${resPerf}" |
        sed 's/seconds.*//' |
        sed 's/ //g'
}

getExclusivity() {
    grep exclusivity "${resStats}" |
        cut -d':' -f2 |
        sed 's/ //g'
}

getUniqueness() {
    grep uniqueness "${resStats}" |
        cut -d':' -f2 |
        sed 's/ //g'
}

getParity() {
    grep parity "${resStats}" |
        cut -d':' -f2 |
        sed 's/ //g'
}

getBacktracking() {
    grep backtracking "${resStats}" |
        cut -d':' -f2 |
        sed 's/ //g'
}

doSugolver() {
    local cmd="$1"
    local opts="$2"

    local cmdOpts=""
    [[ "${opts}" =~ "E" ]] && cmdOpts="${cmdOpts} --exclusivity"
    [[ "${opts}" =~ "U" ]] && cmdOpts="${cmdOpts} --uniqueness"
    [[ "${opts}" =~ "P" ]] && cmdOpts="${cmdOpts} --parity"

    ${cmd} ${cmdOpts} --stats >"${resStats}"
    perf stat -e cycles -r ${NR} ${cmd} ${cmdOpts} >"${resPerf}" 2>&1
}

dumpCSV() {
    local name="$1"
    local opts="$2"
    local difficulty="$3"

    echo -n "$(basename "${name}");${difficulty};"
    [[ "${opts}" =~ "E" ]] && echo -n "true;" || echo -n "false;"
    [[ "${opts}" =~ "U" ]] && echo -n "true;" || echo -n "false;"
    [[ "${opts}" =~ "P" ]] && echo -n "true;" || echo -n "false;"
    echo -n "$(getCycles);$(getTime);"
    echo -n "$(getExclusivity);$(getUniqueness);$(getParity);$(getBacktracking)"
    echo ""
}

doPerf() {
    local grid="$1"
    local difficulty="$2"

    grid="$(head -n1 "${grid}")"
    local cmd="${sugolver} --grid ${grid}"

    # exclusivity uniqueness parity
    for opts in "" "E" "U" "P" "EUP"; do
        doSugolver "${cmd}" "${opts}"
        dumpCSV "$1" "${opts}" "${difficulty}"
    done
}

echo "name;difficulty;E;U;P;opt;cycles;time;nrE;nrU;nrP;nrB"

difficulties="simple easy intermediate expert"

for d in ${difficulties}; do
    while read -r grid; do
        doPerf "${grid}" "${d}"
    done < <(find "${rootDir}/test/grids/" -name "${d}*grid")
done

exit 0

#!/bin/bash

SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
ROOT_DIR=${SCRIPT_DIR}/..

SUGOLVER=${ROOT_DIR}/sugolver

declare -A col=(
    [reset]="\033[0m"
    [red]="\033[0;31m"
    [green]="\033[0;32m"
    [blue]="\033[0;34m"
    [cyan]="\033[0;36m"
)

pushd "${ROOT_DIR}" || exit 1

[ -f "${SUGOLVER}" ] && rm "${SUGOLVER}"
go build
[ ! -f "${SUGOLVER}" ] &&
    echo "Error command: go build" 2>&1 &&
    exit 1

DIFFICULTIES="simple easy intermediate expert"

NR_OK=0
NR_KO=0

do_sugolver() {
    grid="$1"
    expected_solution="$2"
    opt="$3"

    [ "${opt}" == "" ] && opt="no option"
    echo -n "  - ${opt}"

    align=$((30 - ${#opt}))
    for _ in $(seq 1 ${align}); do echo -n " "; done

    solution=$(${SUGOLVER} --grid "${grid}" --dump=solution "${opt}")

    if [ "${solution}" != "${expected_solution}" ]; then
        echo -e "${col[red]}[KO]${col[reset]}"
        NR_KO=$((NR_KO + 1))
        echo "Expected   : ${expected_solution}"
        echo "Instead of : ${solution}"
    else
        echo -e "${col[green]}[OK]${col[reset]}"
        NR_OK=$((NR_OK + 1))
    fi
}

for d in ${DIFFICULTIES}; do
    while read -r line; do
        echo -e "+ Test[${col[cyan]}$(basename "${line}")${col[reset]}]"
        grid=$(head -n1 "${line}")
        expected_solution=$(tail -n1 "${line}")

        do_sugolver "${grid}" "${expected_solution}" ""
        for opt in exclusivity uniqueness parity; do
            do_sugolver "${grid}" "${expected_solution}" "${opt}"
        done
    done < <(find "${SCRIPT_DIR}/grids/" -name "${d}*grid")
done

popd || exit 1

echo -e "\n${col[blue]}[[[ Results ]]]"
echo "Test OK: ${NR_OK}"
echo "Test KO: ${NR_KO}"
echo "${col[reset]}"

[ ${NR_KO} -eq 0 ] && exit 0 || exit 1

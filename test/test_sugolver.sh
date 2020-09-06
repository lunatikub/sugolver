#!/bin/bash

scriptDir=$(dirname "$(readlink -f "$0")")
rootDir=${scriptDir}/..

sugolver=${rootDir}/sugolver

declare -A col=(
    [reset]="\033[0m"
    [red]="\033[0;31m"
    [green]="\033[0;32m"
    [blue]="\033[0;34m"
    [cyan]="\033[0;36m"
)

pushd "${rootDir}" || exit 1

[ -f "${sugolver}" ] && rm "${sugolver}"
go build
[ ! -f "${sugolver}" ] &&
    echo "Error command: go build" 2>&1 &&
    exit 1

difficulties="simple easy intermediate expert"

nrOK=0
nrKO=0

testGrid() {
    grid="$1"
    expected_solution="$2"
    opt="$3"

    [ "${opt}" == "" ] && opt="no option"
    echo -n "  - ${opt}"

    align=$((30 - ${#opt}))
    for _ in $(seq 1 ${align}); do echo -n " "; done

    solution=$(${sugolver} --grid "${grid}" --dump=solution "${opt}")

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

for d in ${difficulties}; do
    while read -r line; do
        echo -e "+ Test[${col[cyan]}$(basename "${line}")${col[reset]}]"
        grid=$(head -n1 "${line}")
        expected_solution=$(tail -n1 "${line}")

        testGrid "${grid}" "${expected_solution}" ""
        for opt in exclusivity uniqueness parity; do
            testGrid "${grid}" "${expected_solution}" "${opt}"
        done
    done < <(find "${scriptDir}/grids/" -name "${d}*grid")
done

popd || exit 1

echo -e "\n${col[blue]}[[[ Results ]]]"
echo -e "Test OK: ${nrOK}"
echo -e "Test KO: ${nrKO}${col[reset]}"

[ ${nrKO} -eq 0 ] && exit 0 || exit 1

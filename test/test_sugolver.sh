#!/bin/bash

SCRIPT_NAME=./$(basename $0)
SCRIPT_DIR=$(dirname $(readlink -f $0))
ROOT_DIR=${SCRIPT_DIR}/..

SUGOLVER=${ROOT_DIR}/sugolver

declare -A col=( \
        [reset]="\033[0m"		\
        [red]="\033[0;31m"	 	\
        [green]="\033[0;32m"	\
        [blue]="\033[0;34m"		\
        [cyan]="\033[0;36m"	 	
)

pushd ${ROOT_DIR}

[ -f ${SUGOLVER} ] && rm ${SUGOLVER}
go build
[ ! -f ${SUGOLVER} ] &&
    echo "Error command: go build" 2>&1 &&
    exit 1

DIFFICULTIES="simple easy intermediate expert"

NR_OK=0
NR_KO=0

for d in ${DIFFICULTIES}
do
    printf "[difficulty] ${col[blue]}${d}${col[reset]}\n"
    for f in $(find ${SCRIPT_DIR}/grids/ -name "${d}*grid")
    do
        printf "  [file] ${col[cyan]}$(basename ${f}) ${col[reset]}"
        grid=$(head -n1 ${f})
        expected_solution=$(tail -n1 ${f})
        solution=$(${SUGOLVER} --grid ${grid} --solution=true)
        if [ "${solution}" = "${expected_solution}" ]
        then
            printf "${colr[green]}[OK]${col[reset]}\n"
            NR_OK=$((NR_OK + 1))
        else
            printf "${col[red]}[KO]${col[reset]}\n"
            NR_KO=$((NR_KO + 1))
            printf "Expected   : ${expected_solution}\n"
            printf "Instead of : ${solution}\n"
            exit 1
        fi
    done
done

popd

echo "Test OK: ${NR_OK}"
echo "Test KO: ${NR_KO}"

[ ${NR_KO} -eq  0 ] && exit 0 || exit 1

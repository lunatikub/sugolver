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

do_sugolver()
{
    difficulty="$1"
    grid="$2"
    expected_solution="$3"
    opt="$4"

    [ "${opt}" == "" ] && opt="no option"
    printf "  - ${opt}"

    align=$((30 - ${#opt}))
    for i in $(seq 1 ${align}); do echo -n " "; done

    solution=$(${SUGOLVER} --grid ${grid} --dump=solution ${opt})

    if [ "${solution}" != "${expected_solution}" ]
    then
        printf "${col[red]}[KO]${col[reset]}\n"
        NR_KO=$((NR_KO + 1))
        printf "Expected   : ${expected_solution}\n"
        printf "Instead of : ${solution}\n"
    else
        printf "${col[green]}[OK]${col[reset]}\n"
        NR_OK=$((NR_OK + 1))
    fi
}

for d in ${DIFFICULTIES}
do
    for f in $(find ${SCRIPT_DIR}/grids/ -name "${d}*grid")
    do
        printf "+ Test[${col[cyan]}$(basename ${f})${col[reset]}]\n"
        grid=$(head -n1 ${f})
        expected_solution=$(tail -n1 ${f})  

        do_sugolver "${d}" "${grid}" "${expected_solution}" ""
        for opt in exclusivity uniqueness
        do
            do_sugolver "${d}" "${grid}" "${expected_solution}" "${opt}"
        done
    done
done

popd

printf "\n${col[blue]}[[[ Results ]]]\n"
echo "Test OK: ${NR_OK}"
echo "Test KO: ${NR_KO}"
printf "${col[reset]}"

[ ${NR_KO} -eq  0 ] && exit 0 || exit 1

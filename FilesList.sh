#!/bin/bash

Directory=$1 

declare -a local FileList
declare -i local index=0

cd "$Directory"
shift #Change Directory
#$ ~/$Directory 

for f in $(ls $pwd 2> /dev/null); do
    if [ ! -d $f ]; then
        RequiredFile=$(readlink -f $f)
        #if [[ $(file $f | awk '{print substr($0, 5)}') = "ASCII text, with CRLF, LF line terminators" ]]; then
        if [ ! $f == *.txt ]; then
            FileList[$index]=$RequiredFile
            let index+=1
       fi
    fi
done

echo ${FileList[@]}

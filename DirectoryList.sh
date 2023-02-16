#!/bin/bash

declare -a local DirectoryList
declare -i local index=0

for directory in $(ls -R ./enron_mail_20110402/ | grep /); do 
    len=$(echo "$directory" | wc -c); let len=$(($len-2));
    directory=$(echo "$directory" | cut -c1-$len)
    DirectoryList[$index]=$directory
    let index+=1
done

echo ${DirectoryList[@]}

#for file in $(ls -R); do 
#    if [ $(echo $file | cut -c1) = "." ]; then
#	len=$(echo "$file" | wc -c); let len=$(($len-2));
#	file=$(echo "$file" | cut -c1-$len)
#        DirectoryList[$index]=$file
#        let index+=1
#    fi
#done

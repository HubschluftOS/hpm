#!/bin/bash

BIN_NAME=pem
function check_uname(){
if [[ $(uname) == "Linux" ]] 
then 
	return 0
else 
	exit
fi
}

check_uname

if [[ $1 == "build" ]]
then
	pwd
	go build -o pem
	echo -e "\n${BIN_NAME} is built!"
	mv pem bin/
	grep -rIL .
elif [[ $1 == "clean" ]]
then
	go clean
else
	echo -e "Usage: ./build <argument>\n"

	echo -e "build		build this project\nclean		clean this project"
fi

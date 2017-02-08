#!/usr/bin/env bash

PROJECT_NAME=domio_api

platform='unknown'
unamestr=`uname -s`

if [[ "$unamestr" == 'CYGWIN_NT-10.0' ]]; then
   platform='cygwin'
elif [[ "$unamestr" == 'Darwin' ]]; then
   platform='mac'
elif [[ "$unamestr" == 'Linux' ]]; then
   platform='linux'
fi

if [ $platform == "cygwin" ]
    then
        echo Running Windows version...
        /cygdrive/c/usr/local/bin/${PROJECT_NAME}_win start
fi

if [ $platform == "mac" ]
    then
        echo Running Mac verision...
        /usr/local/bin/${PROJECT_NAME}_mac start
fi
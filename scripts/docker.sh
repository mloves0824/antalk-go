#!/bin/bash

hostip=`ifconfig eth0 | grep "inet " | awk -F " " '{print $2}'`

if [ "x$hostip" == "x" ]; then
    echo "cann't resolve host ip address"
    exit 1
fi

mkdir -p log

case "$1" in
apigw)
    docker rm -f      "Antalk-A50051" &> /dev/null
    docker run --name "Antalk-A50051" -d \
         -v `realpath log`:/antalk-go/log \
         -p 50051:50051 \
     antalk-image \
     apigw
    ;;

cleanup)
    docker rm -f      "Antalk-A50051" &> /dev/null
    ;;

*)
    echo "wrong argument(s)"
    ;;

esac


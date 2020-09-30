#!/usr/bin/env sh

dirname=$(cd `dirname $0`; pwd)
${dirname}/sealmsd

if [ $1 == "-d" ]; then
    while true; do sleep 1000; done
fi

if [ $1 == "-sh" ]; then
    /bin/sh
fi

if [ $1 == "-ash" ]; then
    /bin/ash
fi

if [ $1 == "-bash" ]; then
    /bin/bash
fi

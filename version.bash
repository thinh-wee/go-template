#!/usr/bin/env bash

BINARY="./bin/application.exec"

if [ -f $BINARY ]; then
    bash -c "$BINARY --version"
else
    echo "Can not excute application. File not found ($BINARY)."
fi
#!/usr/bin/env bash

echo "- 10%  | Get building environment and complier infomation"
COMPLIER_CMD=go
COMPLIER=go_1.18.4
CGO_ENABLED=0 # 0 or 1
OS=$OSTYPE # linux or windows or darwin 
ARCH=$PROCESSOR_ARCHITECTURE # amd64 or arm
BUILD_USER=$(whoami)@$(hostname)
BUILD_DATE=$(date +%FT%T%z)
GIT_URL=LOCAL # $(git config --get remote.origin.url)
GIT_BRANCH=unknow # $(git rev-parse --abbrev-ref HEAD)
GIT_REVISION=unknow # $(git rev-parse --short HEAD)
VERSION=1.0.0.alpha

BINARY="./bin/application.exec"

echo "- 20%  | Replace OS flag"
if [ $OS=="msys" ]; then
    GOOS="windows"
else
    GOOS=$OS
fi

echo "- 30%  | Replace ARCH flag"
if [ $ARCH=="AMD64" ]; then
    GOARCH="amd64"
elif [ $ARCH=="AMD32" ]; then
    GOARCH="amd32"
elif [ $ARCH=="ARM64" ]; then
    GOARCH="arm64"
elif [ $ARCH=="ARM32" ]; then
    GOARCH="arm32"
else
    GOARCH=$ARCH
fi

echo "- 40%  | Set complier flags"
APP_FLAGS="-X app.Binary=$BINARY -X app.Language=golang -X app.LanguageVersion=$COMPLIER -X app.Repository=$GIT_URL -X app.Version=$VERSION -X app.GitBranch=$GIT_BRANCH -X app.GitRevision=$GIT_REVISION -X app.Os=$GOOS -X app.Arch=$GOARCH -X app.Build=$BUILD_USER -X app.BuildDate=$BUILD_DATE"

echo "- 50%  | Remove old binary file"
if [ -f $BINARY ]; then
    rm -rf $BINARY
fi

echo "- 60%  | Building ..."
export GOOS=$GOOS;\
    export GOARCH=$GOARCH;\
    export CGO_ENABLED=$CGO_ENABLED;\
    $COMPLIER_CMD build --ldflags "-extldflags \"-static\" $APP_FLAGS"\
        -o $BINARY\
        -trimpath application/*.go 

echo "- 99%  | Get config path"
if [ $WORKDIR="" ]; then
    CONFIG_PATH="$(pwd)/config/setting.yaml"
else
    CONFIG_PATH="$WORKDIR/config/setting.yaml"
fi

echo "- 100% | Starting ..."
echo "-----------------------------------------------------------------"
if [ -f $BINARY ]; then
    bash -c "$BINARY --environ=TEST --config=\"$CONFIG_PATH\" start"
    # bash -c "$BINARY --start-env=DEVELOP --config=\"$CONFIG_PATH\""
fi

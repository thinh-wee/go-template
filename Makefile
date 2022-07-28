--- Makefile ----
.DEFAULT_GOAL := start
.PHONY: clean install
SHELL = /usr/bin/env bash

# This how we want to name the binary output
BINARY=./bin/gomake

CGO_ENABLED := 0 # 0 or 1
GOOS := linux # linux or windows or darwin 
GOARCH := amd64 # amd64 or arm
GIN_MODE := debug # debug or release

COMPLIER_CMD := go
COMPLIER := go_1.0.18

# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION := $(shell git describe --tags)
BUILD_USER := $(shell whoami)@$(shell hostname)
BUILD_DATE :=  $(shell date +%FT%T%z)
GIT_URL := $(shell git config --get remote.origin.url)
GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

# Setup the -ldflags option for go build here, interpolate the variable values
GO_FLAG := -s -w -X app.Binary=${BINARY} -X app.Language=golang -X app.LanguageVersion=${COMPLIER} -X app.Repository=${GIT_URL} -X app.Version=${VERSION} -X app.GitBranch=${GIT_BRANCH} -X app.GitRevision=${GIT_REVISION} -X app.Os=${GOOS} -X app.Arch=${GOARCH} -X app.Build=${BUILD_USER} -X app.BuildDate=${BUILD_DATE} 
LDFLAGS := -ldflags "-extldflags \"-static\" ${GO_FLAG}"

# Builds the project
build:
	${COMPLIER_CMD} build ${LDFLAGS} -o ${BINARY} -trimpath application/*.go

# Installs our project: copies binaries
install:
	${COMPLIER_CMD} install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
    if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	${BINARY} --environ=DEVELOP start
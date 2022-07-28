#!/usr/bin/env bash

GIT_URL="https://github.com/project/repository"

echo "- Removing old config (.git)"
if [ -d .git ]; then
    rm -rf .git/
fi

echo "- Init new repository"
git init

echo "- Add files to commit"
git add .

echo "- Commit all changed"
git commit -am "First commit"

if [ $? == 0 ]; then
    echo "SUCCESS"
    echo "- Remove this script"
    if [ -f $0 ]; then
        rm -rf $0
    fi
else
    echo -e "\nFAILED\nPlease try again !"
fi
#!/usr/bin/env bash
set -eu

PROBLEM_NUMBER=$1
TITLE=$2

BRANCH_NAME=$1_$2

git checkout main
git pull
git checkout -b $BRANCH_NAME

DIR_NAME=problems/$BRANCH_NAME
mkdir $DIR_NAME
cat > $DIR_NAME/resolve.go << EOF 
package main
EOF
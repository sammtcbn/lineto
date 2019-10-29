#!/bin/bash
INSPATH=~/bin/
mkdir -p ${INSPATH} || exit 1
cp -f lineto ${INSPATH} || exit 1
chmod +x ${INSPATH}/lineto || exit 1
echo "setup ok"

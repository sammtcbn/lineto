#!/bin/bash
INSPATH=~/bin/
#INSPATH=/usr/local/bin/
mkdir -p ${INSPATH} || exit 1
cp -f lineto ${INSPATH} || exit 1
chmod +x ${INSPATH}/lineto || exit 1
echo "Install to ${INSPATH} ... ok"

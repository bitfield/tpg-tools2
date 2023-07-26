#!/bin/bash

set -e
for TOPIC in */; do
  for LISTING in ${TOPIC}[0-9+]; do
    # Update deps:
    # cd $LISTING && go get -t -u && cd -
    ./run_testscript.sh $LISTING
  done
done
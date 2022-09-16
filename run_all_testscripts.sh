#!/bin/bash

for TOPIC in */; do
  for LISTING in ${TOPIC}[0-9+]; do
    ./run_testscript.sh $LISTING
  done
done
#!/bin/bash
set -e

for f in ./tests/*
do
  echo "running ${f}"
  docker run -it shellmock-dev "${f}"
  echo "finished running ${f}"
done

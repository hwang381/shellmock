#!/bin/bash
set -e

echo "---------"
for f in ./tests/*
do
  echo "running ${f}"
  docker run -it shellmock-dev "${f}"
  echo "finished running ${f}"
  echo "---------"
done

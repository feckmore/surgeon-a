#!/bin/sh

if [ "$DEVELOPMENT" == "true" ]; then
  rerun --build github.com/feckmore/surgeon-a/api
else
  api
fi

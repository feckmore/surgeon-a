#!/bin/sh

if [ "$DEVELOPMENT" == "true" ]; then
  rerun --build git.arthrex.io/dschultz/surgeon-a/api
else
  api
fi

#!/usr/bin/env sh

if [ -z $DEVELOPMENT ]; then
    PORT=8080 go-wrapper run
else
    glide install
    gin -p 8080 run
fi

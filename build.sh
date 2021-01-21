#!/bin/bash
export PATH="$PATH:/protoc/bin"
BUFDIR=$(mktemp -d -t buf-XXXXXXXXXX)
buf build -o image.bin
protoc --descriptor_set_in=image.bin --go_out=$BUFDIR $(buf ls-files --input image.bin)
cp -r $BUFDIR/github.com/korjavin/pbuf-example/* ./
ls -la  $BUFDIR/
rm -rf $BUFDIR

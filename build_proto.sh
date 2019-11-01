#!/usr/bin/env bash

echo "Building oip proto files"
pushd ./src/oip
protoc --go_out=$GOPATH/src multipart.proto pubkey.proto signedMessage.proto txid.proto
popd

echo "Building oip5 proto files"
pushd ./src/oip5
protoc --go_out=$GOPATH/src -I=. -I=$GOPATH/src/github.com/oipwg/proto/src/oip \
  edit.proto NormalizeRecord.proto NormalizeRecord.proto oip5.proto Record.proto
popd

echo "Building oip5 template proto files"
pushd ./src/oip5/templates
protoc --go_out=$GOPATH/src  RecordTemplateProto.proto tmpl_433C2783.proto
popd

echo "Building historian proto files"
pushd ./src/historian
protoc --go_out=$GOPATH/src historian.proto
popd
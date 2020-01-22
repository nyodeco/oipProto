#!/usr/bin/env bash

echo "Building oip proto files"
pushd ./src/oip || exit
protoc --go_out="$GOPATH/src" multipart.proto pubkey.proto signedMessage.proto txid.proto
popd || exit

echo "Building oip5 proto files"
pushd ./src/oip5 || exit
protoc --go_out="$GOPATH/src" -I=. -I="$GOPATH/src/github.com/oipwg/proto/src/oip" \
  -I="$GOPATH/src/github.com/bitspill/protoPatch" \
  edit.proto NormalizeRecord.proto NormalizeRecord.proto oip5.proto Record.proto
popd || exit

echo "Building oip5 template proto files"
pushd ./src/oip5/templates || exit
protoc --go_out="$GOPATH/src" -I=. -I="$GOPATH/src/github.com/oipwg/proto/src/oip" \
  -I="$GOPATH/src/github.com/bitspill/protoPatch" RecordTemplateProto.proto tmpl_433C2783.proto \
  tmpl_6E6D471D.proto
popd || exit

echo "Building oip5 template P descriptor"
pushd ./src/oip5/templates || exit
if [[ -f p.proto ]]; then
  protoc --descriptor_set_out=p.dso -I=. -I="$GOPATH/src/github.com/oipwg/proto/src/oip" \
    -I="$GOPATH/src/github.com/bitspill/protoPatch" p.proto
else
  echo "p.proto does not exist"
fi
popd || exit

echo "Building historian proto files"
pushd ./src/historian || exit
protoc --go_out="$GOPATH/src" historian.proto
popd || exit

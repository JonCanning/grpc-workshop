#!/bin/bash
rm -rf proto
rm -rf build
mkdir proto
protoc --ts_out="./proto" --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts -I=../proto hello.proto
bun i
bun build index.ts --target node --outfile build/index.mjs
node ./build/index.mjs

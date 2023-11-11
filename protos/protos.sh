#!/bin/sh
#File Name    : protos.sh
#Author       : aico
#Mail         : 2237616014@qq.com
#Github       : https://github.com/TBBtianbaoboy
#Site         : https://www.lengyangyu520.cn
#Create Time  : 2021-12-09 14:10:21
#Description  :

set -eu
#:check script but not execute,below is example.
#bash -n main.sh

CURRENT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd ${CURRENT_DIR}

if [[ -d build ]]; then
  rm -rf build;
fi

mkdir -p build

# protoc --go_out=plugins=grpc:build ./rpcapi/*.proto
protoc -I=./rpcapi --go_out=plugins=grpc:build ./rpcapi/**/*.proto

cp -rf build/nas-common/rpcapi/* ../nas-common/rpcapi/

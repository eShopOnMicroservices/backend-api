#!/usr/bin/bash

source tools/script/source-env.bash

docker run --rm -i $PROTOC_GEN_GO_GRPC

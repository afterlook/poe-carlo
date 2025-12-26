#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

set -euo pipefail

pushd "${SCRIPT_DIR}"
docker build --build-arg REPOE_SHA="$(git ls-remote https://github.com/repoe-fork/repoe.git master | awk '{print $1}')" \
	--build-arg PYPOE_SHA="$(git ls-remote https://github.com/repoe-fork/pypoe.git master | awk '{print $1}')" \
	-t poedata:latest -f Dockerfile.poedata .
popd
docker create --name poedata poedata:latest

mkdir -p data
docker cp poedata:export/. data
docker rm poedata

go run hack/go-tools/archive/main.go archive data/base_items.min.json data/base_item.min.json.gzip

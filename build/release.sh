#!/bin/bash
build::image() {
  export CGO_ENABLED=0
  go build -o ./cmd/$1/rainbond-$1 ./cmd/$1/
}

build_items=("safety-consumer" "task-plug-producer")

build::image::all() {
  for item in "${build_items[@]}"; do
  		build::image "$item"
  	done
}

if [ "$1" = "all" ]; then
  build::image::all
else
	build::image $1
fi

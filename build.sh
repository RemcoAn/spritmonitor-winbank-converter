#!/bin/bash

script_dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
root_dir=$(realpath $script_dir)

cd $root_dir/cmd/sprit-win-convert
go install
cd $root_dir
 

# ./build.sh && sprit-win-convert.exe fuelings.csv --latest-conversion 25/6/2022
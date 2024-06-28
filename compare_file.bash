#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 file1 file2"
    exit 1
fi

file1=$1
file2=$2

# Check if both files exist
if [ ! -f "$file1" ]; then
    echo "Error: File '$file1' not found!"
    exit 1
fi

if [ ! -f "$file2" ]; then
    echo "Error: File '$file2' not found!"
    exit 1
fi

diff <(tail -n +2 "$file1") <(tail -n +2 "$file2")

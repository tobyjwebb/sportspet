#!/bin/bash

echo 'Checking Go format...'
FMT_RESULT="$(gofmt -l -s .)"
[ -z "$FMT_RESULT" ] || {
    echo "Bad formatting in these files:" >&2
    echo "$FMT_RESULT" >&2
    exit 1
}
echo 'Go format is correct'

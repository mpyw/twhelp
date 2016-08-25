#!/bin/bash

oslist=(darwin windows linux freebsd)
for os in "${oslist[@]}"; do
    [ "$os" = "windows" ] && suffix=".exe" || suffix=""
    dirname="dist/$os"
    filename="twhelp$suffix"
    zipname="../twhelp-x64$os.zip"
    (
        GOARCH=amd64 GOOS="$os" go build -ldflags="-s -w" -o "$dirname/$filename" &&
        cd "$dirname" &&
        zip "$zipname" "$filename"
    ) &
done
wait

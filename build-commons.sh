#!/bin/bash

docker run --rm -it -v "$PWD":/_ -w /_ golang:1.7 bash -c '

    go get -d ./...

    oslist=(darwin windows linux freebsd)
    for os in "${oslist[@]}"; do
        [ "$os" = "windows" ] && suffix=".exe" || suffix=""
        dirname="dist/$os"
        filename="twhelp$suffix"
        (GOARCH=amd64 GOOS="$os" go build -ldflags="-s -w" -o "$dirname/$filename") &
    done

    wait

' && {

    oslist=(darwin windows linux freebsd)
    for os in "${oslist[@]}"; do
        [ "$os" = "windows" ] && suffix=".exe" || suffix=""
        dirname="dist/$os"
        filename="twhelp$suffix"
        zipname="../twhelp-x64$os.zip"
        (cd "$dirname" && zip "$zipname" "$filename") &
    done

    wait

}

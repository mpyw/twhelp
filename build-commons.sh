#!/bin/bash

docker run --rm -it -v "$PWD":/_ -w /_ golang:1.17 bash -c '

    GO111MODULE=off go get -d ./...

    oslist=(darwin windows linux freebsd)
    for os in "${oslist[@]}"; do
        [ "$os" = windows ] && suffix=.exe || suffix=
        dirname="dist/$os/amd64"
        filename="twhelp$suffix"
        (GO111MODULE=off GOARCH=amd64 GOOS="$os" go build -ldflags="-s -w" -o "$dirname/$filename") &
        if [ "$os" = darwin ]; then
            (GO111MODULE=off GOARCH=arm64 GOOS=darwin go build -ldflags="-s -w" -o dist/darwin/arm64/twhelp) &
        fi
    done

    wait

' && {

    oslist=(darwin windows linux freebsd)
    for os in "${oslist[@]}"; do
        [ "$os" = windows ] && suffix=.exe || suffix=
        dirname="dist/$os/amd64"
        filename="twhelp$suffix"
        zipname="../../twhelp-$os-amd64-$1.zip"
        (cd "$dirname" && zip "$zipname" "$filename") &
        if [ "$os" = darwin ]; then
            (cd dist/darwin/arm64 && zip "../../twhelp-darwin-arm64-$1.zip" "twhelp") &
        fi
    done

    wait

}

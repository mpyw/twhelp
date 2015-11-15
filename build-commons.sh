#!/bin/bash

oslist=(darwin windows linux freebsd)
for os in "${oslist[@]}"; do
    [ "$os" = "windows" ] && suffix=".exe" || suffix=""
    dist="dist/twhelp-x64${os}${suffix}"
    $(GOARCH=amd64 GOOS="$os" go build -o "$dist" && zip "$dist.zip" "$dist") &
done

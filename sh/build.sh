#!/bin/bash

export GIT_SHA=$(git log -1 --format=%h)
export GIT_COMMIT=$(git log -1 --date=iso | grep "Date" | sed -E 's/Date: *([0-9-]+) ([0-9:]+) .*/\1_\2/')
echo "build DATE:$GIT_COMMIT;SHA:$GIT_SHA"
CGO_ENABLED=0 go build -ldflags "-X main.version=DATE:$GIT_COMMIT;SHA:$GIT_SHA" -trimpath

#!/bin/sh

# Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

set -e

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/runner_linux_amd64 main.go
curl -L https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz | tar xJ
upx-3.96-amd64_linux/upx bin/runner_linux_amd64 
sudo chmod 764 bin/runner_linux_amd64
rm -r upx-3.96-amd64_linux
go install github.com/GeertJohan/go.rice/rice@v1.0.2
RICEBIN="$GOBIN"
if [ -z "$RICEBIN" ]; then
  if [ -z "$GOPATH" ]; then
    RICEBIN="$HOME"/go/bin
  else
    RICEBIN="$GOPATH"/bin
  fi
fi

"$RICEBIN"/rice embed-go -i github.com/bhojpur/ara/pkg/test/runner

if [ $(ls -l bin/runner_linux_amd64 | cut -d ' ' -f 5) -gt 3437900 ]; then
    echo "runner binary is too big (> gRPC message size)"
    exit 1
fi

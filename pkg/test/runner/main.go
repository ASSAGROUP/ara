//go:build runner
// +build runner

package main

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	testsuites "github.com/bhojpur/ara/pkg/test"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <base64-encoded-spec>\n", os.Args[0])
		os.Exit(1)
	}

	buf, err := base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		fail(fmt.Errorf("cannot decode spec: %w", err))
	}

	var spec testsuites.Spec
	err = json.Unmarshal(buf, &spec)
	if err != nil {
		fail(fmt.Errorf("cannot unmarshal spec: %w", err))
	}

	executor := testsuites.LocalExecutor{}
	res, err := executor.Run(context.Background(), &spec)
	if err != nil {
		res = &testsuites.RunResult{
			Stderr:     []byte(fmt.Sprintf("cannot run command: %+q\nenv: %s\n", err, strings.Join(os.Environ(), "\n\t"))),
			StatusCode: 255,
		}
	}

	err = json.NewEncoder(os.Stdout).Encode(res)
	if err != nil {
		fail(fmt.Errorf("cannot marshal result: %w", err))
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(2)
}

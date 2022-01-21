package fancylog

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
	"sort"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

// Formatter formats log output
type Formatter struct {
	Level int
}

// DefaultIndent is the spacing for any output
const DefaultIndent = "              "

// Format renders a single log entry
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var res []byte
	for i := 0; i < f.Level; i++ {
		res = append(res, []byte("  ")...)
	}

	step, ok := entry.Data["step"]
	if ok {
		res = append(res, []byte(color.Sprintf("<fg=black;bg=white> step %02d </>", step))...)
		res = append(res, ' ')
	} else {
		res = append(res, []byte("          ")...)
	}

	emoji, ok := entry.Data["emoji"]
	if ok {
		res = append(res, []byte(emoji.(string))...)
		res = append(res, []byte("  ")...)
	} else {
		res = append(res, []byte("    ")...)
	}

	var cl *color.Theme
	switch entry.Level {
	case logrus.DebugLevel:
		cl = color.Debug
	case logrus.InfoLevel:
		cl = color.Info
	case logrus.WarnLevel:
		cl = &color.Theme{Name: "warning", Style: color.Style{color.Yellow}}
	case logrus.ErrorLevel:
		cl = color.Error
	case logrus.FatalLevel:
		cl = color.Danger
	}
	res = append(res, []byte(cl.Sprintf("%-44s", entry.Message))...)

	var keys []string
	for k := range entry.Data {
		if k == "step" || k == "emoji" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, k := range keys {
		v := entry.Data[k]

		if _, ok := v.(string); ok {
			res = append(res, []byte(color.FgDarkGray.Sprintf("%s=\"%s\" ", k, v))...)
		} else {
			res = append(res, []byte(color.FgDarkGray.Sprintf("%s=%v ", k, v))...)
		}
	}

	res = append(res, '\n')
	return res, nil
}

// Push increases the level by one
func (f *Formatter) Push() {
	f.Level++
}

// Pop decreases the level by one
func (f *Formatter) Pop() {
	f.Level--
}

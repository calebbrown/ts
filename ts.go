// Copyright 2013 Caleb Brown. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

var timeFormat string

func init() {
	defaultTimeFormat := os.Getenv("TS_TIME_FORMAT")
	if defaultTimeFormat == "" {
		defaultTimeFormat = "2006-01-02 15:04:05"
	}

	flag.StringVar(&timeFormat, "f", defaultTimeFormat, "set the time format")
}

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := time.Now()
		fmt.Printf("[%s] %s\n", t.Format(timeFormat), scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

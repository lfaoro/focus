//===----------------------------------------------------------------------===//
//
// Copyright (c) 2017 Leonardo Faoro
// Licensed under the BSD License
//
// See https://github.com/lfaoro/focus/license.md for license information
// See https://www.lfaoro.com for details about the author
//
//===----------------------------------------------------------------------===//

// Focus is an implementation of the Pomodoro technique
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/lfaoro/focus/session"
)

const (
	// Version of the program
	Version = "1.0.0"
	// Author of the program
	Author = "Leonardo Faoro"
)

var (
	delay     time.Duration
	context   = flag.String("c", "personal", "Context towards which to gather statistics.")
	focusTime = 20 * time.Minute
	breakTime = 5 * time.Minute
)

// TODO(leo): add color with disable switch
// TODO(leo): make the timer pausable by pressing p
// TODO(leo): restart the timer by pressing r
func main() {
	flag.DurationVar(&delay, "d", (5 * time.Second), "Delay the start of the timer by n amount of seconds.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v [flags...] [focus-time] [break-time]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags: \n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	if *context != "personal" {
		// save to DB with new key
	}
	// TODO(leo): Implement stats based on context
	// e.g.: 100 focus hours dedicated to "Self Improvement"
	// pom stats
	// pom stats reset
	if flag.Arg(0) == "stats" {
		fmt.Println("WARNING: Not implemented")
		fmt.Println("100 hours dedicated to \"Self Improvement\"")
		fmt.Println(" 30 hours dedicated to \"Work\"")
		fmt.Println("357 hours dedicated to \"Upwork Clients\"")
		return
	}

	f := flag.NArg()
	switch {
	case f > 2:
		flag.Usage()
	case f == 2:
		focusTime = parseTime(flag.Arg(0))
		breakTime = parseTime(flag.Arg(1))
	default:
		break
	}

	focus := session.New("Focus", focusTime, delay)
	// fmt.Println(focus.String())
	focus.Start() // Blocking

	// focusBreak := session.New("Break", breakTime, delay)
	// fmt.Println(focusBreak.String())
	// focusBreak.Start() // Blocking

}

func parseTime(flag string) time.Duration {
	i, err := strconv.ParseInt(flag, 10, 64)
	if err != nil {
		log.Fatalf("Could not parse %s argument", flag)
	}
	return time.Duration(i) * time.Minute
}

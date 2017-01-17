//===----------------------------------------------------------------------===//
//
// Copyright (c) 2017 Leonardo Faoro
// Licensed under the BSD License
//
// See https://github.com/lfaoro/focus/license.md for license information
// See https://www.lfaoro.com for details about the author
//
//===----------------------------------------------------------------------===//

package session

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/lfaoro/focus/bar"
)

// Session defines all the session components.
type Session struct {
	name     string
	duration time.Duration
	delay    time.Duration
	start    time.Time
	end      time.Time
	C        chan os.Signal
	ticker   *time.Ticker
	timer    *time.Timer
}

// New instances a new session.
func New(name string, duration time.Duration, delay time.Duration) *Session {
	s := &Session{}
	s.name = name
	s.duration = duration
	s.delay = delay
	s.start = time.Now().Add(delay)
	s.end = s.start.Add(s.duration)
	s.C = make(chan os.Signal, 1)
	return s
}

// Start runs the session.
func (s *Session) Start() {
	defer s.Stop()
	go s.restart()

	beep := sync.Once{}
	beep.Do(func() { print("\a") })

	delay := sync.Once{}
	delay.Do(func() { time.Sleep(s.delay) })

	s.ticker = time.NewTicker(time.Second)
	b := bar.New(s.name, 50, s.duration)
	go func() {
		for range s.ticker.C {
			progress := time.Now().Sub(s.start)
			fmt.Println(b.Progress(progress))
		}
	}()

	s.timer = time.NewTimer(s.duration)
	<-s.timer.C
}

// Stop stops everything that could cause a panic during restart.
func (s *Session) Stop() {
	if s.ticker == nil || s.timer == nil {
		os.Exit(1)
	}

	s.ticker.Stop()
	s.timer.Stop()

}

func (s *Session) restart() {
	signal.Notify(s.C, os.Interrupt, os.Kill)

	<-s.C

	s.Stop()

	fmt.Print("\nWould you like to restart the timer? ")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	switch input {
	case "n", "N", "no", "NO":
		os.Exit(0)
	case "y", "Y", "yes", "YES":
		s.start = time.Now()
		s.Start()
	}
}

func (s *Session) String() {
	// make it print pretty
}

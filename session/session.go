//===----------------------------------------------------------------------===//
//
// Copyright (c) 2017 Leonardo Faoro
// Licensed under the BSD License
//
// See https://github.com/lfaoro/focus/blob/master/license.md for license
// information
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
	timer    *time.Timer
	bar      *bar.Bar
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
	s.bar = bar.New(s.name, 50, s.duration)
	return s
}

// Start runs the session.
func (s *Session) Start() {
	defer s.Stop()
	go s.restart()

	delay := sync.Once{}
	delay.Do(func() { time.Sleep(s.delay) })

	beep := sync.Once{}
	beep.Do(func() { Ring(5) })

	s.bar.Start(s.start)

	s.timer = time.NewTimer(s.duration)
	<-s.timer.C
}

// Stop stops everything that could cause a panic during restart.
func (s *Session) Stop() {
	if s.bar == nil || s.timer == nil {
		os.Exit(1)
	}

	s.bar.Stop()
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

func (s *Session) String() string {
	ss := fmt.Sprintf("%s starts %+10s\n", s.name, s.start.Format("15:04:05"))
	sd := fmt.Sprintf("%s delay  %+10s\n", s.name, timeFormat(s.delay))
	se := fmt.Sprintf("%s ends   %+10s\n", s.name, s.end.Format("15:04:05"))

	return "\n" + ss + sd + se
}

func timeFormat(d time.Duration) string {
	t := time.Date(0, time.January, 0, int(d.Hours()), int(d.Minutes()),
		int(d.Seconds()), 0, time.Local)

	return t.Format("15:04:05")
}

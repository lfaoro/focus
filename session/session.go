package session

import (
	"os"
	"os/signal"
	"time"
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
	s.C = make(chan os.Signal)
	return s
}

// Start runs the session.
func (s *Session) Start() {
	defer s.Stop()
	// Once beep
	// Once smth

	s.ticker = time.NewTicker(time.Second)
	// goroutine for ticker

	s.timer = time.NewTimer(s.duration)
	<-s.timer.C

}

// Stop makes sure all goroutines have been stopped.
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
}

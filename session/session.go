package session

import (
	"os"
	"time"
)

// Session defines all the session components.
type Session struct {
	name     string
	duration time.Duration
	delay    time.Duration
	start    time.Duration
	end      time.Duration
	C        os.Signal
	ticker   time.Ticker
	timer    time.Timer
}

// New instances a new sessions.
func New() {

}

// Start runs the session in a goroutine.
func Start() {

}

// Stop makes sure all goroutines have been stopped.
func Stop() {

}

func restart() {

}

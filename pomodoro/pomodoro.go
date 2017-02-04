// Package pomodoro implements the Pomodoro technique according to the original idea.
package pomodoro

type pomodoro struct {
	focusTime int
	breakTime int
	delay     int
}

func New() *pomodoro {
	p := &pomodoro{}
	p.focusTime = 25
	p.breakTime = 5
	p.delay = 5

	return p
}

// Start will start a pomodoro session.
func Start() {

}

// Stop will stop a pomodoro session.
func Stop() {

}

// Pause will pause a pomodoro session.
func Pause() {

}

// Restart will restart a pomodoro session.
func Restart() {

}

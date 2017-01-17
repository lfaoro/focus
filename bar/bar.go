//===----------------------------------------------------------------------===//
//
// Copyright (c) 2017 Leonardo Faoro
// Licensed under the BSD License
//
// See https://github.com/lfaoro/focus/blob/master/license.md for license
// information
//
//===----------------------------------------------------------------------===//

// Package bar provides a tool for creating custom progress bars
// and display them on the screen.
package bar

import (
	"fmt"
	"strings"
	"time"
)

// Bar holds all the data needed.
type Bar struct {
	name     string
	length   float64
	duration time.Duration
	pressure float64
	progress float64
	elapsed  float64
}

// New creates a new Bar
func New(name string, length float64, duration time.Duration) *Bar {
	b := &Bar{}
	b.name = name
	b.length = length
	b.duration = duration
	b.pressure = b.length / b.duration.Minutes()
	b.progress = 0.0
	b.elapsed = 0.0
	return b
}

// Progress displays the current progress.
func (b *Bar) Progress(elapsed time.Duration) string {
	if b.pressure <= 1.0 {
		b.pressure = 1.0
	}

	b.progress = b.pressure * elapsed.Minutes()
	prc := elapsed.Minutes() / b.duration.Minutes() * 100

	dash := int(b.length) - int(b.progress)
	bar := strings.Repeat("#", int(b.progress)) + strings.Repeat("-", dash)

	elapsed = removePrecision(elapsed)
	return fmt.Sprintf("\r%s progress   %s / %s  [%s] (%.f%%)",
		b.name,
		elapsed.String(),
		b.duration.String(),
		bar,
		prc)
}

func removePrecision(t time.Duration) time.Duration {
	f := int64(t.Seconds())
	return time.Duration(f) * time.Second
}

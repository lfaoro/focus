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
	"os/exec"
	"path/filepath"
	"sort"
)

type sound struct {
	completed  int
	shortBreak int
	longBreak  int
}

// The path to macOS sounds library
const libSounds = "/System/Library/Sounds"

// Contains the path to each system library sounds.
var sounds []string

// Ring uses `afplay` to play audio files from the macOS sounds library.
func Ring(sound int) {

	bell := getSound(sound)
	_, err := os.Stat(bell)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	cmd := exec.Command("afplay", bell)
	cmd.Start()
}

func getSound(s int) string {
	if sounds != nil {
		sort.Strings(sounds)
		return sounds[s]
	}

	err := filepath.Walk(libSounds, walkFn)
	if err != nil {
		log.Fatal(err)
	}
	// Discard the first slice element
	sounds = sounds[1:]
	if s > len(sounds) {
		log.Fatal("sound not present")
	}

	sort.Strings(sounds)
	fmt.Println(sounds)

	return sounds[s]
}

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	sounds = append(sounds, path)
	return nil
}

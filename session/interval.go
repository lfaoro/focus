//WIP:
package session

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

// Experiment:
type interval struct {
	interval        time.Duration // 25min
	intervalsTarget int           // 10 intervals per day (4 focus hours)

	shortBreak     time.Duration // 5min
	longBreak      time.Duration // 15-30min
	longBreakAfter int           // 4 intervals
}

type tomlData struct {
	session sessionData
}

type sessionData struct {
	Sessions     int
	TotalSession int
}

func update() {
	path, file := getTemp("daily_sessions.toml")

	d := tomlData{}.session
	_, err := toml.DecodeFile(path, &d)
	if err != nil {
		log.Fatal(err)
	}

	// Reset daily goal if it's a new day
	if d.Date != time.Now().Format("20060902") {
		d.Date = time.Now().Format("20060902")

		d.Sessions = 0
		d.TotalSession = 0
	}

	// take a longerBreak every 4 sessions
	if d.Sessions >= 4 {
		//do stuff
	}
	// productivity goal for the day reached
	if d.TotalSession >= 10 {
		//do stuff
	}

	d.Sessions++
	d.TotalSession++

	e := toml.NewEncoder(file)
	if err := e.Encode(d); err != nil {
		log.Fatal(err)
	}
}

func getTemp(name string) (path string, file *os.File) {
	path = os.TempDir()
	path = filepath.Join(path, name)
	log.Println(path)
	if _, err := os.Stat(path); err != nil {
		file, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

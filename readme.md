# Focus
> CLI implementation of the [Pomodoro][wiki] time-management technique.

[wiki]: https://en.wikipedia.org/wiki/Pomodoro_Technique "Wiki: Pomodoro 
Technique"

## Usage
```
Usage: focus [flags...] [focus-time] [break-time]
Flags: 
  -d duration
        Delay the start of the timer by n amount of seconds. (default 5s)
  -v    Shows the program version.
```
## Example
```bash
# 5sec delay / 20min session / 5min break
focus -d 5s 20 5

Focus starts   02:49:43
Focus delay    00:00:05
Focus ends     03:09:43
Focus progress 14s / 20m0s  [--------------------------------------------------] (1%)
```
## CTRL+C
```
Would you like to restart the timer?
```

## vFuture
- [ ] Context tracking (e.g.: "108 hours dedicated to "Self Improvement")
- [ ] Customize alert bell (13 sounds)
- [ ] Pause timer
- [x] Restart timer
- [ ] Anonymous Cloud sync of your stats
- [ ] Encryption of the data sent over the network

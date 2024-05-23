package utils

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)


type Clock struct {
	Stop chan struct{}
	CountDown string
	mu sync.Mutex
	running bool
}

var currentDuration = 0

func SetDuration(duration int) {
	currentDuration = duration
}

func FormatCountdownTOTimestamp(countdown string) int {
	parts := strings.Split(countdown, ":")
	if len(parts) >= 3 {
		return 0
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0
	}
	return (hours * 3600) + (minutes * 60) + seconds
}

func FormatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func NewClock() *Clock {
	return &Clock{
		Stop: make(chan struct{}, 1),
		CountDown: "04:00:00",
		mu: sync.Mutex{},
	}
}

func StartCountdown(clock *Clock, duration int){
	clock.mu.Lock()
	defer clock.mu.Unlock()

	if clock.running {
		fmt.Println("Clock is already running")
		return
	}

	clock.running = true
    
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		clock.CountDown = FormatDuration(duration - currentDuration)
		currentDuration--
		if currentDuration < 0 {
			currentDuration = duration
		}
		select {
			case <-clock.Stop:
				return
			default:
				fmt.Println(clock.CountDown)
		}
	}
}

func StopCountdown(clock *Clock) {
	clock.Stop <- struct{}{}
}
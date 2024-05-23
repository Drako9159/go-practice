package utils

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
package mainDomain

import (
	"net/http"
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"GoBaby/ui"
	"GoBaby/cmd/web/domain/log"
	"context"
	"log/slog"
)

var duration = 14400

var clockInstance = utils.NewClock()

func GetClockInstance() *utils.Clock {
	return clockInstance
}

func GetDuration() int {
	return duration
}

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)
	utils.ParseTemplateFiles(w, "clock", clockInstance, utils.EmptyFuncMap, ui.Content, "html/pages/main/clock.tmpl.html")
}

func RestartCycle(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.RESTART_CYCLE)
	select {
		case <- clockInstance.Stop: // if channel is already closed, do nothing
		default:
			// save a log inside the database
			err := logDomain.SaveLog(utils.FormatCountdownToTimestamp(clockInstance.CountDown))
			if err != nil {
				slog.Log(context.TODO(), slog.LevelError, err.Message)
				return
			}
			
			utils.StopCountdown(clockInstance)
			clockInstance.CountDown = "04:00:00"
			utils.SetDuration(duration)
			go utils.StartCountdown(clockInstance, duration)

			w.Write([]byte("Cycle restarted"))
	}
}

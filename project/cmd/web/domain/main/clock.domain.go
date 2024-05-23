package mainDomain

import (
	"net/http"
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
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
	utils.ParseTemplateFiles(w, "clock", clockInstance, utils.EmptyFuncMap, "ui/html/pages/main/clock.tmpl.html")

}

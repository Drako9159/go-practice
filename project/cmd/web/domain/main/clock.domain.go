package mainDomain

import (
	"net/http"
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
)

var duration = 14400

var ClockInstance = utils.NewClock()

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)
	utils.ParseTemplateFiles(w, "clock", ClockInstance, utils.EmptyFuncMap, "ui/html/pages/main/clock.tmpl.html")

}

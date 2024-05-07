package mainDomain

import (
	"net/http"
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
)

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)

	utils.ParseTemplateFiles(w, "clock", utils.EmptyStruct, utils.EmptyFuncMap, "ui/html/pages/main/clock.tmlp.html")

}

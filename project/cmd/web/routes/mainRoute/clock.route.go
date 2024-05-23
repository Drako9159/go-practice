package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"GoBaby/internals/models"
	"GoBaby/cmd/web/domain/main"
    "GoBaby/internals/utils"
)


func ClockRender(){
	duration := mainDomain.GetDuration()
	clock := mainDomain.GetClockInstance()

	utils.SetDuration(duration)
	go utils.StartCountdown(clock, duration)

	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesInstance.CLOCK, mainDomain.ClockFragment)
	routes.GetMuxInstance().HandleFunc("POST "+models.RoutesInstance.RESTART_CYCLE, mainDomain.RestartCycle)
}
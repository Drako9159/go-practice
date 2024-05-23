package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"GoBaby/internals/models"
	"GoBaby/cmd/web/domain/main"
)


func ClockRender(){
	duration := mainDomain.GetDuration()
	clock := mainDomain.GetClockInstance()

	utils.SetDuration(duration)

	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesInstance.CLOCK, mainDomain.ClockFragment)
}
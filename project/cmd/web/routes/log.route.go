package routes



func LogRender() {

	mux.HandlerFunc("GET "+models.RoutesInstance.LOGS, logDomain)
}


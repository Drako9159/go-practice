package monitorDomain

import (
	"GoBaby/internals/models"
)

func AddLog(monitorLog models.MonitorLog){
	repository_adapters.AddMonitorLog(monitorLog)
}
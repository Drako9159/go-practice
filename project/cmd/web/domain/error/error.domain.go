package errorDomain

import (
	"net/http"
	"fmt"
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"GoBaby/ui"
)

func ErrorTemplate(w http.ResponseWriter, r *http.Request, errMsg *models.AppError){
	w.Header().Set("HX-Retarget", "#error")
	w.Header().Add("HX-Reswap", "outerHTML")
	file := "html/pages/error/error.html"
	context := struct {ErrorMessage string} { ErrorMessage: fmt.Sprint(errMsg)}
	utils.ParseTemplateFiles(w, "error", context, utils.EmptyFuncMap, ui.Content, file)
	
}
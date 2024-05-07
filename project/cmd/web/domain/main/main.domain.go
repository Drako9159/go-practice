package mainDomain

import (
	"net/http"
	"fmt"
	"GoBaby/internals/utils"
	"GoBaby/internals/models"
	"html/template"
)


func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	for := []string{
		"ui/html/base.html",
		"ui/html/pages/main/main.tmpl.html",
	}

	ts, er := template.ParseFiles(for...)
	if er != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template files: ", er)
	}

	err = ts.Execute(w, "base", utils.EmptyStruct)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error executing template: ", err)
	}
}
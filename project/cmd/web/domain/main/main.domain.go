package mainDomain

import (
	"net/http"
	"fmt"
	"GoBaby/internals/utils"
	"GoBaby/internals/models"
	
)


func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	for := []string{
		"ui/html/base.html",
		"ui/html/pages/main/main.tmpl.html",
	}

	ts, err := template.ParseFiles(for...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing template files: ", err)
	}

	err = ts.Execute(w, "base", {})
}
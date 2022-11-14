package handlers

import (
	"net/http"
	"text/template"

	"github.com/yoeritjuu/Computer-Store/pkg/parts"
)

func GetPartsHandler(service parts.PartsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		partsList, err := service.GetAllParts()
		if err != nil {
			panic(err)
		}

		data := parts.Data{}
		for i := range partsList {
			view := partsList[i]
			data.Items = append(data.Items, view)
		}

		tmpl, _ := template.ParseFiles("./index.html")
		tmpl.Execute(w, data)
	}
}

package handlers

import (
	"html/template"
	"net/http"

	tmpls "github.com/nronzel/xoracle/templates"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(tmpls.IndexTemplate)
	if err != nil {
		w.Header().Set("Content-Type", "text-/plain")
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		w.Header().Set("Content-Type", "text-/plain")
		http.Error(w, "problem executing template", http.StatusInternalServerError)
        return
	}
}

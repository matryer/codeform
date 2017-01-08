package editor

import (
	"html/template"
	"net/http"

	"golang.org/x/net/context"
)

var generalTemplates = []string{"layout"}

func templateHandler(name string) Handler {
	tmplFiles := make([]string, len(generalTemplates)+1)
	for i, t := range generalTemplates {
		tmplFiles[i] = "./templates/" + t + ".tpl.html"
	}
	tmplFiles[len(tmplFiles)-1] = "./templates/" + name + "/content.tpl.html"
	tmpl, err := template.ParseFiles(tmplFiles...)
	return HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		if err != nil {
			return err
		}
		if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
			return err
		}
		return nil
	})
}

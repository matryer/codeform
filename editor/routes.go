package editor

func routes(r *Router) {
	r.HandleFunc("POST", "/preview", previewHandler)
	r.HandleFunc("GET", "/default-source", defaultSourceHandler)
	r.HandleFunc("GET", "/default-template", defaultTemplateHandler)
	r.Handle("GET", "/", templateHandler("editor"))
}

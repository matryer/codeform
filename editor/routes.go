package editor

func routes(r *Router) {
	r.HandleFunc("POST", "/preview", previewHandler)
	r.Handle("GET", "/", templateHandler("editor"))
}

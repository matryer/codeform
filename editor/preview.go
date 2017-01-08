package editor

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/matryer/codeform/source"
	"github.com/matryer/codeform/tool"
	"golang.org/x/net/context"
)

type previewRequest struct {
	Template string
	Source   string
}
type previewResponse struct {
	Output string `json:"output"`
}

func previewHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var req previewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	codeSource := source.Reader("source.go", strings.NewReader(req.Source))
	templateSource := source.Reader("template.tpl", strings.NewReader(req.Template))
	j := tool.Job{
		Code:     codeSource,
		Template: templateSource,
	}
	var buf bytes.Buffer
	if err := j.Execute(&buf); err != nil {
		respond(ctx, w, r, err, http.StatusBadRequest)
		return nil
	}
	res := previewResponse{
		Output: buf.String(),
	}
	return respond(ctx, w, r, res, http.StatusOK)
}

func defaultSourceHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	http.ServeFile(w, r, "./code/default-source.go")
	return nil
}

func defaultTemplateHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Println("serving:", "./code/default-source.go")
	http.ServeFile(w, r, "./code/default-template.tpl")
	return nil
}

func respond(ctx context.Context, w http.ResponseWriter, r *http.Request, data interface{}, code int) error {
	if err, ok := data.(error); ok {
		data = struct {
			Error string `json:"error"`
		}{Error: err.Error()}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}

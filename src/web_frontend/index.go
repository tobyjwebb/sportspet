package web_frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed html
var staticFiles embed.FS

func setupHtmlHandler() {
	var staticFS = fs.FS(staticFiles)
	htmlContent, err := fs.Sub(staticFS, "html")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(htmlContent))

	// Serve static files
	http.Handle("/", fs)
}

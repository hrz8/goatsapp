package assets

import (
	"embed"
	"io"
	"net/http"
)

//go:embed "public/*"
var publicFiles embed.FS

func PublicFile(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/assets/"):]
	file, err := publicFiles.Open("public/" + filePath)
	if err != nil {
		w.Write([]byte("not found"))
		return
	}
	defer file.Close()

	ff, ok := file.(io.ReadSeeker)
	if !ok {
		w.Write([]byte("not found"))
		return
	}

	fileInfo, _ := file.Stat()
	http.ServeContent(w, r, filePath, fileInfo.ModTime(), ff)
}

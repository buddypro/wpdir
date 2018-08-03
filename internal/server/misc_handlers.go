package server

import (
	"io"
	"net/http"
	"strings"

	"github.com/wpdirectory/wpdir/internal/data"
)

// TODO: Rewrite how the index.html file is served
// Need to embed SEO related info too

func (s *Server) static() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		file, err := data.Assets.Open("/index.html")
		if err != nil {
			s.Logger.Fatalf("Failed to open index.html: %s\n", err)
		}
		defer file.Close()

		fileinfo, err := file.Stat()
		if err != nil {
			s.Logger.Fatalf("Failed to get file info index.html: %s\n", err)
		}

		filesize := fileinfo.Size()
		buffer := make([]byte, filesize)

		_, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			s.Logger.Fatalf("Failed to read index.html: %s\n", err)
		}

		// Embed Hostname into HTML, remove trailing slash
		host := strings.TrimRight(s.Config.Host, "/")
		buffer = []byte(strings.Replace(string(buffer), "%HOSTNAME%", host, 1))

		w.Write(buffer)
	}
}

func (s *Server) notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		w.WriteHeader(http.StatusNotFound)

		file, err := data.Assets.Open("/index.html")
		if err != nil {
			s.Logger.Fatalf("Failed to open index.html: %s\n", err)
		}
		defer file.Close()

		fileinfo, err := file.Stat()
		if err != nil {
			s.Logger.Fatalf("Failed to get file info index.html: %s\n", err)
		}

		filesize := fileinfo.Size()
		buffer := make([]byte, filesize)

		_, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			s.Logger.Fatalf("Failed to read index.html: %s\n", err)
		}

		w.Write(buffer)
	}
}

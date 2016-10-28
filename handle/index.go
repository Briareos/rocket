package handle

import (
	"net/http"
	"path"
	"strings"
)

const (
	buildDir  = "build/"
	indexFile = "index.html"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path[1:], buildDir) {
			http.ServeFile(w, r, path.Join("..", "..", r.URL.Path))
		} else {
			http.ServeFile(w, r, path.Join("..", "..", indexFile))
		}
	}
}

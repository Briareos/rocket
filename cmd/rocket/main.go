package main

import (
	"net/http"
	"path/filepath"

	"github.com/Briareos/rocket/container"
)

func makeGroupDays() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func main() {
	c := container.MustLoadFromPath(filepath.Join("..", "..", "config.yml"))
	c.MustWarmUp()

	c.HTTPHandler().Handle("/api/v1/groupDays", makeGroupDays())

	c.HTTPServer().ListenAndServe()
}

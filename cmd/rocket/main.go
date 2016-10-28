package main

import (
	"path/filepath"

	"github.com/Briareos/rocket/container"
)

func main() {
	c := container.MustLoadFromPath(filepath.Join("..", "..", "config.yml"))
	c.MustWarmUp()

	c.HTTPServer().ListenAndServe()
}

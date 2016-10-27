package main

import (
	"github.com/Briareos/rocket/container"
	"path/filepath"
)

func main() {
	c := container.MustLoadFromPath(filepath.Join("..", "..", "config.yml"))
	c.MustWarmUp()

	c.HTTPServer().ListenAndServe()
}

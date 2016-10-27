package main

import (
	"github.com/Briareos/rocket/container"
)

func main() {
	c := container.MustLoadFromPath("config.yml")
	c.MustWarmUp()

	c.HTTPServer().ListenAndServe()
}

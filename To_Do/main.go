package main

import (
	"github.com/imshivang03/Assignment3/To_Do/app"
	"github.com/imshivang03/Assignment3/To_Do/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}

package main

import (
	"os"

	"github.com/lzjluzijie/go-study-notes/notes"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Study Note"
	app.Usage = "Hello go!"
	app.Author = "Halulu"
	app.Version = "0.0.0"

	app.Commands = notes.GetCommands()

	app.Run(os.Args)
}

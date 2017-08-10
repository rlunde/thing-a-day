package main

import (
	"github.com/codegangsta/cli"
	"rlunde/thing-a-day/cmd"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "thing-a-day"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}


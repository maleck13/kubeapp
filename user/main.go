package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/maleck13/kubeapp/user/cmd"
)

func main() {

	app := cli.NewApp()
	app.Name = "user"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}

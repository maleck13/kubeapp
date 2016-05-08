package main

import (
	"github.com/codegangsta/cli"
	"github.com/maleck13/kubeapp/auth/cmd"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "auth"
	commands := []cli.Command{
		cmd.ServeCommand(),
	}
	app.Commands = commands
	app.Run(os.Args)

}


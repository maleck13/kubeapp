package main

import (
	"github.com/codegangsta/cli"
	"github.com/maleck13/kubeapp/user/cmd"
	"os"
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


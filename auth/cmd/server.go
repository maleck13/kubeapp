package cmd

import (
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/maleck13/kubeapp/auth/api"
	"github.com/maleck13/kubeapp/auth/config"

	"github.com/maleck13/kubeapp/auth/data"

	"net/http"
	_ "net/http/pprof"

	"github.com/maleck13/kubeapp/auth/domain"
)

var port, configPath string

func ServeCommand() cli.Command {
	return cli.Command{
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "port",
				Value:       ":3000",
				Usage:       "serves up the json data",
				Destination: &port,
			},
			cli.StringFlag{
				Name:        "config",
				Value:       "./config/config.json",
				Usage:       "config file location",
				Destination: &configPath,
			},
		},
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "start the httpgit web service",
		Action:  serve,
	}
}

func serve(context *cli.Context) {
	config.SetGlobalConfig(configPath)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	data.InitStomp(config.Conf.GetStomp())
	defer data.DestroyStomp()
	loginService := &domain.LoginService{}
	if err := loginService.StartSubscribers(); err != nil {
		logrus.Fatal("failed to set up subcribers for login service")
	}
	defer loginService.StopSubscribers()

	router := api.NewRouter()
	if config.Conf.GetPProfEnabled() {
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
	}
	logrus.Info("starting " + context.App.Name + " port " + port)
	if err := http.ListenAndServe(port, router); err != nil {
		logrus.Fatal(err)
	}
}

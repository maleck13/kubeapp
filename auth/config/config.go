package config

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"

	"io/ioutil"
	"os"
)

//use an interface to limit access to the config object to read only
type Configuration interface {
	GetPProfEnabled() bool
	GetExample() string

	GetStomp() *Stomp_config
}

type config struct {
	PProfEnabled bool   `json:"pprof_enabled"`
	Example      string `json:"example"`

	Stomp *Stomp_config `json:"stomp"`
}

type Stomp_config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Vhost    string `json:"vhost"`
}

func (c *config) GetExample() string {
	return c.Example
}

func (c *config) GetPProfEnabled() bool {
	return c.PProfEnabled
}

func (c *config) GetStomp() *Stomp_config {
	return c.Stomp
}

var Conf Configuration

func SetGlobalConfig(path string) {
	Conf = &config{}
	file, err := os.Open(path)
	if nil != err {
		logrus.Panic("failed to open config file ", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if nil != err {
		logrus.Panic("failed to read config file ", err)
		return
	}
	if err = json.Unmarshal(data, Conf); err != nil {
		logrus.Panic("failed to decode config file ", err)
		return
	}
}

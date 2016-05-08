package data

import (
	"github.com/maleck13/kubeapp/auth/config"
	"net"
	"github.com/Sirupsen/logrus"
	"github.com/maleck13/stompy"
	"time"
)

var stompClient stompy.StompClient
func InitStomp(connectionDetails *config.Stomp_config)  {
	var err error
	address := net.JoinHostPort(connectionDetails.Host,connectionDetails.Port)
	clientOpts := stompy.ClientOpts{
		Vhost:connectionDetails.Vhost,
		HostAndPort:address,
		Timeout:time.Second * 10,
		User:connectionDetails.User,
		PassCode:connectionDetails.Pass,
		Version:connectionDetails.Protocol,
	}
	stompClient =stompy.NewClient(clientOpts)
	err = stompClient.Connect()
	if nil != err{
		logrus.Fatal("failed to connect via stomp ", err)
	}
}

func DestroyStomp(){
	if nil != stompClient{
		if err := stompClient.Disconnect(); err != nil{
			logrus.Error("failed to disconnect from stomp server ", err)
		}
	}
}

func StompClient()stompy.StompClient  {
	return stompClient
}
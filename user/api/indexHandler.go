package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	"github.com/maleck13/kubeapp/user/config"
	
	"github.com/maleck13/kubeapp/user/data"
	
	
	
	"github.com/maleck13/stompy"
	"github.com/Sirupsen/logrus"
	
)

//Example route handler
func IndexHandler(rw http.ResponseWriter, req *http.Request) HttpError {
	encoder := json.NewEncoder(rw)
	resData := make(map[string]string)
	resData["example"] = config.Conf.GetExample()

	val,has := context.GetOk(req,"test")
	if has{
		resData["context"] = val.(string)
	}

	if err := encoder.Encode(resData); err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}




//example of publishing to stomp
func IndexStomp(rw http.ResponseWriter, req *http.Request)HttpError{
	resData := make(map[string]string)
	resData["example"] = config.Conf.GetExample()

	data.Subscribe("test","test",func(msg stompy.Frame){
		jsonData := make(map[string]string)
		if err := json.Unmarshal(msg.Body,&jsonData); err != nil{
			logrus.Error("failed to unmarshal msg ", err.Error())
		}
		logrus.Info("handling msg 1: ", jsonData)

	},nil)

	data.Subscribe("test","test",func(msg stompy.Frame){
		jsonData := make(map[string]string)
		if err := json.Unmarshal(msg.Body,&jsonData); err != nil{
			logrus.Error("failed to unmarshal msg ", err.Error())
		}
		logrus.Info("handling msg 2: ", jsonData)
	},nil)

	for i :=0; i < 10; i++ {
		data.Publish("test", "test",resData, nil)
	}
	return nil
}


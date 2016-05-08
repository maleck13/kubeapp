package domain

import (
	"github.com/maleck13/kubeapp/auth/data"
	"github.com/maleck13/stompy"
	"github.com/Sirupsen/logrus"
	"encoding/json"
)

type LoginService struct {
	subIds []string
}

func (ls *LoginService) StartSubscribers() error  {
	sc := data.StompClient()
	id,err := sc.Subscribe("/login/login", handleLogin,stompy.StompHeaders{},nil)
	if err != nil{
		return err
	}
	ls.subIds = append(ls.subIds,id)
	return nil
}

func (ls *LoginService)StopSubscribers()  {
	sc := data.StompClient()
	for _,id := range ls.subIds{
		if err := sc.Unsubscribe(id,stompy.StompHeaders{},nil); err != nil{
			logrus.Error("failed to unsubscribe" , err)
		}
	}
}

type LoginResponse struct {
	Message string `json:"message"`
	Code int  `json:"code"`
	SessionId string `json:"sessionId"`
}

func handleLogin(frame stompy.Frame){
	sc :=  data.StompClient()
	replyTo := frame.Headers["replyTo"]
	correlationId := frame.Headers["correllationID"]
	message := make(map[string]string)
	var response * LoginResponse
	responseHeaders := stompy.StompHeaders{}
	responseHeaders["correllationID"] = correlationId
	if err := json.Unmarshal(frame.Body,&message); err != nil{
		response = &LoginResponse{}
		logrus.Error("handleLogin: failed to decode message ", err)
		response.Code = 400
		response.Message = "failed to decode message " + err.Error()
	}

	if nil == response{
		//check credentials
		if checkCredentials(message){
			response = &LoginResponse{Message:"login successful", Code:200, SessionId:"test"}
		}else{
			response = &LoginResponse{Message:"login failed ", Code:403}
		}
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil{
		logrus.Error("handleLogin: failed to encode response message ", err)
		return
	}
	if err := sc.Publish(replyTo,"application/json",marshalResponse,responseHeaders,nil); err != nil{
		//failed to publish
	}
}

func checkCredentials(loginDetails map[string]string)bool  {
	if _,ok := loginDetails["user"]; !ok{
		return false
	}
	if _,ok := loginDetails["pass"]; !ok{
		return false
	}

	return loginDetails["pass"] == "test"
}

package middleware

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
)

type ReqLog struct {
	Method string
	Url    string
}

//example middle ware that logs incoming requests
func ExampleMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	context.Set(req, "test", "test")
	logrus.Info("request: ", &ReqLog{
		req.Method,
		req.URL.Path,
	})
	next(rw, req)

}

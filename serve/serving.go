package serve

import (
	"github.com/jacksinn/apifaker/route"
	"net/http"
	"fmt"
)

/*
|--------------------------------------------------------------------------
| Server
|--------------------------------------------------------------------------
|
| Here is where the HTTP Server is setup
|
*/
type Server struct {
	Address string
	Port    int
	Routes  route.Routes
}

func (server Server) getResponse(request *http.Request) route.Response {
	//getting our routes from JSON
	for _, route := range server.Routes {
		//if Path and Method Match
		if route.Request.Path == request.URL.Path && route.Request.Method == request.Method {
			//Respond with route response JSON
			return route.Response
		}
	}
	//else return bad response
	return route.Response{Code: 404, Body: "Not Found"}
}

func (server Server) Handle(writer http.ResponseWriter, request *http.Request) {
	response := server.getResponse(request)
	fmt.Fprintf(writer, response.Body)
}

func (server Server) Run() {
	//Handler Func Setup
	http.HandleFunc("/", server.Handle)
	//Setting up addr to listen & serve on, this feels a little crappy, cleanup s.c.s.a and s.c.s.p
	addr := fmt.Sprintf("%v:%v", server.Address, server.Port)
	panic(http.ListenAndServe(addr, nil))
}

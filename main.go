package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"net/http"
)

/*
 * Parsing of JSON
 */
type Route struct {
	Request  Request `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Response struct {
	Code int
	Body string `json:"body"`
}

/*
 * Slice of Route structs
 */
type Routes []Route

/*
 * Setting up HTTP server
 */

type Server struct {
	Address string
	Port    int
	Routes  Routes
}

func (server Server) getResponse(request *http.Request) Response {
	//getting our routes from JSON
	for _, route := range server.Routes {
		//if Path and Method Match
		if route.Request.Path == request.URL.Path && route.Request.Method == request.Method {
			//Respond with route response JSON
			return route.Response
		}
	}
	//else return bad response
	return Response{Code: 404, Body: "Not Found"}
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

/*
 * All Config
 */
type ParsedRoutes struct {
	Routes Routes `json:"routes"`
}

func main() {
	//Reading Routes FIle
	apiRoutes, err := ioutil.ReadFile("routes.json")

	//Panic at the Disco
	if err != nil {
		panic(err)
	}

	//Route Parsing Setup
	var routes ParsedRoutes

	//Grabbing JSON, storing in Config Struct which parses the JSON automagically
	json.Unmarshal(apiRoutes, &routes)

	//Server Address and Port Setup
	server := Server{Address: "127.0.0.1", Port: 8080, Routes: routes.Routes}

	//Protect and Serve
	server.Run()

}

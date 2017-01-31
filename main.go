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
type ServerConfig struct {
	Address string
	Port    int
}

type Server struct {
	Config Config
}

func (server Server) getResponse(request *http.Request) Response {
	//getting our routes from JSON
	for _, route := range server.Config.Routes {
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
	http.HandleFunc("/", server.Handle)
	addr := fmt.Sprintf("%v:%v", server.Config.Server.Address, server.Config.Server.Port)
	panic(http.ListenAndServe(addr, nil))
}

/*
 * All Config
 */
type Config struct {
	Server ServerConfig
	Routes Routes `json:"routes"`
}

func main() {
	apiRoutes, err := ioutil.ReadFile("routes.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(apiRoutes))

	serverConfig := ServerConfig{Address: "0.0.0.0", Port: 8080 }

	var config Config
	json.Unmarshal(apiRoutes, &config)

	config.Server = serverConfig
	fmt.Println(config)

	//http.HandleFunc("/", Handler)
	//http.ListenAndServe(":8080", nil)
	server := Server{config}
	server.Run()

}

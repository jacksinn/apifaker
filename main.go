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
	//Reading Routes FIle
	apiRoutes, err := ioutil.ReadFile("routes.json")

	//Panic at the Disco
	if err != nil {
		panic(err)
	}

	//Server Address and Port Setup
	serverConfig := ServerConfig{Address: "0.0.0.0", Port: 8080 }

	//General Config Setup
	var config Config

	//Grabbing JSON, storing in Config Struct which parses the JSON automagically
	json.Unmarshal(apiRoutes, &config)

	//Configuring HTTP Server
	config.Server = serverConfig

	server := Server{config}
	server.Run()

}

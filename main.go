package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Server struct {
	Config Config
}

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
	Body string `json:"body"`
}

/*
 * All Config
 */
type Config struct {
	Server ServerConfig
	Routes Routes `json:"routes"`
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
	Port	int
}

func main() {
	apiRoutes, err := ioutil.ReadFile("routes.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(apiRoutes))

	serverConfig := ServerConfig{ Address: "0.0.0.0", Port: 8080 }

	var config Config
	json.Unmarshal(apiRoutes, &config)
	config.Server = serverConfig
	fmt.Println(config)


}

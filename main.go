package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"net/http"
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

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello World, %s!", request.URL.Path[1:])
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

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)


}

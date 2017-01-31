package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

//type Server struct {
//	Config Config
//}

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

type Config struct {
	Routes Routes `json:"routes"`
}

type Routes []Route

func main() {
	apiRoutes, err := ioutil.ReadFile("apiroutes.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(apiRoutes))

	var config Config
	json.Unmarshal(apiRoutes, &config)
	fmt.Println(config)


}

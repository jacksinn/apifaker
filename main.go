package main

import (
	"io/ioutil"
	"encoding/json"
	"github.com/jacksinn/apifaker/route"
	"github.com/jacksinn/apifaker/serve"
)
/*
|--------------------------------------------------------------------------
| Main
|--------------------------------------------------------------------------
|
| The goods. Grabbing Routes file and sending off to be parsed into
| structs. HTTP Server is setup here too.
| @todo Refactor more of this to be cleaner and reusable
|
*/
func main() {
	//Reading Routes FIle
	apiRoutes, err := ioutil.ReadFile("routes.json")

	//Panic at the Disco
	if err != nil {
		panic(err)
	}

	//Route Parsing Setup
	var routes route.JSONRoutes

	//Grabbing JSON, storing in Config Struct which parses the JSON automagically
	json.Unmarshal(apiRoutes, &routes)

	//Server Address and Port Setup
	server := serve.Server{Address: "127.0.0.1", Port: 8080, Routes: routes.Routes}

	//Protect and Serve
	server.Run()

}

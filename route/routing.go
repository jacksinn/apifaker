package route

/*
|--------------------------------------------------------------------------
| Routes
|--------------------------------------------------------------------------
|
| Here is where the structure of the route is parsed
| from the routes.json file
|
*/
type JSONRoutes struct {
	Routes Routes `json:"routes"`
}

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


type Routes []Route

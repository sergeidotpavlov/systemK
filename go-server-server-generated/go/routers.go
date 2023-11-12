/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateUserV1",
		strings.ToUpper("Post"),
		"/v1/user",
		CreateUserV1,
	},

	Route{
		"DeleteUserV1",
		strings.ToUpper("Delete"),
		"/v1/user",
		DeleteUserV1,
	},

	Route{
		"GetUserV1",
		strings.ToUpper("Get"),
		"/v1/user",
		GetUserV1,
	},

	Route{
		"UpdateUserV1",
		strings.ToUpper("Put"),
		"/v1/user",
		UpdateUserV1,
	},

	Route{
		"CreateUserV1",
		strings.ToUpper("Post"),
		"/v1/user",
		CreateUserV1,
	},

	Route{
		"DeleteUserV1",
		strings.ToUpper("Delete"),
		"/v1/user",
		DeleteUserV1,
	},

	Route{
		"GetUserV1",
		strings.ToUpper("Get"),
		"/v1/user",
		GetUserV1,
	},

	Route{
		"UpdateUserV1",
		strings.ToUpper("Put"),
		"/v1/user",
		UpdateUserV1,
	},
}

package routes

import (
	"fmt"
	goji "goji.io"
	"goji.io/pat"
	"net/http"
)

func SetupRouter() *goji.Mux {
	mux := goji.NewMux()
	mux.Handle(pat.Get("/apidoc/*"), http.StripPrefix("/apidoc/", http.FileServer(http.FileSystem(http.Dir("apidoc")))))

	mux.HandleFunc(pat.Get("/hello/:name"), hello)

	return mux
}

/**
 * @api {get} /hello/:name Hello World
 * @apiName HelloWorld
 * @apiGroup Sample
 *
 * @apiParam {Name} name The name to address the greeting to
 *
 * @apiSuccess {String} Hello Name given to endpoint.
 */
func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

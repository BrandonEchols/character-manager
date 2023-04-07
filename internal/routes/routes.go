package routes

import (
	"fmt"
	goji "goji.io"
	"goji.io/pat"
	"net/http"
)

func SetupRouter() *goji.Mux {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/api/hello/:name"), hello)
	mux.Handle(pat.Get("/apidoc/*"), http.StripPrefix("/apidoc/", http.FileServer(http.FileSystem(http.Dir("apidoc")))))
	mux.Handle(pat.Get("/*"), http.FileServer(http.FileSystem(http.Dir("frontend/build"))))

	return mux
}

/**
 * @api {get} /api/hello/:name Hello World
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

// this is a very simple fastcgi web server
package main

import (
	"log"
	"net/http"

	"github.com/yookoala/gofast"
)

func main() {
	// Get fastcgi application server tcp address
	// from env FASTCGI_ADDR. Then configure
	// connection factory for the address.
	address := "127.0.0.1:9000"
	connFactory := gofast.SimpleConnFactory("tcp", address)

	// handles static assets in the assets folder
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.FileSystem(http.Dir("/Users/jasonmccallister/dev/plugins-dev/web/assets")))))

	// route all requests to relevant PHP file
	http.Handle("/", gofast.NewHandler(
		gofast.NewPHPFS("/Users/jasonmccallister/dev/plugins-dev/web")(gofast.BasicSession),
		gofast.SimpleClientFactory(connFactory, 0),
	))

	// serve at 8080 port
	log.Fatal(http.ListenAndServe(":8080", nil))
}

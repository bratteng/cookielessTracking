package main

import (
	"fmt"
    "net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/caarlos0/env"
	"github.com/satori/go.uuid"
)

type config struct {
	Port	int		`env:"PORT" envDefault:"8080"`
}

func Tracking(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u1 := uuid.Must(uuid.NewV4())
	w.Header().Set("content-type", "text/javascript")
	w.Header().Set("last-modified", "Thu, 01 Jan 1970 00:00:00 GMT")
	fmt.Fprintf(w, "var trackingID = \"%s\";\n", u1)
}

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Redirect(w, req, "/tracking.js", 301)
}

func main() {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

    router := httprouter.New()
	router.GET("/", Index)
	router.GET("/tracking.js", Tracking)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}

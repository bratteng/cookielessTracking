package main

import (
	"fmt"
    "net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/caarlos0/env"
	"github.com/matoous/go-nanoid"
)

type config struct {
	Port				int		`env:"PORT" envDefault:"8080"`
	TrackingVariable	string	`env:"TRACKING_VARIABLE" envDefault:"trackingID"`
	TrackingScript		string	`env:"TRACKING_SCRIPT" envDefault:"tracking.js"`
}

func Tracking(cfg config) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.Header().Set("content-type", "text/javascript")
		w.Header().Set("last-modified", "Thu, 01 Jan 1970 00:00:00 GMT")
		fmt.Fprintf(w, "var %s = \"%s\";\n", fmt.Sprintf("%s", cfg.TrackingVariable), gonanoid.Must())
	}
}

func Index(cfg config) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		http.Redirect(w, req, fmt.Sprintf("/%s", cfg.TrackingScript), 301)
	}
}

func main() {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

    router := httprouter.New()
	router.GET("/", Index(cfg))
	router.GET(fmt.Sprintf("/%s", cfg.TrackingScript), Tracking(cfg))

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}

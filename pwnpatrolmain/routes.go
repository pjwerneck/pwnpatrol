package pwnpatrolmain

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	r := httprouter.New()

	r.GET("/range/:prefix", RangeHandler)

	return r
}

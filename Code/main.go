package main

import (
	"net/http"

	"github.com/pluralsight/webservice/models/controllers"
)

const (
	first = iota + 6
	second
	third
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

package main

import (
	"flag"
	"log"
	"plugin"
)

func main() {

	path := flag.String("plugin", "", "Plugin to execute")
	flag.Parse()

	if *path == "" {
		log.Fatal("Path to plugin must be provided")
	}

	p, err := plugin.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("ThingsToDo")

	if err != nil {
		log.Fatal(err)
	}

	thingstodo, ok := f.(func())

	if !ok {
		log.Fatal("Couldn't find the function ThingsToDo")
	}
	thingstodo()
	log.Println("Did the thing")
}

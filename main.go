package main

import (
	"time"

	"github.com/tem-mars/repl-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient : pokeapi.NewClient(time.Hour),
	}
		

	startRepl(&cfg)

}

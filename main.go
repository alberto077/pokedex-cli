package main

import (
	"github.com/alberto077/pokedexcli/internal/pokecache"
)

func main() {
	cfg := GetCommands()
	pcache := pokecache.NewCache()

	startRepl(cfg, pcache)
}

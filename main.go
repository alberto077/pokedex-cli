package main

import (
	"time"

	"github.com/alberto077/pokedexcli/internal/pokecache"
)

func main() {
	cfg := GetCommands()
	pcache := pokecache.NewCache(5 * time.Second)

	startRepl(cfg, pcache)
}

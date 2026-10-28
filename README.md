# Pokemon cli 

## Makes requests to the real pokemon api

### Allows the user to map through differnt areas of the world and explore it to see what pokemon are inside. The catch command will allow you to throw a pokeball and you get a chance to catch it

### Getting Started:

- go build
- ./pokedexcli
- help

## Next improvements

- Optimize reap loop for proper cache clearing
- Caching of recent commands that a user and up and down arrow through
- Proper persistant display of current area you region you are in
- Better Game logic (only allow catching of the area you explore or the explore an area that you have mapped to)
- Persistant storage of data across executions
    - Flags to clear data and start new sessions
- Properly byte allign structs

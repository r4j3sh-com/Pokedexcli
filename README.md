## PokédexCLI

PokédexCLI is a command-line interface Pokédex application written in Go. It allows users to explore Pokémon locations, catch Pokémon, and manage their own Pokédex.

## Features

- Explore Pokémon locations

- Catch Pokémon with a realistic catching mechanism

- View your caught Pokémon in your Pokédex

- Interactive command-line interface

## Installation

To install PokédexCLI, make sure you have Go installed on your system, then run:

```go
go get github.com/r4j3sh-com/pokedexcli
```

## Usage

After installation, you can run the application by typing:

```go
pokedexcl
```

This will start the interactive command-line interface.

## Commands

- help: Display a list of available commands
- map: Get the next page of locations
- mapb: Get the previous page of locations
- explore <location>: Explore Pokémon in a specific location
- catch <pokemon>: Attempt to catch a Pokémon
- pokedex: View your caught Pokémon
- exit: Exit the application

## Project Structure

- main.go: Entry point of the application
- repl.go: Implements the Read-Eval-Print Loop for the CLI
- command_*.go: Individual command implementations
- internal/pokeapi/: Package for interacting with the PokéAPI
- internal/pokecache/: Package for caching API responses

## Configuration

The application uses a configuration struct to manage state:

```go
type config struct {
   pokeapiClient    pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
    caughtPokemon    map[string]pokeapi.Pokemon
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

[PokéAPI](https://pokeapi.co/) for providing the Pokémon data

The Go community for their excellent libraries and tools

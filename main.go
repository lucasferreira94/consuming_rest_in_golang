package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// mapear toda resposta e trazer o nome da região + todos os pokemons dela
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// mapear todos os pokemons
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// mapear as especies Pokemon incluindo seus nomes
type PokemonSpecies struct {
	Name string `json:"name"`
}

const kanto string = "http://pokeapi.co/api/v2/pokedex/kanto/"

func main() {

	// faz a requisição ao endpoint e recebe a resposta.
	// caso a resposta contenha erro, irá printar na tela e sair do sistema
	response, err := http.Get(kanto)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// é feito a conversão da resposta do Body de []bytes para algo mais significativo como string na tela
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// deserializar a resposta e armazenar na variavel responseObject
	// a var responseObject possui o tipo struct Response que declaramo acima
	// o valor será apenas o nome da região >> kanto
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Name)

	// loop que irá iterar sobre todo objeto array da responseObject
	// depois printar desde o 1 pokemon até o último
	for idx, pokemon := range responseObject.Pokemon {
		fmt.Printf("[%d] %s \n", idx, pokemon.Species.Name)
	}

}

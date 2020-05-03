package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphery Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
}

func main() {
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "  ")

	if err != nil {
		log.Fatalf("marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	var moviesRead []Movie
	err = json.Unmarshal(data, &moviesRead)
	if err != nil {
		log.Fatalf("unmarshaling failed: %s", err)
	}
	fmt.Println("---------------")
	fmt.Println(moviesRead)

	var titleActors []struct {
		Title  string
		Actors []string
	}
	err = json.Unmarshal(data, &titleActors)
	if err != nil {
		log.Fatalf("unmarshaling failed: %s", err)
	}
	fmt.Println("---------------")
	fmt.Println(titleActors)

}

package main

import (
	"fmt"
	"legopl/ch4/github"
	"log"
	"os"
)

func main() {
	args := []string{"java", "elasticsearch"} //os.Args[1:]
	result, err := github.SearchIssues(args)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

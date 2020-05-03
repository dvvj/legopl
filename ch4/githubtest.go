package main

import (
	"fmt"
	"legopl/ch4/github"
	"log"
	"os"
	"time"
)

func ageDesc(createdAt time.Time) string {
	//	fmt.Println(createdAt)
	t := createdAt.AddDate(0, 1, 0)
	if t.After(time.Now()) {
		return "< 1month "
	} else {
		t1 := createdAt.AddDate(1, 0, 0)
		if t1.After(time.Now()) {
			return "[1month,  1year)"
		} else {
			return ">  1year "
		}
	}
}

func main() {
	args := os.Args[1:] //[]string{"golang", "elasticsearch", "deep", "learning"} //os.Args[1:]
	result, err := github.SearchIssues(args)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("%20.20s #%-5d %9.9s %.55s\n",
			ageDesc(item.CreatedAt), item.Number, item.User.Login, item.Title)
	}
}

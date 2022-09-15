package main

import (
	"fmt"
	"log"
	"os"
)

//run go build
//then ./github-issue-tracker repo:golang/go is:open json decoder
func main() {

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d Issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %s %s\n", item.Number, item.User.Login, item.Title, item.State, item.CreatedAt)
	}

}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get("http://" + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		s := resp.Status
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()

		fmt.Printf("%s\n", b)
		fmt.Printf("Status HTTP = %s", s)
	}
}

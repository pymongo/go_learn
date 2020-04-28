package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	httpRequest()
}

func httpRequest()  {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	} else {
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(bodyBytes))
		}
	}
}
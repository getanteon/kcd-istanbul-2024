package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("kcd-client running with pid: %d\n", os.Getpid())
	// wait for a trigger to send a request to the server
	for {
		fmt.Println("Press enter to send a request to the server")
		fmt.Scanln()

		response, err := http.Get("http://localhost:8080/")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := io.ReadAll(response.Body)
			fmt.Println(string(data))
			response.Body.Close()
		}
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	x := 0
	y := 1

	for i := 2; i <= n; i++ {
		temp := y
		y += x
		x = temp
	}

	return y
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(w, "%d\n", fib(i))
}

func getPort() string {
	port, found := os.LookupEnv("port")
	if !found {
		port = "8080"
	}
	return port
}

func main() {
	port := getPort()

	http.HandleFunc("/", mainHandler)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

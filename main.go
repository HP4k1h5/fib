package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func handleFib(w http.ResponseWriter, r *http.Request) {
	// handle route param
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil {
		errMsg := fmt.Sprintf(" :: couldn't convert parameter %v to integer:: %s\n", n, err)
		http.Error(w, err.Error()+errMsg, http.StatusBadRequest)
		return
	}

	// check memoized cache

	// otherwise calc val
	fib := Fib(n)

	// respond
	fmt.Fprintf(w, "%d", fib)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib/{n}", handleFib)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}

func Fib(n int) int {
	if n == 0 {
		return 0
	}

	var v = 0
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		v = a + b
		a = b
		b = v
	}

	return a
}

package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
	"strconv"
)

// global db var
var conn *pgx.Conn

func main() {
	connectPG()
	defer conn.Close(context.Background())

	handleRequests()
}

func connectPG() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgresql://hp4k1h5@localhost:5432/fib")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect %v\n", err)
		os.Exit(1)
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib/{n}", handleFib)
	router.HandleFunc("/memoized/{val}", handleMemoized)
	router.HandleFunc("/memoized", handleClearMemoized).Methods("DELETE")
	port := ":7357"
	fmt.Printf("api running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func handleFib(w http.ResponseWriter, r *http.Request) {
	// handle route param
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil {
		errMsg := fmt.Sprintf(" :: couldn't convert parameter %v to integer:: %s\n", n, err)
		http.Error(w, err.Error()+errMsg, http.StatusBadRequest)
		return
	}

	var val int = -1

	// check memoized cache
	rows, _ := conn.Query(context.Background(), "SELECT val FROM memo WHERE n = $1", n)
	for rows.Next() {
		err := rows.Scan(&val)
		if err != nil {
			errMsg := fmt.Sprintf(" :: couldn't query memcache for %d :: %s", n, err)
			http.Error(w, err.Error()+errMsg, http.StatusInternalServerError)
			return
		}
		fmt.Println("memoized", val)
	}

	// if no cache hit, calculate and store fib val
	if val == -1 {
		fmt.Println("no cache hit")
		val = Fib(n)
		_, err := conn.Exec(context.Background(), "INSERT INTO memo (n, val) VALUES ($1, $2)", n, val)
		if err != nil {
			errMsg := fmt.Sprintf(" :: couldn't query memcache for %d :: %s", n, err)
			http.Error(w, err.Error()+errMsg, http.StatusInternalServerError)
			return
		}
	}

	// respond
	fmt.Fprintf(w, "%d", val)
}

func handleMemoized(w http.ResponseWriter, r *http.Request) {
	// handle route param
	vars := mux.Vars(r)
	val, err := strconv.Atoi(vars["val"])
	if err != nil {
		errMsg := fmt.Sprintf(" :: couldn't convert parameter %v to integer:: %s\n", val, err)
		http.Error(w, err.Error()+errMsg, http.StatusBadRequest)
		return
	}

	// check memoized cache
	var count int
	rows, _ := conn.Query(context.Background(), "SELECT COUNT(*) FROM memo WHERE val < $1", val)
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			errMsg := fmt.Sprintf(" :: couldn't retrieve count for %d", val)
			http.Error(w, err.Error()+errMsg, http.StatusInternalServerError)
			return
		}
	}

	// respond
	fmt.Fprintf(w, "%d", count)
}

func handleClearMemoized(w http.ResponseWriter, r *http.Request) {
	// clear memoized cache
	_, err := conn.Exec(context.Background(), "DELETE FROM memo")
	if err != nil {
		errMsg := fmt.Sprintf(" :: couldn't clear memoized cache")
		http.Error(w, err.Error()+errMsg, http.StatusInternalServerError)
		return
	}
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

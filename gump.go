package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type CustomListener struct{}

func (c *CustomListener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	date := time.Now().UTC()
	fmt.Printf("%02d:%02d:%02d %s\n", date.Hour(), date.Minute(), date.Second(), r.Proto)
	fmt.Printf("%s %s\n", r.Method, r.URL)
	for header, values := range r.Header {
		fmt.Printf("%v: %v\n", header, strings.Join(values, ", "))
	}
	// Should only do this if r.Method is something that even accepts a body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if len(body) > 0 {
		fmt.Println(string(body))
	}
	fmt.Fprintln(w, "")
}

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Run Gump, run!")

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), &CustomListener{})
	if err != nil {
		log.Fatal(err)
	}
}

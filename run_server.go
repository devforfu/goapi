package main

import (
    "./src/api"
    "flag"
    "fmt"
)

func main() {
    url := getURL()
    db := api.Database{"socks": 5, "shoes": 50}
    api.RunServer(db, url)
}

func getURL() string {
    host := flag.String("-h", "localhost", "database address")
    port := flag.Int("-p", 8080, "database port")
    url := fmt.Sprintf("%s:%d", *host, *port)
    return url
}

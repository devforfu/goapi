package main

import (
    "./src/server"
    "flag"
    "fmt"
)

func main() {
    url := getURL()
    db := server.Database{"socks": 5, "shoes": 50}
    server.RunServer(db, url)
}

func getURL() string {
    host := flag.String("-h", "localhost", "database address")
    port := flag.Int("-p", 8080, "database port")
    url := fmt.Sprintf("%s:%d", *host, *port)
    return url
}

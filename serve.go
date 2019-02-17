package main

import (
    "./src/api"
    "context"
    "log"
    "net/http"
    "os"
    "sync"
)

func main() {
    conf := api.ParseConfig()
    db := api.Database{"socks": 5, "shoes": 50}
    server := api.CreateServer(db, conf)

    group := sync.WaitGroup{}
    group.Add(1)

    go func() {
        defer group.Done()
        if err := server.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("Server error: %s", err)
        }
    }()

    client := api.Client{BaseURL: conf.URL()}
    prices, _ := client.FetchPriceList()
    api.PrintPrices(os.Stdin, prices)
    server.Shutdown(context.TODO())
    group.Wait()
}



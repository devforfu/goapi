package api

import (
    "flag"
    "fmt"
    "io"
)

type ConnectionConfig struct {
    Host string
    Port int
}

func (c ConnectionConfig) Addr() string { return fmt.Sprintf("%s:%d", c.Host, c.Port) }

func (c ConnectionConfig) URL() string { return fmt.Sprintf("http://%s", c.Addr())}

func ParseConfig() ConnectionConfig {
    host := flag.String("-h", "localhost", "database address")
    port := flag.Int("-p", 8080, "database port")
    return ConnectionConfig{*host,*port}
}

func PrintPrices(w io.Writer, prices Database) {
    w.Write([]byte("List of prices:\n"))
    for k, v := range prices {
        w.Write([]byte(fmt.Sprintf("%s\t%v\n", k, v)))
    }
}
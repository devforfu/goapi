package api

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type Dollars float64
type Database map[string]Dollars
type Response map[string]interface{}

func notFound(w http.ResponseWriter, req *http.Request) {
    logMessage(req,"not found page request")
    json.NewEncoder(w).Encode(Response{"error": "not found"})
}

func (db Database) list(w http.ResponseWriter, req *http.Request) {
    logRequest(req)
    NewJSONResponse(w).SendDatabase(db)
}

func (db Database) price(w http.ResponseWriter, req *http.Request) {
    logRequest(req)
    item := req.URL.Query().Get("item")
    resp := NewJSONResponse(w)
    if price, ok := db[item]; !ok {
        resp.SendError(fmt.Sprintf("no such item: %s", item))
    } else {
        resp.SendPrice(item, price)
    }
}

func (db Database) update(w http.ResponseWriter, req *http.Request) {
    logRequest(req)
    data := make(map[string]string)
    json.NewDecoder(req.Body).Decode(&data)
    resp := NewJSONResponse(w)
    err := checkParameters(data, "item", "price")
    if err != nil {
        logMessage(req, err.Error())
        resp.SendError(err.Error())
        return
    }
    key, price := data["item"], data["price"]
    value, err := strconv.ParseFloat(price, 64)
    if err != nil {
        logMessage(req, err.Error())
        resp.SendError(err.Error())
        return
    }
    oldPrice := db[key]
    var mux sync.Mutex
    mux.Lock()
    db[key] = Dollars(value)
    mux.Unlock()
    resp.SendSuccess(Response{"item": key, "new": value, "old": oldPrice})
}

func getRequestString(req *http.Request) string {
    return fmt.Sprintf("%s: %s", req.RemoteAddr, req.RequestURI)
}

func logRequest(req *http.Request) {
    log.Print(getRequestString(req))
}

func logMessage(req *http.Request, message string) {
    log.Printf("%s: %s", getRequestString(req), message)
}

func checkParameters(data map[string]string, keys ...string) error {
    for _, key := range keys {
        if _, ok := data[key]; !ok {
            return fmt.Errorf("required key '%s' is missing", key)
        }
    }
    return nil
}

func RunServer(db Database, addr string) {
    mux := http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(notFound))
    mux.Handle("/list", http.HandlerFunc(db.list))
    mux.Handle("/price", http.HandlerFunc(db.price))
    mux.Handle("/update", http.HandlerFunc(db.update))
    log.Fatal(http.ListenAndServe(addr, mux))
}

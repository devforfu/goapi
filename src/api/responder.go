// Convenience tool to serialize api responses into valid JSON binary data.
package api

import (
    "encoding/json"
    "net/http"
)

// Responder serializes API-specific data structures into binary format and
// writes them using response writer.
type Responder struct {
    http.ResponseWriter
    *json.Encoder
}

func (r Responder) SendError(message string) {
    r.WriteHeader(http.StatusBadRequest)
    r.Encode(Response{"error": message})
}

func (r Responder) SendDatabase(db Database) {
    r.WriteHeader(http.StatusAccepted)
    r.Encode(db)
}

func (r Responder) SendPrice(item string, price Dollars) {
    r.WriteHeader(http.StatusAccepted)
    r.Encode(Response{"item": item, "price": price})
}

func (r Responder) SendSuccess(resp Response) {
    r.WriteHeader(http.StatusAccepted)
    r.Encode(resp)
}

// Creates a new responder to serialize the data into ResponseWriter.
func NewJSONResponse(w http.ResponseWriter) Responder {
    resp := Responder{}
    resp.ResponseWriter = w
    resp.Encoder = json.NewEncoder(w)
    return resp
}
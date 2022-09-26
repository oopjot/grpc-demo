package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	adder "github.com/oopjot/grpc-demo/adder/client"
	fib "github.com/oopjot/grpc-demo/fibonacci/client"
)

type addPayload struct {
    A int64 `json:"a"`
    B int64 `json:"b"`
}

type addResponse struct {
    Result int64 `json:"result"`
}

type fibResponse struct {
    Position int64 `json:"position"`
    Result int64 `json:"result"`
}

func AdderHandler(c *adder.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var payload addPayload
        if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
            return
        }
        res, err := c.Add(payload.A, payload.B)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }

        data, err := json.Marshal(&addResponse{Result: res})
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }
        w.Write(data)
    }
}


func FibNumberHandler(c *fib.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        n, err := strconv.ParseInt(vars["n"], 10, 64)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
        }

        res, err := c.Number(n)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }

        data, err := json.Marshal(&fibResponse{Result: res, Position: n})
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }
        w.Write(data)
    }
}

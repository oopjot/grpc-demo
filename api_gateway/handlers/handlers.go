package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	adder "github.com/oopjot/grpc-demo/adder/client"
	fib "github.com/oopjot/grpc-demo/fibonacci/client"
	"github.com/oopjot/grpc-demo/fibonacci/fibonacci"
)

type addResponse struct {
    Result int64 `json:"result"`
}

type fibResponse struct {
    Position int64 `json:"position"`
    Result int64 `json:"result"`
}

func AdderHandler(c *adder.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        a := r.FormValue("a")
        A, err := strconv.ParseInt(a, 10, 64)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
            return
        }
        b := r.FormValue("b")
        B, err := strconv.ParseInt(b, 10, 64)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
            return
        }

        res, err := c.Add(A, B)
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

func FibSeqHandler(c *fib.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        n, err := strconv.ParseInt(vars["n"], 10, 64)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
        }

        results := make(chan *fibonacci.FibResponse)

        go func() {
            err = c.Sequence(n, results)
        }()
        go func() {
            for res := range results {
                data, err := json.Marshal(res)
                if err == nil {
                    w.Write(data)
                    w.(http.Flusher).Flush()
                }
            }
        }()

        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }

    }
}

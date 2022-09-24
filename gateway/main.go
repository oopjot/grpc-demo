package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	adder "github.com/oopjot/grpc-demo/adder/client"
	fib "github.com/oopjot/grpc-demo/fibonacci/client"
	"github.com/oopjot/grpc-demo/gateway/handlers"
)

func main() {
    r := mux.NewRouter()

    adder, err := adder.New("adder", 50000)
    if err != nil {
        log.Printf("Adder service unavaliable: %v", err)
    } else {
        log.Println("Adder service loaded")
        r.HandleFunc("/add", handlers.AdderHandler(adder))
    }

    fib, err := fib.New("fibonacci", 50001)
    if err != nil {
        log.Printf("Fibonacci service unavaliable: %v", err)
    } else {
        log.Println("Fibonacci service loaded")
        r.HandleFunc("/fibonacci/{n:[0-9]+}", handlers.FibNumberHandler(fib))
    }
    log.Println("Gateway listening on 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
    

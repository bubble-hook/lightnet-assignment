package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bubble-hook/lightnet-assignment/proxy/handler"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/calculator.sum", handler.Calculate)
	http.HandleFunc("/calculator.mul", handler.Calculate)
	http.HandleFunc("/calculator.div", handler.Calculate)
	http.HandleFunc("/calculator.sub", handler.Calculate)
	fmt.Println("Proxy Server Start @ :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bubble-hook/lightnet-assignment/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/calculator.sum", handler.Calculate("sum"))
	http.HandleFunc("/calculator.mul", handler.Calculate("mul"))
	http.HandleFunc("/calculator.div", handler.Calculate("div"))
	http.HandleFunc("/calculator.sub", handler.Calculate("sub"))
	fmt.Println("Internal Server Start @ :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

package main

import (
	"fmt"
	"github.com/klein2ms/go-monkey-interpreter/lexer"
	"github.com/klein2ms/go-monkey-interpreter/token"
	"io"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
	fmt.Println(lexer.New(""))
	fmt.Println(token.ASSIGN)
}

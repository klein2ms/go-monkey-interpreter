package main

import (
	"fmt"
	"io"
	"lexer"
	"net/http"
	"token"
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

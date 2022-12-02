package main

import (
	"net/http"
)

// Funcao de requesicao
// writer = Response/res
// request = Request/req
func handler(writer http.ResponseWriter, request *http.Request) {
	// Envia arquivo index.html como resposta
	http.ServeFile(writer, request, "./static/index.html")
}

func main() {
	// Chama a funcao handler
	http.HandleFunc("/", handler)
	// Abre o server na porta 3000
	http.ListenAndServe(":3000", nil)
}

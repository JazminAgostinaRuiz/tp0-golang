package main

import (
	"net/http"
	"server/utils"
)

func main() {
	print("Iniciando")
	mux := http.NewServeMux()
	print("Escuchando .... ")
	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	print("Se recibió un paquete")
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)
	print("Se recibió un mensaje")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		print("Hubo error en el servidor")
		panic(err)
	}
}

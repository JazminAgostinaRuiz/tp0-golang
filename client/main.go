package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	log.Println("Soy un Log!")
	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente

	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Mensaje)
	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)
	// leer de la consola el mensaje
	mensajes := utils.LeerConsola()
	log.Println(mensajes)
	utils.GenerarYEnviarPaquete(mensajes, globals.ClientConfig.Ip, globals.ClientConfig.Puerto)
	// generamos un paquete y lo enviamos al servidor
}

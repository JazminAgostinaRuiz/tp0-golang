package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func LeerConsola() []string {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")
	var lineas []string
	//text, _ := reader.ReadString('\n')
	//log.Print(text)
	for {
		text, _ := reader.ReadString('\n')
		//log.Print(len(text))
		//text == "\r\n"
		//len(text) == 1
		if len(text) == 1 {
			log.Print("Carga finalizada")
			break
		}
		log.Print(text)
		lineas = append(lineas, text)
	}
	//len(text) == 0 parece que siempre lee un espacio al menos
	//log.Print(lineas)
	return lineas

}

func GenerarYEnviarPaquete(mensajes []string, ip string, puerto int) {
	paquete := Paquete{}
	// Leemos y cargamos el paquete
	paquete.Valores = mensajes
	log.Printf("paquete a enviar: %+v", paquete)
	// Enviamos el paqute
	EnviarPaquete(ip, puerto, paquete)
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}

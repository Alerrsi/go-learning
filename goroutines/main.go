package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func readFile(ruta string, wg *sync.WaitGroup) {

	defer wg.Done()

	datos, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Printf("Error al leer el archivo %s %v}\n", ruta, err)
		return
	}
	firtsLine := strings.Split(string(datos), "\n")[0]
	fmt.Printf("[%s] -> %s\n", ruta, firtsLine)

}

func main() {
	var wg sync.WaitGroup
	archivos := []string{"/proc/loadavg", "/proc/uptime", "/proc/version", "/var/log"}

	for _, archivo := range archivos {
		wg.Add(1)
		go readFile(archivo, &wg) // con & especificamos que utilize exactamente el mismo wg referenciando la memoria
	}

	wg.Wait() // terminar las rutinas y llegar a 0 antes dec cerrar main
	fmt.Println("recoleccion de datos terminada")
}

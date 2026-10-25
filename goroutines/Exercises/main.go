package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	archivos := []string{"/etc/passwd", "/etc/fstab", "etc/archivo_inexistente"}

	for _, archivo := range archivos {
		wg.Add(1)
		go exists(archivo, &wg) // con & especificamos que utilize exactamente el mismo wg referenciando la memoria
	}

	wg.Wait() // terminar las rutinas y llegar a 0 antes dec cerrar main
	fmt.Println("recoleccion de datos terminada")
}

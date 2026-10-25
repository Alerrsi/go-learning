package main

import (
	"fmt"
	"slices"
)

type Entrada struct {
	Hora  string
	Nivel string // "INFO", "WARN" o "ERROR"
	Msg   string
	Ms    int // tiempo de respuesta en milisegundos
}

func datos() []Entrada {
	return []Entrada{
		{"10:00:01", "INFO", "servidor iniciado", 12},
		{"10:00:05", "INFO", "usuario conectado", 45},
		{"10:00:09", "WARN", "disco al 85%", 230},
		{"10:00:14", "ERROR", "conexion rechazada", 980},
		{"10:00:20", "INFO", "consulta ejecutada", 60},
		{"10:00:27", "ERROR", "timeout en base de datos", 1500},
		{"10:00:31", "WARN", "memoria alta", 310},
		{"10:00:40", "INFO", "usuario desconectado", 20},
		{"10:00:44", "ERROR", "archivo no encontrado", 150},
		{"10:00:50", "INFO", "backup completado", 700},
	}
}

// info, warm, error
func contar(logs []Entrada) (int, int, int) {
	// zero value
	var info int
	var warn int
	var error int

	for _, log := range logs {
		if log.Nivel == "INFO" {
			info++
		} else if log.Nivel == "WARN" {
			warn++
		} else {
			error++
		}
	}

	return info, warn, error

}

func filtrar(logs []Entrada, nivel string) []Entrada {
	var lista []Entrada
	for _, log := range logs {
		if log.Nivel == nivel {
			lista = append(lista, log)
		}
	}

	return lista
}

// estadisticas

func estadisticas(logs []Entrada) []float64 {
	var promedio float64
	var maximo float64
	minimo := 999999999.0
	var suma float64

	for _, log := range logs {
		msI := float64(log.Ms)
		if minimo > msI {
			minimo = msI
		}
		if maximo < msI {
			maximo = msI
		}
		suma = suma + msI
	}

	promedio = suma / float64(len(logs))

	stats := []float64{promedio, maximo, minimo}

	return stats

}

func low_inputs(logs []Entrada) []Entrada {

	var lowMs []Entrada

	for _, log := range logs {
		if log.Ms > 500 {
			lowMs = append(lowMs, log)
		}
	}

	return lowMs
}

func deleteInfo(logs []Entrada) []Entrada {
	wInfo := make([]Entrada, len(logs))
	copy(wInfo, logs)

	wInfo = slices.DeleteFunc(wInfo, func(l Entrada) bool {
		return l.Nivel == "INFO"
	})

	return wInfo
}

// creación de propios tipos basados en existentes
type Mitipo Entrada

// funcion principal
func main() {
	// se imprime al final con el fin de limpiar o cerrar archivos
	defer fmt.Println("Final de programa")

	logs := datos()
	info, warn, error := contar(logs)
	fmt.Println(info, warn, error)

	// filtrar por error
	filtro := filtrar(logs, "ERROR")

	for _, log := range filtro {
		fmt.Println("Hora de incidente", log.Hora, "Mensaje: ", log.Msg)

	}

	// estadisticas
	st := estadisticas(logs)
	fmt.Println(st)

	// entradas lentas
	lentas := low_inputs(logs)

	for i, l := range lentas {
		fmt.Println(fmt.Sprintf("%d . %s - %s", i, l.Hora, l.Msg))
	}

	// Entradas sin nivel INFO
	//
	fmt.Println("-------------------SIN INFOS-------------------")
	wInfo := deleteInfo(logs)
	fmt.Println("tamaño de copia sin INFO: ", len(wInfo), "tamaño original: ", len(logs))

}

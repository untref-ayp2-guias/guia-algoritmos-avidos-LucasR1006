package ejercicios

import "sort"

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	resultado := make(map[Objeto]float32)
	//Ordenamos de mayor a menor
	sort.Slice(objetos, func(i, j int) bool {
		return float32(objetos[i].Valor)/float32(objetos[i].Peso) > float32(objetos[j].Valor)/float32(objetos[j].Peso)
	})

	capacidadRestante := capacidad
	for _, obj := range objetos {
		if capacidadRestante == 0 {
			break
		}
		if obj.Peso <= capacidadRestante {
			capacidadRestante -= obj.Peso
			resultado[obj] = 1.0
		} else {
			fraccion := float32(capacidadRestante) / float32(obj.Peso)
			resultado[obj] = fraccion
			capacidadRestante = 0
		}
	}
	return resultado
}

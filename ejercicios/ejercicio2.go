package ejercicios

import "sort"

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {

	sort.Slice(tareas, func(i, j int) bool { return tareas[i].Tiempo < tareas[j].Tiempo })

}

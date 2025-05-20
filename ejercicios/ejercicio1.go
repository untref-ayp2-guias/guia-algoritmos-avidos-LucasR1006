package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {
	arreglo := make([]Actividad, 0)

	arreglo = append(arreglo, actividades[0])

	return Selector(actividades[1:], arreglo)
}

func Selector(actividades []Actividad, array []Actividad) []Actividad {
	if len(actividades) == 1 {
		return append(array, actividades[0])
	}
	if actividades[0].Inicio >= array[len(array)-1].Fin {
		array = append(array, actividades[0])
	}
	return Selector(actividades[1:], array)
}

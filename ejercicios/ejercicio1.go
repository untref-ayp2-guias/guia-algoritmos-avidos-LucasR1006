package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {
	pepe2 := make([]Actividad, 0)

	pepe2 = append(pepe2, actividades[0])

	return Selector(actividades[1:], pepe2)
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
